package stoporders

import (
	"log"
)

/**
 * Executes a given stop order.
 *
 * @param order The stop order to execute
 */
func Execute(order *StopOrder) error {
	log.Println("Executing order", order)

	// TODO - This does not actually do anything yet
	// but there is where you would place your market order.

	return nil
}
