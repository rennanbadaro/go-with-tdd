package main

const greetingPrefix string = "Hi there, "

func Hello(name string) string {
	if name == "" {
		return greetingPrefix + "stranger"
	}

	return greetingPrefix + name
}
