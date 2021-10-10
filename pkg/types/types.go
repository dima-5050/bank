package types

type Money int64

type Currency string

const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

type PAN string

type Card struct {
	ID       int
	PAN      PAN
	Balance  Money
	Currency Currency
	Color    string
	Name     string
	Active   bool
}
type Category string

type Status string

const (
	StatusOk   Status = "OK"
	StatusFail Status = "FAIL"
	StatusProgrss Status="INPROGRESS"
)

type Payment struct {
	ID       int
	Amount   Money
	Category Category
	Status   Status
}
