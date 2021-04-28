package datastore

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/gilcrest/errs"

	_ "github.com/lib/pq" // pq driver calls for blank identifier
)

// Datastorer is an interface for working with the Database
type Datastorer interface {
	// DB returns a sql.DB
	DB() *sql.DB
	// BeginTx starts a sql.Tx using the input context
	BeginTx(context.Context) (*sql.Tx, error)
	// RollbackTx rolls back the input sql.Tx
	RollbackTx(*sql.Tx, error) error
	// CommitTx commits the Tx
	CommitTx(*sql.Tx) error
}

// NewPGDatasourceName is an initializer for PGDatasourceName, which
// is a struct that holds the PostgreSQL datasource name details.
func NewPGDatasourceName(host, dbname, user, password string, port int) PGDatasourceName {
	return PGDatasourceName{
		DBName:   dbname,
		User:     user,
		Password: password,
		Host:     host,
		Port:     port,
	}
}

// PGDatasourceName is a Postgres datasource name
type PGDatasourceName struct {
	Host     string
	Port     int
	DBName   string
	User     string
	Password string
}

// String returns a formatted PostgreSQL datasource name. If you are
// using a local db with no password, it removes the password from the
// string, otherwise the connection will fail.
func (dsn PGDatasourceName) String() string {
	// Craft string for database connection
	switch dsn.Password {
	case "":
		return fmt.Sprintf("host=%s port=%d dbname=%s user=%s sslmode=disable", dsn.Host, dsn.Port, dsn.DBName, dsn.User)
	default:
		return fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable", dsn.Host, dsn.Port, dsn.DBName, dsn.User, dsn.Password)
	}
}

// NewDefaultDatastore is an initializer for the default Datastore struct
func NewDefaultDatastore(db *sql.DB) DefaultDatastore {
	return DefaultDatastore{db: db}
}

// DefaultDatastore is a concrete implementation for a sql database
type DefaultDatastore struct {
	db *sql.DB
}

// DB returns the sql.Db for the Datastore struct
func (ds DefaultDatastore) DB() *sql.DB {
	return ds.db
}

// BeginTx is a wrapper for sql.DB.BeginTx in order to expose from
// the Datastore interface
func (ds DefaultDatastore) BeginTx(ctx context.Context) (*sql.Tx, error) {
	if ds.db == nil {
		return nil, errs.E(errs.Database, errors.New("DB cannot be nil"))
	}

	tx, err := ds.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errs.E(errs.Database, err)
	}

	return tx, nil
}

// RollbackTx is a wrapper for sql.Tx.Rollback in order to expose from
// the Datastore interface. Proper error handling is also considered.
func (ds DefaultDatastore) RollbackTx(tx *sql.Tx, err error) error {
	if tx == nil {
		return errs.E(errs.Database, errs.Code("nil_tx"), errors.New(fmt.Sprintf("RollbackTx() error = tx cannot be nil: Original error = %s", err.Error())))
	}

	// Attempt to rollback the transaction
	if rollbackErr := tx.Rollback(); rollbackErr != nil {
		return errs.E(errs.Database, errs.Code("rollback_err"), errors.New(fmt.Sprintf("RollbackTx() error = %v: Original error = %s", rollbackErr, err.Error())))
	}

	// If rollback was successful, send back original error
	return err
}

// CommitTx is a wrapper for sql.Tx.Commit in order to expose from
// the Datastore interface. Proper error handling is also considered.
func (ds DefaultDatastore) CommitTx(tx *sql.Tx) error {
	if err := tx.Commit(); err != nil {
		return errs.E(errs.Database, err)
	}

	return nil
}

// NewNullString returns a null if s is empty, otherwise it returns
// the string which was input
func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

// NewNullInt64 returns a null if i == 0, otherwise it returns
// the int64 which was input.
func NewNullInt64(i int64) sql.NullInt64 {
	if i == 0 {
		return sql.NullInt64{}
	}
	return sql.NullInt64{
		Int64: i,
		Valid: true,
	}
}
