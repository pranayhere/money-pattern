package money

import (
	"errors"
	"math"
)

// Amount is a datastructure that stores the amount being used for calculations.
type Amount struct {
	val int64
}

// Money represents monetary value information, stores
// currency and amount value
type Money struct {
	amount   *Amount
	currency *Currency
}

// New creates and returns new instance of Money.
func NewMoney(amount int64, code string) *Money {
	return &Money{
		amount: &Amount{val: amount},
		currency: GetCurrency(code),
	}
}

// Amount returns a copy of the internal monetary value as an int64
func (m *Money) Amount() int64 {
	return m.amount.val
}

// Currency returns the currency used by Money
func (m *Money) Currency() *Currency {
	return m.currency
}

// SameCurrency check if given Money is equals by currency.
func (m *Money) SameCurrency(om *Money) bool {
	return m.currency.equals(om.currency)
}

func (m *Money) compare(om *Money) int {
	switch {
	case m.amount.val > om.amount.val:
		return 1
	case m.amount.val < om.amount.val:
		return -1
	}

	return 0
}

func (m *Money) assertSameCurrency(om *Money) error {
	if !m.SameCurrency(om) {
		return errors.New("currencies don't match")
	}

	return nil
}

func (m *Money) Equals(om *Money) (bool, error) {
	if err := m.assertSameCurrency(om); err != nil {
		return false, err
	}

	return m.compare(om) == 0, nil
}

// GreaterThan checks whether the value of Money is greater than the other.
func (m *Money) GreaterThan(om *Money) (bool, error) {
	if err := m.assertSameCurrency(om); err != nil {
		return false, err
	}

	return m.compare(om) == 1, nil
}

// GreaterThanOrEqual checks whether the value of Money is greater or equal than the other.
func (m *Money) GreaterThanOrEqual(om *Money) (bool, error) {
	if err := m.assertSameCurrency(om); err != nil {
		return false, err
	}

	return m.compare(om) >= 0, nil
}

// LessThan checks whether the value of Money is less than the other.
func (m *Money) LessThan(om *Money) (bool, error) {
	if err := m.assertSameCurrency(om); err != nil {
		return false, err
	}

	return m.compare(om) == -1, nil
}

// LessThanOrEqual checks whether the value of Money is less or equal than the other.
func (m *Money) LessThanOrEqual(om *Money) (bool, error) {
	if err := m.assertSameCurrency(om); err != nil {
		return false, err
	}

	return m.compare(om) <= 0, nil
}

// IsZero returns boolean of whether the value of Money is equals to zero.
func (m *Money) IsZero() bool {
	return m.amount.val == 0
}

func (m *Money) IsPositive() bool {
	return m.amount.val > 0
}

// IsNegative returns boolean of whether the value of Money is negative.
func (m *Money) IsNegative() bool {
	return m.amount.val < 0
}

func (m *Money) Add(om *Money) (*Money, error) {
	if err := m.assertSameCurrency(om); err != nil {
		return nil, err
	}

	return &Money{amount: &Amount{m.amount.val + om.amount.val}, currency: m.currency}, nil
}

func (m *Money) Subtract(om *Money) (*Money, error) {
	if err := m.assertSameCurrency(om); err != nil {
		return nil, err
	}

	return &Money{amount: &Amount{m.amount.val - om.amount.val}, currency: m.currency}, nil
}

// Multiply returns new Money struct with value representing Self multiplied value by multiplier.
func (m *Money) Multiply(mul int64) *Money {
	return &Money{amount: &Amount{m.amount.val * mul}, currency: m.currency}
}

// AsMajorUnits lets represent Money struct as subunits (float64) in given Currency value
func (m *Money) AsMajorUnits() float64 {
	c := m.currency.get()

	if m.currency.Fraction == 0 {
		return float64(m.amount.val)
	}

	return float64(m.amount.val) / float64(math.Pow10(c.Fraction))
}
