package cspc

import (
	"encoding/json"

	"github.com/gilcrest/errs"
)

// USStatesJSON is the JSON array for all US States (including DC and Puerto Rico)
const USStatesJSON string = `[{"state_prov_cd":"AL","state_name":"Alabama","state_fips_cd":"01","latitude_average":"32.318231","longitude_average":"-86.902298"}, 
 {"state_prov_cd":"AK","state_name":"Alaska","state_fips_cd":"02","latitude_average":"63.588753","longitude_average":"-154.493062"}, 
 {"state_prov_cd":"AZ","state_name":"Arizona","state_fips_cd":"04","latitude_average":"34.048928","longitude_average":"-111.093731"}, 
 {"state_prov_cd":"AR","state_name":"Arkansas","state_fips_cd":"05","latitude_average":"35.20105","longitude_average":"-91.831833"}, 
 {"state_prov_cd":"CA","state_name":"California","state_fips_cd":"06","latitude_average":"36.778261","longitude_average":"-119.417932"}, 
 {"state_prov_cd":"CO","state_name":"Colorado","state_fips_cd":"08","latitude_average":"39.550051","longitude_average":"-105.782067"}, 
 {"state_prov_cd":"CT","state_name":"Connecticut","state_fips_cd":"09","latitude_average":"41.603221","longitude_average":"-73.087749"}, 
 {"state_prov_cd":"DE","state_name":"Delaware","state_fips_cd":"10","latitude_average":"38.910832","longitude_average":"-75.52767"}, 
 {"state_prov_cd":"DC","state_name":"District of Columbia","state_fips_cd":"11","latitude_average":"38.905985","longitude_average":"-77.033418"}, 
 {"state_prov_cd":"FL","state_name":"Florida","state_fips_cd":"12","latitude_average":"27.664827","longitude_average":"-81.515754"}, 
 {"state_prov_cd":"GA","state_name":"Georgia","state_fips_cd":"13","latitude_average":"32.157435","longitude_average":"-82.907123"}, 
 {"state_prov_cd":"HI","state_name":"Hawaii","state_fips_cd":"15","latitude_average":"19.898682","longitude_average":"-155.665857"}, 
 {"state_prov_cd":"ID","state_name":"Idaho","state_fips_cd":"16","latitude_average":"44.068202","longitude_average":"-114.742041"}, 
 {"state_prov_cd":"IL","state_name":"Illinois","state_fips_cd":"17","latitude_average":"40.633125","longitude_average":"-89.398528"}, 
 {"state_prov_cd":"IN","state_name":"Indiana","state_fips_cd":"18","latitude_average":"40.551217","longitude_average":"-85.602364"}, 
 {"state_prov_cd":"IA","state_name":"Iowa","state_fips_cd":"19","latitude_average":"41.878003","longitude_average":"-93.097702"}, 
 {"state_prov_cd":"KS","state_name":"Kansas","state_fips_cd":"20","latitude_average":"39.011902","longitude_average":"-98.484246"}, 
 {"state_prov_cd":"KY","state_name":"Kentucky","state_fips_cd":"21","latitude_average":"37.839333","longitude_average":"-84.270018"}, 
 {"state_prov_cd":"LA","state_name":"Louisiana","state_fips_cd":"22","latitude_average":"31.244823","longitude_average":"-92.145024"}, 
 {"state_prov_cd":"ME","state_name":"Maine","state_fips_cd":"23","latitude_average":"45.253783","longitude_average":"-69.445469"}, 
 {"state_prov_cd":"MD","state_name":"Maryland","state_fips_cd":"24","latitude_average":"39.045755","longitude_average":"-76.641271"}, 
 {"state_prov_cd":"MA","state_name":"Massachusetts","state_fips_cd":"25","latitude_average":"42.407211","longitude_average":"-71.382437"}, 
 {"state_prov_cd":"MI","state_name":"Michigan","state_fips_cd":"26","latitude_average":"44.314844","longitude_average":"-85.602364"}, 
 {"state_prov_cd":"MN","state_name":"Minnesota","state_fips_cd":"27","latitude_average":"46.729553","longitude_average":"-94.6859"}, 
 {"state_prov_cd":"MS","state_name":"Mississippi","state_fips_cd":"28","latitude_average":"32.354668","longitude_average":"-89.398528"}, 
 {"state_prov_cd":"MO","state_name":"Missouri","state_fips_cd":"29","latitude_average":"37.964253","longitude_average":"-91.831833"}, 
 {"state_prov_cd":"MT","state_name":"Montana","state_fips_cd":"30","latitude_average":"46.879682","longitude_average":"-110.362566"}, 
 {"state_prov_cd":"NE","state_name":"Nebraska","state_fips_cd":"31","latitude_average":"41.492537","longitude_average":"-99.901813"}, 
 {"state_prov_cd":"NV","state_name":"Nevada","state_fips_cd":"32","latitude_average":"38.80261","longitude_average":"-116.419389"}, 
 {"state_prov_cd":"NH","state_name":"New Hampshire","state_fips_cd":"33","latitude_average":"43.193852","longitude_average":"-71.572395"}, 
 {"state_prov_cd":"NJ","state_name":"New Jersey","state_fips_cd":"34","latitude_average":"40.058324","longitude_average":"-74.405661"}, 
 {"state_prov_cd":"NM","state_name":"New Mexico","state_fips_cd":"35","latitude_average":"34.97273","longitude_average":"-105.032363"}, 
 {"state_prov_cd":"NY","state_name":"New York","state_fips_cd":"36","latitude_average":"43.299428","longitude_average":"-74.217933"}, 
 {"state_prov_cd":"NC","state_name":"North Carolina","state_fips_cd":"37","latitude_average":"35.759573","longitude_average":"-79.0193"}, 
 {"state_prov_cd":"ND","state_name":"North Dakota","state_fips_cd":"38","latitude_average":"47.551493","longitude_average":"-101.002012"}, 
 {"state_prov_cd":"OH","state_name":"Ohio","state_fips_cd":"39","latitude_average":"40.417287","longitude_average":"-82.907123"}, 
 {"state_prov_cd":"OK","state_name":"Oklahoma","state_fips_cd":"40","latitude_average":"35.007752","longitude_average":"-97.092877"}, 
 {"state_prov_cd":"OR","state_name":"Oregon","state_fips_cd":"41","latitude_average":"43.804133","longitude_average":"-120.554201"}, 
 {"state_prov_cd":"PA","state_name":"Pennsylvania","state_fips_cd":"42","latitude_average":"41.203322","longitude_average":"-77.194525"}, 
 {"state_prov_cd":"PR","state_name":"Puerto Rico","state_fips_cd":"72","latitude_average":"18.220833","longitude_average":"-66.590149"}, 
 {"state_prov_cd":"RI","state_name":"Rhode Island","state_fips_cd":"44","latitude_average":"41.580095","longitude_average":"-71.477429"}, 
 {"state_prov_cd":"SC","state_name":"South Carolina","state_fips_cd":"45","latitude_average":"33.836081","longitude_average":"-81.163725"}, 
 {"state_prov_cd":"SD","state_name":"South Dakota","state_fips_cd":"46","latitude_average":"43.969515","longitude_average":"-99.901813"}, 
 {"state_prov_cd":"TN","state_name":"Tennessee","state_fips_cd":"47","latitude_average":"35.517491","longitude_average":"-86.580447"}, 
 {"state_prov_cd":"TX","state_name":"Texas","state_fips_cd":"48","latitude_average":"31.968599","longitude_average":"-99.901813"}, 
 {"state_prov_cd":"UT","state_name":"Utah","state_fips_cd":"49","latitude_average":"39.32098","longitude_average":"-111.093731"}, 
 {"state_prov_cd":"VT","state_name":"Vermont","state_fips_cd":"50","latitude_average":"44.558803","longitude_average":"-72.577841"}, 
 {"state_prov_cd":"VA","state_name":"Virginia","state_fips_cd":"51","latitude_average":"37.431573","longitude_average":"-78.656894"}, 
 {"state_prov_cd":"WA","state_name":"Washington","state_fips_cd":"53","latitude_average":"47.751074","longitude_average":"-120.740139"}, 
 {"state_prov_cd":"WV","state_name":"West Virginia","state_fips_cd":"54","latitude_average":"38.597626","longitude_average":"-80.454903"}, 
 {"state_prov_cd":"WI","state_name":"Wisconsin","state_fips_cd":"55","latitude_average":"43.78444","longitude_average":"-88.787868"}, 
 {"state_prov_cd":"WY","state_name":"Wyoming","state_fips_cd":"56","latitude_average":"43.075968","longitude_average":"-107.290284"}]`

// USStatesNameJSON is the JSON array for all US State Names
// (including DC and Puerto Rico) in alpha order by name
const USStatesNameJSON string = `["Alabama","Alaska","Arizona","Arkansas","California","Colorado","Connecticut","Delaware","District of Columbia","Florida","Georgia","Hawaii","Idaho","Illinois","Indiana","Iowa","Kansas","Kentucky","Louisiana","Maine","Maryland","Massachusetts","Michigan","Minnesota","Mississippi","Missouri","Montana","Nebraska","Nevada","New Hampshire","New Jersey","New Mexico","New York","North Carolina","North Dakota","Ohio","Oklahoma","Oregon","Pennsylvania","Puerto Rico","Rhode Island","South Carolina","South Dakota","Tennessee","Texas","Utah","Vermont","Virginia","Washington","West Virginia","Wisconsin","Wyoming"]`

// USStatesCodeJSON is the JSON array for all US State Codes
// (including DC and Puerto Rico) in alpha order by name
const USStatesCodeJSON string = `["AL","AK","AZ","AR","CA","CO","CT","DE","DC","FL","GA","HI","ID","IL","IN","IA","KS","KY","LA","ME","MD","MA","MI","MN","MS","MO","MT","NE","NV","NH","NJ","NM","NY","NC","ND","OH","OK","OR","PA","PR","RI","SC","SD","TN","TX","UT","VT","VA","WA","WV","WI","WY"]`

// StateProvince is a State or Province of a Country
// that makes up a particular territory
type StateProvince struct {
	Code             string `json:"state_prov_cd"`
	Name             string `json:"state_name"`
	FIPSCode         string `json:"state_fips_cd"`
	LatitudeAverage  string `json:"latitude_average"`
	LongitudeAverage string `json:"longitude_average"`
	Counties         []County
}

// USStates returns all US States
func USStates() ([]StateProvince, error) {
	const op errs.Op = "cspc/USStates"

	var states []StateProvince

	err := json.Unmarshal([]byte(USStatesJSON), &states)
	if err != nil {
		return nil, errs.E(op, err)
	}

	return states, nil
}
