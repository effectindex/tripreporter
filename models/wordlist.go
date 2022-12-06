package models

import (
	"crypto/rand"
	"encoding/json"
	"math/big"
	"os"
	"strings"

	"github.com/effectindex/tripreporter/types"
	"go.uber.org/zap"
)

var (
	Wordlist    WordlistModel
	WordlistLen *big.Int
)

type WordlistModel struct { // This name sucks but I don't see a better option
	types.Context
	Words []string `json:"words,omitempty"`
}

// SetupWordlist creates a new db.Wordlist from the $WORDLIST file
func SetupWordlist(ctx types.Context) {
	ctx.Validate()

	if bytes, err := os.ReadFile(os.Getenv("WORDLIST")); err != nil {
		ctx.Logger.Panicw("Failed to read WORDLIST file", zap.Error(err))
	} else {
		var words []string

		if err := json.Unmarshal(bytes, &words); err != nil {
			ctx.Logger.Panicw("Failed to unmarshal wordlist", zap.Error(err))
		} else {
			Wordlist = WordlistModel{Context: ctx, Words: words}
			WordlistLen = big.NewInt(int64(len(Wordlist.Words) - 1))

			ctx.Logger.Infof("Loaded wordlist with %s keys", WordlistLen)
		}
	}
}

func (w *WordlistModel) Random(amount int) string {
	w.Context.Validate()

	if amount <= 0 {
		return ""
	}

	words := make([]string, amount)

	for i := 0; i < amount; i++ {
		if n, err := rand.Int(rand.Reader, WordlistLen); err != nil {
			w.Logger.Panicw("Failed to make rand.Int", zap.Error(err))
			return ""
		} else {
			words[i] = Wordlist.Words[n.Int64()]
		}
	}

	return strings.Join(words, "-")
}
