package statestore

import (
	"context"
	"database/sql"
	"testing"

	"github.com/matryer/is"

	"github.com/gilcrest/cspc"
	"github.com/gilcrest/cspc/datastore"
	"github.com/gilcrest/cspc/datastore/countrystore"
	"github.com/gilcrest/cspc/datastoretest"
)

func TestDB_FindByStateProvCode(t *testing.T) {
	db, cleanup := datastoretest.NewDB(t)
	defer cleanup()

	type fields struct {
		DB *sql.DB
	}
	f := fields{db}
	ctx := context.Background()
	type args struct {
		ctx context.Context
		c   cspc.Country
		spc string
	}

	countrySelector := countrystore.NewDefaultSelector(datastore.NewDefaultDatastore(db))

	us, err := countrySelector.FindByAlpha2Code(ctx, "US")
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    cspc.StateProvince
		wantErr bool
	}{
		{
			name:   "Massachusetts",
			fields: f,
			args:   args{ctx: ctx, c: us, spc: "MA"},
			want: cspc.StateProvince{
				Code: "MA",
				Name: "Massachusetts",
			},
			wantErr: false,
		},
		{
			name:   "Colorado",
			fields: f,
			args:   args{ctx: ctx, c: us, spc: "CO"},
			want: cspc.StateProvince{
				Code: "CO",
				Name: "Colorado",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			d := &DefaultSelector{
				datastore.NewDefaultDatastore(tt.fields.DB),
			}
			got, err := d.FindByStateProvCode(tt.args.ctx, tt.args.c, tt.args.spc)
			is.NoErr(err)
			is.Equal(tt.want.Name, got.Name)
			is.Equal(tt.want.Code, got.Code)
		})
	}
}
