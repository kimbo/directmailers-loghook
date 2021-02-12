# directmailers-loghook for logrus

Have you ever wanted to get a postcard in the mail the last time your program crashed?

This is a hook for [logrus](https://github.com/sirupsen/logrus) that will send your logs to you (or some other address of your choosing) in a letter or post card in the mail.

# Requirements

As the name suggests, this library uses the directmailers API for sending letters and postcards.
So you'll need an account with them.
Signing up is free, and they'll give you $10 in credit when you do, which should be enough to send about 17 postcards (see https://print.directmailers.com/tour/pricing/).

# Installation and usage

```
go get github.com/kimbo/directmailers-loghook/...
```

Then you can initialize the hook like this:

```go
package main

import (
	"errors"
	"os"

	"github.com/kimbo/directmailers-loghook/api"
	"github.com/kimbo/directmailers-loghook/hook"
	"github.com/sirupsen/logrus"
)

func main() {
    // get these at https://dashboard.directmailers.com/settings
	user := os.Getenv("DIRECTMAILER_USERNAME")
	pass := os.Getenv("DIRECTMAILER_PASSWORD")

	h := hook.New(user, pass, hook.Config{
		DryRun: true,
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
			Name:         "The White House", // why not send you logs to the President?
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
```

# Roadmap

- Improve default template for letters and postcards (right now it's just `<p>log message goes here</p>`)
- Clean way to use custom templates

# Limitations

I don't know this for sure, but I suspect directmailers only will send mail in the United States of America.
