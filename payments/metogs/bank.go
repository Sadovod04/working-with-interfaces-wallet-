package metogs

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type bank struct {
}

func (c bank) Pay(usd int) int {
	fmt.Println("оплата картой")
	fmt.Println("размер оплаты:", usd)
	n, _ := rand.Int(rand.Reader, big.NewInt(1000000)) // случайное число до 999999
	return int(n.Int64())
}

func (c bank) Cancel(id int) {
	fmt.Println("отмена платежа", "ID", id)
}
func NewBank() bank {
	return bank{}
}
