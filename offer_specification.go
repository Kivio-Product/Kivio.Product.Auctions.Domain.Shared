package offer

import "time"

type OfferSpecification struct {
	ExternalID string
	CreatedAt  time.Time
	OfferID    string
}
