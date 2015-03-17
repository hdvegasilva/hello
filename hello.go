package main

import "./greeting"

func main() {

	var s = Salutation{"Bob", "Hello"}
	Greet(s, CreatePrintCustom("?"), false)
}
