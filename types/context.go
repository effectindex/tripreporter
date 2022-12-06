package types

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Context struct {
	Database *pgxpool.Pool
	Logger   *zap.SugaredLogger
}

func (ctx *Context) DB() pgx.Tx {
	ctx.Validate()

	db, err := ctx.Database.Begin(context.Background())
	if err != nil {
		ctx.Logger.Errorw("Failed to begin db transaction", zap.Error(err))
		return nil
	}

	return db
}

func (ctx *Context) Validate() {
	if ctx.Logger == nil {
		panic(ErrorContextNilLogger)
	}

	if ctx.Database == nil {
		panic(ErrorContextNilDatabase)
	}
}
