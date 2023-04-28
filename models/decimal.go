// SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

package models

import (
	"strconv"

	"github.com/shopspring/decimal"
)

type Decimal struct {
	decimal.Decimal
}

func (d *Decimal) Valid() bool {
	return d != nil && !d.IsZero()
}

func (d *Decimal) Set(whole, point int) (*Decimal, error) {
	if d1, err := decimal.NewFromString(strconv.Itoa(whole) + "." + strconv.Itoa(point)); err != nil {
		return nil, err
	} else {
		*d = Decimal{d1}
		return d, nil
	}
}

func (d *Decimal) Get() *Decimal {
	return d
}

func (d *Decimal) Add(value *Decimal) *Decimal {
	*d = Decimal{d.Decimal.Add(value.Decimal)}
	return d
}

func (d *Decimal) Sub(value *Decimal) *Decimal {
	*d = Decimal{d.Decimal.Sub(value.Decimal)}
	return d
}

func (d *Decimal) Mul(value float64) *Decimal {
	*d = Decimal{d.Decimal.Mul(decimal.NewFromFloat(value))}
	return d
}

func (d *Decimal) Div(value float64) *Decimal {
	*d = Decimal{d.Decimal.Div(decimal.NewFromFloat(value))}
	return d
}

func (d *Decimal) Zero() *Decimal {
	*d = Decimal{decimal.New(0, 0)}
	return d
}

func (d *Decimal) Default() {
	func(x **Decimal) {
		*x = nil
	}(&d)
}

func (d *Decimal) UnmarshalJSON(b []byte) error {
	if string(b) == `""` {
		return nil
	}

	return d.Decimal.UnmarshalJSON(b)
}
