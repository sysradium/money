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

func (m Money) MarshalTo(data []byte) (n int, err error) {
	m.Decimal, _ = decimal.NewFromString(string(data))
	return len(data), nil
}

func (m *Money) Size() int {
	if m == nil {
		return 0
	}

	b := []byte(m.Decimal.String())
	if len(b) == 0 {
		return 0
	}
	return len(b)
}

func NewFromString(value string) *Money {
	v, _ := decimal.NewFromString(value)
	return &Money{Decimal: v}
}

func NewFromFloat(value float64) *Money {
	return &Money{Decimal: decimal.NewFromFloat(value)}
}
