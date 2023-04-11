// SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

package models

import (
	"github.com/effectindex/tripreporter/types"
)

type StringRestriction struct {
	MinLength          int          `json:"min_length"`
	MaxLength          int          `json:"max_length"`
	MinUniqueTotal     int          `json:"min_unique_total"`
	MinUniqueSymbol    int          `json:"min_unique_symbol"`
	MinUniqueNonSymbol int          `json:"min_unique_non_symbol"`
	Message            string       `json:"message"`
	Allowed            allowedChars `json:"allowed_chars"`
}

type allowedChars struct {
	Symbol    map[string]bool `json:"symbol"`
	NonSymbol map[string]bool `json:"non_symbol"`
}

func (r StringRestriction) Validate(s string) error {
	if len(s) == 0 {
		return types.ErrorStringEmpty
	}

	if len(s) < r.MinLength {
		return types.ErrorStringShort.ContextError(len(s), r.MinLength)
	}

	if len(s) > r.MaxLength {
		return types.ErrorStringLong.ContextError(len(s), r.MaxLength)
	}

	uniqueTotal := make(map[string]bool, 0)
	uniqueSymbol := make(map[string]bool, 0)
	uniqueNonSymbol := make(map[string]bool, 0)
	invalidChars := make(map[string]bool, 0)

	// Check for non-allowed chars
	for _, character := range s {
		c := string(character)

		allowedS, isSymbol := r.Allowed.Symbol[c]
		allowedN, isNonSymbol := r.Allowed.NonSymbol[c]

		if !isSymbol && !isNonSymbol {
			// Technically, it would be faster to break here before reading the rest of the string.
			// We want to be able to give context for all the invalid chars, so instead we allow continuing to read.
			invalidChars[c] = true
			continue
		}

		uniqueTotal[c] = true

		if allowedS {
			uniqueSymbol[c] = true
		}

		if allowedN {
			uniqueNonSymbol[c] = true
		}
	}

	// Check for any invalid characters.
	if len(invalidChars) > 0 {
		return types.ErrorStringInvalidChar.ContextError(invalidChars)
	}

	// Check for minimum unique chars.
	if len(uniqueTotal) < r.MinUniqueTotal {
		return types.ErrorStringUniqueChar.ContextError(len(uniqueTotal), r.MinUniqueTotal)
	}

	// Check for minimum unique symbols, for example, a config value of 1 means a string must contain at least one symbol.
	if len(uniqueSymbol) < r.MinUniqueSymbol {
		return types.ErrorStringSymbolChar.ContextError(len(uniqueSymbol), r.MinUniqueSymbol)
	}

	// Check for minimum unique non-symbols, for example, a config value of 2 means a string must contain at least two non-symbol chars in it.
	if len(uniqueNonSymbol) < r.MinUniqueNonSymbol {
		return types.ErrorStringNonSymbolChar.ContextError(len(uniqueNonSymbol), r.MinUniqueNonSymbol)
	}

	return nil
}
