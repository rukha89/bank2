package types

// Money представляет собой денежную сумму в минималных единиц (центы, копейки, дирамы и т.д).
type Money int64

// Category представляет собой категорию, в которию был совершен платёж (авто, аптека, рестараны и т.д).
type Category string

// Status представляет статус плятежей.
type Status string

// Предопределеные статусы платежа.
const (
  StatusOk       Status = "OK"
  StatusFail     Status = "FAIL"
  StatusProgress Status = "INPROGRESS"
)

// Payment представляет игформацию о платиже.
type Payment struct {
  ID       int
  Amount   Money
  Category Category
  Status   Status
}