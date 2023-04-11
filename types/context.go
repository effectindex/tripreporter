// SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

package types

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Context struct {
	Logger   *zap.SugaredLogger `json:"-" database:"-"`
	Database *pgxpool.Pool      `json:"-" database:"-"`
	Cache    *redis.Client      `json:"-" database:"-"`
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

	if ctx.Cache == nil {
		panic(ErrorContextNilCache)
	}
}
