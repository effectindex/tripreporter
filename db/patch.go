// SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

package db

import (
	"context"

	"github.com/effectindex/tripreporter/models"
	"github.com/effectindex/tripreporter/types"
	"github.com/georgysavva/scany/v2/pgxscan"
	"go.uber.org/zap"
)

var (
	patches = []PatchFn{
		func(c types.Context) bool {
			db := c.DB()
			defer db.Commit(context.Background())

			var u1 []*models.User
			if err := pgxscan.Select(context.Background(), db, &u1,
				`SELECT accounts.id AS id FROM accounts LEFT JOIN users ON accounts.id = users.account_id WHERE users.account_id IS NULL;`); err != nil {
				c.Logger.Warnw("Patch: Failed to get accounts from DB", zap.Error(err))
				return false
			}

			confirm := true
			for _, u := range u1 {
				if u != nil {
					u.Context = c

					if u, err := u.Post(); err != nil {
						confirm = false
						c.Logger.Warnw("Patch: Failed to create user for account", "user", u, zap.Error(err))
					} else {
						c.Logger.Debugw("Patch: Created missing user for account", "user", u)
					}
				}
			}

			c.Logger.Debugw("Patch: Finished patching for missing users")
			return confirm
		},
		func(c types.Context) bool {
			db := c.DB()
			defer db.Commit(context.Background())

			if _, err := db.Exec(context.Background(), `ALTER TABLE accounts ADD COLUMN display_name varchar(255) NOT NULL DEFAULT '';`); err != nil {
				_ = db.Rollback(context.Background())
				c.Logger.Warnw("Patch: Failed to update key constraint for accounts_display_name", zap.Error(err))
				return false
			}

			c.Logger.Debugw("Patch: Finished patching accounts_display_name")
			return true
		},
		func(c types.Context) bool {
			db := c.DB()
			defer db.Commit(context.Background())

			if _, err := db.Exec(context.Background(), `UPDATE accounts SET display_name = users.display_name FROM users WHERE accounts.id = users.account_id;`); err != nil {
				_ = db.Rollback(context.Background())
				c.Logger.Warnw("Patch: Failed to copy users.display_name to accounts.display_name", zap.Error(err))
				return false
			}

			c.Logger.Debugw("Patch: Finished patching copy users.display_name to accounts.display_name")
			return true
		},
		func(c types.Context) bool {
			db := c.DB()
			defer db.Commit(context.Background())

			if _, err := db.Exec(context.Background(), `ALTER TABLE users DROP COLUMN display_name;`); err != nil {
				_ = db.Rollback(context.Background())
				c.Logger.Warnw("Patch: Failed to drop column display_name", zap.Error(err))
				return false
			}

			c.Logger.Debugw("Patch: Finished patching drop column display_name")
			return true
		},
	}
)

// PatchFn executes a database patch once
type PatchFn func(c types.Context) bool

type Patch struct {
	Index   int  `db:"index"`
	Patched bool `db:"patched"`
}

// Confirm will insert a Patch into the database, with patched being true on a successful patch
func (p *PatchFn) Confirm(c types.Context, index int, exists, patched bool) {
	db := c.DB()
	defer db.Commit(context.Background())

	if exists {
		if _, err := db.Exec(context.Background(), `update db_patches set patched=$2 where index=$1;`, index, patched); err != nil {
			_ = db.Rollback(context.Background())
			c.Logger.Warnw("Patch: Failed to confirm", zap.Error(err))
		}
	} else {
		if _, err := db.Exec(context.Background(), `insert into db_patches(index, patched) values($1, $2);`, index, patched); err != nil {
			_ = db.Rollback(context.Background())
			c.Logger.Warnw("Patch: Failed to confirm", zap.Error(err))
		}
	}
}

// SetupPatches will check the database for which patches have been done and were successful, and run any patches that are new / unsuccessful.
func SetupPatches(c types.Context) {
	db := c.DB()

	var p1 []*Patch
	if err := pgxscan.Select(context.Background(), db, &p1,
		`select * from db_patches;`); err != nil {
		c.Logger.Warnw("Patch: Failed to get patches from DB", zap.Error(err))
		_ = db.Commit(context.Background())
		return
	}

	_ = db.Commit(context.Background())

	patched := make(map[int]bool, 0)
	for _, p := range p1 {
		patched[p.Index] = p.Patched
	}

	for n, p := range patches {
		exists := false

		if p2, ok := patched[n]; ok && p2 {
			continue
		} else {
			exists = ok
		}

		p.Confirm(c, n, exists, p(c))
	}
}
