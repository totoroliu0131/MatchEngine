package engine

type MakerBook struct {
	BuyOrders  []Order
	SellOrders []Order
}

func (book *MakerBook) addBuyOrder(order Order) {
	n := len(book.BuyOrders)
	var i int
	for i := n - 1; i >= 0; i-- {
		buyOrder := book.BuyOrders[i]
		if buyOrder.Price < order.Price {
			break
		}
	}
	if i == n-1 {
		book.BuyOrders = append(book.BuyOrders, order)
	} else {
		copy(book.BuyOrders[i+1:], book.BuyOrders[i:])
		book.BuyOrders[i] = order
	}
}

func (book *MakerBook) addSellOrder(order Order) {
	n := len(book.SellOrders)
	var i int
	for i := n - 1; i >= 0; i-- {
		sellOrder := book.SellOrders[i]
		if sellOrder.Price > order.Price {
			break
		}
	}
	if i == n-1 {
		book.SellOrders = append(book.SellOrders, order)
	} else {
		copy(book.SellOrders[i+1:], book.SellOrders[i:])
		book.SellOrders[i] = order
	}
}

func (book *MakerBook) removeBuyOrder(index int) {
	book.BuyOrders = append(book.BuyOrders[:index], book.BuyOrders[index+1:]...)
}

func (book *MakerBook) removeSellOrder(index int) {
	book.SellOrders = append(book.SellOrders[:index], book.SellOrders[index+1:]...)
}

func (book *MakerBook) takerBuy(order Order) []Trade {
	trades := make([]Trade, 0)
	n := len(book.SellOrders)

	if n != 0 || book.SellOrders[n-1].Price <= order.Price {
		for i := n - 1; i >= 0; i-- {
			sellOrder := book.SellOrders[i]
			if sellOrder.Price > order.Price {
				break
			}
			if sellOrder.Amount >= order.Amount {
				trades = append(trades, Trade{order.Id, sellOrder.Id, sellOrder.Price, order.Amount})
				sellOrder.Amount -= order.Amount
				if sellOrder.Amount == 0 {
					book.removeSellOrder(i)
				}
				return trades
			} else {
				trades = append(trades, Trade{order.Id, sellOrder.Id, sellOrder.Price, sellOrder.Amount})
				order.Amount -= sellOrder.Amount
				book.removeSellOrder(i)
				continue
			}
		}
	}
	book.addBuyOrder(order)

	return trades

}

func (book *MakerBook) takerSell(order Order) []Trade {
	trades := make([]Trade, 0)
	n := len(book.SellOrders)

	if n != 0 || book.BuyOrders[n-1].Price >= order.Price {
		for i := n - 1; i >= 0; i-- {
			buyOrder := book.BuyOrders[i]
			if buyOrder.Price < order.Price {
				break
			}
			if buyOrder.Amount >= order.Amount {
				trades = append(trades, Trade{order.Id, buyOrder.Id, buyOrder.Price, order.Amount})
				buyOrder.Amount -= order.Amount
				if buyOrder.Amount == 0 {
					book.removeBuyOrder(i)
				}
				return trades
			} else {
				trades = append(trades, Trade{order.Id, buyOrder.Id, buyOrder.Price, buyOrder.Amount})
				order.Amount -= buyOrder.Amount
				book.removeBuyOrder(i)
				continue
			}
		}
	}

	book.addSellOrder(order)
	return trades
}
