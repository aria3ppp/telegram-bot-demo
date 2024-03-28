package main

import (
	"fmt"
	"log"
	"time"

	tele "gopkg.in/telebot.v3"
)

func runTeleBot() {
	pref := tele.Settings{
		Token:  "6636321864:AAF1KWLWVTkftegnrWu0BYG_OCFK3VYzwp0",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(c tele.Context) error {
		fmt.Println("new /hello request")

		fname := c.Sender().FirstName
		uname := c.Sender().Username

		return c.Send(fmt.Sprintf("Hi %s (%s)", fname, uname))
	})

	b.Handle("/bye", func(c tele.Context) error {
		fmt.Println("new /bye request")

		fname := c.Sender().FirstName
		uname := c.Sender().Username

		return c.Send(fmt.Sprintf("Bye %s (%s)", fname, uname))
	})

	b.Handle("/foo", func(c tele.Context) error {
		fmt.Println("new /foo request")

		fname := c.Sender().FirstName
		uname := c.Sender().Username

		return c.Send(fmt.Sprintf("This is foo %s (%s)", fname, uname))
	})

	b.Start()
}
