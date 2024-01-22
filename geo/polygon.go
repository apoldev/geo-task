package geo

import (
	geo "github.com/kellydunn/golang-geo"
)

type Point struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type PolygonChecker interface {
	Contains(point Point) bool // проверить, находится ли точка внутри полигона
	Allowed() bool             // разрешено ли входить в полигон
	RandomPoint() Point        // сгенерировать случайную точку внутри полигона
}

type Polygon struct {
	polygon *geo.Polygon
	allowed bool
}

func NewPolygon(points []Point, allowed bool) *Polygon {
	// используем библиотеку golang-geo для создания полигона

	return &Polygon{
		polygon: geo.NewPolygon(geoPoints),
		allowed: allowed,
	}
}

func CheckPointIsAllowed(point Point, allowedZone PolygonChecker, disabledZones []PolygonChecker) bool {
	// проверить, находится ли точка в разрешенной зоне

	return false
}

func GetRandomAllowedLocation(allowedZone PolygonChecker, disabledZones []PolygonChecker) Point {
	var point Point
	// получение случайной точки в разрешенной зоне

	return point
}

func NewDisAllowedZone1() *Polygon {
	// добавить полигон с разрешенной зоной
	// полигоны лежат в /public/js/polygons.js

	return NewPolygon(points, false)
}

func NewDisAllowedZone2() *Polygon {
	// добавить полигон с разрешенной зоной
	// полигоны лежат в /public/js/polygons.js

	return NewPolygon(points, false)
}

func NewAllowedZone() *Polygon {
	// добавить полигон с разрешенной зоной
	// полигоны лежат в /public/js/polygons.js
	return NewPolygon(points, true)
}
