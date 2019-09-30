package implementation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceAny(t *testing.T) {
	f := func(given []T, expected bool) {
		even := func(t T) bool { return (t % 2) == 0 }
		actual := Slice{given}.Any(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, false)
	f([]T{1, 3}, false)
	f([]T{2}, true)
	f([]T{1, 2}, true)
}

func TestSliceAll(t *testing.T) {
	f := func(given []T, expected bool) {
		even := func(t T) bool { return (t % 2) == 0 }
		actual := Slice{given}.All(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, true)
	f([]T{2}, true)
	f([]T{1}, false)
	f([]T{2, 4}, true)
	f([]T{2, 4, 1}, false)
	f([]T{1, 2, 4}, false)
}

func TestSliceChunkBy(t *testing.T) {
	f := func(given []T, expected [][]T) {
		reminder := func(t T) G { return G((t % 2)) }
		actual := Slice{given}.ChunkBy(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, [][]T{})
	f([]T{1}, [][]T{{1}})
	f([]T{1, 2, 3}, [][]T{{1}, {2}, {3}})
	f([]T{1, 3, 2, 4, 5}, [][]T{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkEvery(t *testing.T) {
	f := func(count int, given []T, expected [][]T) {
		actual, _ := Slice{given}.ChunkEvery(count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, []T{}, [][]T{})
	f(2, []T{1}, [][]T{{1}})
	f(2, []T{1, 2, 3, 4}, [][]T{{1, 2}, {3, 4}})
	f(2, []T{1, 2, 3, 4, 5}, [][]T{{1, 2}, {3, 4}, {5}})
}

func TestSliceContains(t *testing.T) {
	f := func(el T, given []T, expected bool) {
		actual := Slice{given}.Contains(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, []T{}, false)
	f(1, []T{1}, true)
	f(1, []T{2}, false)
	f(1, []T{2, 3, 4, 5}, false)
	f(1, []T{2, 3, 1, 4, 5}, true)
	f(1, []T{2, 3, 1, 1, 4, 5}, true)
}

func TestSliceCount(t *testing.T) {
	f := func(el T, given []T, expected int) {
		actual := Slice{given}.Count(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, []T{}, 0)
	f(1, []T{1}, 1)
	f(1, []T{2}, 0)
	f(1, []T{2, 3, 4, 5}, 0)
	f(1, []T{2, 3, 1, 4, 5}, 1)
	f(1, []T{2, 3, 1, 1, 4, 5}, 2)
	f(1, []T{1, 1, 1, 1, 1}, 5)
}

func TestSliceCountBy(t *testing.T) {
	f := func(given []T, expected int) {
		even := func(t T) bool { return (t % 2) == 0 }
		actual := Slice{given}.CountBy(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, 0)
	f([]T{1}, 0)
	f([]T{2}, 1)
	f([]T{1, 2, 3, 4, 5}, 2)
	f([]T{1, 2, 3, 4, 5, 6}, 3)
}

func TestSliceCycle(t *testing.T) {
	f := func(count int, given []T, expected []T) {
		c := Slice{given}.Cycle()
		seq := Channel{c}.Take(count)
		actual := Channel{seq}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(5, []T{}, []T{})
	f(5, []T{1}, []T{1, 1, 1, 1, 1})
	f(5, []T{1, 2}, []T{1, 2, 1, 2, 1})
}

func TestSliceDedup(t *testing.T) {
	f := func(given []T, expected []T) {
		actual := Slice{given}.Dedup()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, []T{})
	f([]T{1}, []T{1})
	f([]T{1, 1}, []T{1})
	f([]T{1, 2}, []T{1, 2})
	f([]T{1, 2, 3}, []T{1, 2, 3})
	f([]T{1, 2, 2, 3}, []T{1, 2, 3})
	f([]T{1, 2, 2, 3, 3, 3, 2, 1, 1}, []T{1, 2, 3, 2, 1})
}

func TestSliceDedupBy(t *testing.T) {
	f := func(given []T, expected []T) {
		even := func(el T) G { return G(el % 2) }
		actual := Slice{given}.DedupBy(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, []T{})
	f([]T{1}, []T{1})
	f([]T{1, 1}, []T{1})
	f([]T{1, 2}, []T{1, 2})
	f([]T{1, 2, 3}, []T{1, 2, 3})
	f([]T{1, 2, 2, 3}, []T{1, 2, 3})
	f([]T{1, 2, 4, 3, 5, 7, 10}, []T{1, 2, 3, 10})
}

func TestSliceDelete(t *testing.T) {
	f := func(given []T, el T, expected []T) {
		actual := Slice{given}.Delete(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, 1, []T{})
	f([]T{1}, 1, []T{})
	f([]T{2}, 1, []T{2})
	f([]T{1, 2}, 1, []T{2})
	f([]T{1, 2, 3}, 2, []T{1, 3})
	f([]T{1, 2, 2, 3, 2}, 2, []T{1, 2, 3, 2})
}

func TestSliceDeleteAll(t *testing.T) {
	f := func(given []T, el T, expected []T) {
		actual := Slice{given}.DeleteAll(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, 1, []T{})
	f([]T{1}, 1, []T{})
	f([]T{2}, 1, []T{2})
	f([]T{1, 2}, 1, []T{2})
	f([]T{1, 2, 3}, 2, []T{1, 3})
	f([]T{1, 2, 2, 3, 2}, 2, []T{1, 3})
}

func TestSliceDeleteAt(t *testing.T) {
	f := func(given []T, indices []int, expected []T) {
		actual, _ := Slice{given}.DeleteAt(indices...)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, []int{}, []T{})
	f([]T{1}, []int{0}, []T{})
	f([]T{1, 2}, []int{0}, []T{2})

	f([]T{1, 2, 3}, []int{0}, []T{2, 3})
	f([]T{1, 2, 3}, []int{1}, []T{1, 3})
	f([]T{1, 2, 3}, []int{2}, []T{1, 2})

	f([]T{1, 2, 3}, []int{0, 1}, []T{3})
	f([]T{1, 2, 3}, []int{0, 2}, []T{2})
	f([]T{1, 2, 3}, []int{1, 2}, []T{1})
}

func TestSliceDropEvery(t *testing.T) {
	f := func(given []T, nth int, expected []T) {
		actual, _ := Slice{given}.DropEvery(nth)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, 1, []T{})
	f([]T{1, 2, 3}, 1, []T{})

	f([]T{1, 2, 3, 4}, 2, []T{1, 3})
	f([]T{1, 2, 3, 4, 5}, 2, []T{1, 3, 5})

	f([]T{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2, []T{1, 3, 5, 7, 9})
	f([]T{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, []T{1, 2, 4, 5, 7, 8, 10})
}

func TestSliceDropWhile(t *testing.T) {
	f := func(given []T, expected []T) {
		even := func(el T) bool { return el%2 == 0 }
		actual := Slice{given}.DropWhile(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, []T{})
	f([]T{2}, []T{})
	f([]T{1}, []T{1})
	f([]T{2, 1}, []T{1})
	f([]T{2, 1, 2}, []T{1, 2})
	f([]T{1, 2}, []T{1, 2})
	f([]T{2, 4, 6, 1, 8}, []T{1, 8})
}

func TestSliceEndsWith(t *testing.T) {
	f := func(given []T, suffix []T, expected bool) {
		actual := Slice{given}.EndsWith(suffix)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, []T{}, true)
	f([]T{1}, []T{1}, true)
	f([]T{1}, []T{2}, false)
	f([]T{2, 3}, []T{1, 2, 3}, false)

	f([]T{1, 2, 3}, []T{3}, true)
	f([]T{1, 2, 3}, []T{2, 3}, true)
	f([]T{1, 2, 3}, []T{1, 2, 3}, true)

	f([]T{1, 2, 3}, []T{1}, false)
	f([]T{1, 2, 3}, []T{2}, false)
	f([]T{1, 2, 3}, []T{1, 2}, false)
	f([]T{1, 2, 3}, []T{3, 2}, false)
}

func TestSliceFilter(t *testing.T) {
	f := func(given []T, expected []T) {
		even := func(t T) bool { return (t % 2) == 0 }
		actual := Slice{given}.Filter(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, []T{})
	f([]T{1, 2, 3, 4}, []T{2, 4})
	f([]T{1, 3}, []T{})
	f([]T{2, 4}, []T{2, 4})
}

func TestSliceGroupBy(t *testing.T) {
	f := func(given []T, expected map[G][]T) {
		reminder := func(t T) G { return G((t % 2)) }
		actual := Slice{given}.GroupBy(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, map[G][]T{})
	f([]T{1}, map[G][]T{1: {1}})
	f([]T{1, 3, 2, 4, 5}, map[G][]T{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceIntersperse(t *testing.T) {
	f := func(el T, given []T, expected []T) {
		actual := Slice{given}.Intersperse(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, []T{}, []T{})
	f(0, []T{1}, []T{1})
	f(0, []T{1, 2}, []T{1, 0, 2})
	f(0, []T{1, 2, 3}, []T{1, 0, 2, 0, 3})
}

func TestSliceMap(t *testing.T) {
	f := func(given []T, expected []G) {
		double := func(t T) G { return G((t * 2)) }
		actual := Slice{given}.Map(double)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]T{}, []G{})
	f([]T{1}, []G{2})
	f([]T{1, 2, 3}, []G{2, 4, 6})
}

func TestSlicesPermutations(t *testing.T) {
	f := func(size int, given []T, expected [][]T) {
		actual := make([][]T, 0)
		i := 0
		s := Slice{given}
		for el := range s.Permutations(size) {
			actual = append(actual, el)
			i++
			if i > 50 {
				t.Fatal("infinite loop")
			}
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, []T{}, [][]T{})
	f(2, []T{1}, [][]T{{1}})
	f(2, []T{1, 2, 3}, [][]T{{1, 2}, {1, 3}, {2, 1}, {2, 3}, {3, 1}, {3, 2}})
}
