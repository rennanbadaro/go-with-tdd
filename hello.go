package main

import "strings"

const defaultLanguage = "english"

var languagesGreetMapping = map[string]string{
	"english": "Hi there, ",
	"spanish": "Hola, ",
	"french":  "Bonjour, ",
}

func Hello(name string, language string) string {
	if name == "" {
		name = "stranger"
	}

	language = strings.ToLower(language)

	greeting := languagesGreetMapping[language]

	if greeting == "" {
		greeting = languagesGreetMapping[defaultLanguage]
	}

	return greeting + name
}
