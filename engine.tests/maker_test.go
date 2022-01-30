package engine_tests

import (
	"MatchEngine/engine"
	"github.com/stretchr/testify/assert"
	"testing"
)

func giveOrderBook(book *engine.MakerBook) {
	book.SellOrders = append(book.SellOrders, engine.Order{Price: 110, Id: 5, Amount: 5, Direction: engine.SellSide})
	book.SellOrders = append(book.SellOrders, engine.Order{Price: 108, Id: 4, Amount: 4, Direction: engine.SellSide})
	book.SellOrders = append(book.SellOrders, engine.Order{Price: 106, Id: 3, Amount: 3, Direction: engine.SellSide})
	book.SellOrders = append(book.SellOrders, engine.Order{Price: 104, Id: 2, Amount: 2, Direction: engine.SellSide})
	book.SellOrders = append(book.SellOrders, engine.Order{Price: 100, Id: 1, Amount: 1, Direction: engine.SellSide})

	book.BuyOrders = append(book.BuyOrders, engine.Order{Price: 1000, Id: 10, Amount: 10, Direction: engine.BuySide})
	book.BuyOrders = append(book.BuyOrders, engine.Order{Price: 1100, Id: 11, Amount: 20, Direction: engine.BuySide})
	book.BuyOrders = append(book.BuyOrders, engine.Order{Price: 1200, Id: 12, Amount: 30, Direction: engine.BuySide})
	book.BuyOrders = append(book.BuyOrders, engine.Order{Price: 1300, Id: 13, Amount: 40, Direction: engine.BuySide})
	book.BuyOrders = append(book.BuyOrders, engine.Order{Price: 1400, Id: 14, Amount: 50, Direction: engine.BuySide})
}

func TestSingleBuySideOrder(t *testing.T) {
	makerBook := &engine.MakerBook{
		BuyOrders:  make([]engine.Order, 0),
		SellOrders: make([]engine.Order, 0),
	}
	giveOrderBook(makerBook)

	actual := makerBook.TakerProcess(engine.Order{Price: 102, Id: 101, Amount: 1, Direction: engine.BuySide})[0]

	expected := engine.Trade{TakerId: 101, MakerId: 1, Price: 100, Amount: 1}

	assert.JSONEq(t, string(expected.Serialize()), string(actual.Serialize()))
}

func TestSingleSellSideOrder(t *testing.T) {
	makerBook := &engine.MakerBook{
		BuyOrders:  make([]engine.Order, 0),
		SellOrders: make([]engine.Order, 0),
	}
	giveOrderBook(makerBook)

	actual := makerBook.TakerProcess(engine.Order{Price: 1000, Id: 101, Amount: 10, Direction: engine.SellSide})[0]

	expected := engine.Trade{TakerId: 101, MakerId: 14, Price: 1400, Amount: 10}

	assert.JSONEq(t, string(expected.Serialize()), string(actual.Serialize()))
}

func TestDoubleBuySideOrderWithNoLeft(t *testing.T) {
	makerBook := &engine.MakerBook{
		BuyOrders:  make([]engine.Order, 0),
		SellOrders: make([]engine.Order, 0),
	}
	giveOrderBook(makerBook)

	actual := makerBook.TakerProcess(engine.Order{Price: 104, Id: 101, Amount: 3, Direction: engine.BuySide})

	expected := []engine.Trade{
		{TakerId: 101, MakerId: 1, Price: 100, Amount: 1},
		{TakerId: 101, MakerId: 2, Price: 104, Amount: 2},
	}

	assert.JSONEq(t, string(expected[0].Serialize()), string(actual[0].Serialize()))
	assert.JSONEq(t, string(expected[1].Serialize()), string(actual[1].Serialize()))
}

func TestDoubleSellSideOrderWithNoLeft(t *testing.T) {
	makerBook := &engine.MakerBook{
		BuyOrders:  make([]engine.Order, 0),
		SellOrders: make([]engine.Order, 0),
	}
	giveOrderBook(makerBook)

	actual := makerBook.TakerProcess(engine.Order{Price: 1000, Id: 101, Amount: 90, Direction: engine.SellSide})

	expected := []engine.Trade{
		{TakerId: 101, MakerId: 14, Price: 1400, Amount: 50},
		{TakerId: 101, MakerId: 13, Price: 1300, Amount: 40},
	}

	assert.JSONEq(t, string(expected[0].Serialize()), string(actual[0].Serialize()))
	assert.JSONEq(t, string(expected[1].Serialize()), string(actual[1].Serialize()))
}

func TestDoubleBuySideOrderWithLeft(t *testing.T) {
	makerBook := &engine.MakerBook{
		BuyOrders:  make([]engine.Order, 0),
		SellOrders: make([]engine.Order, 0),
	}
	giveOrderBook(makerBook)

	actual := makerBook.TakerProcess(engine.Order{Price: 104, Id: 101, Amount: 2, Direction: engine.BuySide})

	expected := []engine.Trade{
		{TakerId: 101, MakerId: 1, Price: 100, Amount: 1},
		{TakerId: 101, MakerId: 2, Price: 104, Amount: 1},
	}

	assert.JSONEq(t, string(expected[0].Serialize()), string(actual[0].Serialize()))
	assert.JSONEq(t, string(expected[1].Serialize()), string(actual[1].Serialize()))
}

func TestDoubleSellSideOrderWithLeft(t *testing.T) {
	makerBook := &engine.MakerBook{
		BuyOrders:  make([]engine.Order, 0),
		SellOrders: make([]engine.Order, 0),
	}
	giveOrderBook(makerBook)

	actual := makerBook.TakerProcess(engine.Order{Price: 1000, Id: 101, Amount: 60, Direction: engine.SellSide})

	expected := []engine.Trade{
		{TakerId: 101, MakerId: 14, Price: 1400, Amount: 50},
		{TakerId: 101, MakerId: 13, Price: 1300, Amount: 10},
	}

	assert.JSONEq(t, string(expected[0].Serialize()), string(actual[0].Serialize()))
	assert.JSONEq(t, string(expected[1].Serialize()), string(actual[1].Serialize()))
}

func TestBuySideOrderWithInSufficientOrder(t *testing.T) {
	makerBook := &engine.MakerBook{
		BuyOrders:  make([]engine.Order, 0),
		SellOrders: make([]engine.Order, 0),
	}
	giveOrderBook(makerBook)

	actual := makerBook.TakerProcess(engine.Order{Price: 104, Id: 101, Amount: 4, Direction: engine.BuySide})
	remainOrder := makerBook.BuyOrders[0].Serialize()

	expected := []engine.Trade{
		{TakerId: 101, MakerId: 1, Price: 100, Amount: 1},
		{TakerId: 101, MakerId: 2, Price: 104, Amount: 2},
	}
	expectedRemainOrder := engine.Order{Price: 104, Id: 101, Amount: 1, Direction: engine.BuySide}

	assert.JSONEq(t, string(expected[0].Serialize()), string(actual[0].Serialize()))
	assert.JSONEq(t, string(expected[1].Serialize()), string(actual[1].Serialize()))
	assert.JSONEq(t, string(expectedRemainOrder.Serialize()), string(remainOrder))
}

func TestSellSideOrderWithInSufficientOrder(t *testing.T) {
	makerBook := &engine.MakerBook{
		BuyOrders:  make([]engine.Order, 0),
		SellOrders: make([]engine.Order, 0),
	}
	giveOrderBook(makerBook)

	actual := makerBook.TakerProcess(engine.Order{Price: 1400, Id: 101, Amount: 60, Direction: engine.SellSide})
	remainOrder := makerBook.SellOrders[0].Serialize()

	expected := []engine.Trade{
		{TakerId: 101, MakerId: 14, Price: 1400, Amount: 50},
	}
	expectedRemainOrder := engine.Order{Price: 1400, Id: 101, Amount: 10, Direction: engine.SellSide}

	assert.JSONEq(t, string(expected[0].Serialize()), string(actual[0].Serialize()))
	assert.JSONEq(t, string(expectedRemainOrder.Serialize()), string(remainOrder))
}
