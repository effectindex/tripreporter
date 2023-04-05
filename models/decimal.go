// SPDX-FileCopyrightText: 2023 froggie <incoming@frogg.ie>
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

func (d *Decimal) Default() {
	func(x **Decimal) {
		*x = nil
	}(&d)
}
