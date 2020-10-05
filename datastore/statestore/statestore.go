package statestore

import (
	"context"
	"database/sql"
	"time"

	"github.com/gilcrest/cspc"
	"github.com/gilcrest/errs"
	"github.com/google/uuid"
)

// Transactor performs DML actions against the DB
type Transactor interface {
	CreateStateProvince(ctx context.Context, countryAlpha2Code string, s cspc.StateProvince) error
}

// NewTx initializes a pointer to a Tx struct that holds a *sql.Tx
func NewTx(tx *sql.Tx) (*Tx, error) {
	const op errs.Op = "datastore/statestore/NewTx"
	if tx == nil {
		return nil, errs.E(op, errs.MissingField("tx"))
	}
	return &Tx{Tx: tx}, nil
}

// Tx stores a sql.Tx which will be used for all DML operations
type Tx struct {
	*sql.Tx
}

// CreateArgs are the arguments for
type CreateArgs struct {
	Country   cspc.Country
	StateProv cspc.StateProvince
	Username  string
}

// NewCreateArgs is an initializer for CreateArgs
func NewCreateArgs(country cspc.Country, stateProv cspc.StateProvince, username string) *CreateArgs {
	return &CreateArgs{Country: country, StateProv: stateProv, Username: username}
}

// CreateStateProvince inserts a record in the lookup.state_prov_cd_lkup table
func (t *Tx) CreateStateProvince(ctx context.Context, args CreateArgs) error {
	const op errs.Op = "datastore/statestore/Tx.CreateStateProvince"

	now := time.Now()

	result, execErr := t.Tx.ExecContext(ctx,
		`INSERT INTO lookup.state_prov_lkup (
                               state_prov_id,
                               country_id,
                               state_prov_cd,
                               state_name,
                               state_fips_cd,
                               latitude_average,
                               longitude_average,
                               create_username, 
                               create_timestamp, 
                               update_username, 
                               update_timestamp) 
                     VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		uuid.New(),                      // $1
		args.Country.ID,                 // $2
		args.StateProv.Code,             // $3
		args.StateProv.Name,             // $4
		args.StateProv.FIPSCode,         // $5
		args.StateProv.LatitudeAverage,  // $6
		args.StateProv.LongitudeAverage, // $7
		args.Username,                   // $8
		now,                             // $9
		args.Username,                   // $10
		now)                             // $11

	if execErr != nil {
		return errs.E(op, errs.Database, execErr)
	}

	// Only 1 row should be inserted, check the result count to
	// ensure this is correct
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errs.E(op, errs.Database, err)
	}
	if rowsAffected == 0 {
		return errs.E(op, errs.Database, "No Rows Updated")
	} else if rowsAffected > 1 {
		return errs.E(op, errs.Database, "Too Many Rows Inserted")
	}

	return nil
}
