package countrystore

import (
	"context"
	"os"
	"testing"

	"github.com/rs/zerolog"

	"github.com/gilcrest/cspc/logger"

	"github.com/gilcrest/cspc/datastoretest"

	"github.com/gilcrest/cspc"
	"github.com/matryer/is"
)

func TestDB_FindByAlpha2Code(t *testing.T) {

	lgr := logger.NewLogger(os.Stdout, zerolog.DebugLevel, true)

	ds, cleanup := datastoretest.NewDefaultDatastore(t, lgr)
	t.Cleanup(cleanup)

	ctx := context.Background()
	type args struct {
		ctx context.Context
		a2c string
	}
	tests := []struct {
		name string
		args args
		want cspc.Country
	}{
		{
			name: "United Kingdom",
			args: args{ctx, "GB"},
			want: cspc.Country{
				Name:       "United Kingdom",
				Alpha2Code: "GB",
			},
		},
		{
			name: "United States",
			args: args{ctx, "US"},
			want: cspc.Country{
				Name:       "United States",
				Alpha2Code: "US",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)

			d := NewDefaultSelector(ds)
			got, err := d.FindByAlpha2Code(tt.args.ctx, tt.args.a2c)
			is.NoErr(err)
			is.Equal(tt.want.Name, got.Name)
			is.Equal(tt.want.Alpha2Code, got.Alpha2Code)
		})
	}
}
