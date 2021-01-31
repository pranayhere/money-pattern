package money

type Currency struct {
	Code     string
	Fraction int
}

var currencies = map[string]*Currency{
	"INR": {Code: "INR", Fraction: 2},
	"USD": {Code: "USD", Fraction: 3},
}

func GetCurrency(code string) *Currency {
	return currencies[code]
}

func (c *Currency) get() *Currency {
	if curr, ok := currencies[c.Code]; ok {
		return curr
	}

	return c.getDefault()
}

func (c *Currency) getDefault() *Currency {
	return &Currency{Code: c.Code, Fraction: 2}
}

func (c *Currency) equals(oc *Currency) bool {
	return c.Code == oc.Code
}