package money

import "github.com/shopspring/decimal"

func init() {
}

type Money struct {
	decimal.Decimal
}

func (m *Money) Div(o *Money) *Money {
	return &Money{Decimal: m.Decimal.Div(o.Decimal)}
}

func (m *Money) Marshal() ([]byte, error) {
	return []byte(m.Decimal.String()), nil
}

func (m *Money) Unmarshal(data []byte) error {
	if len(data) == 0 {
		m = nil
		return nil
	}

	m.Decimal, _ = decimal.NewFromString(string(data))
	return nil
}

func NewFromString(value string) *Money {
	v, _ := decimal.NewFromString(value)
	return &Money{Decimal: v}
}

func NewFromFloat(value float64) *Money {
	return &Money{Decimal: decimal.NewFromFloat(value)}
}
