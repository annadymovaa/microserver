package main

import (
	"log"

	"github.com/avito-test/http-rest-api/cmd/inetrnal/app/apiserver"
)

func main() {
	s := apiserver.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
