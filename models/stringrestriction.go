package models

import "github.com/effectindex/tripreporter/types"

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
		return types.ErrorStringShort
	}

	if len(s) > r.MaxLength {
		return types.ErrorStringLong
	}

	uniqueTotal := make(map[string]bool, 0)
	uniqueSymbol := make(map[string]bool, 0)
	uniqueNonSymbol := make(map[string]bool, 0)

	// Check for non-allowed chars
	for _, character := range s {
		c := string(character)

		allowedS, isSymbol := r.Allowed.Symbol[c]
		allowedN, isNonSymbol := r.Allowed.NonSymbol[c]

		if !isSymbol && !isNonSymbol {
			return types.ErrorStringInvalidChar
		}

		uniqueTotal[c] = true

		if allowedS {
			uniqueSymbol[c] = true
		}

		if allowedN {
			uniqueNonSymbol[c] = true
		}
	}

	// Check for minimum unique chars
	if len(uniqueTotal) < r.MinUniqueTotal {
		return types.ErrorStringUniqueChar
	}

	// Check for minimum unique symbols, for example, a config value of 1 means a string must contain at least one symbol.
	if len(uniqueSymbol) < r.MinUniqueSymbol {
		return types.ErrorStringSymbolChar
	}

	// Check for minimum unique non-symbols, for example, a config value of 2 means a string must contain at least two non-symbol chars in it.
	if len(uniqueNonSymbol) < r.MinUniqueNonSymbol {
		return types.ErrorStringNonSymbolChar
	}

	return nil
}
