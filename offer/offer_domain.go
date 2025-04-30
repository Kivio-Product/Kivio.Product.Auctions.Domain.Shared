package domain

import (
	"errors"
	"time"
)

type Offer struct {
	OfferId     string
	CreatedAt   time.Time
	Description string
	Name        string
	PosId       string
	State       string
	Type        string
	AuctionTime int64
}

var (
	StateInactive = "Created"
	StateActive   = "Offered"
)

func (o *Offer) Update(name, description string) error {
	if name == "" {
		return errors.New("El nombre no puede estar vacío")
	}
	if description == "" {
		return errors.New("La descripcion no puede estar vacía")
	}
	o.Name = name
	o.Description = description
	return nil
}

func (o *Offer) UpdateState(state string) error {
	if state == "" {
		return errors.New("El estado no puede estar vacío")
	}
	o.State = state
	return nil
}
