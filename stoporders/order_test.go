package stoporders

import "testing"

func TestStopOrder(t *testing.T) {
	order := StopOrder{
		Side:         "BUY",
		Quantity:     "10",
		TriggerPrice: 100.0,
		Symbol:       "AAPL",
		OrderId:      "123",
	}

	if order.Side != "BUY" {
		t.Errorf("Expected Side to be 'BUY', but got %s", order.Side)
	}

	if order.Quantity != "10" {
		t.Errorf("Expected Quantity to be '10', but got %s", order.Quantity)
	}

	if order.TriggerPrice != 100.0 {
		t.Errorf("Expected TriggerPrice to be 100.0, but got %f", order.TriggerPrice)
	}

	if order.Symbol != "AAPL" {
		t.Errorf("Expected Symbol to be 'AAPL', but got %s", order.Symbol)
	}

	if order.OrderId != "123" {
		t.Errorf("Expected OrderId to be '123', but got %s", order.OrderId)
	}
}
