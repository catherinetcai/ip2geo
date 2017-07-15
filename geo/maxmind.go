package geo

import "github.com/jinzhu/gorm"

type GeoLiteBlock struct {
	gorm.Model
}

//func (g *GeoLiteBlock) TableName() {
//	return "geolite_city_blocks"
//}
//
//type GeoLiteLocation struct {
//	gorm.Model
//}
//
//func (g *GeoLiteLocation) TableName() {
//	return "geolite_city_locations"
//}
