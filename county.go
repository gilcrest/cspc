package cspc

// County is a geographical region of a country used for administrative
// or other purposes, in certain modern nations. For the US, the Code
// is the Federal Information Processing System (FIPS) Codes for States
// and Counties. FIPS codes are numbers which uniquely identify geographic
// areas.
type County struct {
	Code             string `json:"county_cd"`
	Name             string `json:"county_name"`
	LatitudeAverage  string `json:"latitude_average"`
	LongitudeAverage string `json:"longitude_average"`
}
