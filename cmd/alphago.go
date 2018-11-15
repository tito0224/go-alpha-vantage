package main

import (
	"github.com/tito0224/go-alpha-vantage"
)

func main() {
	println("Hello world")

	client := alphago.NewDefaultClient("")
	q, err := client.GetQuote("MSFT")
	if err != nil {
		panic(err)
	}

	println(q.PrevClose)

	r, err := client.Search("micro")
	if err != nil {
		panic(err)
	}

	println(r[0].Name)
}
