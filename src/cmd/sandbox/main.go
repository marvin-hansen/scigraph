package main

import "github.com/marvin-hansen/goC8"

func main() {
	println("Main ")

	client := goC8.NewClient(nil)
	client.PubFunc()
}
