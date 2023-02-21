package stoporders

import "fmt"

/**
 * Executes all stop orders for the given symbol.
 *
 * @param price The current price of the symbol
 * @param symbol The symbol to execute stop orders for (ex. BTC_USDT)
 * @param stopOrders The map of all stop orders in memory
 * @return The list of order ids that were executed
 */
func ExecuteOrders(price float64, symbol string, stopOrders map[string][]StopOrder) []string {
	buyKey := fmt.Sprintf("%v:BUY", symbol)
	sellKey := fmt.Sprintf("%v:SELL", symbol)

	executedBuyOrderIds := checkBuyOrders(price, stopOrders[buyKey])
	executedSellOrderIds := checkSellOrders(price, stopOrders[sellKey])

	executedOrderIds := append(executedBuyOrderIds, executedSellOrderIds...)

	return executedOrderIds
}

/**
 * Checks all stop orders for the given symbol that are buy orders
 * and executes them if the price is above the trigger price.
 *
 * @param price The current price of the symbol
 * @param orders The list of stop orders for the symbol
 * @return The list of order ids that were executed
 */
func checkBuyOrders(price float64, orders []StopOrder) []string {
	orderIdsExecuted := searchForOrdersToExecute(price, orders, ShouldTriggerBuy, true)
	return orderIdsExecuted
}

/**
 * Checks all stop orders for the given symbol that are sell orders
 * and executes them if the price is below the trigger price.
 *
 * @param price The current price of the symbol
 * @param orders The list of stop orders for the symbol
 * @return The list of order ids that were executed
 */
func checkSellOrders(price float64, orders []StopOrder) []string {
	orderIdsExecuted := searchForOrdersToExecute(price, orders, ShouldTriggerSell, false)
	return orderIdsExecuted
}

/**
 * Searches for all orders that should be executed given the current price. This uses a binary search
 * to find the orders that should be executed. It is important that he orders are sorted by trigger price.
 *
 * @param price The current price of the symbol
 * @param orders The list of stop orders for the symbol
 * @param shouldTrigger The function that determines if an order should be triggered
 * @param isBuy Whether or not the orders are buy orders
 * @return The list of order ids that were executed
 */
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
				// Sell
				ordersToExecute = orders[midIndex:]

				// Reassign the orders we should now check
				orders = orders[:midIndex]
			}

			for _, order := range ordersToExecute {
				Execute(&order)
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
