package order

import (
	"github.com/annadymovaa/avito-test/inetrnal/app/account"
	"github.com/annadymovaa/avito-test/inetrnal/app/service"
)

type Order struct {
	Id_order uint64          `json:'id_order`
	Service  service.Service `json:"id_service"`
	Account  account.Account `json:"id_account"`
}
