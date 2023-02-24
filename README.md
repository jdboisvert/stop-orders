# Stop Market Orders v1.0.0

An implementation showcasing how one can manage orders in memory to execute [Stop Market Orders](https://www.investopedia.com/terms/s/stoporder.asp). This proof of concept illustrates how one can manage orders in a sorted list and find the orders to execute on every tick giving a price. This is a proof of concept and is not intended for production use.

## Usage

You can review the unit tests in `stoporders` to see how to use the package but here is a quick way you can test it out yourself in a `main.go``

```go
import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jdboisvert/stop-orders/stoporders"
)

/**
 * Just a quick function used to help you generate some orders in memory to test with :)
 *
 * @param numberOfOrdersToGenerate The number of orders you want to generate for the test.
 * @return A map of stop orders keyed by the symbol and side.
 */
func generateOrders(numberOfOrdersToGenerate int) map[string][]stoporders.StopOrder {
	stopOrders := map[string][]stoporders.StopOrder{}

	// // Negatives
	for i := numberOfOrdersToGenerate; i > 0; i-- {
		buyStopOrder := stoporders.StopOrder{Side: "BUY", Quantity: "80123.00", TriggerPrice: 20000.0 - float64(i), Symbol: "AAPL", OrderId: uuid.New().String()}
		sellStopOrder := stoporders.StopOrder{Side: "SELL", Quantity: "80123.00", TriggerPrice: 20000.0 - float64(i), Symbol: "AAPL", OrderId: uuid.New().String()}

		stopOrders["AAPL:BUY"] = append(stopOrders["AAPL:BUY"], buyStopOrder)
		stopOrders["AAPL:SELL"] = append(stopOrders["AAPL:SELL"], sellStopOrder)
	}

	// Positives
	for i := 0; i < numberOfOrdersToGenerate; i++ {
		buyStopOrder := stoporders.StopOrder{Side: "BUY", Quantity: "80123.00", TriggerPrice: 20000.0 + float64(i), Symbol: "AAPL", OrderId: uuid.New().String()}
		sellStopOrder := stoporders.StopOrder{Side: "SELL", Quantity: "80123.00", TriggerPrice: 20000.0 + float64(i), Symbol: "AAPL", OrderId: uuid.New().String()}

		stopOrders["AAPL:BUY"] = append(stopOrders["AAPL:BUY"], buyStopOrder)
		stopOrders["AAPL:SELL"] = append(stopOrders["AAPL:SELL"], sellStopOrder)
	}

	return stopOrders
}

func main() {
    numberOfOrdersToGenerate := 1000 // Change this to whatever number you wish to test with (this will ensure there is this amount of orders in both directions. ex: 1000 will generate 1000 buy and 1000 sell orders in both directions).

    // Note that the generating of the orders can take a while depending on the number of orders you generate which is why it is not included in the start and end time.
	stopOrders := generateOrders(numberOfOrdersToGenerate)
	start := time.Now()
	stoporders.ExecuteOrders(20000.00, "AAPL", stopOrders)

    // This will print out the time it took to execute the orders in memory.
	fmt.Printf("Took %s", time.Since(start))
}

```

## Development

### Getting Started

    # install golang
    brew install golang

    # install the golangci linter
    # more details: https://golangci-lint.run/
    brew install golangci-lint

    # install pre-commit
    pip install pre-commit
    pre-commit install

    # install dependencies
    go mod download

### Pre-commit

A number of pre-commit hooks are set up to ensure all commits meet basic code quality standards.

If one of the hooks changes a file, you will need to `git add` that file and re-run `git commit` before being able to continue.


### Testing

All test files are named *_test.go. Github workflow automatically run the tests when code is pushed and will return a report with results when finished.

You can also run the tests locally:

    go test ./...

To run the tests with coverage:

    go test -coverprofile=coverage.out ./...
    go tool cover -html=coverage.out
