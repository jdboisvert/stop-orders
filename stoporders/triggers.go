package stoporders

import "fmt"

type CheckTriggerOrderFunction func(float64, *StopOrder) (bool, error)

/**
 * Determines if a stop order should be executed given the current price.
 * This is true when the price is below or equal to the trigger price for a buy order.
 * This will always return false if the order is not a buy order.
 *
 * @param price The current price to be checked against
 * @param order The stop order to check if it should be executed
 * @return Whether or not the order should be executed and an error if applicable
 */
func ShouldTriggerBuy(price float64, order *StopOrder) (bool, error) {
	if order.Side != "BUY" {
		// This function is only for buy orders
		return false, fmt.Errorf("Expected order side to be BUY, but got %s", order.Side)
	}

	return order.TriggerPrice <= price, nil
}

/**
 * Determines if a stop order should be executed given the current price.
 * This is true when the price is above or equal to the trigger price for a sell order.
 * This will always return false if the order is not a sell order.
 *
 * @param price The current price to be checked against
 * @param order The stop order to check if it should be executed
 * @return Whether or not the order should be executed and an error if applicable.
 */
func ShouldTriggerSell(price float64, order *StopOrder) (bool, error) {
	if order.Side != "SELL" {
		// This function is only for sell orders
		return false, fmt.Errorf("Expected order side to be SELL, but got %s", order.Side)
	}

	return order.TriggerPrice >= price, nil
}
