package geo

import (
	"github.com/shopspring/decimal"
)

/*
{
   "status":"success",
   "description":"Data successfully received.",
   "data":{
      "geo":{
         "host":"www.google.com",
         "ip":"74.125.29.147",
         "rdns":"qg-in-f147.1e100.net",
         "asn":"AS15169",
         "isp":"Google Inc. ",
         "country_name":"United States",
         "country_code":"US",
         "region":"CA",
         "city":"Mountain View",
         "postal_code":"94043",
         "continent_code":"NA",
         "latitude":37.419200897217,
         "longitude":-122.05740356445,
         "dma_code":807,
         "area_code":650,
         "timezone":"America\/Los_Angeles",
         "datetime":"2015-05-22 09:10:35"
      }
   }
 }
*/
// KeyCDNResponse wraps the KeyCDN response json
type KeyCDNResponse struct {
	Status      string `json:"success"`
	Description string `json:"description"`
	Data        struct {
		Geo struct {
			Host          string          `json:"host"`
			IP            string          `json:"ip"`
			RDNS          string          `json:"rdns"`
			ASN           string          `json:"asn"`
			ISP           string          `json:"isp"`
			Country       string          `json:"country_name"`
			CountryCode   string          `json:"country_code"`
			City          string          `json:"city"`
			Postal        string          `json:"postal_code"`
			ContinentCode string          `json:"continent_code"`
			Latitude      decimal.Decimal `json:"latitude"`
			Longitude     decimal.Decimal `json:"longitude"`
			DMACode       string          `json:"dma_code"`
			AreaCode      string          `json:"area_code"`
			Timezone      string          `json:"timezone"`
			DateTime      string          `json:"datetime"`
		}
	}
}
