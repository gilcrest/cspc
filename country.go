package cspc

import (
	"encoding/json"

	"github.com/gilcrest/errs"
)

// CountryFullJSON string is a JSON string with all country information
// in JSON format, ordered by country_name
const CountryFullJSON string = `[{"country_alpha_2_cd":"AF","country_alpha_3_cd":"AFG","country_un_m49_cd":4,"country_name":"Afghanistan","latitude_average":"33","longitude_average":"65"}, 
 {"country_alpha_2_cd":"AL","country_alpha_3_cd":"ALB","country_un_m49_cd":8,"country_name":"Albania","latitude_average":"41","longitude_average":"20"}, 
 {"country_alpha_2_cd":"DZ","country_alpha_3_cd":"DZA","country_un_m49_cd":12,"country_name":"Algeria","latitude_average":"28","longitude_average":"3"}, 
 {"country_alpha_2_cd":"AS","country_alpha_3_cd":"ASM","country_un_m49_cd":16,"country_name":"American Samoa","latitude_average":"-14.3333","longitude_average":"-170"}, 
 {"country_alpha_2_cd":"AD","country_alpha_3_cd":"AND","country_un_m49_cd":20,"country_name":"Andorra","latitude_average":"42.5","longitude_average":"1.6"}, 
 {"country_alpha_2_cd":"AO","country_alpha_3_cd":"AGO","country_un_m49_cd":24,"country_name":"Angola","latitude_average":"-12.5","longitude_average":"18.5"}, 
 {"country_alpha_2_cd":"AI","country_alpha_3_cd":"AIA","country_un_m49_cd":660,"country_name":"Anguilla","latitude_average":"18.25","longitude_average":"-63.1667"}, 
 {"country_alpha_2_cd":"AQ","country_alpha_3_cd":"ATA","country_un_m49_cd":10,"country_name":"Antarctica","latitude_average":"-90","longitude_average":"0"}, 
 {"country_alpha_2_cd":"AG","country_alpha_3_cd":"ATG","country_un_m49_cd":28,"country_name":"Antigua and Barbuda","latitude_average":"17.05","longitude_average":"-61.8"}, 
 {"country_alpha_2_cd":"AR","country_alpha_3_cd":"ARG","country_un_m49_cd":32,"country_name":"Argentina","latitude_average":"-34","longitude_average":"-64"}, 
 {"country_alpha_2_cd":"AM","country_alpha_3_cd":"ARM","country_un_m49_cd":51,"country_name":"Armenia","latitude_average":"40","longitude_average":"45"}, 
 {"country_alpha_2_cd":"AW","country_alpha_3_cd":"ABW","country_un_m49_cd":533,"country_name":"Aruba","latitude_average":"12.5","longitude_average":"-69.9667"}, 
 {"country_alpha_2_cd":"AU","country_alpha_3_cd":"AUS","country_un_m49_cd":36,"country_name":"Australia","latitude_average":"-27","longitude_average":"133"}, 
 {"country_alpha_2_cd":"AT","country_alpha_3_cd":"AUT","country_un_m49_cd":40,"country_name":"Austria","latitude_average":"47.3333","longitude_average":"13.3333"}, 
 {"country_alpha_2_cd":"AZ","country_alpha_3_cd":"AZE","country_un_m49_cd":31,"country_name":"Azerbaijan","latitude_average":"40.5","longitude_average":"47.5"}, 
 {"country_alpha_2_cd":"BS","country_alpha_3_cd":"BHS","country_un_m49_cd":44,"country_name":"Bahamas","latitude_average":"24.25","longitude_average":"-76"}, 
 {"country_alpha_2_cd":"BH","country_alpha_3_cd":"BHR","country_un_m49_cd":48,"country_name":"Bahrain","latitude_average":"26","longitude_average":"50.55"}, 
 {"country_alpha_2_cd":"BD","country_alpha_3_cd":"BGD","country_un_m49_cd":50,"country_name":"Bangladesh","latitude_average":"24","longitude_average":"90"}, 
 {"country_alpha_2_cd":"BB","country_alpha_3_cd":"BRB","country_un_m49_cd":52,"country_name":"Barbados","latitude_average":"13.1667","longitude_average":"-59.5333"}, 
 {"country_alpha_2_cd":"BY","country_alpha_3_cd":"BLR","country_un_m49_cd":112,"country_name":"Belarus","latitude_average":"53","longitude_average":"28"}, 
 {"country_alpha_2_cd":"BE","country_alpha_3_cd":"BEL","country_un_m49_cd":56,"country_name":"Belgium","latitude_average":"50.8333","longitude_average":"4"}, 
 {"country_alpha_2_cd":"BZ","country_alpha_3_cd":"BLZ","country_un_m49_cd":84,"country_name":"Belize","latitude_average":"17.25","longitude_average":"-88.75"}, 
 {"country_alpha_2_cd":"BJ","country_alpha_3_cd":"BEN","country_un_m49_cd":204,"country_name":"Benin","latitude_average":"9.5","longitude_average":"2.25"}, 
 {"country_alpha_2_cd":"BM","country_alpha_3_cd":"BMU","country_un_m49_cd":60,"country_name":"Bermuda","latitude_average":"32.3333","longitude_average":"-64.75"}, 
 {"country_alpha_2_cd":"BT","country_alpha_3_cd":"BTN","country_un_m49_cd":64,"country_name":"Bhutan","latitude_average":"27.5","longitude_average":"90.5"}, 
 {"country_alpha_2_cd":"BO","country_alpha_3_cd":"BOL","country_un_m49_cd":68,"country_name":"Bolivia","latitude_average":"-17","longitude_average":"-65"}, 
 {"country_alpha_2_cd":"BA","country_alpha_3_cd":"BIH","country_un_m49_cd":70,"country_name":"Bosnia and Herzegovina","latitude_average":"44","longitude_average":"18"}, 
 {"country_alpha_2_cd":"BW","country_alpha_3_cd":"BWA","country_un_m49_cd":72,"country_name":"Botswana","latitude_average":"-22","longitude_average":"24"}, 
 {"country_alpha_2_cd":"BV","country_alpha_3_cd":"BVT","country_un_m49_cd":74,"country_name":"Bouvet Island","latitude_average":"-54.4333","longitude_average":"3.4"}, 
 {"country_alpha_2_cd":"BR","country_alpha_3_cd":"BRA","country_un_m49_cd":76,"country_name":"Brazil","latitude_average":"-10","longitude_average":"-55"}, 
 {"country_alpha_2_cd":"IO","country_alpha_3_cd":"IOT","country_un_m49_cd":86,"country_name":"British Indian Ocean Territory","latitude_average":"-6","longitude_average":"71.5"}, 
 {"country_alpha_2_cd":"BN","country_alpha_3_cd":"BRN","country_un_m49_cd":96,"country_name":"Brunei","latitude_average":"4.5","longitude_average":"114.6667"}, 
 {"country_alpha_2_cd":"BG","country_alpha_3_cd":"BGR","country_un_m49_cd":100,"country_name":"Bulgaria","latitude_average":"43","longitude_average":"25"}, 
 {"country_alpha_2_cd":"BF","country_alpha_3_cd":"BFA","country_un_m49_cd":854,"country_name":"Burkina Faso","latitude_average":"13","longitude_average":"-2"}, 
 {"country_alpha_2_cd":"BI","country_alpha_3_cd":"BDI","country_un_m49_cd":108,"country_name":"Burundi","latitude_average":"-3.5","longitude_average":"30"}, 
 {"country_alpha_2_cd":"KH","country_alpha_3_cd":"KHM","country_un_m49_cd":116,"country_name":"Cambodia","latitude_average":"13","longitude_average":"105"}, 
 {"country_alpha_2_cd":"CM","country_alpha_3_cd":"CMR","country_un_m49_cd":120,"country_name":"Cameroon","latitude_average":"6","longitude_average":"12"}, 
 {"country_alpha_2_cd":"CA","country_alpha_3_cd":"CAN","country_un_m49_cd":124,"country_name":"Canada","latitude_average":"60","longitude_average":"-95"}, 
 {"country_alpha_2_cd":"CV","country_alpha_3_cd":"CPV","country_un_m49_cd":132,"country_name":"Cape Verde","latitude_average":"16","longitude_average":"-24"}, 
 {"country_alpha_2_cd":"KY","country_alpha_3_cd":"CYM","country_un_m49_cd":136,"country_name":"Cayman Islands","latitude_average":"19.5","longitude_average":"-80.5"}, 
 {"country_alpha_2_cd":"CF","country_alpha_3_cd":"CAF","country_un_m49_cd":140,"country_name":"Central African Republic","latitude_average":"7","longitude_average":"21"}, 
 {"country_alpha_2_cd":"TD","country_alpha_3_cd":"TCD","country_un_m49_cd":148,"country_name":"Chad","latitude_average":"15","longitude_average":"19"}, 
 {"country_alpha_2_cd":"CL","country_alpha_3_cd":"CHL","country_un_m49_cd":152,"country_name":"Chile","latitude_average":"-30","longitude_average":"-71"}, 
 {"country_alpha_2_cd":"CN","country_alpha_3_cd":"CHN","country_un_m49_cd":156,"country_name":"China","latitude_average":"35","longitude_average":"105"}, 
 {"country_alpha_2_cd":"CX","country_alpha_3_cd":"CXR","country_un_m49_cd":162,"country_name":"Christmas Island","latitude_average":"-10.5","longitude_average":"105.6667"}, 
 {"country_alpha_2_cd":"CC","country_alpha_3_cd":"CCK","country_un_m49_cd":166,"country_name":"Cocos (Keeling) Islands","latitude_average":"-12.5","longitude_average":"96.8333"}, 
 {"country_alpha_2_cd":"CO","country_alpha_3_cd":"COL","country_un_m49_cd":170,"country_name":"Colombia","latitude_average":"4","longitude_average":"-72"}, 
 {"country_alpha_2_cd":"KM","country_alpha_3_cd":"COM","country_un_m49_cd":174,"country_name":"Comoros","latitude_average":"-12.1667","longitude_average":"44.25"}, 
 {"country_alpha_2_cd":"CG","country_alpha_3_cd":"COG","country_un_m49_cd":178,"country_name":"Congo","latitude_average":"-1","longitude_average":"15"}, 
 {"country_alpha_2_cd":"CD","country_alpha_3_cd":"COD","country_un_m49_cd":180,"country_name":"Congo, the Democratic Republic of the","latitude_average":"0","longitude_average":"25"}, 
 {"country_alpha_2_cd":"CK","country_alpha_3_cd":"COK","country_un_m49_cd":184,"country_name":"Cook Islands","latitude_average":"-21.2333","longitude_average":"-159.7667"}, 
 {"country_alpha_2_cd":"CR","country_alpha_3_cd":"CRI","country_un_m49_cd":188,"country_name":"Costa Rica","latitude_average":"10","longitude_average":"-84"}, 
 {"country_alpha_2_cd":"HR","country_alpha_3_cd":"HRV","country_un_m49_cd":191,"country_name":"Croatia","latitude_average":"45.1667","longitude_average":"15.5"}, 
 {"country_alpha_2_cd":"CU","country_alpha_3_cd":"CUB","country_un_m49_cd":192,"country_name":"Cuba","latitude_average":"21.5","longitude_average":"-80"}, 
 {"country_alpha_2_cd":"CY","country_alpha_3_cd":"CYP","country_un_m49_cd":196,"country_name":"Cyprus","latitude_average":"35","longitude_average":"33"}, 
 {"country_alpha_2_cd":"CZ","country_alpha_3_cd":"CZE","country_un_m49_cd":203,"country_name":"Czech Republic","latitude_average":"49.75","longitude_average":"15.5"}, 
 {"country_alpha_2_cd":"DK","country_alpha_3_cd":"DNK","country_un_m49_cd":208,"country_name":"Denmark","latitude_average":"56","longitude_average":"10"}, 
 {"country_alpha_2_cd":"DJ","country_alpha_3_cd":"DJI","country_un_m49_cd":262,"country_name":"Djibouti","latitude_average":"11.5","longitude_average":"43"}, 
 {"country_alpha_2_cd":"DM","country_alpha_3_cd":"DMA","country_un_m49_cd":212,"country_name":"Dominica","latitude_average":"15.4167","longitude_average":"-61.3333"}, 
 {"country_alpha_2_cd":"DO","country_alpha_3_cd":"DOM","country_un_m49_cd":214,"country_name":"Dominican Republic","latitude_average":"19","longitude_average":"-70.6667"}, 
 {"country_alpha_2_cd":"EC","country_alpha_3_cd":"ECU","country_un_m49_cd":218,"country_name":"Ecuador","latitude_average":"-2","longitude_average":"-77.5"}, 
 {"country_alpha_2_cd":"EG","country_alpha_3_cd":"EGY","country_un_m49_cd":818,"country_name":"Egypt","latitude_average":"27","longitude_average":"30"}, 
 {"country_alpha_2_cd":"SV","country_alpha_3_cd":"SLV","country_un_m49_cd":222,"country_name":"El Salvador","latitude_average":"13.8333","longitude_average":"-88.9167"}, 
 {"country_alpha_2_cd":"GQ","country_alpha_3_cd":"GNQ","country_un_m49_cd":226,"country_name":"Equatorial Guinea","latitude_average":"2","longitude_average":"10"}, 
 {"country_alpha_2_cd":"ER","country_alpha_3_cd":"ERI","country_un_m49_cd":232,"country_name":"Eritrea","latitude_average":"15","longitude_average":"39"}, 
 {"country_alpha_2_cd":"EE","country_alpha_3_cd":"EST","country_un_m49_cd":233,"country_name":"Estonia","latitude_average":"59","longitude_average":"26"}, 
 {"country_alpha_2_cd":"ET","country_alpha_3_cd":"ETH","country_un_m49_cd":231,"country_name":"Ethiopia","latitude_average":"8","longitude_average":"38"}, 
 {"country_alpha_2_cd":"FK","country_alpha_3_cd":"FLK","country_un_m49_cd":238,"country_name":"Falkland Islands (Malvinas)","latitude_average":"-51.75","longitude_average":"-59"}, 
 {"country_alpha_2_cd":"FO","country_alpha_3_cd":"FRO","country_un_m49_cd":234,"country_name":"Faroe Islands","latitude_average":"62","longitude_average":"-7"}, 
 {"country_alpha_2_cd":"FJ","country_alpha_3_cd":"FJI","country_un_m49_cd":242,"country_name":"Fiji","latitude_average":"-18","longitude_average":"175"}, 
 {"country_alpha_2_cd":"FI","country_alpha_3_cd":"FIN","country_un_m49_cd":246,"country_name":"Finland","latitude_average":"64","longitude_average":"26"}, 
 {"country_alpha_2_cd":"FR","country_alpha_3_cd":"FRA","country_un_m49_cd":250,"country_name":"France","latitude_average":"46","longitude_average":"2"}, 
 {"country_alpha_2_cd":"GF","country_alpha_3_cd":"GUF","country_un_m49_cd":254,"country_name":"French Guiana","latitude_average":"4","longitude_average":"-53"}, 
 {"country_alpha_2_cd":"PF","country_alpha_3_cd":"PYF","country_un_m49_cd":258,"country_name":"French Polynesia","latitude_average":"-15","longitude_average":"-140"}, 
 {"country_alpha_2_cd":"TF","country_alpha_3_cd":"ATF","country_un_m49_cd":260,"country_name":"French Southern Territories","latitude_average":"-43","longitude_average":"67"}, 
 {"country_alpha_2_cd":"GA","country_alpha_3_cd":"GAB","country_un_m49_cd":266,"country_name":"Gabon","latitude_average":"-1","longitude_average":"11.75"}, 
 {"country_alpha_2_cd":"GM","country_alpha_3_cd":"GMB","country_un_m49_cd":270,"country_name":"Gambia","latitude_average":"13.4667","longitude_average":"-16.5667"}, 
 {"country_alpha_2_cd":"GE","country_alpha_3_cd":"GEO","country_un_m49_cd":268,"country_name":"Georgia","latitude_average":"42","longitude_average":"43.5"}, 
 {"country_alpha_2_cd":"DE","country_alpha_3_cd":"DEU","country_un_m49_cd":276,"country_name":"Germany","latitude_average":"51","longitude_average":"9"}, 
 {"country_alpha_2_cd":"GH","country_alpha_3_cd":"GHA","country_un_m49_cd":288,"country_name":"Ghana","latitude_average":"8","longitude_average":"-2"}, 
 {"country_alpha_2_cd":"GI","country_alpha_3_cd":"GIB","country_un_m49_cd":292,"country_name":"Gibraltar","latitude_average":"36.1833","longitude_average":"-5.3667"}, 
 {"country_alpha_2_cd":"GR","country_alpha_3_cd":"GRC","country_un_m49_cd":300,"country_name":"Greece","latitude_average":"39","longitude_average":"22"}, 
 {"country_alpha_2_cd":"GL","country_alpha_3_cd":"GRL","country_un_m49_cd":304,"country_name":"Greenland","latitude_average":"72","longitude_average":"-40"}, 
 {"country_alpha_2_cd":"GD","country_alpha_3_cd":"GRD","country_un_m49_cd":308,"country_name":"Grenada","latitude_average":"12.1167","longitude_average":"-61.6667"}, 
 {"country_alpha_2_cd":"GP","country_alpha_3_cd":"GLP","country_un_m49_cd":312,"country_name":"Guadeloupe","latitude_average":"16.25","longitude_average":"-61.5833"}, 
 {"country_alpha_2_cd":"GU","country_alpha_3_cd":"GUM","country_un_m49_cd":316,"country_name":"Guam","latitude_average":"13.4667","longitude_average":"144.7833"}, 
 {"country_alpha_2_cd":"GT","country_alpha_3_cd":"GTM","country_un_m49_cd":320,"country_name":"Guatemala","latitude_average":"15.5","longitude_average":"-90.25"}, 
 {"country_alpha_2_cd":"GG","country_alpha_3_cd":"GGY","country_un_m49_cd":831,"country_name":"Guernsey","latitude_average":"49.5","longitude_average":"-2.56"}, 
 {"country_alpha_2_cd":"GN","country_alpha_3_cd":"GIN","country_un_m49_cd":324,"country_name":"Guinea","latitude_average":"11","longitude_average":"-10"}, 
 {"country_alpha_2_cd":"GW","country_alpha_3_cd":"GNB","country_un_m49_cd":624,"country_name":"Guinea-Bissau","latitude_average":"12","longitude_average":"-15"}, 
 {"country_alpha_2_cd":"GY","country_alpha_3_cd":"GUY","country_un_m49_cd":328,"country_name":"Guyana","latitude_average":"5","longitude_average":"-59"}, 
 {"country_alpha_2_cd":"HT","country_alpha_3_cd":"HTI","country_un_m49_cd":332,"country_name":"Haiti","latitude_average":"19","longitude_average":"-72.4167"}, 
 {"country_alpha_2_cd":"HM","country_alpha_3_cd":"HMD","country_un_m49_cd":334,"country_name":"Heard Island and McDonald Islands","latitude_average":"-53.1","longitude_average":"72.5167"}, 
 {"country_alpha_2_cd":"VA","country_alpha_3_cd":"VAT","country_un_m49_cd":336,"country_name":"Holy See (Vatican City State)","latitude_average":"41.9","longitude_average":"12.45"}, 
 {"country_alpha_2_cd":"HN","country_alpha_3_cd":"HND","country_un_m49_cd":340,"country_name":"Honduras","latitude_average":"15","longitude_average":"-86.5"}, 
 {"country_alpha_2_cd":"HK","country_alpha_3_cd":"HKG","country_un_m49_cd":344,"country_name":"Hong Kong","latitude_average":"22.25","longitude_average":"114.1667"}, 
 {"country_alpha_2_cd":"HU","country_alpha_3_cd":"HUN","country_un_m49_cd":348,"country_name":"Hungary","latitude_average":"47","longitude_average":"20"}, 
 {"country_alpha_2_cd":"IS","country_alpha_3_cd":"ISL","country_un_m49_cd":352,"country_name":"Iceland","latitude_average":"65","longitude_average":"-18"}, 
 {"country_alpha_2_cd":"IN","country_alpha_3_cd":"IND","country_un_m49_cd":356,"country_name":"India","latitude_average":"20","longitude_average":"77"}, 
 {"country_alpha_2_cd":"ID","country_alpha_3_cd":"IDN","country_un_m49_cd":360,"country_name":"Indonesia","latitude_average":"-5","longitude_average":"120"}, 
 {"country_alpha_2_cd":"IR","country_alpha_3_cd":"IRN","country_un_m49_cd":364,"country_name":"Iran, Islamic Republic of","latitude_average":"32","longitude_average":"53"}, 
 {"country_alpha_2_cd":"IQ","country_alpha_3_cd":"IRQ","country_un_m49_cd":368,"country_name":"Iraq","latitude_average":"33","longitude_average":"44"}, 
 {"country_alpha_2_cd":"IE","country_alpha_3_cd":"IRL","country_un_m49_cd":372,"country_name":"Ireland","latitude_average":"53","longitude_average":"-8"}, 
 {"country_alpha_2_cd":"IM","country_alpha_3_cd":"IMN","country_un_m49_cd":833,"country_name":"Isle of Man","latitude_average":"54.23","longitude_average":"-4.55"}, 
 {"country_alpha_2_cd":"IL","country_alpha_3_cd":"ISR","country_un_m49_cd":376,"country_name":"Israel","latitude_average":"31.5","longitude_average":"34.75"}, 
 {"country_alpha_2_cd":"IT","country_alpha_3_cd":"ITA","country_un_m49_cd":380,"country_name":"Italy","latitude_average":"42.8333","longitude_average":"12.8333"}, 
 {"country_alpha_2_cd":"CI","country_alpha_3_cd":"CIV","country_un_m49_cd":384,"country_name":"Ivory Coast","latitude_average":"8","longitude_average":"-5"}, 
 {"country_alpha_2_cd":"JM","country_alpha_3_cd":"JAM","country_un_m49_cd":388,"country_name":"Jamaica","latitude_average":"18.25","longitude_average":"-77.5"}, 
 {"country_alpha_2_cd":"JP","country_alpha_3_cd":"JPN","country_un_m49_cd":392,"country_name":"Japan","latitude_average":"36","longitude_average":"138"}, 
 {"country_alpha_2_cd":"JE","country_alpha_3_cd":"JEY","country_un_m49_cd":832,"country_name":"Jersey","latitude_average":"49.21","longitude_average":"-2.13"}, 
 {"country_alpha_2_cd":"JO","country_alpha_3_cd":"JOR","country_un_m49_cd":400,"country_name":"Jordan","latitude_average":"31","longitude_average":"36"}, 
 {"country_alpha_2_cd":"KZ","country_alpha_3_cd":"KAZ","country_un_m49_cd":398,"country_name":"Kazakhstan","latitude_average":"48","longitude_average":"68"}, 
 {"country_alpha_2_cd":"KE","country_alpha_3_cd":"KEN","country_un_m49_cd":404,"country_name":"Kenya","latitude_average":"1","longitude_average":"38"}, 
 {"country_alpha_2_cd":"KI","country_alpha_3_cd":"KIR","country_un_m49_cd":296,"country_name":"Kiribati","latitude_average":"1.4167","longitude_average":"173"}, 
 {"country_alpha_2_cd":"KP","country_alpha_3_cd":"PRK","country_un_m49_cd":408,"country_name":"Korea, Democratic People's Republic of","latitude_average":"40","longitude_average":"127"}, 
 {"country_alpha_2_cd":"KR","country_alpha_3_cd":"KOR","country_un_m49_cd":410,"country_name":"Korea, Republic of","latitude_average":"37","longitude_average":"127.5"}, 
 {"country_alpha_2_cd":"KW","country_alpha_3_cd":"KWT","country_un_m49_cd":414,"country_name":"Kuwait","latitude_average":"29.3375","longitude_average":"47.6581"}, 
 {"country_alpha_2_cd":"KG","country_alpha_3_cd":"KGZ","country_un_m49_cd":417,"country_name":"Kyrgyzstan","latitude_average":"41","longitude_average":"75"}, 
 {"country_alpha_2_cd":"LA","country_alpha_3_cd":"LAO","country_un_m49_cd":418,"country_name":"Lao People's Democratic Republic","latitude_average":"18","longitude_average":"105"}, 
 {"country_alpha_2_cd":"LV","country_alpha_3_cd":"LVA","country_un_m49_cd":428,"country_name":"Latvia","latitude_average":"57","longitude_average":"25"}, 
 {"country_alpha_2_cd":"LB","country_alpha_3_cd":"LBN","country_un_m49_cd":422,"country_name":"Lebanon","latitude_average":"33.8333","longitude_average":"35.8333"}, 
 {"country_alpha_2_cd":"LS","country_alpha_3_cd":"LSO","country_un_m49_cd":426,"country_name":"Lesotho","latitude_average":"-29.5","longitude_average":"28.5"}, 
 {"country_alpha_2_cd":"LR","country_alpha_3_cd":"LBR","country_un_m49_cd":430,"country_name":"Liberia","latitude_average":"6.5","longitude_average":"-9.5"}, 
 {"country_alpha_2_cd":"LY","country_alpha_3_cd":"LBY","country_un_m49_cd":434,"country_name":"Libya","latitude_average":"25","longitude_average":"17"}, 
 {"country_alpha_2_cd":"LI","country_alpha_3_cd":"LIE","country_un_m49_cd":438,"country_name":"Liechtenstein","latitude_average":"47.1667","longitude_average":"9.5333"}, 
 {"country_alpha_2_cd":"LT","country_alpha_3_cd":"LTU","country_un_m49_cd":440,"country_name":"Lithuania","latitude_average":"56","longitude_average":"24"}, 
 {"country_alpha_2_cd":"LU","country_alpha_3_cd":"LUX","country_un_m49_cd":442,"country_name":"Luxembourg","latitude_average":"49.75","longitude_average":"6.1667"}, 
 {"country_alpha_2_cd":"MO","country_alpha_3_cd":"MAC","country_un_m49_cd":446,"country_name":"Macao","latitude_average":"22.1667","longitude_average":"113.55"}, 
 {"country_alpha_2_cd":"MK","country_alpha_3_cd":"MKD","country_un_m49_cd":807,"country_name":"Macedonia, the former Yugoslav Republic of","latitude_average":"41.8333","longitude_average":"22"}, 
 {"country_alpha_2_cd":"MG","country_alpha_3_cd":"MDG","country_un_m49_cd":450,"country_name":"Madagascar","latitude_average":"-20","longitude_average":"47"}, 
 {"country_alpha_2_cd":"MW","country_alpha_3_cd":"MWI","country_un_m49_cd":454,"country_name":"Malawi","latitude_average":"-13.5","longitude_average":"34"}, 
 {"country_alpha_2_cd":"MY","country_alpha_3_cd":"MYS","country_un_m49_cd":458,"country_name":"Malaysia","latitude_average":"2.5","longitude_average":"112.5"}, 
 {"country_alpha_2_cd":"MV","country_alpha_3_cd":"MDV","country_un_m49_cd":462,"country_name":"Maldives","latitude_average":"3.25","longitude_average":"73"}, 
 {"country_alpha_2_cd":"ML","country_alpha_3_cd":"MLI","country_un_m49_cd":466,"country_name":"Mali","latitude_average":"17","longitude_average":"-4"}, 
 {"country_alpha_2_cd":"MT","country_alpha_3_cd":"MLT","country_un_m49_cd":470,"country_name":"Malta","latitude_average":"35.8333","longitude_average":"14.5833"}, 
 {"country_alpha_2_cd":"MH","country_alpha_3_cd":"MHL","country_un_m49_cd":584,"country_name":"Marshall Islands","latitude_average":"9","longitude_average":"168"}, 
 {"country_alpha_2_cd":"MQ","country_alpha_3_cd":"MTQ","country_un_m49_cd":474,"country_name":"Martinique","latitude_average":"14.6667","longitude_average":"-61"}, 
 {"country_alpha_2_cd":"MR","country_alpha_3_cd":"MRT","country_un_m49_cd":478,"country_name":"Mauritania","latitude_average":"20","longitude_average":"-12"}, 
 {"country_alpha_2_cd":"MU","country_alpha_3_cd":"MUS","country_un_m49_cd":480,"country_name":"Mauritius","latitude_average":"-20.2833","longitude_average":"57.55"}, 
 {"country_alpha_2_cd":"YT","country_alpha_3_cd":"MYT","country_un_m49_cd":175,"country_name":"Mayotte","latitude_average":"-12.8333","longitude_average":"45.1667"}, 
 {"country_alpha_2_cd":"MX","country_alpha_3_cd":"MEX","country_un_m49_cd":484,"country_name":"Mexico","latitude_average":"23","longitude_average":"-102"}, 
 {"country_alpha_2_cd":"FM","country_alpha_3_cd":"FSM","country_un_m49_cd":583,"country_name":"Micronesia, Federated States of","latitude_average":"6.9167","longitude_average":"158.25"}, 
 {"country_alpha_2_cd":"MD","country_alpha_3_cd":"MDA","country_un_m49_cd":498,"country_name":"Moldova, Republic of","latitude_average":"47","longitude_average":"29"}, 
 {"country_alpha_2_cd":"MC","country_alpha_3_cd":"MCO","country_un_m49_cd":492,"country_name":"Monaco","latitude_average":"43.7333","longitude_average":"7.4"}, 
 {"country_alpha_2_cd":"MN","country_alpha_3_cd":"MNG","country_un_m49_cd":496,"country_name":"Mongolia","latitude_average":"46","longitude_average":"105"}, 
 {"country_alpha_2_cd":"ME","country_alpha_3_cd":"MNE","country_un_m49_cd":499,"country_name":"Montenegro","latitude_average":"42","longitude_average":"19"}, 
 {"country_alpha_2_cd":"MS","country_alpha_3_cd":"MSR","country_un_m49_cd":500,"country_name":"Montserrat","latitude_average":"16.75","longitude_average":"-62.2"}, 
 {"country_alpha_2_cd":"MA","country_alpha_3_cd":"MAR","country_un_m49_cd":504,"country_name":"Morocco","latitude_average":"32","longitude_average":"-5"}, 
 {"country_alpha_2_cd":"MZ","country_alpha_3_cd":"MOZ","country_un_m49_cd":508,"country_name":"Mozambique","latitude_average":"-18.25","longitude_average":"35"}, 
 {"country_alpha_2_cd":"MM","country_alpha_3_cd":"MMR","country_un_m49_cd":104,"country_name":"Myanmar","latitude_average":"22","longitude_average":"98"}, 
 {"country_alpha_2_cd":"NA","country_alpha_3_cd":"NAM","country_un_m49_cd":516,"country_name":"Namibia","latitude_average":"-22","longitude_average":"17"}, 
 {"country_alpha_2_cd":"NR","country_alpha_3_cd":"NRU","country_un_m49_cd":520,"country_name":"Nauru","latitude_average":"-0.5333","longitude_average":"166.9167"}, 
 {"country_alpha_2_cd":"NP","country_alpha_3_cd":"NPL","country_un_m49_cd":524,"country_name":"Nepal","latitude_average":"28","longitude_average":"84"}, 
 {"country_alpha_2_cd":"NL","country_alpha_3_cd":"NLD","country_un_m49_cd":528,"country_name":"Netherlands","latitude_average":"52.5","longitude_average":"5.75"}, 
 {"country_alpha_2_cd":"AN","country_alpha_3_cd":"ANT","country_un_m49_cd":530,"country_name":"Netherlands Antilles","latitude_average":"12.25","longitude_average":"-68.75"}, 
 {"country_alpha_2_cd":"NC","country_alpha_3_cd":"NCL","country_un_m49_cd":540,"country_name":"New Caledonia","latitude_average":"-21.5","longitude_average":"165.5"}, 
 {"country_alpha_2_cd":"NZ","country_alpha_3_cd":"NZL","country_un_m49_cd":554,"country_name":"New Zealand","latitude_average":"-41","longitude_average":"174"}, 
 {"country_alpha_2_cd":"NI","country_alpha_3_cd":"NIC","country_un_m49_cd":558,"country_name":"Nicaragua","latitude_average":"13","longitude_average":"-85"}, 
 {"country_alpha_2_cd":"NE","country_alpha_3_cd":"NER","country_un_m49_cd":562,"country_name":"Niger","latitude_average":"16","longitude_average":"8"}, 
 {"country_alpha_2_cd":"NG","country_alpha_3_cd":"NGA","country_un_m49_cd":566,"country_name":"Nigeria","latitude_average":"10","longitude_average":"8"}, 
 {"country_alpha_2_cd":"NU","country_alpha_3_cd":"NIU","country_un_m49_cd":570,"country_name":"Niue","latitude_average":"-19.0333","longitude_average":"-169.8667"}, 
 {"country_alpha_2_cd":"NF","country_alpha_3_cd":"NFK","country_un_m49_cd":574,"country_name":"Norfolk Island","latitude_average":"-29.0333","longitude_average":"167.95"}, 
 {"country_alpha_2_cd":"MP","country_alpha_3_cd":"MNP","country_un_m49_cd":580,"country_name":"Northern Mariana Islands","latitude_average":"15.2","longitude_average":"145.75"}, 
 {"country_alpha_2_cd":"NO","country_alpha_3_cd":"NOR","country_un_m49_cd":578,"country_name":"Norway","latitude_average":"62","longitude_average":"10"}, 
 {"country_alpha_2_cd":"OM","country_alpha_3_cd":"OMN","country_un_m49_cd":512,"country_name":"Oman","latitude_average":"21","longitude_average":"57"}, 
 {"country_alpha_2_cd":"PK","country_alpha_3_cd":"PAK","country_un_m49_cd":586,"country_name":"Pakistan","latitude_average":"30","longitude_average":"70"}, 
 {"country_alpha_2_cd":"PW","country_alpha_3_cd":"PLW","country_un_m49_cd":585,"country_name":"Palau","latitude_average":"7.5","longitude_average":"134.5"}, 
 {"country_alpha_2_cd":"PS","country_alpha_3_cd":"PSE","country_un_m49_cd":275,"country_name":"Palestinian Territory, Occupied","latitude_average":"32","longitude_average":"35.25"}, 
 {"country_alpha_2_cd":"PA","country_alpha_3_cd":"PAN","country_un_m49_cd":591,"country_name":"Panama","latitude_average":"9","longitude_average":"-80"}, 
 {"country_alpha_2_cd":"PG","country_alpha_3_cd":"PNG","country_un_m49_cd":598,"country_name":"Papua New Guinea","latitude_average":"-6","longitude_average":"147"}, 
 {"country_alpha_2_cd":"PY","country_alpha_3_cd":"PRY","country_un_m49_cd":600,"country_name":"Paraguay","latitude_average":"-23","longitude_average":"-58"}, 
 {"country_alpha_2_cd":"PE","country_alpha_3_cd":"PER","country_un_m49_cd":604,"country_name":"Peru","latitude_average":"-10","longitude_average":"-76"}, 
 {"country_alpha_2_cd":"PH","country_alpha_3_cd":"PHL","country_un_m49_cd":608,"country_name":"Philippines","latitude_average":"13","longitude_average":"122"}, 
 {"country_alpha_2_cd":"PN","country_alpha_3_cd":"PCN","country_un_m49_cd":612,"country_name":"Pitcairn","latitude_average":"-24.7","longitude_average":"-127.4"}, 
 {"country_alpha_2_cd":"PL","country_alpha_3_cd":"POL","country_un_m49_cd":616,"country_name":"Poland","latitude_average":"52","longitude_average":"20"}, 
 {"country_alpha_2_cd":"PT","country_alpha_3_cd":"PRT","country_un_m49_cd":620,"country_name":"Portugal","latitude_average":"39.5","longitude_average":"-8"}, 
 {"country_alpha_2_cd":"PR","country_alpha_3_cd":"PRI","country_un_m49_cd":630,"country_name":"Puerto Rico","latitude_average":"18.25","longitude_average":"-66.5"}, 
 {"country_alpha_2_cd":"QA","country_alpha_3_cd":"QAT","country_un_m49_cd":634,"country_name":"Qatar","latitude_average":"25.5","longitude_average":"51.25"}, 
 {"country_alpha_2_cd":"RE","country_alpha_3_cd":"REU","country_un_m49_cd":638,"country_name":"Reunion","latitude_average":"-21.1","longitude_average":"55.6"}, 
 {"country_alpha_2_cd":"RO","country_alpha_3_cd":"ROU","country_un_m49_cd":642,"country_name":"Romania","latitude_average":"46","longitude_average":"25"}, 
 {"country_alpha_2_cd":"RU","country_alpha_3_cd":"RUS","country_un_m49_cd":643,"country_name":"Russia","latitude_average":"60","longitude_average":"100"}, 
 {"country_alpha_2_cd":"RW","country_alpha_3_cd":"RWA","country_un_m49_cd":646,"country_name":"Rwanda","latitude_average":"-2","longitude_average":"30"}, 
 {"country_alpha_2_cd":"SH","country_alpha_3_cd":"SHN","country_un_m49_cd":654,"country_name":"Saint Helena, Ascension and Tristan da Cunha","latitude_average":"-15.9333","longitude_average":"-5.7"}, 
 {"country_alpha_2_cd":"KN","country_alpha_3_cd":"KNA","country_un_m49_cd":659,"country_name":"Saint Kitts and Nevis","latitude_average":"17.3333","longitude_average":"-62.75"}, 
 {"country_alpha_2_cd":"LC","country_alpha_3_cd":"LCA","country_un_m49_cd":662,"country_name":"Saint Lucia","latitude_average":"13.8833","longitude_average":"-61.1333"}, 
 {"country_alpha_2_cd":"PM","country_alpha_3_cd":"SPM","country_un_m49_cd":666,"country_name":"Saint Pierre and Miquelon","latitude_average":"46.8333","longitude_average":"-56.3333"}, 
 {"country_alpha_2_cd":"WS","country_alpha_3_cd":"WSM","country_un_m49_cd":882,"country_name":"Samoa","latitude_average":"-13.5833","longitude_average":"-172.3333"}, 
 {"country_alpha_2_cd":"SM","country_alpha_3_cd":"SMR","country_un_m49_cd":674,"country_name":"San Marino","latitude_average":"43.7667","longitude_average":"12.4167"}, 
 {"country_alpha_2_cd":"ST","country_alpha_3_cd":"STP","country_un_m49_cd":678,"country_name":"Sao Tome and Principe","latitude_average":"1","longitude_average":"7"}, 
 {"country_alpha_2_cd":"SA","country_alpha_3_cd":"SAU","country_un_m49_cd":682,"country_name":"Saudi Arabia","latitude_average":"25","longitude_average":"45"}, 
 {"country_alpha_2_cd":"SN","country_alpha_3_cd":"SEN","country_un_m49_cd":686,"country_name":"Senegal","latitude_average":"14","longitude_average":"-14"}, 
 {"country_alpha_2_cd":"RS","country_alpha_3_cd":"SRB","country_un_m49_cd":688,"country_name":"Serbia","latitude_average":"44","longitude_average":"21"}, 
 {"country_alpha_2_cd":"SC","country_alpha_3_cd":"SYC","country_un_m49_cd":690,"country_name":"Seychelles","latitude_average":"-4.5833","longitude_average":"55.6667"}, 
 {"country_alpha_2_cd":"SL","country_alpha_3_cd":"SLE","country_un_m49_cd":694,"country_name":"Sierra Leone","latitude_average":"8.5","longitude_average":"-11.5"}, 
 {"country_alpha_2_cd":"SG","country_alpha_3_cd":"SGP","country_un_m49_cd":702,"country_name":"Singapore","latitude_average":"1.3667","longitude_average":"103.8"}, 
 {"country_alpha_2_cd":"SK","country_alpha_3_cd":"SVK","country_un_m49_cd":703,"country_name":"Slovakia","latitude_average":"48.6667","longitude_average":"19.5"}, 
 {"country_alpha_2_cd":"SI","country_alpha_3_cd":"SVN","country_un_m49_cd":705,"country_name":"Slovenia","latitude_average":"46","longitude_average":"15"}, 
 {"country_alpha_2_cd":"SB","country_alpha_3_cd":"SLB","country_un_m49_cd":90,"country_name":"Solomon Islands","latitude_average":"-8","longitude_average":"159"}, 
 {"country_alpha_2_cd":"SO","country_alpha_3_cd":"SOM","country_un_m49_cd":706,"country_name":"Somalia","latitude_average":"10","longitude_average":"49"}, 
 {"country_alpha_2_cd":"ZA","country_alpha_3_cd":"ZAF","country_un_m49_cd":710,"country_name":"South Africa","latitude_average":"-29","longitude_average":"24"}, 
 {"country_alpha_2_cd":"GS","country_alpha_3_cd":"SGS","country_un_m49_cd":239,"country_name":"South Georgia and the South Sandwich Islands","latitude_average":"-54.5","longitude_average":"-37"}, 
 {"country_alpha_2_cd":"ES","country_alpha_3_cd":"ESP","country_un_m49_cd":724,"country_name":"Spain","latitude_average":"40","longitude_average":"-4"}, 
 {"country_alpha_2_cd":"LK","country_alpha_3_cd":"LKA","country_un_m49_cd":144,"country_name":"Sri Lanka","latitude_average":"7","longitude_average":"81"}, 
 {"country_alpha_2_cd":"VC","country_alpha_3_cd":"VCT","country_un_m49_cd":670,"country_name":"St. Vincent and the Grenadines","latitude_average":"13.25","longitude_average":"-61.2"}, 
 {"country_alpha_2_cd":"SD","country_alpha_3_cd":"SDN","country_un_m49_cd":736,"country_name":"Sudan","latitude_average":"15","longitude_average":"30"}, 
 {"country_alpha_2_cd":"SR","country_alpha_3_cd":"SUR","country_un_m49_cd":740,"country_name":"Suriname","latitude_average":"4","longitude_average":"-56"}, 
 {"country_alpha_2_cd":"SJ","country_alpha_3_cd":"SJM","country_un_m49_cd":744,"country_name":"Svalbard and Jan Mayen","latitude_average":"78","longitude_average":"20"}, 
 {"country_alpha_2_cd":"SZ","country_alpha_3_cd":"SWZ","country_un_m49_cd":748,"country_name":"Swaziland","latitude_average":"-26.5","longitude_average":"31.5"}, 
 {"country_alpha_2_cd":"SE","country_alpha_3_cd":"SWE","country_un_m49_cd":752,"country_name":"Sweden","latitude_average":"62","longitude_average":"15"}, 
 {"country_alpha_2_cd":"CH","country_alpha_3_cd":"CHE","country_un_m49_cd":756,"country_name":"Switzerland","latitude_average":"47","longitude_average":"8"}, 
 {"country_alpha_2_cd":"SY","country_alpha_3_cd":"SYR","country_un_m49_cd":760,"country_name":"Syrian Arab Republic","latitude_average":"35","longitude_average":"38"}, 
 {"country_alpha_2_cd":"TW","country_alpha_3_cd":"TWN","country_un_m49_cd":158,"country_name":"Taiwan, Province of China","latitude_average":"23.5","longitude_average":"121"}, 
 {"country_alpha_2_cd":"TJ","country_alpha_3_cd":"TJK","country_un_m49_cd":762,"country_name":"Tajikistan","latitude_average":"39","longitude_average":"71"}, 
 {"country_alpha_2_cd":"TZ","country_alpha_3_cd":"TZA","country_un_m49_cd":834,"country_name":"Tanzania, United Republic of","latitude_average":"-6","longitude_average":"35"}, 
 {"country_alpha_2_cd":"TH","country_alpha_3_cd":"THA","country_un_m49_cd":764,"country_name":"Thailand","latitude_average":"15","longitude_average":"100"}, 
 {"country_alpha_2_cd":"TL","country_alpha_3_cd":"TLS","country_un_m49_cd":626,"country_name":"Timor-Leste","latitude_average":"-8.55","longitude_average":"125.5167"}, 
 {"country_alpha_2_cd":"TG","country_alpha_3_cd":"TGO","country_un_m49_cd":768,"country_name":"Togo","latitude_average":"8","longitude_average":"1.1667"}, 
 {"country_alpha_2_cd":"TK","country_alpha_3_cd":"TKL","country_un_m49_cd":772,"country_name":"Tokelau","latitude_average":"-9","longitude_average":"-172"}, 
 {"country_alpha_2_cd":"TO","country_alpha_3_cd":"TON","country_un_m49_cd":776,"country_name":"Tonga","latitude_average":"-20","longitude_average":"-175"}, 
 {"country_alpha_2_cd":"TT","country_alpha_3_cd":"TTO","country_un_m49_cd":780,"country_name":"Trinidad and Tobago","latitude_average":"11","longitude_average":"-61"}, 
 {"country_alpha_2_cd":"TN","country_alpha_3_cd":"TUN","country_un_m49_cd":788,"country_name":"Tunisia","latitude_average":"34","longitude_average":"9"}, 
 {"country_alpha_2_cd":"TR","country_alpha_3_cd":"TUR","country_un_m49_cd":792,"country_name":"Turkey","latitude_average":"39","longitude_average":"35"}, 
 {"country_alpha_2_cd":"TM","country_alpha_3_cd":"TKM","country_un_m49_cd":795,"country_name":"Turkmenistan","latitude_average":"40","longitude_average":"60"}, 
 {"country_alpha_2_cd":"TC","country_alpha_3_cd":"TCA","country_un_m49_cd":796,"country_name":"Turks and Caicos Islands","latitude_average":"21.75","longitude_average":"-71.5833"}, 
 {"country_alpha_2_cd":"TV","country_alpha_3_cd":"TUV","country_un_m49_cd":798,"country_name":"Tuvalu","latitude_average":"-8","longitude_average":"178"}, 
 {"country_alpha_2_cd":"UG","country_alpha_3_cd":"UGA","country_un_m49_cd":800,"country_name":"Uganda","latitude_average":"1","longitude_average":"32"}, 
 {"country_alpha_2_cd":"UA","country_alpha_3_cd":"UKR","country_un_m49_cd":804,"country_name":"Ukraine","latitude_average":"49","longitude_average":"32"}, 
 {"country_alpha_2_cd":"AE","country_alpha_3_cd":"ARE","country_un_m49_cd":784,"country_name":"United Arab Emirates","latitude_average":"24","longitude_average":"54"}, 
 {"country_alpha_2_cd":"GB","country_alpha_3_cd":"GBR","country_un_m49_cd":826,"country_name":"United Kingdom","latitude_average":"54","longitude_average":"-2"}, 
 {"country_alpha_2_cd":"US","country_alpha_3_cd":"USA","country_un_m49_cd":840,"country_name":"United States","latitude_average":"38","longitude_average":"-97"}, 
 {"country_alpha_2_cd":"UM","country_alpha_3_cd":"UMI","country_un_m49_cd":581,"country_name":"United States Minor Outlying Islands","latitude_average":"19.2833","longitude_average":"166.6"}, 
 {"country_alpha_2_cd":"UY","country_alpha_3_cd":"URY","country_un_m49_cd":858,"country_name":"Uruguay","latitude_average":"-33","longitude_average":"-56"}, 
 {"country_alpha_2_cd":"UZ","country_alpha_3_cd":"UZB","country_un_m49_cd":860,"country_name":"Uzbekistan","latitude_average":"41","longitude_average":"64"}, 
 {"country_alpha_2_cd":"VU","country_alpha_3_cd":"VUT","country_un_m49_cd":548,"country_name":"Vanuatu","latitude_average":"-16","longitude_average":"167"}, 
 {"country_alpha_2_cd":"VE","country_alpha_3_cd":"VEN","country_un_m49_cd":862,"country_name":"Venezuela","latitude_average":"8","longitude_average":"-66"}, 
 {"country_alpha_2_cd":"VN","country_alpha_3_cd":"VNM","country_un_m49_cd":704,"country_name":"Vietnam","latitude_average":"16","longitude_average":"106"}, 
 {"country_alpha_2_cd":"VG","country_alpha_3_cd":"VGB","country_un_m49_cd":92,"country_name":"Virgin Islands, British","latitude_average":"18.5","longitude_average":"-64.5"}, 
 {"country_alpha_2_cd":"VI","country_alpha_3_cd":"VIR","country_un_m49_cd":850,"country_name":"Virgin Islands, U.S.","latitude_average":"18.3333","longitude_average":"-64.8333"}, 
 {"country_alpha_2_cd":"WF","country_alpha_3_cd":"WLF","country_un_m49_cd":876,"country_name":"Wallis and Futuna","latitude_average":"-13.3","longitude_average":"-176.2"}, 
 {"country_alpha_2_cd":"EH","country_alpha_3_cd":"ESH","country_un_m49_cd":732,"country_name":"Western Sahara","latitude_average":"24.5","longitude_average":"-13"}, 
 {"country_alpha_2_cd":"YE","country_alpha_3_cd":"YEM","country_un_m49_cd":887,"country_name":"Yemen","latitude_average":"15","longitude_average":"48"}, 
 {"country_alpha_2_cd":"ZM","country_alpha_3_cd":"ZMB","country_un_m49_cd":894,"country_name":"Zambia","latitude_average":"-15","longitude_average":"30"}, 
 {"country_alpha_2_cd":"ZW","country_alpha_3_cd":"ZWE","country_un_m49_cd":716,"country_name":"Zimbabwe","latitude_average":"-20","longitude_average":"30"}]`

// CountryNamesJSON string is a JSON string with all country names
// in JSON format, in alphabetical order
const CountryNamesJSON string = `["Afghanistan","Albania","Algeria","American Samoa","Andorra","Angola","Anguilla","Antarctica","Antigua and Barbuda","Argentina","Armenia","Aruba","Australia","Austria","Azerbaijan","Bahamas","Bahrain","Bangladesh","Barbados","Belarus","Belgium","Belize","Benin","Bermuda","Bhutan","Bolivia","Bosnia and Herzegovina","Botswana","Bouvet Island","Brazil","British Indian Ocean Territory","Brunei","Bulgaria","Burkina Faso","Burundi","Cambodia","Cameroon","Canada","Cape Verde","Cayman Islands","Central African Republic","Chad","Chile","China","Christmas Island","Cocos (Keeling) Islands","Colombia","Comoros","Congo","Congo, the Democratic Republic of the","Cook Islands","Costa Rica","Croatia","Cuba","Cyprus","Czech Republic","Denmark","Djibouti","Dominica","Dominican Republic","Ecuador","Egypt","El Salvador","Equatorial Guinea","Eritrea","Estonia","Ethiopia","Falkland Islands (Malvinas)","Faroe Islands","Fiji","Finland","France","French Guiana","French Polynesia","French Southern Territories","Gabon","Gambia","Georgia","Germany","Ghana","Gibraltar","Greece","Greenland","Grenada","Guadeloupe","Guam","Guatemala","Guernsey","Guinea","Guinea-Bissau","Guyana","Haiti","Heard Island and McDonald Islands","Holy See (Vatican City State)","Honduras","Hong Kong","Hungary","Iceland","India","Indonesia","Iran, Islamic Republic of","Iraq","Ireland","Isle of Man","Israel","Italy","Ivory Coast","Jamaica","Japan","Jersey","Jordan","Kazakhstan","Kenya","Kiribati","Korea, Democratic People's Republic of","Korea, Republic of","Kuwait","Kyrgyzstan","Lao People's Democratic Republic","Latvia","Lebanon","Lesotho","Liberia","Libya","Liechtenstein","Lithuania","Luxembourg","Macao","Macedonia, the former Yugoslav Republic of","Madagascar","Malawi","Malaysia","Maldives","Mali","Malta","Marshall Islands","Martinique","Mauritania","Mauritius","Mayotte","Mexico","Micronesia, Federated States of","Moldova, Republic of","Monaco","Mongolia","Montenegro","Montserrat","Morocco","Mozambique","Myanmar","Namibia","Nauru","Nepal","Netherlands","Netherlands Antilles","New Caledonia","New Zealand","Nicaragua","Niger","Nigeria","Niue","Norfolk Island","Northern Mariana Islands","Norway","Oman","Pakistan","Palau","Palestinian Territory, Occupied","Panama","Papua New Guinea","Paraguay","Peru","Philippines","Pitcairn","Poland","Portugal","Puerto Rico","Qatar","Reunion","Romania","Russia","Rwanda","Saint Helena, Ascension and Tristan da Cunha","Saint Kitts and Nevis","Saint Lucia","Saint Pierre and Miquelon","Samoa","San Marino","Sao Tome and Principe","Saudi Arabia","Senegal","Serbia","Seychelles","Sierra Leone","Singapore","Slovakia","Slovenia","Solomon Islands","Somalia","South Africa","South Georgia and the South Sandwich Islands","Spain","Sri Lanka","St. Vincent and the Grenadines","Sudan","Suriname","Svalbard and Jan Mayen","Swaziland","Sweden","Switzerland","Syrian Arab Republic","Taiwan, Province of China","Tajikistan","Tanzania, United Republic of","Thailand","Timor-Leste","Togo","Tokelau","Tonga","Trinidad and Tobago","Tunisia","Turkey","Turkmenistan","Turks and Caicos Islands","Tuvalu","Uganda","Ukraine","United Arab Emirates","United Kingdom","United States","United States Minor Outlying Islands","Uruguay","Uzbekistan","Vanuatu","Venezuela","Vietnam","Virgin Islands, British","Virgin Islands, U.S.","Wallis and Futuna","Western Sahara","Yemen","Zambia","Zimbabwe"]`

// CountryAlpha2CodeJSON string is a JSON string with all country
// ISO 3166-1 Alpha 2 Code in JSON format, in alphabetical order
const CountryAlpha2CodeJSON string = `["AF","AL","DZ","AS","AD","AO","AI","AQ","AG","AR","AM","AW","AU","AT","AZ","BS","BH","BD","BB","BY","BE","BZ","BJ","BM","BT","BO","BA","BW","BV","BR","IO","BN","BG","BF","BI","KH","CM","CA","CV","KY","CF","TD","CL","CN","CX","CC","CO","KM","CG","CD","CK","CR","HR","CU","CY","CZ","DK","DJ","DM","DO","EC","EG","SV","GQ","ER","EE","ET","FK","FO","FJ","FI","FR","GF","PF","TF","GA","GM","GE","DE","GH","GI","GR","GL","GD","GP","GU","GT","GG","GN","GW","GY","HT","HM","VA","HN","HK","HU","IS","IN","ID","IR","IQ","IE","IM","IL","IT","CI","JM","JP","JE","JO","KZ","KE","KI","KP","KR","KW","KG","LA","LV","LB","LS","LR","LY","LI","LT","LU","MO","MK","MG","MW","MY","MV","ML","MT","MH","MQ","MR","MU","YT","MX","FM","MD","MC","MN","ME","MS","MA","MZ","MM","NA","NR","NP","NL","AN","NC","NZ","NI","NE","NG","NU","NF","MP","NO","OM","PK","PW","PS","PA","PG","PY","PE","PH","PN","PL","PT","PR","QA","RE","RO","RU","RW","SH","KN","LC","PM","WS","SM","ST","SA","SN","RS","SC","SL","SG","SK","SI","SB","SO","ZA","GS","ES","LK","VC","SD","SR","SJ","SZ","SE","CH","SY","TW","TJ","TZ","TH","TL","TG","TK","TO","TT","TN","TR","TM","TC","TV","UG","UA","AE","GB","US","UM","UY","UZ","VU","VE","VN","VG","VI","WF","EH","YE","ZM","ZW"]`

// Country represents a nation with its own government, occupying
// a particular territory.
type Country struct {
	Name             string `json:"country_name"`
	Alpha2Code       string `json:"country_alpha_2_cd"`
	Alpha3Code       string `json:"country_alpha_3_cd"`
	UNM49Code        int    `json:"country_un_m49_cd"`
	LatitudeAverage  string `json:"latitude_average"`
	LongitudeAverage string `json:"longitude_average"`
	States           []StateProvince
}

// Names returns a slice of strings containing the Country names
// in order by Country Name
func Names() []string {
	return []string{"Afghanistan", "Albania", "Algeria", "American Samoa", "Andorra", "Angola", "Anguilla", "Antarctica", "Antigua and Barbuda", "Argentina", "Armenia", "Aruba", "Australia", "Austria", "Azerbaijan", "Bahamas", "Bahrain", "Bangladesh", "Barbados", "Belarus", "Belgium", "Belize", "Benin", "Bermuda", "Bhutan", "Bolivia", "Bosnia and Herzegovina", "Botswana", "Bouvet Island", "Brazil", "British Indian Ocean Territory", "Brunei", "Bulgaria", "Burkina Faso", "Burundi", "Cambodia", "Cameroon", "Canada", "Cape Verde", "Cayman Islands", "Central African Republic", "Chad", "Chile", "China", "Christmas Island", "Cocos (Keeling) Islands", "Colombia", "Comoros", "Congo", "Congo, the Democratic Republic of the", "Cook Islands", "Costa Rica", "Croatia", "Cuba", "Cyprus", "Czech Republic", "Denmark", "Djibouti", "Dominica", "Dominican Republic", "Ecuador", "Egypt", "El Salvador", "Equatorial Guinea", "Eritrea", "Estonia", "Ethiopia", "Falkland Islands (Malvinas)", "Faroe Islands", "Fiji", "Finland", "France", "French Guiana", "French Polynesia", "French Southern Territories", "Gabon", "Gambia", "Georgia", "Germany", "Ghana", "Gibraltar", "Greece", "Greenland", "Grenada", "Guadeloupe", "Guam", "Guatemala", "Guernsey", "Guinea", "Guinea-Bissau", "Guyana", "Haiti", "Heard Island and McDonald Islands", "Holy See (Vatican City State)", "Honduras", "Hong Kong", "Hungary", "Iceland", "India", "Indonesia", "Iran, Islamic Republic of", "Iraq", "Ireland", "Isle of Man", "Israel", "Italy", "Ivory Coast", "Jamaica", "Japan", "Jersey", "Jordan", "Kazakhstan", "Kenya", "Kiribati", "Korea, Democratic People's Republic of", "Korea, Republic of", "Kuwait", "Kyrgyzstan", "Lao People's Democratic Republic", "Latvia", "Lebanon", "Lesotho", "Liberia", "Libya", "Liechtenstein", "Lithuania", "Luxembourg", "Macao", "Macedonia, the former Yugoslav Republic of", "Madagascar", "Malawi", "Malaysia", "Maldives", "Mali", "Malta", "Marshall Islands", "Martinique", "Mauritania", "Mauritius", "Mayotte", "Mexico", "Micronesia, Federated States of", "Moldova, Republic of", "Monaco", "Mongolia", "Montenegro", "Montserrat", "Morocco", "Mozambique", "Myanmar", "Namibia", "Nauru", "Nepal", "Netherlands", "Netherlands Antilles", "New Caledonia", "New Zealand", "Nicaragua", "Niger", "Nigeria", "Niue", "Norfolk Island", "Northern Mariana Islands", "Norway", "Oman", "Pakistan", "Palau", "Palestinian Territory, Occupied", "Panama", "Papua New Guinea", "Paraguay", "Peru", "Philippines", "Pitcairn", "Poland", "Portugal", "Puerto Rico", "Qatar", "Reunion", "Romania", "Russia", "Rwanda", "Saint Helena, Ascension and Tristan da Cunha", "Saint Kitts and Nevis", "Saint Lucia", "Saint Pierre and Miquelon", "Samoa", "San Marino", "Sao Tome and Principe", "Saudi Arabia", "Senegal", "Serbia", "Seychelles", "Sierra Leone", "Singapore", "Slovakia", "Slovenia", "Solomon Islands", "Somalia", "South Africa", "South Georgia and the South Sandwich Islands", "Spain", "Sri Lanka", "St. Vincent and the Grenadines", "Sudan", "Suriname", "Svalbard and Jan Mayen", "Swaziland", "Sweden", "Switzerland", "Syrian Arab Republic", "Taiwan, Province of China", "Tajikistan", "Tanzania, United Republic of", "Thailand", "Timor-Leste", "Togo", "Tokelau", "Tonga", "Trinidad and Tobago", "Tunisia", "Turkey", "Turkmenistan", "Turks and Caicos Islands", "Tuvalu", "Uganda", "Ukraine", "United Arab Emirates", "United Kingdom", "United States", "United States Minor Outlying Islands", "Uruguay", "Uzbekistan", "Vanuatu", "Venezuela", "Vietnam", "Virgin Islands, British", "Virgin Islands, U.S.", "Wallis and Futuna", "Western Sahara", "Yemen", "Zambia", "Zimbabwe"}
}

// Alpha2s returns a slice of strings containing the ISO 3166-1 Alpha 2 Codes
// in order by Country Name
func Alpha2s() []string {
	return []string{"AF", "AL", "DZ", "AS", "AD", "AO", "AI", "AQ", "AG", "AR", "AM", "AW", "AU", "AT", "AZ", "BS", "BH", "BD", "BB", "BY", "BE", "BZ", "BJ", "BM", "BT", "BO", "BA", "BW", "BV", "BR", "IO", "BN", "BG", "BF", "BI", "KH", "CM", "CA", "CV", "KY", "CF", "TD", "CL", "CN", "CX", "CC", "CO", "KM", "CG", "CD", "CK", "CR", "HR", "CU", "CY", "CZ", "DK", "DJ", "DM", "DO", "EC", "EG", "SV", "GQ", "ER", "EE", "ET", "FK", "FO", "FJ", "FI", "FR", "GF", "PF", "TF", "GA", "GM", "GE", "DE", "GH", "GI", "GR", "GL", "GD", "GP", "GU", "GT", "GG", "GN", "GW", "GY", "HT", "HM", "VA", "HN", "HK", "HU", "IS", "IN", "ID", "IR", "IQ", "IE", "IM", "IL", "IT", "CI", "JM", "JP", "JE", "JO", "KZ", "KE", "KI", "KP", "KR", "KW", "KG", "LA", "LV", "LB", "LS", "LR", "LY", "LI", "LT", "LU", "MO", "MK", "MG", "MW", "MY", "MV", "ML", "MT", "MH", "MQ", "MR", "MU", "YT", "MX", "FM", "MD", "MC", "MN", "ME", "MS", "MA", "MZ", "MM", "NA", "NR", "NP", "NL", "AN", "NC", "NZ", "NI", "NE", "NG", "NU", "NF", "MP", "NO", "OM", "PK", "PW", "PS", "PA", "PG", "PY", "PE", "PH", "PN", "PL", "PT", "PR", "QA", "RE", "RO", "RU", "RW", "SH", "KN", "LC", "PM", "WS", "SM", "ST", "SA", "SN", "RS", "SC", "SL", "SG", "SK", "SI", "SB", "SO", "ZA", "GS", "ES", "LK", "VC", "SD", "SR", "SJ", "SZ", "SE", "CH", "SY", "TW", "TJ", "TZ", "TH", "TL", "TG", "TK", "TO", "TT", "TN", "TR", "TM", "TC", "TV", "UG", "UA", "AE", "GB", "US", "UM", "UY", "UZ", "VU", "VE", "VN", "VG", "VI", "WF", "EH", "YE", "ZM", "ZW"}
}

// NewCountry is an initializer for a Country struct
// given a Country Code (cc). All states for the Country will also
// be initialized and added
func NewCountry(alpha2code string) (Country, error) {
	const op errs.Op = "cspc/NewCountry"

	var (
		countries []Country
	)
	err := json.Unmarshal([]byte(CountryFullJSON), &countries)
	if err != nil {
		return Country{}, errs.E(op, err)
	}

	for _, country := range countries {
		if alpha2code == country.Alpha2Code {
			if alpha2code == "US" {
				states, err := USStates()
				if err != nil {
					return Country{}, errs.E(op, err)
				}
				country.States = states
			}
			return country, nil
		}
	}

	return Country{}, errs.E(op, "Unknown Country Code")
}

// FindStateByCode returns a State/Province for a country given it's
// State or Province code
func (c Country) FindStateByCode(stateCode string) (StateProvince, error) {
	const op errs.Op = "cspc/Country.FindStatebyCode"

	for _, state := range c.States {
		if state.Code == stateCode {
			return state, nil
		}
	}
	return StateProvince{}, errs.E(op, "Unknown State Code")
}
