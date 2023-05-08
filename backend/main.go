package main

import (
	"webservice/data"
	"webservice/web"
)

func main() {
	err := data.Init()
	if err != nil {
		panic(err)
	}

	web.Serve()
}
