package stoporders

import (
	"testing"
)

func TestShouldTriggerBuy_False(t *testing.T) {
	order := StopOrder{Side: "BUY", Quantity: "10", TriggerPrice: 100.0, Symbol: "AAPL", OrderId: "1"}
	price := 95.0

	shouldTrigger, err := ShouldTriggerBuy(price, &order)
	if shouldTrigger {
		t.Error("Expected ShouldTriggerBuy to be false, but was true")
	}

	if err != nil {
		t.Errorf("Expected ShouldTriggerBuy to return nil error, but got %s", err)
	}
}

func TestShouldTriggerBuy_True(t *testing.T) {
	order := StopOrder{Side: "BUY", Quantity: "10", TriggerPrice: 100.0, Symbol: "AAPL", OrderId: "1"}
	price := 105.0

	shouldTrigger, err := ShouldTriggerBuy(price, &order)
	if !shouldTrigger {
		t.Error("Expected ShouldTriggerBuy to be true, but was false")
	}

	if err != nil {
		t.Errorf("Expected ShouldTriggerBuy to return nil error, but got %s", err)
	}
}

func TestShouldTriggerBuy_FalseWithSellOrder(t *testing.T) {
	order := StopOrder{Side: "SELL", Quantity: "10", TriggerPrice: 100.0, Symbol: "AAPL", OrderId: "1"}
	price := 105.0

	shouldTrigger, err := ShouldTriggerBuy(price, &order)
	if shouldTrigger {
		t.Error("Expected ShouldTriggerBuy to be false, but was true")
	}

	if err == nil {
		t.Error("Expected ShouldTriggerBuy to return error, but got nil")
	}
}

func TestShouldTriggerSell_False(t *testing.T) {
	order := StopOrder{Side: "SELL", Quantity: "10", TriggerPrice: 100.0, Symbol: "AAPL", OrderId: "1"}
	price := 105.0

	shouldTrigger, err := ShouldTriggerSell(price, &order)
	if shouldTrigger {
		t.Error("Expected ShouldTriggerSell to be false, but was true")
	}

	if err != nil {
		t.Errorf("Expected ShouldTriggerSell to return nil error, but got %s", err)
	}
}

func TestShouldTriggerSell_True(t *testing.T) {
	order := StopOrder{Side: "SELL", Quantity: "10", TriggerPrice: 100.0, Symbol: "AAPL", OrderId: "1"}
	price := 95.0

	shouldTrigger, err := ShouldTriggerSell(price, &order)
	if !shouldTrigger {
		t.Error("Expected ShouldTriggerSell to be true, but was false")
	}

	if err != nil {
		t.Errorf("Expected ShouldTriggerSell to return nil error, but got %s", err)
	}
}

func TestShouldTriggerSell_FalseWithSellOrder(t *testing.T) {
	order := StopOrder{Side: "BUY", Quantity: "10", TriggerPrice: 100.0, Symbol: "AAPL", OrderId: "1"}
	price := 105.0

	shouldTrigger, err := ShouldTriggerSell(price, &order)
	if shouldTrigger {
		t.Error("Expected ShouldTriggerSell to be false, but was true")
	}

	if err == nil {
		t.Error("Expected ShouldTriggerSell to return error, but got nil")
	}
}
