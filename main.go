package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func generateOrders(numberOfOrdersToGenerate int) map[string][]StopOrder {
	stopOrders := map[string][]StopOrder{}

	// // Negatives
	for i := numberOfOrdersToGenerate; i > 0; i-- {
		buyStopOrder := StopOrder{"BUY", "80123.00", 20000.0 - float64(i), "BTC_USD", uuid.New().String()}
		sellStopOrder := StopOrder{"SELL", "80123.00", 20000.0 - float64(i), "BTC_USD", uuid.New().String()}

		stopOrders["BTC_USD:BUY"] = append(stopOrders["BTC_USD:BUY"], buyStopOrder)
		stopOrders["BTC_USD:SELL"] = append(stopOrders["BTC_USD:SELL"], sellStopOrder)
	}

	// Positives
	for i := 0; i < numberOfOrdersToGenerate; i++ {
		buyStopOrder := StopOrder{"BUY", "80123.00", 20000.0 + float64(i), "BTC_USD", uuid.New().String()}
		sellStopOrder := StopOrder{"SELL", "80123.00", 20000.0 + float64(i), "BTC_USD", uuid.New().String()}

		stopOrders["BTC_USD:BUY"] = append(stopOrders["BTC_USD:BUY"], buyStopOrder)
		stopOrders["BTC_USD:SELL"] = append(stopOrders["BTC_USD:SELL"], sellStopOrder)
	}

	return stopOrders
}

func main() {
	stopOrders := generateOrders(125000)
	start := time.Now()
	ExecuteOrders(20000.00, "BTC_USD", stopOrders)
	fmt.Printf("Took %s", time.Since(start))

}
