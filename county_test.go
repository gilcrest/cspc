package cspc

import (
	"reflect"
	"testing"
)

func TestFindCounties(t *testing.T) {

	ma := StateProvince{
		Code:             "MA",
		Name:             "Massachusetts",
		FIPSCode:         "25",
		LatitudeAverage:  "42.407211",
		LongitudeAverage: "-71.382437",
		Counties:         nil,
	}

	var counties []County

	counties = append(counties,
		County{
			Code: "25001",
			Name: "Barnstable County",
		},
		County{
			Code: "25003",
			Name: "Berkshire County",
		},
		County{
			Code: "25005",
			Name: "Bristol County",
		},
		County{
			Code: "25007",
			Name: "Dukes County",
		},
		County{
			Code: "25009",
			Name: "Essex County",
		},
		County{
			Code: "25011",
			Name: "Franklin County",
		},
		County{
			Code: "25013",
			Name: "Hampden County",
		},
		County{
			Code: "25015",
			Name: "Hampshire County",
		},
		County{
			Code: "25017",
			Name: "Middlesex County",
		},
		County{
			Code: "25019",
			Name: "Nantucket County",
		},
		County{
			Code: "25021",
			Name: "Norfolk County",
		},
		County{
			Code: "25023",
			Name: "Plymouth County",
		},
		County{
			Code: "25025",
			Name: "Suffolk County",
		},
		County{
			Code: "25027",
			Name: "Worcester County",
		})

	type args struct {
		sp StateProvince
	}
	tests := []struct {
		name    string
		args    args
		want    []County
		wantErr bool
	}{
		{name: "Test 1", args: args{sp: ma}, want: counties, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindCounties(tt.args.sp)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindCounties() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindCounties() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindUSCountyNames(t *testing.T) {
	ma := StateProvince{
		Code:             "MA",
		Name:             "Massachusetts",
		FIPSCode:         "25",
		LatitudeAverage:  "42.407211",
		LongitudeAverage: "-71.382437",
		Counties:         nil,
	}

	var names []string

	names = append(names,
		"Barnstable County",
		"Berkshire County",
		"Bristol County",
		"Dukes County",
		"Essex County",
		"Franklin County",
		"Hampden County",
		"Hampshire County",
		"Middlesex County",
		"Nantucket County",
		"Norfolk County",
		"Plymouth County",
		"Suffolk County",
		"Worcester County")

	type args struct {
		sp StateProvince
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "Test 1", args: args{sp: ma}, want: names, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindUSCountyNames(tt.args.sp)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindUSCountyNames() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindUSCountyNames() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindUSCountyCodes(t *testing.T) {
	ma := StateProvince{
		Code:             "MA",
		Name:             "Massachusetts",
		FIPSCode:         "25",
		LatitudeAverage:  "42.407211",
		LongitudeAverage: "-71.382437",
		Counties:         nil,
	}

	var codes []string

	codes = append(codes,
		"25001",
		"25003",
		"25005",
		"25007",
		"25009",
		"25011",
		"25013",
		"25015",
		"25017",
		"25019",
		"25021",
		"25023",
		"25025",
		"25027")

	type args struct {
		sp StateProvince
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "Test 1", args: args{sp: ma}, want: codes, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindUSCountyCodes(tt.args.sp)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindUSCountyCodes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindUSCountyCodes() got = %v, want %v", got, tt.want)
			}
		})
	}
}
