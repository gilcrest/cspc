package countrystore

import (
	"context"
	"database/sql"
	"testing"

	"github.com/gilcrest/cspc"
	"github.com/gilcrest/cspc/app"
	"github.com/gilcrest/cspc/datastore"
	"github.com/matryer/is"
	"github.com/rs/zerolog"
)

func TestDB_FindByAlpha2Code(t *testing.T) {

	logger := app.NewLogger(zerolog.DebugLevel)
	db, cleanup, err := datastore.NewDB(datastore.LocalDatastore, logger)
	if err != nil {
		t.Fail()
	}
	defer cleanup()

	type fields struct {
		DB *sql.DB
	}
	f := fields{db}
	ctx := context.Background()
	type args struct {
		ctx context.Context
		a2c string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   cspc.Country
	}{
		{
			name:   "United Kingdom",
			fields: f,
			args:   args{ctx, "GB"},
			want: cspc.Country{
				Name:       "United Kingdom",
				Alpha2Code: "GB",
			},
		},
		{
			name:   "United States",
			fields: f,
			args:   args{ctx, "US"},
			want: cspc.Country{
				Name:       "United States",
				Alpha2Code: "US",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)

			d := &DB{
				DB: tt.fields.DB,
			}
			got, err := d.FindByAlpha2Code(tt.args.ctx, tt.args.a2c)
			is.NoErr(err)
			is.Equal(tt.want.Name, got.Name)
			is.Equal(tt.want.Alpha2Code, got.Alpha2Code)
		})
	}
}
