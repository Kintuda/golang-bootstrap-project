package main

import (
	"log"

	cmd "github.com/Kintuda/golang-bootstrap-project/cmd/appname"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
