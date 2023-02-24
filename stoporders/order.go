package stoporders

type StopOrder struct {
	Side         string
	Quantity     string
	TriggerPrice float64
	Symbol       string
	OrderId      string
}
