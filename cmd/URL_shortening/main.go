package main

import "URL_shortening/internal/app"

func main() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}
