package engine

func (book *MakerBook) TakerProcess(order Order) []Trade {
	if order.Direction == BuySide {
		return book.takerBuy(order)
	} else {
		return book.takerSell(order)
	}
}
