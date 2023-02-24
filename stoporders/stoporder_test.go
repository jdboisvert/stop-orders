package stoporders

import (
	"testing"
)

func TestExecuteOrdersBuy(t *testing.T) {
	orders := map[string][]StopOrder{
		"AAPL:BUY": {
			{Side: "BUY", Quantity: "10", TriggerPrice: 100.0, Symbol: "AAPL", OrderId: "1"},
			{Side: "BUY", Quantity: "5", TriggerPrice: 110.0, Symbol: "AAPL", OrderId: "2"},
			{Side: "BUY", Quantity: "20", TriggerPrice: 120.0, Symbol: "AAPL", OrderId: "3"},
			{Side: "BUY", Quantity: "20", TriggerPrice: 125.0, Symbol: "AAPL", OrderId: "4"},
			{Side: "BUY", Quantity: "20", TriggerPrice: 125.001, Symbol: "AAPL", OrderId: "5"}, // Should not execute
			{Side: "BUY", Quantity: "20", TriggerPrice: 130.0, Symbol: "AAPL", OrderId: "6"},   // Should not execute
		},
	}

	executedIds := ExecuteOrders(125.0, "AAPL", orders) // Everything that equals or is less than 125.0 should execute
	expectedIds := []string{"1", "2", "3", "4"}

	if len(executedIds) != len(expectedIds) {
		t.Errorf("Expected %d orders to be executed, but got %d", len(expectedIds), len(executedIds))
	}

	for i, id := range expectedIds {
		if executedIds[i] != id {
			t.Errorf("Expected order ID %s, but got %s", id, executedIds[i])
		}
	}
}

func TestExecuteOrdersSell(t *testing.T) {
	orders := map[string][]StopOrder{
		"AAPL:SELL": {
			{Side: "SELL", Quantity: "10", TriggerPrice: 100.0, Symbol: "AAPL", OrderId: "5"},  // Should not execute
			{Side: "SELL", Quantity: "10", TriggerPrice: 120.0, Symbol: "AAPL", OrderId: "6"},  // Should not execute
			{Side: "SELL", Quantity: "10", TriggerPrice: 124.01, Symbol: "AAPL", OrderId: "7"}, // Should not execute
			{Side: "SELL", Quantity: "15", TriggerPrice: 125, Symbol: "AAPL", OrderId: "8"},
			{Side: "SELL", Quantity: "15", TriggerPrice: 130.0, Symbol: "AAPL", OrderId: "9"},
			{Side: "SELL", Quantity: "8", TriggerPrice: 140.0, Symbol: "AAPL", OrderId: "10"},
			{Side: "SELL", Quantity: "25", TriggerPrice: 150.0, Symbol: "AAPL", OrderId: "11"},
		},
	}

	executedIds := ExecuteOrders(125.0, "AAPL", orders) // Everything that equals or is greater than 125.0 should execute
	expectedIds := []string{"8", "9", "10", "11"}

	if len(executedIds) != len(expectedIds) {
		t.Errorf("Expected %d orders to be executed, but got %d", len(expectedIds), len(executedIds))
	}

	for i, id := range expectedIds {
		if executedIds[i] != id {
			t.Errorf("Expected order ID %s, but got %s", id, executedIds[i])
		}
	}
}

func TestExecuteOrdersBuyAndSell(t *testing.T) {
	orders := map[string][]StopOrder{
		"AAPL:BUY": {
			{Side: "BUY", Quantity: "10", TriggerPrice: 100.0, Symbol: "AAPL", OrderId: "1"},
			{Side: "BUY", Quantity: "5", TriggerPrice: 110.0, Symbol: "AAPL", OrderId: "2"},
			{Side: "BUY", Quantity: "20", TriggerPrice: 120.0, Symbol: "AAPL", OrderId: "3"},
			{Side: "BUY", Quantity: "20", TriggerPrice: 130.0, Symbol: "AAPL", OrderId: "4"}, // Should not execute
		},
		"AAPL:SELL": {
			{Side: "SELL", Quantity: "10", TriggerPrice: 100.0, Symbol: "AAPL", OrderId: "5"}, // Should not execute
			{Side: "SELL", Quantity: "10", TriggerPrice: 120.0, Symbol: "AAPL", OrderId: "6"}, // Should not execute
			{Side: "SELL", Quantity: "15", TriggerPrice: 130.0, Symbol: "AAPL", OrderId: "7"},
			{Side: "SELL", Quantity: "8", TriggerPrice: 140.0, Symbol: "AAPL", OrderId: "9"},
			{Side: "SELL", Quantity: "25", TriggerPrice: 150.0, Symbol: "AAPL", OrderId: "10"},
		},
	}

	executedIds := ExecuteOrders(125.0, "AAPL", orders)
	expectedIds := []string{"1", "2", "3", "7", "9", "10"}

	if len(executedIds) != len(expectedIds) {
		t.Errorf("Expected %d orders to be executed, but got %d", len(expectedIds), len(executedIds))
	}

	for i, id := range expectedIds {
		if executedIds[i] != id {
			t.Errorf("Expected order ID %s, but got %s", id, executedIds[i])
		}
	}
}

func TestExecuteOrders_NoOrders(t *testing.T) {
	orders := map[string][]StopOrder{}

	executedIds := ExecuteOrders(125.0, "AAPL", orders)
	expectedIds := []string{}

	if len(executedIds) != len(expectedIds) {
		t.Errorf("Expected %d orders to be executed, but got %d", len(expectedIds), len(executedIds))
	}
}
