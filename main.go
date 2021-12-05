package main

import (
	"log"

	cmd "github.com/Kintuda/golang-bootstrap-project/cmd/appname"
)

func main() {
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
