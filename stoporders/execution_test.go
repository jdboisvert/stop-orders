package stoporders

import "testing"

func TestExecute(t *testing.T) {
	order := &StopOrder{
		Side:         "BUY",
		Quantity:     "10",
		TriggerPrice: 100.0,
		Symbol:       "AAPL",
		OrderId:      "123",
	}

	err := Execute(order)

	if err != nil {
		t.Errorf("Expected Execute to return nil error, but got %s", err)
	}
}
