package geo

import "github.com/shopspring/decimal"

/*
{
    "status": "success",
    "country": "COUNTRY",
    "countryCode": "COUNTRY CODE",
    "region": "REGION CODE",
    "regionName": "REGION NAME",
    "city": "CITY",
    "zip": "ZIP CODE",
    "lat": LATITUDE,
    "lon": LONGITUDE,
    "timezone": "TIME ZONE",
    "isp": "ISP NAME",
    "org": "ORGANIZATION NAME",
    "as": "AS NUMBER / NAME",
    "query": "IP ADDRESS USED FOR QUERY"
}
*/

type IPAPIResponse struct {
	Status       string          `json:"status"`
	Country      string          `json:"country"`
	CountryCode  string          `json:"countryCode"`
	Region       string          `json:"region"`
	RegionName   string          `json:"regionName"`
	City         string          `json:"city"`
	Postal       string          `json:"zip"`
	Latitude     decimal.Decimal `json:"lat"`
	Longitude    decimal.Decimal `json:"lon"`
	Timezone     string          `json:"timezone"`
	ISP          string          `json:"isp"`
	Organization string          `json:"org"`
	ASN          string          `json:"as"`
	IP           string          `json:"query"`
	Message      string          `json:"message,omitempty"` // Only used for error messages
}
