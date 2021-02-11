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
		Data:            "",
	}
}

func DefaultPostcard() api.PostcardRequest {
	return api.PostcardRequest{
		Description:     "Direct mailer log",
		Size:            "4.25x6",
		DryRun:          true, // dryrun by default to avoid costing $$
		WaitForRender:   true,
		VariablePayload: nil,
		Front:           "",
		Back:            "",
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
}

func New(username, password string, config Config) *DirectMailerHook {
	return &DirectMailerHook{
		api:      api.New(username, password),
		Letter:   DefaultLetter(),
		Postcard: DefaultPostcard(),
		Config:   config,
	}
}

type DirectMailerHook struct {
	Config   Config
	api      *api.API
	Letter   api.LetterRequest
	Postcard api.PostcardRequest
}

func (h *DirectMailerHook) fireLetter(entry *logrus.Entry) error {
	s, err := entry.String()
	if err != nil {
		return err
	}
	h.Letter.Data = fmt.Sprintf("<p>%s</p>", s)
	h.Letter.To = h.Config.To
	h.Letter.From = h.Config.From
	lres, err := h.api.CreateLetter(h.Letter)
	if err != nil {
		return err
	}
	// TODO: remove this or print .RenderedPdf?
	fmt.Printf("%#v\n", lres)
	return nil
}

func (h *DirectMailerHook) firePostcard(entry *logrus.Entry) error {
	h.Postcard.Front = "<p>logrus directmailer hook wishes you the best of luck with your logs!</p>"
	s, err := entry.String()
	if err != nil {
		return err
	}
	h.Postcard.Back = fmt.Sprintf("<p>%s</p>", s)
	h.Postcard.To = h.Config.To
	h.Postcard.From = h.Config.From
	pres, err := h.api.CreatePostcard(h.Postcard)
	if err != nil {
		return err
	}
	// TODO: remove this or print .RenderedPdf?
	fmt.Printf("%#v\n", pres)
	return nil
}

func (h *DirectMailerHook) Fire(entry *logrus.Entry) error {
	if entry.Level > h.Config.MaxLevel {
		return nil
	}
	if h.Config.MailType == Letter {
		return h.fireLetter(entry)
	}
	if h.Config.MailType == Postcard {
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
