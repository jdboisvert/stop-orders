package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jdboisvert/stop-orders/stoporders"
)

func generateOrders(numberOfOrdersToGenerate int) map[string][]stoporders.StopOrder {
	stopOrders := map[string][]stoporders.StopOrder{}

	// // Negatives
	for i := numberOfOrdersToGenerate; i > 0; i-- {
		buyStopOrder := stoporders.StopOrder{Side: "BUY", Quantity: "80123.00", TriggerPrice: 20000.0 - float64(i), Symbol: "BTC_USD", OrderId: uuid.New().String()}
		sellStopOrder := stoporders.StopOrder{Side: "SELL", Quantity: "80123.00", TriggerPrice: 20000.0 - float64(i), Symbol: "BTC_USD", OrderId: uuid.New().String()}

		stopOrders["BTC_USD:BUY"] = append(stopOrders["BTC_USD:BUY"], buyStopOrder)
		stopOrders["BTC_USD:SELL"] = append(stopOrders["BTC_USD:SELL"], sellStopOrder)
	}

	// Positives
	for i := 0; i < numberOfOrdersToGenerate; i++ {
		buyStopOrder := stoporders.StopOrder{Side: "BUY", Quantity: "80123.00", TriggerPrice: 20000.0 + float64(i), Symbol: "BTC_USD", OrderId: uuid.New().String()}
		sellStopOrder := stoporders.StopOrder{Side: "SELL", Quantity: "80123.00", TriggerPrice: 20000.0 + float64(i), Symbol: "BTC_USD", OrderId: uuid.New().String()}

		stopOrders["BTC_USD:BUY"] = append(stopOrders["BTC_USD:BUY"], buyStopOrder)
		stopOrders["BTC_USD:SELL"] = append(stopOrders["BTC_USD:SELL"], sellStopOrder)
	}

	return stopOrders
}

func main() {
	stopOrders := generateOrders(1000)
	start := time.Now()
	stoporders.ExecuteOrders(20000.00, "BTC_USD", stopOrders)
	fmt.Printf("Took %s", time.Since(start))
}
