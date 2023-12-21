package main

import (
	"fmt"
	"log"
	"os"

	"training.go/scriptPixelPerfect/config"
	"training.go/scriptPixelPerfect/script"
	"training.go/scriptPixelPerfect/store"
)

func main() {
	cfgFile := config.MustOpenConfigFile(".config.json")
	configModule := config.New(cfgFile, config.GetParserFromFileExtension(".config.json"))

	storer := store.New(*configModule.Store())
	if err := storer.Open(); err != nil {
		log.Fatal(err)
	}

	defer storer.Close()

	scripter := script.New(storer)

	resultat, err := scripter.GetScript()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile("script.sql", os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	_, err = file.WriteString(resultat)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Insert done")
}
