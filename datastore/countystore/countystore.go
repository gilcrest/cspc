package countystore

import (
	"context"
	"database/sql"

	"github.com/gilcrest/cspc/datastore"

	"github.com/gilcrest/cspc"
	"github.com/gilcrest/cspc/errs"
)

// Transactor performs DML actions against the DB
type Transactor interface {
	CreateCounty(ctx context.Context, stateProv cspc.StateProvince, county cspc.County, tx *sql.Tx) error
}

// NewDefaultTransactor is an initializer for DefaultTransactor
func NewDefaultTransactor(ds datastore.Datastorer) DefaultTransactor {
	return DefaultTransactor{ds}
}

// DefaultTransactor is the default database implementation
// for DML operations
type DefaultTransactor struct {
	datastorer datastore.Datastorer
}

// CreateCounty inserts a record in the lookup.county_lkup table
func (dt DefaultTransactor) CreateCounty(ctx context.Context, stateProv cspc.StateProvince, county cspc.County, tx *sql.Tx) error {
	const op errs.Op = "datastore/countystore/DefaultTransactor.CreateCounty"

	result, execErr := tx.ExecContext(ctx,
		`INSERT INTO lookup.county_lkup (
                               county_id,
                               state_prov_id,
                               county_cd,
                               county_name,
                               latitude_average,
                               longitude_average,
                               create_username, 
                               create_timestamp, 
                               update_username, 
                               update_timestamp) 
                     VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		county.ID,               // $1
		stateProv.ID,            // $2
		county.Code,             // $3
		county.Name,             // $4
		county.LatitudeAverage,  // $5
		county.LongitudeAverage, // $6
		county.CreateUsername,   // $7
		county.CreateTimestamp,  // $8
		county.UpdateUsername,   // $9
		county.UpdateTimestamp)  // $10

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
	FindByCountyCode(ctx context.Context, s cspc.StateProvince, cc string) (cspc.County, error)
}

// NewDefaultSelector is an initializer for DefaultSelector
func NewDefaultSelector(ds datastore.Datastorer) DefaultSelector {
	return DefaultSelector{ds}
}

// DefaultSelector is the database implementation for reading Filings
type DefaultSelector struct {
	datastore.Datastorer
}

// FindByCountyCode returns a County struct given a StateProvince and a County Code
func (d DefaultSelector) FindByCountyCode(ctx context.Context, s cspc.StateProvince, cc string) (cspc.County, error) {
	const op errs.Op = "datastore/countystore/DB.FindByCountyCode"

	db := d.Datastorer.DB()

	// Prepare the sql statement using bind variables
	row := db.QueryRowContext(ctx,
		`select 	l.county_id,
                       	l.county_cd,
       					l.county_name,
       					l.latitude_average,
       					l.longitude_average,
                       	l.create_username,
                       	l.create_timestamp,
                       	l.update_username,
                       	l.update_timestamp
                  from lookup.county_lkup l
                 where l.state_prov_id =  $1
                   and l.county_cd = $2`, s.ID, cc)

	c := new(cspc.County)
	err := row.Scan(
		&c.ID,
		&c.Code,
		&c.Name,
		&c.LatitudeAverage,
		&c.LongitudeAverage,
		&c.CreateUsername,
		&c.CreateTimestamp,
		&c.UpdateUsername,
		&c.UpdateTimestamp)

	if err == sql.ErrNoRows {
		return cspc.County{}, errs.E(op, errs.NotExist, "No record found for given County Code")
	} else if err != nil {
		return cspc.County{}, errs.E(op, err)
	}

	return *c, nil
}
