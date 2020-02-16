// The user only needs an interface that provides him
// this value. By delegating this decision to a Factory, this Factory can provide an interface that
// fits the user needs. It also eases the process of downgrading or upgrading of the
// implementation of the underlying type if needed.

// . Imagine that you want to
// organize your holidays using a trip agency. You don't deal with hotels and traveling and
// you just tell the agency the destination you are interested in so that they provide you with
// everything you need. The trip agency represents a Factory of trips.

package factory

import (
	"errors"
	"fmt"

)
const (
	Cash      = 1
	DebitCard = 2
)

type PaymentMethod interface {
	Pay(amount float32) string
}

func GetPaymentMethod(m int) (PaymentMethod, error) {
	// return nil, errors.New("Not implemented yet")
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	case CreditCardPM:
		return new(CreditCardPM), nil
	default:
		return nil, errors.New("Not a valid payment method")
	}

}
type CashPM struct{}
func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash\n", amount)
}


type DebitCardPM struct{}
func (c *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using debit card\n", amount)

}

type CreditCardPM struct{}

func (d *CreditCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using new credit card implementation\n",amount)
}
