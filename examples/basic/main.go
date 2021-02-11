//
// Example program that uses the logrus log hook from directmailers-loghook
//
package main

import (
	"errors"
	"os"

	"github.com/kimbo/directmailers-loghook/api"
	"github.com/kimbo/directmailers-loghook/hook"
	"github.com/sirupsen/logrus"
)

func main() {
	user := os.Getenv("DIRECTMAILER_USERNAME")
	pass := os.Getenv("DIRECTMAILER_PASSWORD")

	h := hook.New(user, pass, hook.Config{
		MaxLevel: logrus.ErrorLevel,
		MailType: hook.Postcard, // or hook.Letter
		From: api.SenderDetails{
			Name:         "my-application",
			AddressLine1: "123 N 456 W",
			AddressLine2: "Apt 789",
			City:         "San Francisco",
			State:        "CA",
			Zip:          "12345",
		},
		To: api.RecipientDetails{
			Name:         "The White House",
			AddressLine1: "1600 Pennsylvania Ave",
			AddressLine2: "",
			City:         "Washington",
			State:        "DC",
			Zip:          "20006",
		},
	})
	logrus.AddHook(h)

	doSomething := func() error {
		return errors.New("I was running and tripped over the waste basket and now there's garbage everywhere")
	}

	if err := doSomething(); err != nil {
		// that's another $0.56, errors get expensive quickly
		logrus.Errorf("Oh no, we have a problem: %v", err)
	}
}
