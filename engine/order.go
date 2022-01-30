package engine

import (
	"encoding/json"
)

type Direction int

type Order struct {
	Id        uint64    `json:"id"`
	Price     uint64    `json:"price"`
	Amount    uint64    `json:"amount"`
	Direction Direction `json:"side"`
}

func (Order *Order) DeSerialize(message []byte) error {
	return json.Unmarshal(message, Order)
}

func (Order *Order) Serialize() []byte {
	str, _ := json.Marshal(Order)
	return str
}

const (
	BuySide  Direction = 1
	SellSide           = 2
)
