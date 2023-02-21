package stoporders

import "fmt"

/**
 * Executes a given stop order.
 *
 * @param order The stop order to execute
 */
func Execute(order *StopOrder) {
	fmt.Println("Executed Order")
	fmt.Println(order)
	return
}
