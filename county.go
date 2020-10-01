package cspc

// County is a geographical region of a country used for administrative
// or other purposes, in certain modern nations.
type County struct {
	Code             string `json:"county_cd"`
	Name             string `json:"county_name"`
	FIPSCode         int    `json:"county_fips_cd"`
	LatitudeAverage  string `json:"latitude_average"`
	LongitudeAverage string `json:"longitude_average"`
}
