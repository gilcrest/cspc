package statestore

import (
	"context"
	"database/sql"

	"github.com/gilcrest/cspc/datastore"

	"github.com/gilcrest/cspc"
	"github.com/gilcrest/errs"
)

// Transactor performs DML actions against the DB
type Transactor interface {
	CreateStateProvince(ctx context.Context, args *CreateArgs) error
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
}

// NewCreateArgs is an initializer for CreateArgs
func NewCreateArgs(country cspc.Country, stateProv cspc.StateProvince, username string) *CreateArgs {
	return &CreateArgs{Country: country, StateProv: stateProv}
}

// CreateStateProvince inserts a record in the lookup.state_prov_cd_lkup table
func (t *Tx) CreateStateProvince(ctx context.Context, args *CreateArgs) error {
	const op errs.Op = "datastore/statestore/Tx.CreateStateProvince"

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
		args.StateProv.ID,               // $1
		args.Country.ID,                 // $2
		args.StateProv.Code,             // $3
		args.StateProv.Name,             // $4
		args.StateProv.FIPSCode,         // $5
		args.StateProv.LatitudeAverage,  // $6
		args.StateProv.LongitudeAverage, // $7
		args.StateProv.CreateUsername,   // $8
		args.StateProv.CreateTimestamp,  // $9
		args.StateProv.UpdateUsername,   // $10
		args.StateProv.UpdateTimestamp)  // $11

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
	FindByStateProvCode(ctx context.Context, c cspc.Country, spc string) (cspc.StateProvince, error)
}

// NewDefaultSelector is an initializer for DefaultSelector
func NewDefaultSelector(ds datastore.Datastorer) DefaultSelector {
	return DefaultSelector{ds}
}

// DefaultSelector is the database implementation for reading Filings
type DefaultSelector struct {
	datastore.Datastorer
}

// FindByStateProvCode returns a StateProvince struct given an Alpha 2 Code
func (d DefaultSelector) FindByStateProvCode(ctx context.Context, c cspc.Country, spc string) (cspc.StateProvince, error) {
	const op errs.Op = "datastore/statestore/DB.FindByStateProvCode"

	db := d.Datastorer.DB()

	// Prepare the sql statement using bind variables
	row := db.QueryRowContext(ctx,
		`select 	l.state_prov_id,
                       	l.state_prov_cd,
                   		l.state_name,
       					l.state_fips_cd,
       					l.latitude_average,
       					l.longitude_average,
                       	l.create_username,
                       	l.create_timestamp,
                       	l.update_username,
                       	l.update_timestamp
                  from lookup.state_prov_lkup l
                 where l.country_id = $1
                   and l.state_prov_cd =  $2`, c.ID, spc)

	sp := new(cspc.StateProvince)
	err := row.Scan(
		&sp.ID,
		&sp.Code,
		&sp.Name,
		&sp.FIPSCode,
		&sp.LatitudeAverage,
		&sp.LongitudeAverage,
		&sp.CreateUsername,
		&sp.CreateTimestamp,
		&sp.UpdateUsername,
		&sp.UpdateTimestamp)

	if err == sql.ErrNoRows {
		return cspc.StateProvince{}, errs.E(op, errs.NotExist, "No record found for given State/Province Code")
	} else if err != nil {
		return cspc.StateProvince{}, errs.E(op, err)
	}

	return *sp, nil
}
