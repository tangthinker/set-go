package set_go

import "testing"

func TestSet(t *testing.T) {

	set1 := NewSet[int]()
	set2 := NewSet[int]()

	set1.Add(-1, 1, 2, 3, 4, 5)

	set2.Add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	t.Log(set1.Intersection(set2)) // [1 2 3 4 5]

	t.Log(set1.Union(set2)) // [1 2 3 4 5 -1 6 7 8 9 10]

	t.Log(set1.Difference(set2)) // [-1]

	t.Log(set2.Difference(set1)) // [6 7 8 9 10]

	t.Log(set1.IsSubset(set2)) // false

	t.Log(set1.IsSubset(set1)) // true

}
