package offer

import (
	"errors"
	"time"
)

type Offer struct {
	OfferId     string
	CreatedAt   time.Time
	Description string
	ExpireAt    time.Time
	Name        string
	PosId       string
	State       string
	Type        string
}

var (
	StateInactive = "Created"
	StateActive   = "Offered"
)

func (o *Offer) Update(name, description string, expireAt time.Time) error {
	if name == "" {
		return errors.New("El nombre no puede estar vacío")
	}
	if description == "" {
		return errors.New("La descripcion no puede estar vacía")
	}
	if expireAt.IsZero() {
		return errors.New("La fecha de expiración no puede estar vacía")
	}
	o.Name = name
	o.Description = description
	o.ExpireAt = expireAt
	return nil
}

func (o *Offer) UpdateState(state string) error {
	if state == "" {
		return errors.New("El estado no puede estar vacío")
	}
	o.State = state
	return nil
}
