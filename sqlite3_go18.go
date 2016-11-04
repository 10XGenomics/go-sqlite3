package sqlite3

import (
	"database/sql/driver"
	"errors"

	"golang.org/x/net/context"
)

// Ping implement Pinger.
func (c *SQLiteConn) Ping(ctx context.Context) error {
	if c.db == nil {
		return errors.New("Connection was closed")
	}
	return nil
}

func (c *SQLiteConn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	list := make([]namedValue, len(args))
	for i, nv := range args {
		list[i] = namedValue(nv)
	}
	return c.query(ctx, query, list)
}

func (s *SQLiteStmt) QueryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
	list := make([]namedValue, len(args))
	for i, nv := range args {
		list[i] = namedValue(nv)
	}
	return s.query(ctx, list)
}
