package cspc

// USStatesJSON is the JSON array for all US States (including DC and Puerto Rico)
const USStatesJSON string = `[{"state_prov_cd":"AL","state_name":"Alabama","latitude_average":"32.318231","longitude_average":"-86.902298"},
{"state_prov_cd":"AK","state_name":"Alaska","latitude_average":"63.588753","longitude_average":"-154.493062"},
{"state_prov_cd":"AZ","state_name":"Arizona","latitude_average":"34.048928","longitude_average":"-111.093731"},
{"state_prov_cd":"AR","state_name":"Arkansas","latitude_average":"35.20105","longitude_average":"-91.831833"},
{"state_prov_cd":"CA","state_name":"California","latitude_average":"36.778261","longitude_average":"-119.417932"},
{"state_prov_cd":"CO","state_name":"Colorado","latitude_average":"39.550051","longitude_average":"-105.782067"},
{"state_prov_cd":"CT","state_name":"Connecticut","latitude_average":"41.603221","longitude_average":"-73.087749"},
{"state_prov_cd":"DE","state_name":"Delaware","latitude_average":"38.910832","longitude_average":"-75.52767"},
{"state_prov_cd":"DC","state_name":"District of Columbia","latitude_average":"38.905985","longitude_average":"-77.033418"},
{"state_prov_cd":"FL","state_name":"Florida","latitude_average":"27.664827","longitude_average":"-81.515754"},
{"state_prov_cd":"GA","state_name":"Georgia","latitude_average":"32.157435","longitude_average":"-82.907123"},
{"state_prov_cd":"HI","state_name":"Hawaii","latitude_average":"19.898682","longitude_average":"-155.665857"},
{"state_prov_cd":"ID","state_name":"Idaho","latitude_average":"44.068202","longitude_average":"-114.742041"},
{"state_prov_cd":"IL","state_name":"Illinois","latitude_average":"40.633125","longitude_average":"-89.398528"},
{"state_prov_cd":"IN","state_name":"Indiana","latitude_average":"40.551217","longitude_average":"-85.602364"},
{"state_prov_cd":"IA","state_name":"Iowa","latitude_average":"41.878003","longitude_average":"-93.097702"},
{"state_prov_cd":"KS","state_name":"Kansas","latitude_average":"39.011902","longitude_average":"-98.484246"},
{"state_prov_cd":"KY","state_name":"Kentucky","latitude_average":"37.839333","longitude_average":"-84.270018"},
{"state_prov_cd":"LA","state_name":"Louisiana","latitude_average":"31.244823","longitude_average":"-92.145024"},
{"state_prov_cd":"ME","state_name":"Maine","latitude_average":"45.253783","longitude_average":"-69.445469"},
{"state_prov_cd":"MD","state_name":"Maryland","latitude_average":"39.045755","longitude_average":"-76.641271"},
{"state_prov_cd":"MA","state_name":"Massachusetts","latitude_average":"42.407211","longitude_average":"-71.382437"},
{"state_prov_cd":"MI","state_name":"Michigan","latitude_average":"44.314844","longitude_average":"-85.602364"},
{"state_prov_cd":"MN","state_name":"Minnesota","latitude_average":"46.729553","longitude_average":"-94.6859"},
{"state_prov_cd":"MS","state_name":"Mississippi","latitude_average":"32.354668","longitude_average":"-89.398528"},
{"state_prov_cd":"MO","state_name":"Missouri","latitude_average":"37.964253","longitude_average":"-91.831833"},
{"state_prov_cd":"MT","state_name":"Montana","latitude_average":"46.879682","longitude_average":"-110.362566"},
{"state_prov_cd":"NE","state_name":"Nebraska","latitude_average":"41.492537","longitude_average":"-99.901813"},
{"state_prov_cd":"NV","state_name":"Nevada","latitude_average":"38.80261","longitude_average":"-116.419389"},
{"state_prov_cd":"NH","state_name":"New Hampshire","latitude_average":"43.193852","longitude_average":"-71.572395"},
{"state_prov_cd":"NJ","state_name":"New Jersey","latitude_average":"40.058324","longitude_average":"-74.405661"},
{"state_prov_cd":"NM","state_name":"New Mexico","latitude_average":"34.97273","longitude_average":"-105.032363"},
{"state_prov_cd":"NY","state_name":"New York","latitude_average":"43.299428","longitude_average":"-74.217933"},
{"state_prov_cd":"NC","state_name":"North Carolina","latitude_average":"35.759573","longitude_average":"-79.0193"},
{"state_prov_cd":"ND","state_name":"North Dakota","latitude_average":"47.551493","longitude_average":"-101.002012"},
{"state_prov_cd":"OH","state_name":"Ohio","latitude_average":"40.417287","longitude_average":"-82.907123"},
{"state_prov_cd":"OK","state_name":"Oklahoma","latitude_average":"35.007752","longitude_average":"-97.092877"},
{"state_prov_cd":"OR","state_name":"Oregon","latitude_average":"43.804133","longitude_average":"-120.554201"},
{"state_prov_cd":"PA","state_name":"Pennsylvania","latitude_average":"41.203322","longitude_average":"-77.194525"},
{"state_prov_cd":"PR","state_name":"Puerto Rico","latitude_average":"18.220833","longitude_average":"-66.590149"},
{"state_prov_cd":"RI","state_name":"Rhode Island","latitude_average":"41.580095","longitude_average":"-71.477429"},
{"state_prov_cd":"SC","state_name":"South Carolina","latitude_average":"33.836081","longitude_average":"-81.163725"},
{"state_prov_cd":"SD","state_name":"South Dakota","latitude_average":"43.969515","longitude_average":"-99.901813"},
{"state_prov_cd":"TN","state_name":"Tennessee","latitude_average":"35.517491","longitude_average":"-86.580447"},
{"state_prov_cd":"TX","state_name":"Texas","latitude_average":"31.968599","longitude_average":"-99.901813"},
{"state_prov_cd":"UT","state_name":"Utah","latitude_average":"39.32098","longitude_average":"-111.093731"},
{"state_prov_cd":"VT","state_name":"Vermont","latitude_average":"44.558803","longitude_average":"-72.577841"},
{"state_prov_cd":"VA","state_name":"Virginia","latitude_average":"37.431573","longitude_average":"-78.656894"},
{"state_prov_cd":"WA","state_name":"Washington","latitude_average":"47.751074","longitude_average":"-120.740139"},
{"state_prov_cd":"WV","state_name":"West Virginia","latitude_average":"38.597626","longitude_average":"-80.454903"},
{"state_prov_cd":"WI","state_name":"Wisconsin","latitude_average":"43.78444","longitude_average":"-88.787868"},
{"state_prov_cd":"WY","state_name":"Wyoming","latitude_average":"43.075968","longitude_average":"-107.290284"}]`

// USStatesNameJSON is the JSON array for all US State Names
// (including DC and Puerto Rico) in alpha order by name
const USStatesNameJSON string = `["Alabama","Alaska","Arizona","Arkansas","California","Colorado","Connecticut","Delaware","District of Columbia","Florida","Georgia","Hawaii","Idaho","Illinois","Indiana","Iowa","Kansas","Kentucky","Louisiana","Maine","Maryland","Massachusetts","Michigan","Minnesota","Mississippi","Missouri","Montana","Nebraska","Nevada","New Hampshire","New Jersey","New Mexico","New York","North Carolina","North Dakota","Ohio","Oklahoma","Oregon","Pennsylvania","Puerto Rico","Rhode Island","South Carolina","South Dakota","Tennessee","Texas","Utah","Vermont","Virginia","Washington","West Virginia","Wisconsin","Wyoming"]`

// StateProvince is a State or Province of a Country
// that makes up a particular territory
type StateProvince struct {
	Code             string `json:"state_prov_cd"`
	Name             string `json:"state_name"`
	LatitudeAverage  string `json:"latitude_average"`
	LongitudeAverage string `json:"longitude_average"`
}
