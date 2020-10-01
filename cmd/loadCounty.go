package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gilcrest/cspc/datastore/countystore"

	"github.com/gilcrest/cspc/app"
	"github.com/gilcrest/cspc/datastore"
	"github.com/rs/zerolog"

	"github.com/gilcrest/cspc"

	"github.com/gilcrest/errs"
)

type countyInit struct {
	CountyCode string `json:"county_cd"`
	CountyName string `json:"county_name"`
}

func main() {
	s, err := initLoadCounty()
	if err != nil {
		fmt.Println(err)
	}
	ctx := context.Background()
	err = loadCounties2db(ctx, s)
	if err != nil {
		fmt.Println(err)
	}
}

func initLoadCounty() ([]*cspc.StateProvince, error) {
	const op errs.Op = "main/initLoadCounty"

	var (
		counties []countyInit
		states   []*cspc.StateProvince
	)

	err := json.Unmarshal([]byte(cspc.USCountyJSON), &counties)
	if err != nil {
		return nil, errs.E(op, err)
	}

	err = json.Unmarshal([]byte(cspc.USStatesJSON), &states)
	if err != nil {
		return nil, errs.E(op, err)
	}

	for _, state := range states {
		for _, ic := range counties {
			statefip := ic.CountyCode[:2]

			if state.FIPSCode == statefip {

				c := cspc.County{
					Code:             ic.CountyCode,
					Name:             ic.CountyName,
					LatitudeAverage:  "",
					LongitudeAverage: "",
				}

				state.Counties = append(state.Counties, c)
			}
		}
	}

	//for _, s1 := range states {
	//	fmt.Printf("state = %s\n", s1.Name)
	//	for _, c1 := range s1.Counties {
	//		fmt.Printf("\tcounty = %s\n", c1.Name)
	//	}
	//}

	return states, nil
}

func loadCounties2db(ctx context.Context, states []*cspc.StateProvince) error {
	const op errs.Op = "main/loadCounties2db"

	logger := app.NewLogger(zerolog.DebugLevel)

	db, cleanup, err := datastore.NewDB(datastore.LocalDatastore, logger)
	if err != nil {
		return errs.E(op, err)
	}
	defer cleanup()

	// declare variable as the Transactor interface
	var datastorer datastore.Datastorer

	datastorer = datastore.NewDatastore(db)

	a := app.NewApplication(app.Local, datastorer, logger)

	tx, err := a.Datastorer.BeginTx(ctx)
	if err != nil {
		return errs.E(op, err)
	}

	// declare variable as the Transactor interface
	var transactor countystore.Transactor

	transactor, err = countystore.NewTx(tx)
	if err != nil {
		return errs.E(op, err)
	}

	for _, s := range states {
		for _, c := range s.Counties {
			err = transactor.CreateCounty(ctx, "US", s.Code, c)
			if err != nil {
				return errs.E(op, a.Datastorer.RollbackTx(tx, err))
			}
		}
	}

	// Commit the Transaction
	if err = a.Datastorer.CommitTx(tx); err != nil {
		return errs.E(op, err)
	}

	return nil
}
