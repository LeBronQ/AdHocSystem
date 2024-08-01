package AdHocSystem

type Position struct {
	Lat float64
	Lon float64
	Alt float64
}

type Node struct {
	ID  int
	Pos Position
}
