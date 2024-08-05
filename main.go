package main

import (
	"maguro-alternative/go-go-go/route"
)

func main() {
	r := route.Routes()
	r.Run(":8080")
}
