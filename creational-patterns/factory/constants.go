package factory

import (
	"errors"
	"fmt"
)

type PaymentMethodType int

const (
	Cash PaymentMethodType = iota
	DebitCard
)

func GetPaymentMethod(m PaymentMethodType) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	default:
		return nil, errors.New(fmt.Sprintf("Payment method %d not recognized\n", m))
	}
}
