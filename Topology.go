package AdHocSystem

func GenerateRNGGraph(points []Node) [][]int {
	n := len(points)
	adjMatrix := make([][]int, n)
	for i := range adjMatrix {
		adjMatrix[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			adjMatrix[i][j] = 1
			adjMatrix[j][i] = 1
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if adjMatrix[i][j] == 1 {
				for k := 0; k < n; k++ {
					if k != i && k != j && adjMatrix[i][k] == 1 && adjMatrix[j][k] == 1 {
						if Distance(points[i], points[j]) >= max(Distance(points[i], points[k]), Distance(points[j], points[k])) {
							adjMatrix[i][j] = 0
							adjMatrix[j][i] = 0
							break
						}
					}
				}
			}
		}
	}

	return adjMatrix
}
