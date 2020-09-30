package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gilcrest/cspc/datastore/statestore"

	"github.com/gilcrest/cspc"
	"github.com/gilcrest/cspc/app"
	"github.com/gilcrest/cspc/datastore"
	"github.com/gilcrest/errs"
	"github.com/rs/zerolog"
)

func main() {
	ctx := context.Background()

	err := loadUSStates(ctx)
	if err != nil {
		fmt.Println(err)
	}

	//// one-time use - not needed going forward
	//names, err := createUSStatesNameJSON()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//n, err := json.Marshal(names)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(string(n))

	// one-time use - not needed going forward
	//usj, err := createUSStatesCodeJSON()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//u, err := json.Marshal(usj)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(string(u))
}

// createUSStatesNameJSON was used to create the USStatesNameJSON constant
func createUSStatesNameJSON() ([]string, error) {
	const op errs.Op = "main/createStateProvinceNameJSON"

	var (
		states []cspc.StateProvince
		names  []string
	)

	err := json.Unmarshal([]byte(cspc.USStatesJSON), &states)
	if err != nil {
		return nil, errs.E(op, err)
	}

	for _, s := range states {
		names = append(names, s.Name)
	}

	return names, nil
}

// createUSStatesCodeJSON was used to create the USStatesCodeJSON constant
func createUSStatesCodeJSON() ([]string, error) {
	const op errs.Op = "main/createUSStatesCodeJSON"

	var (
		states []cspc.StateProvince
		codes  []string
	)

	err := json.Unmarshal([]byte(cspc.USStatesJSON), &states)
	if err != nil {
		return nil, errs.E(op, err)
	}

	for _, s := range states {
		codes = append(codes, s.Code)
	}

	return codes, nil
}

func loadUSStates(ctx context.Context) error {
	const op errs.Op = "main/loadStates"

	var states []cspc.StateProvince

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
	var transactor statestore.Transactor

	transactor, err = statestore.NewTx(tx)
	if err != nil {
		return errs.E(op, err)
	}

	err = json.Unmarshal([]byte(cspc.USStatesJSON), &states)
	if err != nil {
		return errs.E(op, err)
	}

	for _, s := range states {
		fmt.Printf("State Name = %s, Alpha 2 Code = %s\n", s.Name, s.Code)
		err = transactor.CreateStateProvince(ctx, "US", s)
		if err != nil {
			return errs.E(op, err)
		}
	}

	// Commit the Transaction
	if err = a.Datastorer.CommitTx(tx); err != nil {
		return errs.E(op, err)
	}

	return nil
}
