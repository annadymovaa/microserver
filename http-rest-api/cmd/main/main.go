package apiserver

import (
	"log"

	"github.com/annadymovaa/avito-test/tree/main/http-rest-api/cmd/inetrnal/app/apiserver"
	//"github.com/annadymovaa/avito-test/tree/start/http-rest-api/cmd/inetrnal/app/apiserver"
)

func main() {
	s := apiserver.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
