package main

import (
	"wallet/payments"
	"wallet/payments/metogs"

	"github.com/k0kubun/pp"
)

func main() {
	paypal := metogs.NewPayPal()
	paymentModel := payments.NewPaymentModel(paypal)
	cok := paymentModel.Pay("сок", 100)
	paymentModel.Pay("учеба ", 270)
	paymentModel.Pay("мак купил наконецто", 1100)
	paymentModel.Cancel(cok)
	AllInfo := paymentModel.AllInfo()
	pp.Println(AllInfo)

}
