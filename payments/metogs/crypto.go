package metogs

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type CryptoPayment struct {
}

func (c CryptoPayment) Pay(usd int) int {
	fmt.Println("оплата криптой")
	fmt.Println("размер оплаты:", usd)
	n, _ := rand.Int(rand.Reader, big.NewInt(1000000)) // случайное число до 999999
	return int(n.Int64())
}

func (c CryptoPayment) Cancel(id int) {
	fmt.Println("отмена крипто-платежа", "ID", id)
}
func NewCrypto() CryptoPayment {
	return CryptoPayment{}
}
