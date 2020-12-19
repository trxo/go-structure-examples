package adding

import (
	"github.com/trxo/go-structure-examples/domain/beers"
)

// Service provides beer adding operations
type Service interface {
	AddBeer(...beers.Beer)
	AddSampleBeers()
}

type service struct {
	bR beers.Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(bR beers.Repository) Service {
	return &service{bR}
}

// AddBeer adds the given beer(s) to the database
func (s *service) AddBeer(b ...beers.Beer) {
	for _, beer := range b {
		_ = s.bR.Add(beer) // error handling omitted for simplicity
	}
}

// AddSampleBeers adds some sample beers to the database
func (s *service) AddSampleBeers() {
	for _, b := range beers.DefaultBeers {
		_ = s.bR.Add(b) // error handling omitted for simplicity
	}
}
