package stoporders

import (
	"testing"
)

func TestShouldTriggerBuy_True(t *testing.T) {
	order := StopOrder{Side: "BUY", Quantity: "10", TriggerPrice: 100.0, Symbol: "AAPL", OrderId: "1"}
	price := 95.0

	if ShouldTriggerBuy(price, &order) {
		t.Error("Expected ShouldTriggerBuy to be false, but was true")
	}
}

func TestShouldTriggerBuy_False(t *testing.T) {
	order := StopOrder{Side: "BUY", Quantity: "10", TriggerPrice: 100.0, Symbol: "AAPL", OrderId: "1"}
	price := 105.0

	if !ShouldTriggerBuy(price, &order) {
		t.Error("Expected ShouldTriggerBuy to be true, but was false")
	}
}

func TestShouldTriggerBuy_FalseWithSellOrder(t *testing.T) {
	order := StopOrder{Side: "SELL", Quantity: "10", TriggerPrice: 100.0, Symbol: "AAPL", OrderId: "1"}
	price := 105.0

	if ShouldTriggerBuy(price, &order) {
		t.Error("Expected ShouldTriggerBuy to be false, but was true")
	}
}

func TestShouldTriggerSell_True(t *testing.T) {
	order := StopOrder{Side: "SELL", Quantity: "10", TriggerPrice: 100.0, Symbol: "AAPL", OrderId: "1"}
	price := 105.0

	if ShouldTriggerSell(price, &order) {
		t.Error("Expected ShouldTriggerSell to be false, but was true")
	}
}

func TestShouldTriggerSell_False(t *testing.T) {
	order := StopOrder{Side: "SELL", Quantity: "10", TriggerPrice: 100.0, Symbol: "AAPL", OrderId: "1"}
	price := 95.0

	if !ShouldTriggerSell(price, &order) {
		t.Error("Expected ShouldTriggerSell to be true, but was false")
	}
}

func TestShouldTriggerSell_FalseWithSellOrder(t *testing.T) {
	order := StopOrder{Side: "BUY", Quantity: "10", TriggerPrice: 100.0, Symbol: "AAPL", OrderId: "1"}
	price := 105.0

	if ShouldTriggerSell(price, &order) {
		t.Error("Expected ShouldTriggerSell to be false, but was true")
	}
}
