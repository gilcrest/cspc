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

// CreateStateProvince inserts a record in the lookup.state_prov_cd_lkup table
func (t *Tx) CreateStateProvince(ctx context.Context, countryAlpha2Code string, s cspc.StateProvince) error {
	const op errs.Op = "datastore/statestore/Tx.CreateStateProvince"

	result, execErr := t.Tx.ExecContext(ctx,
		`INSERT INTO lookup.state_prov_cd_lkup (
                               state_prov_id,
                               country_alpha_2_cd,
                               state_prov_cd,
                               state_name,
                               state_fips_cd,
                               latitude_average,
                               longitude_average,
                               create_username, 
                               create_timestamp, 
                               update_username, 
                               update_timestamp) 
                     VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		uuid.New(),         // $1
		countryAlpha2Code,  // $2
		s.Code,             // $3
		s.Name,             // $4
		s.FIPSCode,         // $5
		s.LatitudeAverage,  // $6
		s.LongitudeAverage, // $7
		"gilcrest",         // $8
		time.Now(),         // $9
		"gilcrest",         // $10
		time.Now())         // $11

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
