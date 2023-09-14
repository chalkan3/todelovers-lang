package engine

import (
	"math"
	"math/rand"
)

type AdressGenerator struct {
	generatedAddresses map[int]bool
}

func NewAdressGenerator() *AdressGenerator {
	return &AdressGenerator{
		generatedAddresses: make(map[int]bool),
	}
}

func (ad *AdressGenerator) SetAdress(sbl Symblo) {
	address := rand.Intn(math.MaxInt64)
	if _, ok := ad.generatedAddresses[address]; !ok {
		ad.generatedAddresses[address] = true

		sbl.Address = byte(address)
	}

}
