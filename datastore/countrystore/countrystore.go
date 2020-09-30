package countrystore

import (
	"context"
	"database/sql"
	"time"

	"github.com/gilcrest/cspc"
	"github.com/google/uuid"

	"github.com/gilcrest/errs"
)

// Transactor performs DML actions against the DB
type Transactor interface {
	CreateCountry(ctx context.Context, c cspc.Country) error
}

// NewTx initializes a pointer to a Tx struct that holds a *sql.Tx
func NewTx(tx *sql.Tx) (*Tx, error) {
	const op errs.Op = "datastore/jurisdictionstore/NewTx"
	if tx == nil {
		return nil, errs.E(op, errs.MissingField("tx"))
	}
	return &Tx{Tx: tx}, nil
}

// Tx stores a sql.Tx which will be used for all DML operations
type Tx struct {
	*sql.Tx
}

// CreateCountry inserts a record in the lookup.country_cd_lkup table
func (t *Tx) CreateCountry(ctx context.Context, c cspc.Country) error {
	const op errs.Op = "datastore/jurisdictionstore/Tx.Create"

	result, execErr := t.Tx.ExecContext(ctx,
		`INSERT INTO lookup.country_cd_lkup (
                               country_id,
                               country_alpha_2_cd, 
                               country_alpha_3_cd, 
                               country_un_m49_cd, 
                               country_name,
                               latitude_average,
                               longitude_average,
                               create_username, 
                               create_timestamp, 
                               update_username, 
                               update_timestamp) 
                     VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		uuid.New(),         // $1
		c.Alpha2Code,       // $2
		c.Alpha3Code,       // $3
		c.UNM49Code,        // $4
		c.Name,             // $5
		c.LatitudeAverage,  // $6
		c.LongitudeAverage, // $7
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
