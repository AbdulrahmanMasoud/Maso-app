package main

import (
	"github.com/AbdulrahanMasoud/maso"
	"log"
	"os"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	//init maso
	mso := &maso.Maso{}
	err = mso.New(path)
	if err != nil {
		log.Fatal(err)
	}
	mso.AppName = "masoApp"
	mso.Debug = true

	app := &application{
		App: mso,
	}
	return app
}
