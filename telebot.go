package main

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"log"
	"time"

	"github.com/disintegration/imaging"
	"gopkg.in/telebot.v3"
)

func runTeleBot() {
	pref := telebot.Settings{
		Token:  "6636321864:AAF1KWLWVTkftegnrWu0BYG_OCFK3VYzwp0",
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("shaba", func(ctx telebot.Context) error {
		shaba := ctx.Message().Payload

		// validate shaba

		return ctx.Send("your shaba is ", shaba)
	})

	b.Handle(telebot.OnPhoto, func(ctx telebot.Context) error {
		photo := ctx.Message().Photo
		if photo == nil {
			return ctx.Send("This is not a photo")
		}

		jpegImage, err := jpeg.Decode(photo.FileReader)
		if err != nil {
			return ctx.Send("server error: jpeg.Decode: ", err.Error())
		}

		inverted := imaging.Invert(jpegImage)

		buf := bytes.Buffer{}
		if err := jpeg.Encode(&buf, inverted, nil); err != nil {
			return ctx.Send("server error: jpeg.Encode: ", err.Error())
		}

		tp := telebot.Photo{
			File: telebot.FromReader(&buf),
		}

		return ctx.Send(tp)
	})

	b.Handle("/hello", func(c telebot.Context) error {
		fmt.Println("new /hello request")

		fname := c.Sender().FirstName
		uname := c.Sender().Username

		return c.Send(fmt.Sprintf("Hi %s (%s)", fname, uname))
	})

	b.Handle("/bye", func(c telebot.Context) error {
		fmt.Println("new /bye request")

		fname := c.Sender().FirstName
		uname := c.Sender().Username

		return c.Send(fmt.Sprintf("Bye %s (%s)", fname, uname))
	})

	b.Handle("/foo", func(c telebot.Context) error {
		fmt.Println("new /foo request")

		fname := c.Sender().FirstName
		uname := c.Sender().Username

		return c.Send(fmt.Sprintf("This is foo %s (%s)", fname, uname))
	})

	b.Start()
}
