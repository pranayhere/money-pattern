package money_test

import (
	"fmt"
	money "money-pattern"
	"testing"
)

func TestMoney_AsMajorUnits(t *testing.T) {
	m := money.NewMoney(159, "INR")
	fmt.Println(m.AsMajorUnits())
}
