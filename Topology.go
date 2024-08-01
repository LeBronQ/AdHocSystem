package AdHocSystem

import "sort"

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

type Edge struct {
	src      int
	dst      int
	distance float64
}

func MSTWithDegreeLimit(points []Node, maxDegree int) [][]int {
	n := len(points)
	cnt := 0
	var edges []Edge
	adjMatrix := make([][]int, n)
	var set DisjointSet
	set.InitTree(n)
	degree := make([]int, n)
	for i := 0; i < n; i++ {
		degree[i] = 0
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dis := Distance(points[i], points[j])
			edges = append(edges, Edge{src: points[i].ID, dst: points[j].ID, distance: dis})
		}
	}
	sortEdges(edges)
	for _, edge := range edges {
		u, v := edge.src, edge.dst
		if degree[u] < maxDegree && degree[v] < maxDegree {
			if set.Search(u) != set.Search(v) {
				cnt++
				set.Union(u, v)
				degree[u]++
				degree[v]++
				adjMatrix[u] = append(adjMatrix[u], v)
				adjMatrix[v] = append(adjMatrix[v], u)
			}
		}
		if cnt == n-1 {
			break
		}
	}
	return adjMatrix
}

func sortEdges(edges []Edge) {
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].distance < edges[j].distance
	})
}
