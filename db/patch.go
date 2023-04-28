package db

import (
	"context"

	"github.com/effectindex/tripreporter/models"
	"github.com/effectindex/tripreporter/types"
	"github.com/georgysavva/scany/v2/pgxscan"
	"go.uber.org/zap"
)

var (
	patches = []PatchFn{func(c types.Context) bool {
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
	}}
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
		if p != nil && p.Patched {
			patched[p.Index] = true
		}
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
