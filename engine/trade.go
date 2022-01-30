package engine

import "encoding/json"

type Trade struct {
	TakerId uint64 `json:"taker_id"`
	MakerId uint64 `json:"maker_id"`
	Price   uint64 `json:"price"`
	Amount  uint64 `json:"amount"`
}

func (Trade *Trade) DeSerialize(message []byte) error {
	return json.Unmarshal(message, Trade)
}

func (Trade *Trade) Serialize() []byte {
	str, _ := json.Marshal(Trade)
	return str
}
