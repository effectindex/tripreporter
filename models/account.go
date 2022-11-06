package models

import (
	"context"
	"crypto/rand"
	"math/big"
	"strings"

	"go.uber.org/zap"
)

var (
	wordlist = []string{"an", "example", "wordlist", "whee", "whoo", "swag"} // TODO: better wordlist
	wlLen    = big.NewInt(int64(len(wordlist) - 1))
)

func (a *Account) Submit(logger *zap.SugaredLogger) error {
	db := a.DB(logger)
	defer db.Commit(context.Background())

	if err := a.InitUUID(logger); err != nil {
		return err
	}

	if _, err := db.Exec(context.Background(),
		`insert into accounts(
			id,
			email,
			username,
			password_hash
		)
		values(
			$1,
			$2,
			$3,
			$4
		);`,
		a.ID,
		a.Email,
		a.Username,
		a.Password, // TODO: Salt in DB?
	); err != nil {
		logger.Warnw("Failed to write account to DB", zap.Error(err))
		db.Rollback(context.Background())
		return err
	}

	return nil
}

func Test(logger *zap.SugaredLogger, ctx Context) error {
	if username, err := genDefaultUsername(logger); err != nil {
		return err
	} else {
		a := Account{
			Context:  ctx,
			Type:     "Account",
			Email:    "user@email.com",
			Username: username,
			Password: "examplePword",
		}

		if err := a.Submit(logger); err != nil {
			logger.Warnw("Failed to make test account", "account", a, zap.Error(err))
			return err
		}

		return nil
	}
}

func genDefaultUsername(logger *zap.SugaredLogger) (string, error) { // TODO: Ensure functional
	words := make([]string, 3)

	for i := 0; i < 3; i++ {
		if n, err := rand.Int(rand.Reader, wlLen); err != nil {
			logger.DPanicw("failed to make rand.Int", zap.Error(err))
			return "", err
		} else {
			words[i] = wordlist[n.Int64()]
		}
	}

	return strings.Join(words, "-"), nil
}
