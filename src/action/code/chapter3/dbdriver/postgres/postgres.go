package postgres

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
)

type PostgresDriver struct {
}

func (d PostgresDriver) Open(string) (driver.Conn, error) {
	fmt.Println("Open a postgres driver")
	return nil, errors.New("UnImplemented")
}

var d *PostgresDriver

func init() {
	d = new(PostgresDriver)
	sql.Register("postgres", d)
}
