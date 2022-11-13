package main

import (
	"log"

	"github.com/annadymovaa/avito-test/blob/start/http-rest-api/cmd/inetrnal/app/apiserver"
)

func main() {
	s := apiserver.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
