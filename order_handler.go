package main

import (
	"fmt"
)

type StopOrder struct {
	Side         string
	Quantity     string
	TriggerPrice float64
	Symbol       string
	OrderId      string
}

type trigger func(float64, *StopOrder) bool

func ExecuteOrders(price float64, symbol string, stopOrders map[string][]StopOrder) {
	buyKey := fmt.Sprintf("%v:BUY", symbol)
	sellKey := fmt.Sprintf("%v:SELL", symbol)

	checkBuyOrders(price, stopOrders[buyKey])
	checkSellOrders(price, stopOrders[sellKey])
}

func placeOrder(order *StopOrder) {
	fmt.Println("Executed Order")
	fmt.Println(order)
	return
}

func shouldTriggerBuy(price float64, order *StopOrder) bool {
	return order.TriggerPrice <= price
}

func shouldTriggerSell(price float64, order *StopOrder) bool {
	return order.TriggerPrice >= price
}

func searchForOrdersToExecute(price float64, orders []StopOrder, shouldTrigger trigger, isBuy bool) []string {
	lowIndex := 0
	highIndex := len(orders) - 1

	var orderIdsExecuted []string

	for lowIndex <= highIndex {
		midIndex := (highIndex + lowIndex) / 2
		stopOrder := orders[midIndex]

		if shouldTrigger(price, &stopOrder) {
			var ordersToExecute []StopOrder

			if isBuy {
				ordersToExecute = orders[:midIndex+1]

				// Reassign the orders we should now check
				orders = orders[midIndex+1:]

			} else {
				ordersToExecute = orders[midIndex:]

				// Reassign the orders we should now check
				orders = orders[:midIndex]
			}

			// Just for demo would not actually perform this in the loop to move to goroutine
			for _, order := range ordersToExecute {
				placeOrder(&order)
				orderIdsExecuted = append(orderIdsExecuted, order.OrderId)
			}

			// New slice should be evaluated as a whole.
			lowIndex = 0
			highIndex = len(orders) - 1

			continue
		}

		// Could not find any orders to trigger have to restrict search
		if isBuy {
			highIndex = midIndex - 1
		} else {
			lowIndex = midIndex + 1
		}
	}

	return orderIdsExecuted
}

func checkBuyOrders(price float64, orders []StopOrder) []string {
	orderIdsExecuted := searchForOrdersToExecute(price, orders, shouldTriggerBuy, true)

	return orderIdsExecuted
}

func checkSellOrders(price float64, orders []StopOrder) []string {
	orderIdsExecuted := searchForOrdersToExecute(price, orders, shouldTriggerSell, false)

	return orderIdsExecuted
}
