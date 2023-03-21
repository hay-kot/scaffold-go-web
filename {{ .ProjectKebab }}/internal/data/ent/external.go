package ent

import (
	"database/sql"

	entsql "entgo.io/ent/dialect/sql"
)

// SQL exposes the underlying database connection in the ent client
// so that we can use it to perform custom queries.
func (c *Client) SQL() *sql.DB {
	return c.driver.(*entsql.Driver).DB()
}