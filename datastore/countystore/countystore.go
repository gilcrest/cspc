package countystore

import (
	"context"
	"database/sql"

	"github.com/gilcrest/cspc"
	"github.com/gilcrest/errs"
)

// Transactor performs DML actions against the DB
type Transactor interface {
	CreateCounty(ctx context.Context, args *CreateArgs) error
}

// NewTx initializes a pointer to a Tx struct that holds a *sql.Tx
func NewTx(tx *sql.Tx) (*Tx, error) {
	const op errs.Op = "datastore/countystore/NewTx"
	if tx == nil {
		return nil, errs.E(op, errs.MissingField("tx"))
	}
	return &Tx{Tx: tx}, nil
}

// Tx stores a sql.Tx which will be used for all DML operations
type Tx struct {
	*sql.Tx
}

// CreateArgs are the arguments for CreateCounty
type CreateArgs struct {
	StateProv cspc.StateProvince
	County    cspc.County
}

// NewCreateArgs is an initializer for the CreateArgs struct
func NewCreateArgs(stateProv cspc.StateProvince, county cspc.County) *CreateArgs {
	return &CreateArgs{StateProv: stateProv, County: county}
}

// CreateCounty inserts a record in the lookup.county_lkup table
func (t *Tx) CreateCounty(ctx context.Context, args *CreateArgs) error {
	const op errs.Op = "datastore/countystore/Tx.CreateCounty"

	result, execErr := t.Tx.ExecContext(ctx,
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
		args.County.ID,               // $1
		args.StateProv.ID,            // $2
		args.County.Code,             // $3
		args.County.Name,             // $4
		args.County.LatitudeAverage,  // $5
		args.County.LongitudeAverage, // $6
		args.County.CreateUsername,   // $7
		args.County.CreateTimestamp,  // $8
		args.County.UpdateUsername,   // $9
		args.County.UpdateTimestamp)  // $10

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

// NewDB is an initializer for DB
func NewDB(db *sql.DB) (*DB, error) {
	const op errs.Op = "datastore/countystore/NewDB"
	if db == nil {
		return nil, errs.E(op, errs.MissingField("db"))
	}
	return &DB{DB: db}, nil
}

// DB  struct holds a pointer to a sql database
type DB struct {
	*sql.DB
}

// FindByCountyCode returns a County struct given a StateProvince and a County Code
func (d *DB) FindByCountyCode(ctx context.Context, s cspc.StateProvince, cc string) (cspc.County, error) {
	const op errs.Op = "datastore/countystore/DB.FindByCountyCode"

	// Prepare the sql statement using bind variables
	row := d.DB.QueryRowContext(ctx,
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
