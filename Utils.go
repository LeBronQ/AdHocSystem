package AdHocSystem

import (
	"github.com/umahmood/haversine"
	"math"
)

func Distance(node1 Node, node2 Node) float64 {
	lat1 := node1.Pos.Lat
	lon1 := node1.Pos.Lon

	lat2 := node2.Pos.Lat
	lon2 := node2.Pos.Lon

	_, levelDistanceKm := haversine.Distance(haversine.Coord{Lat: lat1, Lon: lon1}, haversine.Coord{Lat: lat2, Lon: lon2})
	levelDistanceM := levelDistanceKm * 1000
	altDiff := node1.Pos.Alt - node2.Pos.Alt
	distance := math.Sqrt(altDiff*altDiff + levelDistanceM*levelDistanceM)

	return distance
}
