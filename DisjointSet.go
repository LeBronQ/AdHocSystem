package AdHocSystem

type DisjointSet struct {
	Tree []int
}

func (set *DisjointSet) InitTree(n int) {
	for i := 0; i < n; i++ {
		set.Tree[i] = i
	}
}

func (set *DisjointSet) Search(a int) int {
	if a == set.Tree[a] {
		return a
	} else {
		//压缩路径
		set.Tree[a] = set.Search(set.Tree[a])
		return set.Tree[a]
	}
}

func (set *DisjointSet) Union(a, b int) {
	rootA := set.Search(a)
	rootB := set.Search(b)
	if rootA == rootB {
		//a和b已经在同一颗树上
		return
	}
	if rootA > rootB {
		set.Tree[rootB] = rootA
	} else {
		set.Tree[rootA] = rootB
	}
}
