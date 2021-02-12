package hook

import (
	"fmt"

	"github.com/kimbo/directmailers-loghook/api"
	"github.com/sirupsen/logrus"
)

func DefaultLetter() api.LetterRequest {
	return api.LetterRequest{
		Description:     "Direct mailer log",
		Size:            "8.5x11",
		Duplex:          false,
		DryRun:          true, // dryrun by default to avoid costing $$
		WaitForRender:   true,
		BlankFirstPage:  false,
		PostalClass:     "First Class",
		VariablePayload: nil,
		Data:            letterDefault,
	}
}

func DefaultPostcard() api.PostcardRequest {
	return api.PostcardRequest{
		Description:     "Direct mailer log",
		Size:            "4.25x6",
		DryRun:          true, // dryrun by default to avoid costing $$
		WaitForRender:   true,
		VariablePayload: nil,
		Front:           postcardDefaultFront,
		Back:            postcardDefaultBack,
	}
}

type MailType int

const (
	Postcard MailType = iota
	Letter
)

type Config struct {
	To                    api.RecipientDetails
	From                  api.SenderDetails
	MaxLevel              logrus.Level
	MailType              MailType
	DirectmailersUsername string
	DirectmailersPassword string
	DryRun                bool
}

func New(username, password string, config Config) *DirectMailerHook {
	h := &DirectMailerHook{
		api:      api.New(username, password),
		Letter:   DefaultLetter(),
		Postcard: DefaultPostcard(),
		config:   config,
	}
	h.Letter.To = h.config.To
	h.Letter.From = h.config.From
	h.Letter.DryRun = h.config.DryRun
	h.Postcard.To = h.config.To
	h.Postcard.From = h.config.From
	h.Postcard.DryRun = h.config.DryRun
	return h
}

type DirectMailerHook struct {
	config   Config
	api      *api.API
	Letter   api.LetterRequest
	Postcard api.PostcardRequest
}

func (h *DirectMailerHook) fireLetter(entry *logrus.Entry) error {
	logMsg, err := entry.String()
	if err != nil {
		return err
	}
	h.Letter.VariablePayload = map[string]string{
		"appName":       h.config.From.Name,
		"recipientName": h.config.To.Name,
		"logMessage":    logMsg,
	}
	lres, err := h.api.CreateLetter(h.Letter)
	if err != nil {
		return err
	}
	fmt.Printf("Rendered PDF: %v\n", lres.RenderedPdf)
	return nil
}

func (h *DirectMailerHook) firePostcard(entry *logrus.Entry) error {
	logMsg, err := entry.String()
	if err != nil {
		return err
	}
	h.Postcard.VariablePayload = map[string]string{
		"appName":       h.config.From.Name,
		"recipientName": h.config.To.Name,
		"logMessage":    logMsg,
	}
	pres, err := h.api.CreatePostcard(h.Postcard)
	if err != nil {
		return err
	}
	fmt.Printf("Rendered PDF: %v\n", pres.RenderedPdf)
	return nil
}

func (h *DirectMailerHook) Fire(entry *logrus.Entry) error {
	if entry.Level > h.config.MaxLevel {
		return nil
	}
	if h.config.MailType == Letter {
		return h.fireLetter(entry)
	}
	if h.config.MailType == Postcard {
		return h.firePostcard(entry)
	}
	// do nothing
	return nil
}

func (h *DirectMailerHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}
}
