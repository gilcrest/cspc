package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gilcrest/errs"

	"github.com/gilcrest/cspc/app"
	"github.com/gilcrest/cspc/datastore"
	"github.com/rs/zerolog"

	"github.com/gilcrest/cspc/datastore/countrystore"

	"github.com/gilcrest/cspc"
)

func main() {
	ctx := context.Background()

	err := loadCountries(ctx)
	if err != nil {
		fmt.Println(err)
	}

	// one-time use - not needed going forward
	//names, err := createNameJSON()
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
	//a2j, err := createAlpha2JSON()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//a2, err := json.Marshal(a2j)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(string(a2))
}

// createNameJSON was used to create the CountryNamesJSON constant
func createNameJSON() ([]string, error) {
	const op errs.Op = "main/createNameJSON"

	var (
		countries []cspc.Country
		names     []string
	)

	err := json.Unmarshal([]byte(cspc.CountryFullJSON), &countries)
	if err != nil {
		return nil, errs.E(op, err)
	}

	for _, c := range countries {
		names = append(names, c.Name)
	}

	return names, nil
}

// createAlpha2JSON was used to create the CountryAlpha2CodeJSON constant
func createAlpha2JSON() ([]string, error) {
	const op errs.Op = "main/createNameJSON"

	var (
		countries []cspc.Country
		alpha2s   []string
	)

	err := json.Unmarshal([]byte(cspc.CountryFullJSON), &countries)
	if err != nil {
		return nil, errs.E(op, err)
	}

	for _, c := range countries {
		alpha2s = append(alpha2s, c.Alpha2Code)
	}

	return alpha2s, nil
}

func loadCountries(ctx context.Context) error {
	const op errs.Op = "main/loadCountries"

	var countries []cspc.Country

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
	var transactor countrystore.Transactor

	transactor, err = countrystore.NewTx(tx)
	if err != nil {
		return errs.E(op, err)
	}

	err = json.Unmarshal([]byte(cspc.CountryFullJSON), &countries)
	if err != nil {
		return errs.E(op, err)
	}

	for _, c := range countries {
		fmt.Printf("Country Name = %s, Alpha 2 Code = %s\n", c.Name, c.Alpha2Code)
		err = transactor.CreateCountry(ctx, c)
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
