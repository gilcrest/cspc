package countrystore

import (
	"context"
	"database/sql"

	"github.com/gilcrest/cspc/datastore"

	"github.com/gilcrest/cspc"
	"github.com/gilcrest/errs"
)

// Transactor performs DML actions against the DB
type Transactor interface {
	CreateCountry(ctx context.Context, c *cspc.Country) error
}

// NewTx initializes a pointer to a Tx struct that holds a *sql.Tx
func NewTx(tx *sql.Tx) (*Tx, error) {
	const op errs.Op = "datastore/countrystore/NewTx"
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
func (t *Tx) CreateCountry(ctx context.Context, c *cspc.Country) error {
	const op errs.Op = "datastore/countrystore/Tx.Create"

	result, execErr := t.Tx.ExecContext(ctx,
		`INSERT INTO lookup.country_lkup (
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
		c.ID,               // $1
		c.Alpha2Code,       // $2
		c.Alpha3Code,       // $3
		c.UNM49Code,        // $4
		c.Name,             // $5
		c.LatitudeAverage,  // $6
		c.LongitudeAverage, // $7
		c.CreateUsername,   // $8
		c.CreateTimestamp,  // $9
		c.UpdateUsername,   // $10
		c.UpdateTimestamp)  // $11

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

// Selector reads records from the db
type Selector interface {
	FindByAlpha2Code(ctx context.Context, a2c string) (cspc.Country, error)
}

// NewDefaultSelector is an initializer for DefaultSelector
func NewDefaultSelector(ds datastore.Datastorer) DefaultSelector {
	return DefaultSelector{ds}
}

// DefaultSelector is the database implementation for reading Filings
type DefaultSelector struct {
	datastore.Datastorer
}

// FindByAlpha2Code returns a Country struct given an Alpha 2 Code
func (d DefaultSelector) FindByAlpha2Code(ctx context.Context, a2c string) (cspc.Country, error) {
	const op errs.Op = "datastore/countrystore/DB.FindByAlpha2Code"

	db := d.Datastorer.DB()

	// Prepare the sql statement using bind variables
	row := db.QueryRowContext(ctx,
		`select 	l.country_id,
                       	l.country_alpha_2_cd,
                       	l.country_alpha_3_cd,
                   		l.country_un_m49_cd,
       					l.country_name,
       					l.latitude_average,
       					l.longitude_average,
                       	l.create_username,
                       	l.create_timestamp,
                       	l.update_username,
                       	l.update_timestamp
                  from lookup.country_lkup l
                 where l.country_alpha_2_cd =  $1`, a2c)

	c := new(cspc.Country)
	err := row.Scan(
		&c.ID,
		&c.Alpha2Code,
		&c.Alpha3Code,
		&c.UNM49Code,
		&c.Name,
		&c.LatitudeAverage,
		&c.LongitudeAverage,
		&c.CreateUsername,
		&c.CreateTimestamp,
		&c.UpdateUsername,
		&c.UpdateTimestamp)

	if err == sql.ErrNoRows {
		return cspc.Country{}, errs.E(op, errs.NotExist, "No record found for given Alpha 2 Code")
	} else if err != nil {
		return cspc.Country{}, errs.E(op, err)
	}

	return *c, nil
}
