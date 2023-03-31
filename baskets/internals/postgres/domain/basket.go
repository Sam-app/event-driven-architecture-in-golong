package domain

import "github.com/stackus/errors"

var (
	ErrorBasketHasNoItems       = errors.Wrap(errors.ErrBadRequest, "the basket has no items")
	ErrBasketCannotBeModified   = errors.Wrap(errors.ErrBadRequest, "the basket cannot be modified")
	ErrBasketCannotBeCancelled  = errors.Wrap(errors.ErrBadRequest, "the basket cannot be modified")
	ErrQuantityCannotBeNegative = errors.Wrap(errors.ErrBadRequest, "the item quantity cannot be negative")
	ErrBasketIDCannotBeBlank    = errors.Wrap(errors.ErrBadRequest, "the basket id cannot be black")
	ErrPaymentIDCannotBeBlank   = errors.Wrap(errors.ErrBadRequest, "the payment id cannot be blank")
	ErrCustomerIDCannotBeBlank  = errors.Wrap(errors.ErrBadRequest, "the customer id cannot be blank")
)

type BasketStatus string

const (
	BasketUnknown    BasketStatus = ""
	BasketOpen       BasketStatus = "open"
	BasketCancelled  BasketStatus = "cancelled"
	BasketCheckedOut BasketStatus = "checked_out"
)

func (s BasketStatus) String() string {
	switch s {
	case BasketOpen, BasketCancelled, BasketCheckedOut:
		return string(s)
	default:
		return ""
	}
}

type Basket struct {
	ID        string
	CustomID  string
	PaymentID string
	Items     []Item
	Status    BasketStatus
}

func StartBasket(id, customerID string) (*Basket, error) {
	if id == "" {
		return nil, ErrBasketIDCannotBeBlank
	}

	if customerID == "" {
		return nil, ErrCustomerIDCannotBeBlank
	}

	basket := &Basket{
		ID:       id,
		CustomID: customerID,
		Status:   BasketOpen,
		Items:    []Item{},
	}

	return basket, nil
}

func (b *Basket) isOpen() bool {
	return b.Status == BasketOpen
}

func (b *Basket) isCancellable() bool {
	return b.Status == BasketOpen
}

func (b *Basket) Checkout(PaymentID string) error {
	if !b.isOpen() {
		return ErrBasketCannotBeModified
	}

	if len(b.Items) == 0 {
		return ErrorBasketHasNoItems
	}

	if PaymentID == "" {
		return ErrPaymentIDCannotBeBlank
	}

	b.PaymentID = PaymentID
	b.Status = BasketCheckedOut

	return nil
}
