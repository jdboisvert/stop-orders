package stoporders

type trigger func(float64, *StopOrder) bool

/**
 * Determines if a stop order should be executed given the current price.
 * This is true when the price is below or equal to the trigger price for a buy order.
 *
 * @param price The current price to be checked against
 * @param order The stop order to check if it should be executed
 * @return Whether or not the order should be executed
 */
func ShouldTriggerBuy(price float64, order *StopOrder) bool {
	return order.TriggerPrice <= price
}

/**
 * Determines if a stop order should be executed given the current price.
 * This is true when the price is above or equal to the trigger price for a sell order.
 *
 * @param price The current price to be checked against
 * @param order The stop order to check if it should be executed
 * @return Whether or not the order should be executed
 */
func ShouldTriggerSell(price float64, order *StopOrder) bool {
	return order.TriggerPrice >= price
}