package genesis

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing")

func TestSlicesProductInt(t *testing.T) {
	f := func(given [][]int, expected [][]int) {
		actual := make([][]int, 0)
		i := 0
		s := SlicesInt{given}
		for el := range s.Product() {
			actual = append(actual, el)
			i++
			if i > 50 {
				t.Fatal("infinite loop")
			}
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([][]int{{1, 2}, {3, 4}}, [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}})
	f([][]int{{1, 2}, {3}, {4, 5}}, [][]int{{1, 3, 4}, {1, 3, 5}, {2, 3, 4}, {2, 3, 5}})
}

func TestSequenceCountInt(t *testing.T) {
	f := func(start int, step int, count int, expected []int) {
		ctx, cancel := context.WithCancel(context.Background())
		s := SequenceInt{ctx: ctx}
		seq := s.Count(start, step)
		seq2 := ChannelInt{seq}.Take(count)
		actual := ChannelInt{seq2}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 2, 4, []int{1, 3, 5, 7})
}

func TestSequenceExponentialInt(t *testing.T) {
	f := func(start int, factor int, count int, expected []int) {
		ctx, cancel := context.WithCancel(context.Background())
		s := SequenceInt{ctx: ctx}
		seq := s.Exponential(start, factor)
		seq2 := ChannelInt{seq}.Take(count)
		actual := ChannelInt{seq2}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 1, 4, []int{1, 1, 1, 1})
	f(1, 2, 4, []int{1, 2, 4, 8})
}

func TestSequenceIterateInt(t *testing.T) {
	f := func(start int, count int, expected []int) {
		ctx, cancel := context.WithCancel(context.Background())
		s := SequenceInt{ctx: ctx}
		double := func(val int) int { return val * 2 }
		seq := s.Iterate(start, double)
		seq2 := ChannelInt{seq}.Take(count)
		actual := ChannelInt{seq2}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 4, []int{1, 2, 4, 8})
}

func TestSequenceRangeInt(t *testing.T) {
	f := func(start int, stop int, step int, expected []int) {
		ctx, cancel := context.WithCancel(context.Background())
		s := SequenceInt{ctx: ctx}
		seq := s.Range(start, stop, step)
		actual := ChannelInt{seq}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 4, 1, []int{1, 2, 3})
	f(3, 0, -1, []int{3, 2, 1})
	f(1, 1, 1, []int{})
	f(1, 2, 1, []int{1})
}

func TestSequenceRepeatInt(t *testing.T) {
	f := func(count int, given int, expected []int) {
		ctx, cancel := context.WithCancel(context.Background())
		s := SequenceInt{ctx: ctx}
		seq := s.Repeat(given)
		seq2 := ChannelInt{seq}.Take(count)
		actual := ChannelInt{seq2}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, 1, []int{1, 1})
}

func TestSequenceReplicateInt(t *testing.T) {
	f := func(count int, given int, expected []int) {
		ctx, cancel := context.WithCancel(context.Background())
		s := SequenceInt{ctx: ctx}
		seq := s.Replicate(given, count)
		actual := ChannelInt{seq}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, 1, []int{})
	f(1, 1, []int{1})
	f(5, 1, []int{1, 1, 1, 1, 1})
}

func TestSliceAnyInt(t *testing.T) {
	f := func(given []int, expected bool) {
		even := func(t int) bool { return (t % 2) == 0 }
		actual := SliceInt{given}.Any(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, false)
	f([]int{1, 3}, false)
	f([]int{2}, true)
	f([]int{1, 2}, true)
}

func TestSliceAllInt(t *testing.T) {
	f := func(given []int, expected bool) {
		even := func(t int) bool { return (t % 2) == 0 }
		actual := SliceInt{given}.All(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, true)
	f([]int{2}, true)
	f([]int{1}, false)
	f([]int{2, 4}, true)
	f([]int{2, 4, 1}, false)
	f([]int{1, 2, 4}, false)
}

func TestSliceChunkByIntInt(t *testing.T) {
	f := func(given []int, expected [][]int) {
		reminder := func(t int) int { return int((t % 2)) }
		actual := SliceInt{given}.ChunkByInt(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, [][]int{})
	f([]int{1}, [][]int{{1}})
	f([]int{1, 2, 3}, [][]int{{1}, {2}, {3}})
	f([]int{1, 3, 2, 4, 5}, [][]int{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkByIntInt8(t *testing.T) {
	f := func(given []int, expected [][]int) {
		reminder := func(t int) int8 { return int8((t % 2)) }
		actual := SliceInt{given}.ChunkByInt8(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, [][]int{})
	f([]int{1}, [][]int{{1}})
	f([]int{1, 2, 3}, [][]int{{1}, {2}, {3}})
	f([]int{1, 3, 2, 4, 5}, [][]int{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkByIntInt16(t *testing.T) {
	f := func(given []int, expected [][]int) {
		reminder := func(t int) int16 { return int16((t % 2)) }
		actual := SliceInt{given}.ChunkByInt16(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, [][]int{})
	f([]int{1}, [][]int{{1}})
	f([]int{1, 2, 3}, [][]int{{1}, {2}, {3}})
	f([]int{1, 3, 2, 4, 5}, [][]int{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkByIntInt32(t *testing.T) {
	f := func(given []int, expected [][]int) {
		reminder := func(t int) int32 { return int32((t % 2)) }
		actual := SliceInt{given}.ChunkByInt32(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, [][]int{})
	f([]int{1}, [][]int{{1}})
	f([]int{1, 2, 3}, [][]int{{1}, {2}, {3}})
	f([]int{1, 3, 2, 4, 5}, [][]int{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkByIntInt64(t *testing.T) {
	f := func(given []int, expected [][]int) {
		reminder := func(t int) int64 { return int64((t % 2)) }
		actual := SliceInt{given}.ChunkByInt64(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, [][]int{})
	f([]int{1}, [][]int{{1}})
	f([]int{1, 2, 3}, [][]int{{1}, {2}, {3}})
	f([]int{1, 3, 2, 4, 5}, [][]int{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkEveryInt(t *testing.T) {
	f := func(count int, given []int, expected [][]int) {
		actual, _ := SliceInt{given}.ChunkEvery(count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, []int{}, [][]int{})
	f(2, []int{1}, [][]int{{1}})
	f(2, []int{1, 2, 3, 4}, [][]int{{1, 2}, {3, 4}})
	f(2, []int{1, 2, 3, 4, 5}, [][]int{{1, 2}, {3, 4}, {5}})
}

func TestSliceContainsInt(t *testing.T) {
	f := func(el int, given []int, expected bool) {
		actual := SliceInt{given}.Contains(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, []int{}, false)
	f(1, []int{1}, true)
	f(1, []int{2}, false)
	f(1, []int{2, 3, 4, 5}, false)
	f(1, []int{2, 3, 1, 4, 5}, true)
	f(1, []int{2, 3, 1, 1, 4, 5}, true)
}

func TestSliceCountInt(t *testing.T) {
	f := func(el int, given []int, expected int) {
		actual := SliceInt{given}.Count(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, []int{}, 0)
	f(1, []int{1}, 1)
	f(1, []int{2}, 0)
	f(1, []int{2, 3, 4, 5}, 0)
	f(1, []int{2, 3, 1, 4, 5}, 1)
	f(1, []int{2, 3, 1, 1, 4, 5}, 2)
	f(1, []int{1, 1, 1, 1, 1}, 5)
}

func TestSliceCountByInt(t *testing.T) {
	f := func(given []int, expected int) {
		even := func(t int) bool { return (t % 2) == 0 }
		actual := SliceInt{given}.CountBy(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, 0)
	f([]int{1}, 0)
	f([]int{2}, 1)
	f([]int{1, 2, 3, 4, 5}, 2)
	f([]int{1, 2, 3, 4, 5, 6}, 3)
}

func TestSliceCycleInt(t *testing.T) {
	f := func(count int, given []int, expected []int) {
		c := SliceInt{given}.Cycle()
		seq := ChannelInt{c}.Take(count)
		actual := ChannelInt{seq}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(5, []int{}, []int{})
	f(5, []int{1}, []int{1, 1, 1, 1, 1})
	f(5, []int{1, 2}, []int{1, 2, 1, 2, 1})
}

func TestSliceDedupInt(t *testing.T) {
	f := func(given []int, expected []int) {
		actual := SliceInt{given}.Dedup()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, []int{})
	f([]int{1}, []int{1})
	f([]int{1, 1}, []int{1})
	f([]int{1, 2}, []int{1, 2})
	f([]int{1, 2, 3}, []int{1, 2, 3})
	f([]int{1, 2, 2, 3}, []int{1, 2, 3})
	f([]int{1, 2, 2, 3, 3, 3, 2, 1, 1}, []int{1, 2, 3, 2, 1})
}

func TestSliceDedupByIntInt(t *testing.T) {
	f := func(given []int, expected []int) {
		even := func(el int) int { return int(el % 2) }
		actual := SliceInt{given}.DedupByInt(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, []int{})
	f([]int{1}, []int{1})
	f([]int{1, 1}, []int{1})
	f([]int{1, 2}, []int{1, 2})
	f([]int{1, 2, 3}, []int{1, 2, 3})
	f([]int{1, 2, 2, 3}, []int{1, 2, 3})
	f([]int{1, 2, 4, 3, 5, 7, 10}, []int{1, 2, 3, 10})
}

func TestSliceDedupByIntInt8(t *testing.T) {
	f := func(given []int, expected []int) {
		even := func(el int) int8 { return int8(el % 2) }
		actual := SliceInt{given}.DedupByInt8(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, []int{})
	f([]int{1}, []int{1})
	f([]int{1, 1}, []int{1})
	f([]int{1, 2}, []int{1, 2})
	f([]int{1, 2, 3}, []int{1, 2, 3})
	f([]int{1, 2, 2, 3}, []int{1, 2, 3})
	f([]int{1, 2, 4, 3, 5, 7, 10}, []int{1, 2, 3, 10})
}

func TestSliceDedupByIntInt16(t *testing.T) {
	f := func(given []int, expected []int) {
		even := func(el int) int16 { return int16(el % 2) }
		actual := SliceInt{given}.DedupByInt16(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, []int{})
	f([]int{1}, []int{1})
	f([]int{1, 1}, []int{1})
	f([]int{1, 2}, []int{1, 2})
	f([]int{1, 2, 3}, []int{1, 2, 3})
	f([]int{1, 2, 2, 3}, []int{1, 2, 3})
	f([]int{1, 2, 4, 3, 5, 7, 10}, []int{1, 2, 3, 10})
}

func TestSliceDedupByIntInt32(t *testing.T) {
	f := func(given []int, expected []int) {
		even := func(el int) int32 { return int32(el % 2) }
		actual := SliceInt{given}.DedupByInt32(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, []int{})
	f([]int{1}, []int{1})
	f([]int{1, 1}, []int{1})
	f([]int{1, 2}, []int{1, 2})
	f([]int{1, 2, 3}, []int{1, 2, 3})
	f([]int{1, 2, 2, 3}, []int{1, 2, 3})
	f([]int{1, 2, 4, 3, 5, 7, 10}, []int{1, 2, 3, 10})
}

func TestSliceDedupByIntInt64(t *testing.T) {
	f := func(given []int, expected []int) {
		even := func(el int) int64 { return int64(el % 2) }
		actual := SliceInt{given}.DedupByInt64(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, []int{})
	f([]int{1}, []int{1})
	f([]int{1, 1}, []int{1})
	f([]int{1, 2}, []int{1, 2})
	f([]int{1, 2, 3}, []int{1, 2, 3})
	f([]int{1, 2, 2, 3}, []int{1, 2, 3})
	f([]int{1, 2, 4, 3, 5, 7, 10}, []int{1, 2, 3, 10})
}

func TestSliceDeleteInt(t *testing.T) {
	f := func(given []int, el int, expected []int) {
		actual := SliceInt{given}.Delete(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, 1, []int{})
	f([]int{1}, 1, []int{})
	f([]int{2}, 1, []int{2})
	f([]int{1, 2}, 1, []int{2})
	f([]int{1, 2, 3}, 2, []int{1, 3})
	f([]int{1, 2, 2, 3, 2}, 2, []int{1, 2, 3, 2})
}

func TestSliceDeleteAllInt(t *testing.T) {
	f := func(given []int, el int, expected []int) {
		actual := SliceInt{given}.DeleteAll(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, 1, []int{})
	f([]int{1}, 1, []int{})
	f([]int{2}, 1, []int{2})
	f([]int{1, 2}, 1, []int{2})
	f([]int{1, 2, 3}, 2, []int{1, 3})
	f([]int{1, 2, 2, 3, 2}, 2, []int{1, 3})
}

func TestSliceDeleteAtInt(t *testing.T) {
	f := func(given []int, indices []int, expected []int) {
		actual, _ := SliceInt{given}.DeleteAt(indices...)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, []int{}, []int{})
	f([]int{1}, []int{0}, []int{})
	f([]int{1, 2}, []int{0}, []int{2})

	f([]int{1, 2, 3}, []int{0}, []int{2, 3})
	f([]int{1, 2, 3}, []int{1}, []int{1, 3})
	f([]int{1, 2, 3}, []int{2}, []int{1, 2})

	f([]int{1, 2, 3}, []int{0, 1}, []int{3})
	f([]int{1, 2, 3}, []int{0, 2}, []int{2})
	f([]int{1, 2, 3}, []int{1, 2}, []int{1})
}

func TestSliceDropEveryInt(t *testing.T) {
	f := func(given []int, nth int, from int, expected []int) {
		actual, _ := SliceInt{given}.DropEvery(nth, from)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, 1, 1, []int{})
	f([]int{1, 2, 3}, 1, 1, []int{})

	f([]int{1, 2, 3, 4}, 2, 1, []int{1, 3})
	f([]int{1, 2, 3, 4, 5}, 2, 1, []int{1, 3, 5})

	f([]int{1, 2, 3, 4}, 2, 0, []int{2, 4})
	f([]int{1, 2, 3, 4, 5}, 2, 0, []int{2, 4})

	f([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2, 1, []int{1, 3, 5, 7, 9})
	f([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, 1, []int{1, 2, 4, 5, 7, 8, 10})
}

func TestSliceDropWhileInt(t *testing.T) {
	f := func(given []int, expected []int) {
		even := func(el int) bool { return el%2 == 0 }
		actual := SliceInt{given}.DropWhile(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, []int{})
	f([]int{2}, []int{})
	f([]int{1}, []int{1})
	f([]int{2, 1}, []int{1})
	f([]int{2, 1, 2}, []int{1, 2})
	f([]int{1, 2}, []int{1, 2})
	f([]int{2, 4, 6, 1, 8}, []int{1, 8})
}

func TestSliceEndsWithInt(t *testing.T) {
	f := func(given []int, suffix []int, expected bool) {
		actual := SliceInt{given}.EndsWith(suffix)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, []int{}, true)
	f([]int{1}, []int{1}, true)
	f([]int{1}, []int{2}, false)
	f([]int{2, 3}, []int{1, 2, 3}, false)

	f([]int{1, 2, 3}, []int{3}, true)
	f([]int{1, 2, 3}, []int{2, 3}, true)
	f([]int{1, 2, 3}, []int{1, 2, 3}, true)

	f([]int{1, 2, 3}, []int{1}, false)
	f([]int{1, 2, 3}, []int{2}, false)
	f([]int{1, 2, 3}, []int{1, 2}, false)
	f([]int{1, 2, 3}, []int{3, 2}, false)
}

func TestSliceEqualInt(t *testing.T) {
	f := func(left []int, right []int, expected bool) {
		actual := SliceInt{left}.Equal(right)
		assert.Equal(t, expected, actual, "they should be equal")

		actual = SliceInt{right}.Equal(left)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, []int{}, true)
	f([]int{1}, []int{1}, true)
	f([]int{1}, []int{2}, false)
	f([]int{1, 2, 3, 3}, []int{1, 2, 3, 3}, true)
	f([]int{1, 2, 3, 3}, []int{1, 2, 2, 3}, false)
	f([]int{1, 2, 3, 3}, []int{1, 2, 4, 3}, false)

	// different len
	f([]int{1, 2, 3}, []int{1, 2}, false)
	f([]int{1, 2}, []int{1, 2, 3}, false)
	f([]int{}, []int{1, 2, 3}, false)
	f([]int{1, 2, 3}, []int{}, false)
}

func TestSliceFilterInt(t *testing.T) {
	f := func(given []int, expected []int) {
		even := func(t int) bool { return (t % 2) == 0 }
		actual := SliceInt{given}.Filter(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, []int{})
	f([]int{1, 2, 3, 4}, []int{2, 4})
	f([]int{1, 3}, []int{})
	f([]int{2, 4}, []int{2, 4})
}

func TestSliceFindInt(t *testing.T) {
	f := func(given []int, expectedEl int, expectedErr error) {
		even := func(t int) bool { return (t % 2) == 0 }
		el, err := SliceInt{given}.Find(even)
		assert.Equal(t, expectedEl, el, "they should be equal")
		assert.Equal(t, expectedErr, err, "they should be equal")
	}
	f([]int{}, 0, ErrNotFound)
	f([]int{1}, 0, ErrNotFound)
	f([]int{1}, 0, ErrNotFound)
	f([]int{2}, 2, nil)
	f([]int{1, 2}, 2, nil)
	f([]int{1, 2, 3}, 2, nil)
	f([]int{1, 3, 5}, 0, ErrNotFound)
}

func TestSliceFindIndexInt(t *testing.T) {
	f := func(given []int, expectedInd int, expectedErr error) {
		even := func(t int) bool { return (t % 2) == 0 }
		index, err := SliceInt{given}.FindIndex(even)
		assert.Equal(t, expectedInd, index, "they should be equal")
		assert.Equal(t, expectedErr, err, "they should be equal")
	}
	f([]int{}, 0, ErrNotFound)
	f([]int{1}, 0, ErrNotFound)
	f([]int{1}, 0, ErrNotFound)
	f([]int{2}, 0, nil)
	f([]int{1, 2}, 1, nil)
	f([]int{1, 2, 3}, 1, nil)
	f([]int{1, 3, 5, 7, 9, 2}, 5, nil)
	f([]int{1, 3, 5}, 0, ErrNotFound)
}

func TestSliceJoinInt(t *testing.T) {
	f := func(given []int, sep string, expected string) {
		actual := SliceInt{given}.Join(sep)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, "", "")
	f([]int{}, "|", "")

	f([]int{1}, "", "1")
	f([]int{1}, "|", "1")

	f([]int{1, 2, 3}, "", "123")
	f([]int{1, 2, 3}, "|", "1|2|3")
	f([]int{1, 2, 3}, "<int>", "1<int>2<int>3")
}

func TestSliceGroupByIntInt(t *testing.T) {
	f := func(given []int, expected map[int][]int) {
		reminder := func(t int) int { return int((t % 2)) }
		actual := SliceInt{given}.GroupByInt(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, map[int][]int{})
	f([]int{1}, map[int][]int{1: {1}})
	f([]int{1, 3, 2, 4, 5}, map[int][]int{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceGroupByIntInt8(t *testing.T) {
	f := func(given []int, expected map[int8][]int) {
		reminder := func(t int) int8 { return int8((t % 2)) }
		actual := SliceInt{given}.GroupByInt8(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, map[int8][]int{})
	f([]int{1}, map[int8][]int{1: {1}})
	f([]int{1, 3, 2, 4, 5}, map[int8][]int{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceGroupByIntInt16(t *testing.T) {
	f := func(given []int, expected map[int16][]int) {
		reminder := func(t int) int16 { return int16((t % 2)) }
		actual := SliceInt{given}.GroupByInt16(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, map[int16][]int{})
	f([]int{1}, map[int16][]int{1: {1}})
	f([]int{1, 3, 2, 4, 5}, map[int16][]int{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceGroupByIntInt32(t *testing.T) {
	f := func(given []int, expected map[int32][]int) {
		reminder := func(t int) int32 { return int32((t % 2)) }
		actual := SliceInt{given}.GroupByInt32(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, map[int32][]int{})
	f([]int{1}, map[int32][]int{1: {1}})
	f([]int{1, 3, 2, 4, 5}, map[int32][]int{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceGroupByIntInt64(t *testing.T) {
	f := func(given []int, expected map[int64][]int) {
		reminder := func(t int) int64 { return int64((t % 2)) }
		actual := SliceInt{given}.GroupByInt64(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, map[int64][]int{})
	f([]int{1}, map[int64][]int{1: {1}})
	f([]int{1, 3, 2, 4, 5}, map[int64][]int{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceInsertAtInt(t *testing.T) {
	f := func(given []int, index int, expected []int, expectedErr error) {
		actual, err := SliceInt{given}.InsertAt(index, 10)
		assert.Equal(t, expected, actual, "they should be equal")
		assert.Equal(t, expectedErr, err, "they should be equal")
	}
	f([]int{}, -1, []int{}, ErrNegativeValue)
	f([]int{}, 0, []int{10}, nil)
	f([]int{}, 1, []int{}, ErrOutOfRange)

	f([]int{1, 2, 3}, -1, []int{1, 2, 3}, ErrNegativeValue)
	f([]int{1, 2, 3}, 0, []int{10, 1, 2, 3}, nil)
	f([]int{1, 2, 3}, 1, []int{1, 10, 2, 3}, nil)
	f([]int{1, 2, 3}, 3, []int{1, 2, 3, 10}, nil)
	f([]int{1, 2, 3}, 4, []int{1, 2, 3}, ErrOutOfRange)
}

func TestSliceIntersperseInt(t *testing.T) {
	f := func(el int, given []int, expected []int) {
		actual := SliceInt{given}.Intersperse(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, []int{}, []int{})
	f(0, []int{1}, []int{1})
	f(0, []int{1, 2}, []int{1, 0, 2})
	f(0, []int{1, 2, 3}, []int{1, 0, 2, 0, 3})
}

func TestSliceLastInt(t *testing.T) {
	f := func(given []int, expectedEl int, expectedErr error) {
		el, err := SliceInt{given}.Last()
		assert.Equal(t, expectedEl, el, "they should be equal")
		assert.Equal(t, expectedErr, err, "they should be equal")
	}
	f([]int{}, 0, ErrEmpty)
	f([]int{1}, 1, nil)
	f([]int{1, 2, 3}, 3, nil)
}

func TestSliceMapIntInt(t *testing.T) {
	f := func(given []int, expected []int) {
		double := func(t int) int { return int((t * 2)) }
		actual := SliceInt{given}.MapInt(double)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, []int{})
	f([]int{1}, []int{2})
	f([]int{1, 2, 3}, []int{2, 4, 6})
}

func TestSliceMapIntInt8(t *testing.T) {
	f := func(given []int, expected []int8) {
		double := func(t int) int8 { return int8((t * 2)) }
		actual := SliceInt{given}.MapInt8(double)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, []int8{})
	f([]int{1}, []int8{2})
	f([]int{1, 2, 3}, []int8{2, 4, 6})
}

func TestSliceMapIntInt16(t *testing.T) {
	f := func(given []int, expected []int16) {
		double := func(t int) int16 { return int16((t * 2)) }
		actual := SliceInt{given}.MapInt16(double)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, []int16{})
	f([]int{1}, []int16{2})
	f([]int{1, 2, 3}, []int16{2, 4, 6})
}

func TestSliceMapIntInt32(t *testing.T) {
	f := func(given []int, expected []int32) {
		double := func(t int) int32 { return int32((t * 2)) }
		actual := SliceInt{given}.MapInt32(double)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, []int32{})
	f([]int{1}, []int32{2})
	f([]int{1, 2, 3}, []int32{2, 4, 6})
}

func TestSliceMapIntInt64(t *testing.T) {
	f := func(given []int, expected []int64) {
		double := func(t int) int64 { return int64((t * 2)) }
		actual := SliceInt{given}.MapInt64(double)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, []int64{})
	f([]int{1}, []int64{2})
	f([]int{1, 2, 3}, []int64{2, 4, 6})
}

func TestSliceMaxInt(t *testing.T) {
	f := func(given []int, expectedEl int, expectedErr error) {
		el, err := SliceInt{given}.Max()
		assert.Equal(t, expectedEl, el, "they should be equal")
		assert.Equal(t, expectedErr, err, "they should be equal")
	}
	f([]int{}, 0, ErrEmpty)
	f([]int{1}, 1, nil)
	f([]int{1, 2, 3}, 3, nil)
	f([]int{1, 3, 2}, 3, nil)
	f([]int{3, 2, 1}, 3, nil)
}

func TestSliceMinInt(t *testing.T) {
	f := func(given []int, expectedEl int, expectedErr error) {
		el, err := SliceInt{given}.Min()
		assert.Equal(t, expectedEl, el, "they should be equal")
		assert.Equal(t, expectedErr, err, "they should be equal")
	}
	f([]int{}, 0, ErrEmpty)
	f([]int{1}, 1, nil)
	f([]int{1, 2, 3}, 1, nil)
	f([]int{2, 1, 3}, 1, nil)
	f([]int{3, 2, 1}, 1, nil)
}

func TestSlicesPermutationsInt(t *testing.T) {
	f := func(size int, given []int, expected [][]int) {
		actual := make([][]int, 0)
		i := 0
		s := SliceInt{given}
		for el := range s.Permutations(size) {
			actual = append(actual, el)
			i++
			if i > 50 {
				t.Fatal("infinite loop")
			}
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, []int{}, [][]int{})
	f(2, []int{1}, [][]int{{1}})
	f(2, []int{1, 2, 3}, [][]int{{1, 2}, {1, 3}, {2, 1}, {2, 3}, {3, 1}, {3, 2}})
}

func TestSliceProductInt(t *testing.T) {
	f := func(given []int, repeat int, expected [][]int) {
		actual := make([][]int, 0)
		i := 0
		s := SliceInt{given}
		for el := range s.Product(repeat) {
			actual = append(actual, el)
			i++
			if i > 50 {
				t.Fatal("infinite loop")
			}
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}

	f([]int{1, 2}, 0, [][]int{})
	f([]int{}, 2, [][]int{})
	f([]int{1}, 2, [][]int{{1, 1}})

	f([]int{1, 2}, 1, [][]int{{1}, {2}})
	f([]int{1, 2}, 2, [][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}})
	f([]int{1, 2}, 3, [][]int{
		{1, 1, 1}, {1, 1, 2}, {1, 2, 1}, {1, 2, 2},
		{2, 1, 1}, {2, 1, 2}, {2, 2, 1}, {2, 2, 2},
	})
}

func TestSliceReduceIntInt(t *testing.T) {
	f := func(given []int, expected int) {
		sum := func(el int, acc int) int { return int(el) + acc }
		actual := SliceInt{given}.ReduceInt(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, 0)
	f([]int{1}, 1)
	f([]int{1, 2}, 3)
	f([]int{1, 2, 3}, 6)
}

func TestSliceReduceIntInt8(t *testing.T) {
	f := func(given []int, expected int8) {
		sum := func(el int, acc int8) int8 { return int8(el) + acc }
		actual := SliceInt{given}.ReduceInt8(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, 0)
	f([]int{1}, 1)
	f([]int{1, 2}, 3)
	f([]int{1, 2, 3}, 6)
}

func TestSliceReduceIntInt16(t *testing.T) {
	f := func(given []int, expected int16) {
		sum := func(el int, acc int16) int16 { return int16(el) + acc }
		actual := SliceInt{given}.ReduceInt16(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, 0)
	f([]int{1}, 1)
	f([]int{1, 2}, 3)
	f([]int{1, 2, 3}, 6)
}

func TestSliceReduceIntInt32(t *testing.T) {
	f := func(given []int, expected int32) {
		sum := func(el int, acc int32) int32 { return int32(el) + acc }
		actual := SliceInt{given}.ReduceInt32(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, 0)
	f([]int{1}, 1)
	f([]int{1, 2}, 3)
	f([]int{1, 2, 3}, 6)
}

func TestSliceReduceIntInt64(t *testing.T) {
	f := func(given []int, expected int64) {
		sum := func(el int, acc int64) int64 { return int64(el) + acc }
		actual := SliceInt{given}.ReduceInt64(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, 0)
	f([]int{1}, 1)
	f([]int{1, 2}, 3)
	f([]int{1, 2, 3}, 6)
}

func TestSliceReduceWhileIntInt(t *testing.T) {
	f := func(given []int, expected int) {
		sum := func(el int, acc int) (int, error) {
			if el == 0 {
				return acc, ErrEmpty
			}
			return int(el) + acc, nil
		}
		actual, _ := SliceInt{given}.ReduceWhileInt(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, 0)
	f([]int{1}, 1)
	f([]int{1, 2}, 3)
	f([]int{1, 2, 3}, 6)
	f([]int{1, 2, 0, 3}, 3)
}

func TestSliceReduceWhileIntInt8(t *testing.T) {
	f := func(given []int, expected int8) {
		sum := func(el int, acc int8) (int8, error) {
			if el == 0 {
				return acc, ErrEmpty
			}
			return int8(el) + acc, nil
		}
		actual, _ := SliceInt{given}.ReduceWhileInt8(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, 0)
	f([]int{1}, 1)
	f([]int{1, 2}, 3)
	f([]int{1, 2, 3}, 6)
	f([]int{1, 2, 0, 3}, 3)
}

func TestSliceReduceWhileIntInt16(t *testing.T) {
	f := func(given []int, expected int16) {
		sum := func(el int, acc int16) (int16, error) {
			if el == 0 {
				return acc, ErrEmpty
			}
			return int16(el) + acc, nil
		}
		actual, _ := SliceInt{given}.ReduceWhileInt16(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, 0)
	f([]int{1}, 1)
	f([]int{1, 2}, 3)
	f([]int{1, 2, 3}, 6)
	f([]int{1, 2, 0, 3}, 3)
}

func TestSliceReduceWhileIntInt32(t *testing.T) {
	f := func(given []int, expected int32) {
		sum := func(el int, acc int32) (int32, error) {
			if el == 0 {
				return acc, ErrEmpty
			}
			return int32(el) + acc, nil
		}
		actual, _ := SliceInt{given}.ReduceWhileInt32(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, 0)
	f([]int{1}, 1)
	f([]int{1, 2}, 3)
	f([]int{1, 2, 3}, 6)
	f([]int{1, 2, 0, 3}, 3)
}

func TestSliceReduceWhileIntInt64(t *testing.T) {
	f := func(given []int, expected int64) {
		sum := func(el int, acc int64) (int64, error) {
			if el == 0 {
				return acc, ErrEmpty
			}
			return int64(el) + acc, nil
		}
		actual, _ := SliceInt{given}.ReduceWhileInt64(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, 0)
	f([]int{1}, 1)
	f([]int{1, 2}, 3)
	f([]int{1, 2, 3}, 6)
	f([]int{1, 2, 0, 3}, 3)
}

func TestChannelToSliceInt(t *testing.T) {
	f := func(given []int) {
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := ChannelInt{c}.ToSlice()
		assert.Equal(t, given, actual, "they should be equal")
	}
	f([]int{})
	f([]int{1})
	f([]int{1, 2, 3, 1, 2})
}

func TestChannelAnyInt(t *testing.T) {
	f := func(given []int, expected bool) {
		even := func(t int) bool { return t%2 == 0 }
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := ChannelInt{c}.Any(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, false)
	f([]int{1}, false)
	f([]int{2}, true)
	f([]int{1, 2}, true)
	f([]int{1, 2, 3}, true)
	f([]int{1, 3, 5}, false)
	f([]int{1, 3, 5, 7, 9, 11}, false)
	f([]int{1, 3, 5, 7, 10, 11}, true)
}

func TestChannelAllInt(t *testing.T) {
	f := func(given []int, expected bool) {
		even := func(t int) bool { return t%2 == 0 }
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := ChannelInt{c}.All(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, true)
	f([]int{1}, false)
	f([]int{2}, true)
	f([]int{1, 2}, false)
	f([]int{2, 4}, true)
	f([]int{2, 4, 6, 8, 10, 12}, true)
	f([]int{2, 4, 6, 8, 11, 12}, false)
}

func TestChannelEachInt(t *testing.T) {
	f := func(given []int) {
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := make(chan int, len(given))
		mapper := func(t int) { result <- t }
		ChannelInt{c}.Each(mapper)
		close(result)
		actual := ChannelInt{result}.ToSlice()
		assert.Equal(t, given, actual, "they should be equal")
	}

	f([]int{})
	f([]int{1})
	f([]int{1, 2, 3})
	f([]int{1, 2, 3, 4, 5, 6, 7})
}

func TestChannelChunkEveryInt(t *testing.T) {
	f := func(size int, given []int, expected [][]int) {
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := ChannelInt{c}.ChunkEvery(size)
		actual := make([][]int, 0)
		for el := range result {
			actual = append(actual, el)
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, []int{}, [][]int{})
	f(2, []int{1}, [][]int{{1}})
	f(2, []int{1, 2}, [][]int{{1, 2}})
	f(2, []int{1, 2, 3, 4}, [][]int{{1, 2}, {3, 4}})
	f(2, []int{1, 2, 3, 4, 5}, [][]int{{1, 2}, {3, 4}, {5}})
}

func TestChannelCountInt(t *testing.T) {
	f := func(element int, given []int, expected int) {
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := ChannelInt{c}.Count(element)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, []int{}, 0)
	f(1, []int{1}, 1)
	f(1, []int{2}, 0)
	f(1, []int{1, 2, 3, 1, 4}, 2)
}

func TestChannelDropInt(t *testing.T) {
	f := func(count int, given []int, expected []int) {
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := ChannelInt{c}.Drop(count)
		actual := make([]int, 0)
		for el := range result {
			actual = append(actual, el)
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, []int{}, []int{})
	f(1, []int{2}, []int{})
	f(1, []int{2, 3}, []int{3})
	f(1, []int{1, 2, 3}, []int{2, 3})
	f(0, []int{1, 2, 3}, []int{1, 2, 3})
	f(3, []int{1, 2, 3, 4, 5, 6}, []int{4, 5, 6})
	f(1, []int{1, 2, 3, 4, 5, 6}, []int{2, 3, 4, 5, 6})
}

func TestChannelFilterInt(t *testing.T) {
	f := func(given []int, expected []int) {
		even := func(t int) bool { return t%2 == 0 }
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := ChannelInt{c}.Filter(even)
		actual := ChannelInt{result}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, []int{})
	f([]int{1}, []int{})
	f([]int{2}, []int{2})
	f([]int{1, 2, 3, 4}, []int{2, 4})
}

func TestChannelMapIntInt(t *testing.T) {
	f := func(given []int, expected []int) {
		double := func(el int) int { return int(el * 2) }
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := ChannelInt{c}.MapInt(double)

		// convert chan int to chan int
		c2 := make(chan int, 1)
		go func() {
			for el := range result {
				c2 <- int(el)
			}
			close(c2)
		}()

		actual := ChannelInt{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, []int{})
	f([]int{1}, []int{2})
	f([]int{1, 2, 3}, []int{2, 4, 6})
}

func TestChannelMapIntInt8(t *testing.T) {
	f := func(given []int, expected []int) {
		double := func(el int) int8 { return int8(el * 2) }
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := ChannelInt{c}.MapInt8(double)

		// convert chan int to chan int8
		c2 := make(chan int, 1)
		go func() {
			for el := range result {
				c2 <- int(el)
			}
			close(c2)
		}()

		actual := ChannelInt{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, []int{})
	f([]int{1}, []int{2})
	f([]int{1, 2, 3}, []int{2, 4, 6})
}

func TestChannelMapIntInt16(t *testing.T) {
	f := func(given []int, expected []int) {
		double := func(el int) int16 { return int16(el * 2) }
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := ChannelInt{c}.MapInt16(double)

		// convert chan int to chan int16
		c2 := make(chan int, 1)
		go func() {
			for el := range result {
				c2 <- int(el)
			}
			close(c2)
		}()

		actual := ChannelInt{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, []int{})
	f([]int{1}, []int{2})
	f([]int{1, 2, 3}, []int{2, 4, 6})
}

func TestChannelMapIntInt32(t *testing.T) {
	f := func(given []int, expected []int) {
		double := func(el int) int32 { return int32(el * 2) }
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := ChannelInt{c}.MapInt32(double)

		// convert chan int to chan int32
		c2 := make(chan int, 1)
		go func() {
			for el := range result {
				c2 <- int(el)
			}
			close(c2)
		}()

		actual := ChannelInt{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, []int{})
	f([]int{1}, []int{2})
	f([]int{1, 2, 3}, []int{2, 4, 6})
}

func TestChannelMapIntInt64(t *testing.T) {
	f := func(given []int, expected []int) {
		double := func(el int) int64 { return int64(el * 2) }
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := ChannelInt{c}.MapInt64(double)

		// convert chan int to chan int64
		c2 := make(chan int, 1)
		go func() {
			for el := range result {
				c2 <- int(el)
			}
			close(c2)
		}()

		actual := ChannelInt{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, []int{})
	f([]int{1}, []int{2})
	f([]int{1, 2, 3}, []int{2, 4, 6})
}

func TestChannelMaxInt(t *testing.T) {
	f := func(given []int, expected int, expectedErr error) {
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual, actualErr := ChannelInt{c}.Max()
		assert.Equal(t, expected, actual, "they should be equal")
		assert.Equal(t, expectedErr, actualErr, "they should be equal")
	}
	f([]int{}, 0, ErrEmpty)
	f([]int{1, 4, 2}, 4, nil)
	f([]int{1, 2, 4}, 4, nil)
	f([]int{4, 2, 1}, 4, nil)
}

func TestChannelMinInt(t *testing.T) {
	f := func(given []int, expected int, expectedErr error) {
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual, actualErr := ChannelInt{c}.Min()
		assert.Equal(t, expected, actual, "they should be equal")
		assert.Equal(t, expectedErr, actualErr, "they should be equal")
	}
	f([]int{}, 0, ErrEmpty)
	f([]int{4, 1, 2}, 1, nil)
	f([]int{1, 2, 4}, 1, nil)
	f([]int{4, 2, 1}, 1, nil)
}

func TestChannelReduceIntInt(t *testing.T) {
	f := func(given []int, expected int) {
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int, acc int) int { return int(el) + acc }
		actual := ChannelInt{c}.ReduceInt(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, 0)
	f([]int{1}, 1)
	f([]int{1, 2}, 3)
	f([]int{1, 2, 3, 4, 5}, 15)
}

func TestChannelReduceIntInt8(t *testing.T) {
	f := func(given []int, expected int8) {
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int, acc int8) int8 { return int8(el) + acc }
		actual := ChannelInt{c}.ReduceInt8(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, 0)
	f([]int{1}, 1)
	f([]int{1, 2}, 3)
	f([]int{1, 2, 3, 4, 5}, 15)
}

func TestChannelReduceIntInt16(t *testing.T) {
	f := func(given []int, expected int16) {
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int, acc int16) int16 { return int16(el) + acc }
		actual := ChannelInt{c}.ReduceInt16(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, 0)
	f([]int{1}, 1)
	f([]int{1, 2}, 3)
	f([]int{1, 2, 3, 4, 5}, 15)
}

func TestChannelReduceIntInt32(t *testing.T) {
	f := func(given []int, expected int32) {
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int, acc int32) int32 { return int32(el) + acc }
		actual := ChannelInt{c}.ReduceInt32(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, 0)
	f([]int{1}, 1)
	f([]int{1, 2}, 3)
	f([]int{1, 2, 3, 4, 5}, 15)
}

func TestChannelReduceIntInt64(t *testing.T) {
	f := func(given []int, expected int64) {
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int, acc int64) int64 { return int64(el) + acc }
		actual := ChannelInt{c}.ReduceInt64(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, 0)
	f([]int{1}, 1)
	f([]int{1, 2}, 3)
	f([]int{1, 2, 3, 4, 5}, 15)
}

func TestChannelScanIntInt(t *testing.T) {
	f := func(given []int, expected []int) {
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int, acc int) int { return int(el) + acc }
		result := ChannelInt{c}.ScanInt(0, sum)

		// convert chan int to chan int
		c2 := make(chan int, 1)
		go func() {
			for el := range result {
				c2 <- int(el)
			}
			close(c2)
		}()

		actual := ChannelInt{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, []int{})
	f([]int{1}, []int{1})
	f([]int{1, 2}, []int{1, 3})
	f([]int{1, 2, 3, 4, 5}, []int{1, 3, 6, 10, 15})
}

func TestChannelScanIntInt8(t *testing.T) {
	f := func(given []int, expected []int) {
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int, acc int8) int8 { return int8(el) + acc }
		result := ChannelInt{c}.ScanInt8(0, sum)

		// convert chan int to chan int8
		c2 := make(chan int, 1)
		go func() {
			for el := range result {
				c2 <- int(el)
			}
			close(c2)
		}()

		actual := ChannelInt{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, []int{})
	f([]int{1}, []int{1})
	f([]int{1, 2}, []int{1, 3})
	f([]int{1, 2, 3, 4, 5}, []int{1, 3, 6, 10, 15})
}

func TestChannelScanIntInt16(t *testing.T) {
	f := func(given []int, expected []int) {
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int, acc int16) int16 { return int16(el) + acc }
		result := ChannelInt{c}.ScanInt16(0, sum)

		// convert chan int to chan int16
		c2 := make(chan int, 1)
		go func() {
			for el := range result {
				c2 <- int(el)
			}
			close(c2)
		}()

		actual := ChannelInt{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, []int{})
	f([]int{1}, []int{1})
	f([]int{1, 2}, []int{1, 3})
	f([]int{1, 2, 3, 4, 5}, []int{1, 3, 6, 10, 15})
}

func TestChannelScanIntInt32(t *testing.T) {
	f := func(given []int, expected []int) {
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int, acc int32) int32 { return int32(el) + acc }
		result := ChannelInt{c}.ScanInt32(0, sum)

		// convert chan int to chan int32
		c2 := make(chan int, 1)
		go func() {
			for el := range result {
				c2 <- int(el)
			}
			close(c2)
		}()

		actual := ChannelInt{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, []int{})
	f([]int{1}, []int{1})
	f([]int{1, 2}, []int{1, 3})
	f([]int{1, 2, 3, 4, 5}, []int{1, 3, 6, 10, 15})
}

func TestChannelScanIntInt64(t *testing.T) {
	f := func(given []int, expected []int) {
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int, acc int64) int64 { return int64(el) + acc }
		result := ChannelInt{c}.ScanInt64(0, sum)

		// convert chan int to chan int64
		c2 := make(chan int, 1)
		go func() {
			for el := range result {
				c2 <- int(el)
			}
			close(c2)
		}()

		actual := ChannelInt{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, []int{})
	f([]int{1}, []int{1})
	f([]int{1, 2}, []int{1, 3})
	f([]int{1, 2, 3, 4, 5}, []int{1, 3, 6, 10, 15})
}

func TestChannelSumInt(t *testing.T) {
	f := func(given []int, expected int) {
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := ChannelInt{c}.Sum()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int{}, 0)
	f([]int{1}, 1)
	f([]int{1, 2}, 3)
	f([]int{1, 2, 3, 4, 5}, 15)
}

func TestChannelTakeInt(t *testing.T) {
	f := func(count int, given int, expected []int) {
		ctx, cancel := context.WithCancel(context.Background())
		s := SequenceInt{ctx: ctx}
		seq := s.Repeat(given)
		seq2 := ChannelInt{seq}.Take(count)
		actual := ChannelInt{seq2}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, 1, []int{})
	f(1, 1, []int{1})
	f(2, 1, []int{1, 1})
}

func TestChannelTeeInt(t *testing.T) {
	f := func(count int, given []int) {
		c := make(chan int, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		channels := ChannelInt{c}.Tee(count)
		for _, ch := range channels {
			go func(ch chan int) {
				actual := ChannelInt{ch}.ToSlice()
				assert.Equal(t, given, actual, "they should be equal")
			}(ch)
		}
	}
	f(1, []int{})
	f(1, []int{1})
	f(1, []int{1, 2})
	f(1, []int{1, 2, 3})
	f(1, []int{1, 2, 3, 1, 2})

	f(2, []int{})
	f(2, []int{1})
	f(2, []int{1, 2})
	f(2, []int{1, 2, 3})
	f(2, []int{1, 2, 3, 1, 2})

	f(10, []int{1, 2, 3, 1, 2})
}

func TestAsyncSliceAnyInt(t *testing.T) {
	f := func(check func(t int) bool, given []int, expected bool) {
		s := AsyncSliceInt{Data: given, Workers: 2}
		actual := s.Any(check)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	isEven := func(t int) bool { return (t % 2) == 0 }

	f(isEven, []int{}, false)
	f(isEven, []int{1}, false)
	f(isEven, []int{1, 3}, false)
	f(isEven, []int{2}, true)
	f(isEven, []int{1, 2}, true)
	f(isEven, []int{1, 3, 5, 7, 9, 11}, false)
	f(isEven, []int{1, 3, 5, 7, 9, 12}, true)
}

func TestAsyncSliceAllInt(t *testing.T) {
	f := func(check func(t int) bool, given []int, expected bool) {
		s := AsyncSliceInt{Data: given, Workers: 2}
		actual := s.All(check)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	isEven := func(t int) bool { return (t % 2) == 0 }

	f(isEven, []int{}, true)
	f(isEven, []int{1}, false)
	f(isEven, []int{1, 3}, false)
	f(isEven, []int{2}, true)
	f(isEven, []int{2, 4}, true)
	f(isEven, []int{2, 3}, false)
	f(isEven, []int{2, 4, 6, 8, 10, 12}, true)
	f(isEven, []int{2, 4, 6, 8, 10, 11}, false)
}

func TestAsyncSliceEachInt(t *testing.T) {
	f := func(given []int) {
		s := AsyncSliceInt{Data: given, Workers: 2}
		result := make(chan int, len(given))
		mapper := func(t int) { result <- t }
		s.Each(mapper)
		close(result)
		actual := ChannelInt{result}.ToSlice()
		sorted := SliceInt{actual}.Sort()
		assert.Equal(t, given, sorted, "they should be equal")
	}

	f([]int{})
	f([]int{1})
	f([]int{1, 2, 3})
	f([]int{1, 2, 3, 4, 5, 6, 7})
}

func TestAsyncSliceFilterInt(t *testing.T) {
	f := func(given []int, expected []int) {
		filter := func(t int) bool { return t > 10 }
		s := AsyncSliceInt{Data: given, Workers: 2}
		actual := s.Filter(filter)
		assert.Equal(t, expected, actual, "they should be equal")
	}

	f([]int{}, []int{})
	f([]int{5}, []int{})
	f([]int{15}, []int{15})
	f([]int{9, 11, 12, 13, 6}, []int{11, 12, 13})
}

func TestAsyncSliceMapIntInt(t *testing.T) {
	f := func(mapper func(t int) int, given []int, expected []int) {
		s := AsyncSliceInt{Data: given, Workers: 2}
		actual := s.MapInt(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int) int { return int((t * 2)) }

	f(double, []int{}, []int{})
	f(double, []int{1}, []int{2})
	f(double, []int{1, 2, 3}, []int{2, 4, 6})
}

func TestAsyncSliceMapIntInt8(t *testing.T) {
	f := func(mapper func(t int) int8, given []int, expected []int8) {
		s := AsyncSliceInt{Data: given, Workers: 2}
		actual := s.MapInt8(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int) int8 { return int8((t * 2)) }

	f(double, []int{}, []int8{})
	f(double, []int{1}, []int8{2})
	f(double, []int{1, 2, 3}, []int8{2, 4, 6})
}

func TestAsyncSliceMapIntInt16(t *testing.T) {
	f := func(mapper func(t int) int16, given []int, expected []int16) {
		s := AsyncSliceInt{Data: given, Workers: 2}
		actual := s.MapInt16(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int) int16 { return int16((t * 2)) }

	f(double, []int{}, []int16{})
	f(double, []int{1}, []int16{2})
	f(double, []int{1, 2, 3}, []int16{2, 4, 6})
}

func TestAsyncSliceMapIntInt32(t *testing.T) {
	f := func(mapper func(t int) int32, given []int, expected []int32) {
		s := AsyncSliceInt{Data: given, Workers: 2}
		actual := s.MapInt32(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int) int32 { return int32((t * 2)) }

	f(double, []int{}, []int32{})
	f(double, []int{1}, []int32{2})
	f(double, []int{1, 2, 3}, []int32{2, 4, 6})
}

func TestAsyncSliceMapIntInt64(t *testing.T) {
	f := func(mapper func(t int) int64, given []int, expected []int64) {
		s := AsyncSliceInt{Data: given, Workers: 2}
		actual := s.MapInt64(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int) int64 { return int64((t * 2)) }

	f(double, []int{}, []int64{})
	f(double, []int{1}, []int64{2})
	f(double, []int{1, 2, 3}, []int64{2, 4, 6})
}

func TestAsyncSliceReduceInt(t *testing.T) {
	f := func(reducer func(a int, b int) int, given []int, expected int) {
		s := AsyncSliceInt{Data: given, Workers: 4}
		actual := s.Reduce(reducer)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	sum := func(a int, b int) int { return a + b }

	f(sum, []int{}, 0)
	f(sum, []int{1}, 1)
	f(sum, []int{1, 2}, 3)
	f(sum, []int{1, 2, 3}, 6)
	f(sum, []int{1, 2, 3, 4}, 10)
	f(sum, []int{1, 2, 3, 4, 5}, 15)
}

func TestSlicesProductInt8(t *testing.T) {
	f := func(given [][]int8, expected [][]int8) {
		actual := make([][]int8, 0)
		i := 0
		s := SlicesInt8{given}
		for el := range s.Product() {
			actual = append(actual, el)
			i++
			if i > 50 {
				t.Fatal("infinite loop")
			}
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([][]int8{{1, 2}, {3, 4}}, [][]int8{{1, 3}, {1, 4}, {2, 3}, {2, 4}})
	f([][]int8{{1, 2}, {3}, {4, 5}}, [][]int8{{1, 3, 4}, {1, 3, 5}, {2, 3, 4}, {2, 3, 5}})
}

func TestSequenceCountInt8(t *testing.T) {
	f := func(start int8, step int8, count int, expected []int8) {
		ctx, cancel := context.WithCancel(context.Background())
		s := SequenceInt8{ctx: ctx}
		seq := s.Count(start, step)
		seq2 := ChannelInt8{seq}.Take(count)
		actual := ChannelInt8{seq2}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 2, 4, []int8{1, 3, 5, 7})
}

func TestSequenceExponentialInt8(t *testing.T) {
	f := func(start int8, factor int8, count int, expected []int8) {
		ctx, cancel := context.WithCancel(context.Background())
		s := SequenceInt8{ctx: ctx}
		seq := s.Exponential(start, factor)
		seq2 := ChannelInt8{seq}.Take(count)
		actual := ChannelInt8{seq2}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 1, 4, []int8{1, 1, 1, 1})
	f(1, 2, 4, []int8{1, 2, 4, 8})
}

func TestSequenceIterateInt8(t *testing.T) {
	f := func(start int8, count int, expected []int8) {
		ctx, cancel := context.WithCancel(context.Background())
		s := SequenceInt8{ctx: ctx}
		double := func(val int8) int8 { return val * 2 }
		seq := s.Iterate(start, double)
		seq2 := ChannelInt8{seq}.Take(count)
		actual := ChannelInt8{seq2}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 4, []int8{1, 2, 4, 8})
}

func TestSequenceRangeInt8(t *testing.T) {
	f := func(start int8, stop int8, step int8, expected []int8) {
		ctx, cancel := context.WithCancel(context.Background())
		s := SequenceInt8{ctx: ctx}
		seq := s.Range(start, stop, step)
		actual := ChannelInt8{seq}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 4, 1, []int8{1, 2, 3})
	f(3, 0, -1, []int8{3, 2, 1})
	f(1, 1, 1, []int8{})
	f(1, 2, 1, []int8{1})
}

func TestSequenceRepeatInt8(t *testing.T) {
	f := func(count int, given int8, expected []int8) {
		ctx, cancel := context.WithCancel(context.Background())
		s := SequenceInt8{ctx: ctx}
		seq := s.Repeat(given)
		seq2 := ChannelInt8{seq}.Take(count)
		actual := ChannelInt8{seq2}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, 1, []int8{1, 1})
}

func TestSequenceReplicateInt8(t *testing.T) {
	f := func(count int, given int8, expected []int8) {
		ctx, cancel := context.WithCancel(context.Background())
		s := SequenceInt8{ctx: ctx}
		seq := s.Replicate(given, count)
		actual := ChannelInt8{seq}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, 1, []int8{})
	f(1, 1, []int8{1})
	f(5, 1, []int8{1, 1, 1, 1, 1})
}

func TestSliceAnyInt8(t *testing.T) {
	f := func(given []int8, expected bool) {
		even := func(t int8) bool { return (t % 2) == 0 }
		actual := SliceInt8{given}.Any(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, false)
	f([]int8{1, 3}, false)
	f([]int8{2}, true)
	f([]int8{1, 2}, true)
}

func TestSliceAllInt8(t *testing.T) {
	f := func(given []int8, expected bool) {
		even := func(t int8) bool { return (t % 2) == 0 }
		actual := SliceInt8{given}.All(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, true)
	f([]int8{2}, true)
	f([]int8{1}, false)
	f([]int8{2, 4}, true)
	f([]int8{2, 4, 1}, false)
	f([]int8{1, 2, 4}, false)
}

func TestSliceChunkByInt8Int(t *testing.T) {
	f := func(given []int8, expected [][]int8) {
		reminder := func(t int8) int { return int((t % 2)) }
		actual := SliceInt8{given}.ChunkByInt(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, [][]int8{})
	f([]int8{1}, [][]int8{{1}})
	f([]int8{1, 2, 3}, [][]int8{{1}, {2}, {3}})
	f([]int8{1, 3, 2, 4, 5}, [][]int8{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkByInt8Int8(t *testing.T) {
	f := func(given []int8, expected [][]int8) {
		reminder := func(t int8) int8 { return int8((t % 2)) }
		actual := SliceInt8{given}.ChunkByInt8(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, [][]int8{})
	f([]int8{1}, [][]int8{{1}})
	f([]int8{1, 2, 3}, [][]int8{{1}, {2}, {3}})
	f([]int8{1, 3, 2, 4, 5}, [][]int8{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkByInt8Int16(t *testing.T) {
	f := func(given []int8, expected [][]int8) {
		reminder := func(t int8) int16 { return int16((t % 2)) }
		actual := SliceInt8{given}.ChunkByInt16(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, [][]int8{})
	f([]int8{1}, [][]int8{{1}})
	f([]int8{1, 2, 3}, [][]int8{{1}, {2}, {3}})
	f([]int8{1, 3, 2, 4, 5}, [][]int8{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkByInt8Int32(t *testing.T) {
	f := func(given []int8, expected [][]int8) {
		reminder := func(t int8) int32 { return int32((t % 2)) }
		actual := SliceInt8{given}.ChunkByInt32(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, [][]int8{})
	f([]int8{1}, [][]int8{{1}})
	f([]int8{1, 2, 3}, [][]int8{{1}, {2}, {3}})
	f([]int8{1, 3, 2, 4, 5}, [][]int8{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkByInt8Int64(t *testing.T) {
	f := func(given []int8, expected [][]int8) {
		reminder := func(t int8) int64 { return int64((t % 2)) }
		actual := SliceInt8{given}.ChunkByInt64(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, [][]int8{})
	f([]int8{1}, [][]int8{{1}})
	f([]int8{1, 2, 3}, [][]int8{{1}, {2}, {3}})
	f([]int8{1, 3, 2, 4, 5}, [][]int8{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkEveryInt8(t *testing.T) {
	f := func(count int, given []int8, expected [][]int8) {
		actual, _ := SliceInt8{given}.ChunkEvery(count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, []int8{}, [][]int8{})
	f(2, []int8{1}, [][]int8{{1}})
	f(2, []int8{1, 2, 3, 4}, [][]int8{{1, 2}, {3, 4}})
	f(2, []int8{1, 2, 3, 4, 5}, [][]int8{{1, 2}, {3, 4}, {5}})
}

func TestSliceContainsInt8(t *testing.T) {
	f := func(el int8, given []int8, expected bool) {
		actual := SliceInt8{given}.Contains(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, []int8{}, false)
	f(1, []int8{1}, true)
	f(1, []int8{2}, false)
	f(1, []int8{2, 3, 4, 5}, false)
	f(1, []int8{2, 3, 1, 4, 5}, true)
	f(1, []int8{2, 3, 1, 1, 4, 5}, true)
}

func TestSliceCountInt8(t *testing.T) {
	f := func(el int8, given []int8, expected int) {
		actual := SliceInt8{given}.Count(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, []int8{}, 0)
	f(1, []int8{1}, 1)
	f(1, []int8{2}, 0)
	f(1, []int8{2, 3, 4, 5}, 0)
	f(1, []int8{2, 3, 1, 4, 5}, 1)
	f(1, []int8{2, 3, 1, 1, 4, 5}, 2)
	f(1, []int8{1, 1, 1, 1, 1}, 5)
}

func TestSliceCountByInt8(t *testing.T) {
	f := func(given []int8, expected int) {
		even := func(t int8) bool { return (t % 2) == 0 }
		actual := SliceInt8{given}.CountBy(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, 0)
	f([]int8{1}, 0)
	f([]int8{2}, 1)
	f([]int8{1, 2, 3, 4, 5}, 2)
	f([]int8{1, 2, 3, 4, 5, 6}, 3)
}

func TestSliceCycleInt8(t *testing.T) {
	f := func(count int, given []int8, expected []int8) {
		c := SliceInt8{given}.Cycle()
		seq := ChannelInt8{c}.Take(count)
		actual := ChannelInt8{seq}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(5, []int8{}, []int8{})
	f(5, []int8{1}, []int8{1, 1, 1, 1, 1})
	f(5, []int8{1, 2}, []int8{1, 2, 1, 2, 1})
}

func TestSliceDedupInt8(t *testing.T) {
	f := func(given []int8, expected []int8) {
		actual := SliceInt8{given}.Dedup()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, []int8{})
	f([]int8{1}, []int8{1})
	f([]int8{1, 1}, []int8{1})
	f([]int8{1, 2}, []int8{1, 2})
	f([]int8{1, 2, 3}, []int8{1, 2, 3})
	f([]int8{1, 2, 2, 3}, []int8{1, 2, 3})
	f([]int8{1, 2, 2, 3, 3, 3, 2, 1, 1}, []int8{1, 2, 3, 2, 1})
}

func TestSliceDedupByInt8Int(t *testing.T) {
	f := func(given []int8, expected []int8) {
		even := func(el int8) int { return int(el % 2) }
		actual := SliceInt8{given}.DedupByInt(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, []int8{})
	f([]int8{1}, []int8{1})
	f([]int8{1, 1}, []int8{1})
	f([]int8{1, 2}, []int8{1, 2})
	f([]int8{1, 2, 3}, []int8{1, 2, 3})
	f([]int8{1, 2, 2, 3}, []int8{1, 2, 3})
	f([]int8{1, 2, 4, 3, 5, 7, 10}, []int8{1, 2, 3, 10})
}

func TestSliceDedupByInt8Int8(t *testing.T) {
	f := func(given []int8, expected []int8) {
		even := func(el int8) int8 { return int8(el % 2) }
		actual := SliceInt8{given}.DedupByInt8(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, []int8{})
	f([]int8{1}, []int8{1})
	f([]int8{1, 1}, []int8{1})
	f([]int8{1, 2}, []int8{1, 2})
	f([]int8{1, 2, 3}, []int8{1, 2, 3})
	f([]int8{1, 2, 2, 3}, []int8{1, 2, 3})
	f([]int8{1, 2, 4, 3, 5, 7, 10}, []int8{1, 2, 3, 10})
}

func TestSliceDedupByInt8Int16(t *testing.T) {
	f := func(given []int8, expected []int8) {
		even := func(el int8) int16 { return int16(el % 2) }
		actual := SliceInt8{given}.DedupByInt16(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, []int8{})
	f([]int8{1}, []int8{1})
	f([]int8{1, 1}, []int8{1})
	f([]int8{1, 2}, []int8{1, 2})
	f([]int8{1, 2, 3}, []int8{1, 2, 3})
	f([]int8{1, 2, 2, 3}, []int8{1, 2, 3})
	f([]int8{1, 2, 4, 3, 5, 7, 10}, []int8{1, 2, 3, 10})
}

func TestSliceDedupByInt8Int32(t *testing.T) {
	f := func(given []int8, expected []int8) {
		even := func(el int8) int32 { return int32(el % 2) }
		actual := SliceInt8{given}.DedupByInt32(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, []int8{})
	f([]int8{1}, []int8{1})
	f([]int8{1, 1}, []int8{1})
	f([]int8{1, 2}, []int8{1, 2})
	f([]int8{1, 2, 3}, []int8{1, 2, 3})
	f([]int8{1, 2, 2, 3}, []int8{1, 2, 3})
	f([]int8{1, 2, 4, 3, 5, 7, 10}, []int8{1, 2, 3, 10})
}

func TestSliceDedupByInt8Int64(t *testing.T) {
	f := func(given []int8, expected []int8) {
		even := func(el int8) int64 { return int64(el % 2) }
		actual := SliceInt8{given}.DedupByInt64(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, []int8{})
	f([]int8{1}, []int8{1})
	f([]int8{1, 1}, []int8{1})
	f([]int8{1, 2}, []int8{1, 2})
	f([]int8{1, 2, 3}, []int8{1, 2, 3})
	f([]int8{1, 2, 2, 3}, []int8{1, 2, 3})
	f([]int8{1, 2, 4, 3, 5, 7, 10}, []int8{1, 2, 3, 10})
}

func TestSliceDeleteInt8(t *testing.T) {
	f := func(given []int8, el int8, expected []int8) {
		actual := SliceInt8{given}.Delete(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, 1, []int8{})
	f([]int8{1}, 1, []int8{})
	f([]int8{2}, 1, []int8{2})
	f([]int8{1, 2}, 1, []int8{2})
	f([]int8{1, 2, 3}, 2, []int8{1, 3})
	f([]int8{1, 2, 2, 3, 2}, 2, []int8{1, 2, 3, 2})
}

func TestSliceDeleteAllInt8(t *testing.T) {
	f := func(given []int8, el int8, expected []int8) {
		actual := SliceInt8{given}.DeleteAll(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, 1, []int8{})
	f([]int8{1}, 1, []int8{})
	f([]int8{2}, 1, []int8{2})
	f([]int8{1, 2}, 1, []int8{2})
	f([]int8{1, 2, 3}, 2, []int8{1, 3})
	f([]int8{1, 2, 2, 3, 2}, 2, []int8{1, 3})
}

func TestSliceDeleteAtInt8(t *testing.T) {
	f := func(given []int8, indices []int, expected []int8) {
		actual, _ := SliceInt8{given}.DeleteAt(indices...)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, []int{}, []int8{})
	f([]int8{1}, []int{0}, []int8{})
	f([]int8{1, 2}, []int{0}, []int8{2})

	f([]int8{1, 2, 3}, []int{0}, []int8{2, 3})
	f([]int8{1, 2, 3}, []int{1}, []int8{1, 3})
	f([]int8{1, 2, 3}, []int{2}, []int8{1, 2})

	f([]int8{1, 2, 3}, []int{0, 1}, []int8{3})
	f([]int8{1, 2, 3}, []int{0, 2}, []int8{2})
	f([]int8{1, 2, 3}, []int{1, 2}, []int8{1})
}

func TestSliceDropEveryInt8(t *testing.T) {
	f := func(given []int8, nth int, from int, expected []int8) {
		actual, _ := SliceInt8{given}.DropEvery(nth, from)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, 1, 1, []int8{})
	f([]int8{1, 2, 3}, 1, 1, []int8{})

	f([]int8{1, 2, 3, 4}, 2, 1, []int8{1, 3})
	f([]int8{1, 2, 3, 4, 5}, 2, 1, []int8{1, 3, 5})

	f([]int8{1, 2, 3, 4}, 2, 0, []int8{2, 4})
	f([]int8{1, 2, 3, 4, 5}, 2, 0, []int8{2, 4})

	f([]int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2, 1, []int8{1, 3, 5, 7, 9})
	f([]int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, 1, []int8{1, 2, 4, 5, 7, 8, 10})
}

func TestSliceDropWhileInt8(t *testing.T) {
	f := func(given []int8, expected []int8) {
		even := func(el int8) bool { return el%2 == 0 }
		actual := SliceInt8{given}.DropWhile(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, []int8{})
	f([]int8{2}, []int8{})
	f([]int8{1}, []int8{1})
	f([]int8{2, 1}, []int8{1})
	f([]int8{2, 1, 2}, []int8{1, 2})
	f([]int8{1, 2}, []int8{1, 2})
	f([]int8{2, 4, 6, 1, 8}, []int8{1, 8})
}

func TestSliceEndsWithInt8(t *testing.T) {
	f := func(given []int8, suffix []int8, expected bool) {
		actual := SliceInt8{given}.EndsWith(suffix)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, []int8{}, true)
	f([]int8{1}, []int8{1}, true)
	f([]int8{1}, []int8{2}, false)
	f([]int8{2, 3}, []int8{1, 2, 3}, false)

	f([]int8{1, 2, 3}, []int8{3}, true)
	f([]int8{1, 2, 3}, []int8{2, 3}, true)
	f([]int8{1, 2, 3}, []int8{1, 2, 3}, true)

	f([]int8{1, 2, 3}, []int8{1}, false)
	f([]int8{1, 2, 3}, []int8{2}, false)
	f([]int8{1, 2, 3}, []int8{1, 2}, false)
	f([]int8{1, 2, 3}, []int8{3, 2}, false)
}

func TestSliceEqualInt8(t *testing.T) {
	f := func(left []int8, right []int8, expected bool) {
		actual := SliceInt8{left}.Equal(right)
		assert.Equal(t, expected, actual, "they should be equal")

		actual = SliceInt8{right}.Equal(left)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, []int8{}, true)
	f([]int8{1}, []int8{1}, true)
	f([]int8{1}, []int8{2}, false)
	f([]int8{1, 2, 3, 3}, []int8{1, 2, 3, 3}, true)
	f([]int8{1, 2, 3, 3}, []int8{1, 2, 2, 3}, false)
	f([]int8{1, 2, 3, 3}, []int8{1, 2, 4, 3}, false)

	// different len
	f([]int8{1, 2, 3}, []int8{1, 2}, false)
	f([]int8{1, 2}, []int8{1, 2, 3}, false)
	f([]int8{}, []int8{1, 2, 3}, false)
	f([]int8{1, 2, 3}, []int8{}, false)
}

func TestSliceFilterInt8(t *testing.T) {
	f := func(given []int8, expected []int8) {
		even := func(t int8) bool { return (t % 2) == 0 }
		actual := SliceInt8{given}.Filter(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, []int8{})
	f([]int8{1, 2, 3, 4}, []int8{2, 4})
	f([]int8{1, 3}, []int8{})
	f([]int8{2, 4}, []int8{2, 4})
}

func TestSliceFindInt8(t *testing.T) {
	f := func(given []int8, expectedEl int8, expectedErr error) {
		even := func(t int8) bool { return (t % 2) == 0 }
		el, err := SliceInt8{given}.Find(even)
		assert.Equal(t, expectedEl, el, "they should be equal")
		assert.Equal(t, expectedErr, err, "they should be equal")
	}
	f([]int8{}, 0, ErrNotFound)
	f([]int8{1}, 0, ErrNotFound)
	f([]int8{1}, 0, ErrNotFound)
	f([]int8{2}, 2, nil)
	f([]int8{1, 2}, 2, nil)
	f([]int8{1, 2, 3}, 2, nil)
	f([]int8{1, 3, 5}, 0, ErrNotFound)
}

func TestSliceFindIndexInt8(t *testing.T) {
	f := func(given []int8, expectedInd int, expectedErr error) {
		even := func(t int8) bool { return (t % 2) == 0 }
		index, err := SliceInt8{given}.FindIndex(even)
		assert.Equal(t, expectedInd, index, "they should be equal")
		assert.Equal(t, expectedErr, err, "they should be equal")
	}
	f([]int8{}, 0, ErrNotFound)
	f([]int8{1}, 0, ErrNotFound)
	f([]int8{1}, 0, ErrNotFound)
	f([]int8{2}, 0, nil)
	f([]int8{1, 2}, 1, nil)
	f([]int8{1, 2, 3}, 1, nil)
	f([]int8{1, 3, 5, 7, 9, 2}, 5, nil)
	f([]int8{1, 3, 5}, 0, ErrNotFound)
}

func TestSliceJoinInt8(t *testing.T) {
	f := func(given []int8, sep string, expected string) {
		actual := SliceInt8{given}.Join(sep)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, "", "")
	f([]int8{}, "|", "")

	f([]int8{1}, "", "1")
	f([]int8{1}, "|", "1")

	f([]int8{1, 2, 3}, "", "123")
	f([]int8{1, 2, 3}, "|", "1|2|3")
	f([]int8{1, 2, 3}, "<int8>", "1<int8>2<int8>3")
}

func TestSliceGroupByInt8Int(t *testing.T) {
	f := func(given []int8, expected map[int][]int8) {
		reminder := func(t int8) int { return int((t % 2)) }
		actual := SliceInt8{given}.GroupByInt(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, map[int][]int8{})
	f([]int8{1}, map[int][]int8{1: {1}})
	f([]int8{1, 3, 2, 4, 5}, map[int][]int8{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceGroupByInt8Int8(t *testing.T) {
	f := func(given []int8, expected map[int8][]int8) {
		reminder := func(t int8) int8 { return int8((t % 2)) }
		actual := SliceInt8{given}.GroupByInt8(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, map[int8][]int8{})
	f([]int8{1}, map[int8][]int8{1: {1}})
	f([]int8{1, 3, 2, 4, 5}, map[int8][]int8{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceGroupByInt8Int16(t *testing.T) {
	f := func(given []int8, expected map[int16][]int8) {
		reminder := func(t int8) int16 { return int16((t % 2)) }
		actual := SliceInt8{given}.GroupByInt16(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, map[int16][]int8{})
	f([]int8{1}, map[int16][]int8{1: {1}})
	f([]int8{1, 3, 2, 4, 5}, map[int16][]int8{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceGroupByInt8Int32(t *testing.T) {
	f := func(given []int8, expected map[int32][]int8) {
		reminder := func(t int8) int32 { return int32((t % 2)) }
		actual := SliceInt8{given}.GroupByInt32(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, map[int32][]int8{})
	f([]int8{1}, map[int32][]int8{1: {1}})
	f([]int8{1, 3, 2, 4, 5}, map[int32][]int8{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceGroupByInt8Int64(t *testing.T) {
	f := func(given []int8, expected map[int64][]int8) {
		reminder := func(t int8) int64 { return int64((t % 2)) }
		actual := SliceInt8{given}.GroupByInt64(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, map[int64][]int8{})
	f([]int8{1}, map[int64][]int8{1: {1}})
	f([]int8{1, 3, 2, 4, 5}, map[int64][]int8{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceInsertAtInt8(t *testing.T) {
	f := func(given []int8, index int, expected []int8, expectedErr error) {
		actual, err := SliceInt8{given}.InsertAt(index, 10)
		assert.Equal(t, expected, actual, "they should be equal")
		assert.Equal(t, expectedErr, err, "they should be equal")
	}
	f([]int8{}, -1, []int8{}, ErrNegativeValue)
	f([]int8{}, 0, []int8{10}, nil)
	f([]int8{}, 1, []int8{}, ErrOutOfRange)

	f([]int8{1, 2, 3}, -1, []int8{1, 2, 3}, ErrNegativeValue)
	f([]int8{1, 2, 3}, 0, []int8{10, 1, 2, 3}, nil)
	f([]int8{1, 2, 3}, 1, []int8{1, 10, 2, 3}, nil)
	f([]int8{1, 2, 3}, 3, []int8{1, 2, 3, 10}, nil)
	f([]int8{1, 2, 3}, 4, []int8{1, 2, 3}, ErrOutOfRange)
}

func TestSliceIntersperseInt8(t *testing.T) {
	f := func(el int8, given []int8, expected []int8) {
		actual := SliceInt8{given}.Intersperse(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, []int8{}, []int8{})
	f(0, []int8{1}, []int8{1})
	f(0, []int8{1, 2}, []int8{1, 0, 2})
	f(0, []int8{1, 2, 3}, []int8{1, 0, 2, 0, 3})
}

func TestSliceLastInt8(t *testing.T) {
	f := func(given []int8, expectedEl int8, expectedErr error) {
		el, err := SliceInt8{given}.Last()
		assert.Equal(t, expectedEl, el, "they should be equal")
		assert.Equal(t, expectedErr, err, "they should be equal")
	}
	f([]int8{}, 0, ErrEmpty)
	f([]int8{1}, 1, nil)
	f([]int8{1, 2, 3}, 3, nil)
}

func TestSliceMapInt8Int(t *testing.T) {
	f := func(given []int8, expected []int) {
		double := func(t int8) int { return int((t * 2)) }
		actual := SliceInt8{given}.MapInt(double)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, []int{})
	f([]int8{1}, []int{2})
	f([]int8{1, 2, 3}, []int{2, 4, 6})
}

func TestSliceMapInt8Int8(t *testing.T) {
	f := func(given []int8, expected []int8) {
		double := func(t int8) int8 { return int8((t * 2)) }
		actual := SliceInt8{given}.MapInt8(double)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, []int8{})
	f([]int8{1}, []int8{2})
	f([]int8{1, 2, 3}, []int8{2, 4, 6})
}

func TestSliceMapInt8Int16(t *testing.T) {
	f := func(given []int8, expected []int16) {
		double := func(t int8) int16 { return int16((t * 2)) }
		actual := SliceInt8{given}.MapInt16(double)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, []int16{})
	f([]int8{1}, []int16{2})
	f([]int8{1, 2, 3}, []int16{2, 4, 6})
}

func TestSliceMapInt8Int32(t *testing.T) {
	f := func(given []int8, expected []int32) {
		double := func(t int8) int32 { return int32((t * 2)) }
		actual := SliceInt8{given}.MapInt32(double)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, []int32{})
	f([]int8{1}, []int32{2})
	f([]int8{1, 2, 3}, []int32{2, 4, 6})
}

func TestSliceMapInt8Int64(t *testing.T) {
	f := func(given []int8, expected []int64) {
		double := func(t int8) int64 { return int64((t * 2)) }
		actual := SliceInt8{given}.MapInt64(double)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, []int64{})
	f([]int8{1}, []int64{2})
	f([]int8{1, 2, 3}, []int64{2, 4, 6})
}

func TestSliceMaxInt8(t *testing.T) {
	f := func(given []int8, expectedEl int8, expectedErr error) {
		el, err := SliceInt8{given}.Max()
		assert.Equal(t, expectedEl, el, "they should be equal")
		assert.Equal(t, expectedErr, err, "they should be equal")
	}
	f([]int8{}, 0, ErrEmpty)
	f([]int8{1}, 1, nil)
	f([]int8{1, 2, 3}, 3, nil)
	f([]int8{1, 3, 2}, 3, nil)
	f([]int8{3, 2, 1}, 3, nil)
}

func TestSliceMinInt8(t *testing.T) {
	f := func(given []int8, expectedEl int8, expectedErr error) {
		el, err := SliceInt8{given}.Min()
		assert.Equal(t, expectedEl, el, "they should be equal")
		assert.Equal(t, expectedErr, err, "they should be equal")
	}
	f([]int8{}, 0, ErrEmpty)
	f([]int8{1}, 1, nil)
	f([]int8{1, 2, 3}, 1, nil)
	f([]int8{2, 1, 3}, 1, nil)
	f([]int8{3, 2, 1}, 1, nil)
}

func TestSlicesPermutationsInt8(t *testing.T) {
	f := func(size int, given []int8, expected [][]int8) {
		actual := make([][]int8, 0)
		i := 0
		s := SliceInt8{given}
		for el := range s.Permutations(size) {
			actual = append(actual, el)
			i++
			if i > 50 {
				t.Fatal("infinite loop")
			}
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, []int8{}, [][]int8{})
	f(2, []int8{1}, [][]int8{{1}})
	f(2, []int8{1, 2, 3}, [][]int8{{1, 2}, {1, 3}, {2, 1}, {2, 3}, {3, 1}, {3, 2}})
}

func TestSliceProductInt8(t *testing.T) {
	f := func(given []int8, repeat int, expected [][]int8) {
		actual := make([][]int8, 0)
		i := 0
		s := SliceInt8{given}
		for el := range s.Product(repeat) {
			actual = append(actual, el)
			i++
			if i > 50 {
				t.Fatal("infinite loop")
			}
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}

	f([]int8{1, 2}, 0, [][]int8{})
	f([]int8{}, 2, [][]int8{})
	f([]int8{1}, 2, [][]int8{{1, 1}})

	f([]int8{1, 2}, 1, [][]int8{{1}, {2}})
	f([]int8{1, 2}, 2, [][]int8{{1, 1}, {1, 2}, {2, 1}, {2, 2}})
	f([]int8{1, 2}, 3, [][]int8{
		{1, 1, 1}, {1, 1, 2}, {1, 2, 1}, {1, 2, 2},
		{2, 1, 1}, {2, 1, 2}, {2, 2, 1}, {2, 2, 2},
	})
}

func TestSliceReduceInt8Int(t *testing.T) {
	f := func(given []int8, expected int) {
		sum := func(el int8, acc int) int { return int(el) + acc }
		actual := SliceInt8{given}.ReduceInt(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, 0)
	f([]int8{1}, 1)
	f([]int8{1, 2}, 3)
	f([]int8{1, 2, 3}, 6)
}

func TestSliceReduceInt8Int8(t *testing.T) {
	f := func(given []int8, expected int8) {
		sum := func(el int8, acc int8) int8 { return int8(el) + acc }
		actual := SliceInt8{given}.ReduceInt8(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, 0)
	f([]int8{1}, 1)
	f([]int8{1, 2}, 3)
	f([]int8{1, 2, 3}, 6)
}

func TestSliceReduceInt8Int16(t *testing.T) {
	f := func(given []int8, expected int16) {
		sum := func(el int8, acc int16) int16 { return int16(el) + acc }
		actual := SliceInt8{given}.ReduceInt16(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, 0)
	f([]int8{1}, 1)
	f([]int8{1, 2}, 3)
	f([]int8{1, 2, 3}, 6)
}

func TestSliceReduceInt8Int32(t *testing.T) {
	f := func(given []int8, expected int32) {
		sum := func(el int8, acc int32) int32 { return int32(el) + acc }
		actual := SliceInt8{given}.ReduceInt32(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, 0)
	f([]int8{1}, 1)
	f([]int8{1, 2}, 3)
	f([]int8{1, 2, 3}, 6)
}

func TestSliceReduceInt8Int64(t *testing.T) {
	f := func(given []int8, expected int64) {
		sum := func(el int8, acc int64) int64 { return int64(el) + acc }
		actual := SliceInt8{given}.ReduceInt64(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, 0)
	f([]int8{1}, 1)
	f([]int8{1, 2}, 3)
	f([]int8{1, 2, 3}, 6)
}

func TestSliceReduceWhileInt8Int(t *testing.T) {
	f := func(given []int8, expected int) {
		sum := func(el int8, acc int) (int, error) {
			if el == 0 {
				return acc, ErrEmpty
			}
			return int(el) + acc, nil
		}
		actual, _ := SliceInt8{given}.ReduceWhileInt(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, 0)
	f([]int8{1}, 1)
	f([]int8{1, 2}, 3)
	f([]int8{1, 2, 3}, 6)
	f([]int8{1, 2, 0, 3}, 3)
}

func TestSliceReduceWhileInt8Int8(t *testing.T) {
	f := func(given []int8, expected int8) {
		sum := func(el int8, acc int8) (int8, error) {
			if el == 0 {
				return acc, ErrEmpty
			}
			return int8(el) + acc, nil
		}
		actual, _ := SliceInt8{given}.ReduceWhileInt8(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, 0)
	f([]int8{1}, 1)
	f([]int8{1, 2}, 3)
	f([]int8{1, 2, 3}, 6)
	f([]int8{1, 2, 0, 3}, 3)
}

func TestSliceReduceWhileInt8Int16(t *testing.T) {
	f := func(given []int8, expected int16) {
		sum := func(el int8, acc int16) (int16, error) {
			if el == 0 {
				return acc, ErrEmpty
			}
			return int16(el) + acc, nil
		}
		actual, _ := SliceInt8{given}.ReduceWhileInt16(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, 0)
	f([]int8{1}, 1)
	f([]int8{1, 2}, 3)
	f([]int8{1, 2, 3}, 6)
	f([]int8{1, 2, 0, 3}, 3)
}

func TestSliceReduceWhileInt8Int32(t *testing.T) {
	f := func(given []int8, expected int32) {
		sum := func(el int8, acc int32) (int32, error) {
			if el == 0 {
				return acc, ErrEmpty
			}
			return int32(el) + acc, nil
		}
		actual, _ := SliceInt8{given}.ReduceWhileInt32(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, 0)
	f([]int8{1}, 1)
	f([]int8{1, 2}, 3)
	f([]int8{1, 2, 3}, 6)
	f([]int8{1, 2, 0, 3}, 3)
}

func TestSliceReduceWhileInt8Int64(t *testing.T) {
	f := func(given []int8, expected int64) {
		sum := func(el int8, acc int64) (int64, error) {
			if el == 0 {
				return acc, ErrEmpty
			}
			return int64(el) + acc, nil
		}
		actual, _ := SliceInt8{given}.ReduceWhileInt64(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, 0)
	f([]int8{1}, 1)
	f([]int8{1, 2}, 3)
	f([]int8{1, 2, 3}, 6)
	f([]int8{1, 2, 0, 3}, 3)
}

func TestChannelToSliceInt8(t *testing.T) {
	f := func(given []int8) {
		c := make(chan int8, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := ChannelInt8{c}.ToSlice()
		assert.Equal(t, given, actual, "they should be equal")
	}
	f([]int8{})
	f([]int8{1})
	f([]int8{1, 2, 3, 1, 2})
}

func TestChannelAnyInt8(t *testing.T) {
	f := func(given []int8, expected bool) {
		even := func(t int8) bool { return t%2 == 0 }
		c := make(chan int8, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := ChannelInt8{c}.Any(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, false)
	f([]int8{1}, false)
	f([]int8{2}, true)
	f([]int8{1, 2}, true)
	f([]int8{1, 2, 3}, true)
	f([]int8{1, 3, 5}, false)
	f([]int8{1, 3, 5, 7, 9, 11}, false)
	f([]int8{1, 3, 5, 7, 10, 11}, true)
}

func TestChannelAllInt8(t *testing.T) {
	f := func(given []int8, expected bool) {
		even := func(t int8) bool { return t%2 == 0 }
		c := make(chan int8, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := ChannelInt8{c}.All(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, true)
	f([]int8{1}, false)
	f([]int8{2}, true)
	f([]int8{1, 2}, false)
	f([]int8{2, 4}, true)
	f([]int8{2, 4, 6, 8, 10, 12}, true)
	f([]int8{2, 4, 6, 8, 11, 12}, false)
}

func TestChannelEachInt8(t *testing.T) {
	f := func(given []int8) {
		c := make(chan int8, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := make(chan int8, len(given))
		mapper := func(t int8) { result <- t }
		ChannelInt8{c}.Each(mapper)
		close(result)
		actual := ChannelInt8{result}.ToSlice()
		assert.Equal(t, given, actual, "they should be equal")
	}

	f([]int8{})
	f([]int8{1})
	f([]int8{1, 2, 3})
	f([]int8{1, 2, 3, 4, 5, 6, 7})
}

func TestChannelChunkEveryInt8(t *testing.T) {
	f := func(size int, given []int8, expected [][]int8) {
		c := make(chan int8, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := ChannelInt8{c}.ChunkEvery(size)
		actual := make([][]int8, 0)
		for el := range result {
			actual = append(actual, el)
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, []int8{}, [][]int8{})
	f(2, []int8{1}, [][]int8{{1}})
	f(2, []int8{1, 2}, [][]int8{{1, 2}})
	f(2, []int8{1, 2, 3, 4}, [][]int8{{1, 2}, {3, 4}})
	f(2, []int8{1, 2, 3, 4, 5}, [][]int8{{1, 2}, {3, 4}, {5}})
}

func TestChannelCountInt8(t *testing.T) {
	f := func(element int8, given []int8, expected int) {
		c := make(chan int8, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := ChannelInt8{c}.Count(element)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, []int8{}, 0)
	f(1, []int8{1}, 1)
	f(1, []int8{2}, 0)
	f(1, []int8{1, 2, 3, 1, 4}, 2)
}

func TestChannelDropInt8(t *testing.T) {
	f := func(count int, given []int8, expected []int8) {
		c := make(chan int8, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := ChannelInt8{c}.Drop(count)
		actual := make([]int8, 0)
		for el := range result {
			actual = append(actual, el)
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, []int8{}, []int8{})
	f(1, []int8{2}, []int8{})
	f(1, []int8{2, 3}, []int8{3})
	f(1, []int8{1, 2, 3}, []int8{2, 3})
	f(0, []int8{1, 2, 3}, []int8{1, 2, 3})
	f(3, []int8{1, 2, 3, 4, 5, 6}, []int8{4, 5, 6})
	f(1, []int8{1, 2, 3, 4, 5, 6}, []int8{2, 3, 4, 5, 6})
}

func TestChannelFilterInt8(t *testing.T) {
	f := func(given []int8, expected []int8) {
		even := func(t int8) bool { return t%2 == 0 }
		c := make(chan int8, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := ChannelInt8{c}.Filter(even)
		actual := ChannelInt8{result}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, []int8{})
	f([]int8{1}, []int8{})
	f([]int8{2}, []int8{2})
	f([]int8{1, 2, 3, 4}, []int8{2, 4})
}

func TestChannelMapInt8Int(t *testing.T) {
	f := func(given []int8, expected []int8) {
		double := func(el int8) int { return int(el * 2) }
		c := make(chan int8, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := ChannelInt8{c}.MapInt(double)

		// convert chan int8 to chan int
		c2 := make(chan int8, 1)
		go func() {
			for el := range result {
				c2 <- int8(el)
			}
			close(c2)
		}()

		actual := ChannelInt8{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, []int8{})
	f([]int8{1}, []int8{2})
	f([]int8{1, 2, 3}, []int8{2, 4, 6})
}

func TestChannelMapInt8Int8(t *testing.T) {
	f := func(given []int8, expected []int8) {
		double := func(el int8) int8 { return int8(el * 2) }
		c := make(chan int8, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := ChannelInt8{c}.MapInt8(double)

		// convert chan int8 to chan int8
		c2 := make(chan int8, 1)
		go func() {
			for el := range result {
				c2 <- int8(el)
			}
			close(c2)
		}()

		actual := ChannelInt8{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, []int8{})
	f([]int8{1}, []int8{2})
	f([]int8{1, 2, 3}, []int8{2, 4, 6})
}

func TestChannelMapInt8Int16(t *testing.T) {
	f := func(given []int8, expected []int8) {
		double := func(el int8) int16 { return int16(el * 2) }
		c := make(chan int8, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := ChannelInt8{c}.MapInt16(double)

		// convert chan int8 to chan int16
		c2 := make(chan int8, 1)
		go func() {
			for el := range result {
				c2 <- int8(el)
			}
			close(c2)
		}()

		actual := ChannelInt8{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, []int8{})
	f([]int8{1}, []int8{2})
	f([]int8{1, 2, 3}, []int8{2, 4, 6})
}

func TestChannelMapInt8Int32(t *testing.T) {
	f := func(given []int8, expected []int8) {
		double := func(el int8) int32 { return int32(el * 2) }
		c := make(chan int8, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := ChannelInt8{c}.MapInt32(double)

		// convert chan int8 to chan int32
		c2 := make(chan int8, 1)
		go func() {
			for el := range result {
				c2 <- int8(el)
			}
			close(c2)
		}()

		actual := ChannelInt8{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, []int8{})
	f([]int8{1}, []int8{2})
	f([]int8{1, 2, 3}, []int8{2, 4, 6})
}

func TestChannelMapInt8Int64(t *testing.T) {
	f := func(given []int8, expected []int8) {
		double := func(el int8) int64 { return int64(el * 2) }
		c := make(chan int8, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := ChannelInt8{c}.MapInt64(double)

		// convert chan int8 to chan int64
		c2 := make(chan int8, 1)
		go func() {
			for el := range result {
				c2 <- int8(el)
			}
			close(c2)
		}()

		actual := ChannelInt8{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, []int8{})
	f([]int8{1}, []int8{2})
	f([]int8{1, 2, 3}, []int8{2, 4, 6})
}

func TestChannelMaxInt8(t *testing.T) {
	f := func(given []int8, expected int8, expectedErr error) {
		c := make(chan int8, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual, actualErr := ChannelInt8{c}.Max()
		assert.Equal(t, expected, actual, "they should be equal")
		assert.Equal(t, expectedErr, actualErr, "they should be equal")
	}
	f([]int8{}, 0, ErrEmpty)
	f([]int8{1, 4, 2}, 4, nil)
	f([]int8{1, 2, 4}, 4, nil)
	f([]int8{4, 2, 1}, 4, nil)
}

func TestChannelMinInt8(t *testing.T) {
	f := func(given []int8, expected int8, expectedErr error) {
		c := make(chan int8, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual, actualErr := ChannelInt8{c}.Min()
		assert.Equal(t, expected, actual, "they should be equal")
		assert.Equal(t, expectedErr, actualErr, "they should be equal")
	}
	f([]int8{}, 0, ErrEmpty)
	f([]int8{4, 1, 2}, 1, nil)
	f([]int8{1, 2, 4}, 1, nil)
	f([]int8{4, 2, 1}, 1, nil)
}

func TestChannelReduceInt8Int(t *testing.T) {
	f := func(given []int8, expected int) {
		c := make(chan int8, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int8, acc int) int { return int(el) + acc }
		actual := ChannelInt8{c}.ReduceInt(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, 0)
	f([]int8{1}, 1)
	f([]int8{1, 2}, 3)
	f([]int8{1, 2, 3, 4, 5}, 15)
}

func TestChannelReduceInt8Int8(t *testing.T) {
	f := func(given []int8, expected int8) {
		c := make(chan int8, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int8, acc int8) int8 { return int8(el) + acc }
		actual := ChannelInt8{c}.ReduceInt8(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, 0)
	f([]int8{1}, 1)
	f([]int8{1, 2}, 3)
	f([]int8{1, 2, 3, 4, 5}, 15)
}

func TestChannelReduceInt8Int16(t *testing.T) {
	f := func(given []int8, expected int16) {
		c := make(chan int8, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int8, acc int16) int16 { return int16(el) + acc }
		actual := ChannelInt8{c}.ReduceInt16(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, 0)
	f([]int8{1}, 1)
	f([]int8{1, 2}, 3)
	f([]int8{1, 2, 3, 4, 5}, 15)
}

func TestChannelReduceInt8Int32(t *testing.T) {
	f := func(given []int8, expected int32) {
		c := make(chan int8, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int8, acc int32) int32 { return int32(el) + acc }
		actual := ChannelInt8{c}.ReduceInt32(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, 0)
	f([]int8{1}, 1)
	f([]int8{1, 2}, 3)
	f([]int8{1, 2, 3, 4, 5}, 15)
}

func TestChannelReduceInt8Int64(t *testing.T) {
	f := func(given []int8, expected int64) {
		c := make(chan int8, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int8, acc int64) int64 { return int64(el) + acc }
		actual := ChannelInt8{c}.ReduceInt64(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, 0)
	f([]int8{1}, 1)
	f([]int8{1, 2}, 3)
	f([]int8{1, 2, 3, 4, 5}, 15)
}

func TestChannelScanInt8Int(t *testing.T) {
	f := func(given []int8, expected []int8) {
		c := make(chan int8, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int8, acc int) int { return int(el) + acc }
		result := ChannelInt8{c}.ScanInt(0, sum)

		// convert chan int8 to chan int
		c2 := make(chan int8, 1)
		go func() {
			for el := range result {
				c2 <- int8(el)
			}
			close(c2)
		}()

		actual := ChannelInt8{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, []int8{})
	f([]int8{1}, []int8{1})
	f([]int8{1, 2}, []int8{1, 3})
	f([]int8{1, 2, 3, 4, 5}, []int8{1, 3, 6, 10, 15})
}

func TestChannelScanInt8Int8(t *testing.T) {
	f := func(given []int8, expected []int8) {
		c := make(chan int8, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int8, acc int8) int8 { return int8(el) + acc }
		result := ChannelInt8{c}.ScanInt8(0, sum)

		// convert chan int8 to chan int8
		c2 := make(chan int8, 1)
		go func() {
			for el := range result {
				c2 <- int8(el)
			}
			close(c2)
		}()

		actual := ChannelInt8{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, []int8{})
	f([]int8{1}, []int8{1})
	f([]int8{1, 2}, []int8{1, 3})
	f([]int8{1, 2, 3, 4, 5}, []int8{1, 3, 6, 10, 15})
}

func TestChannelScanInt8Int16(t *testing.T) {
	f := func(given []int8, expected []int8) {
		c := make(chan int8, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int8, acc int16) int16 { return int16(el) + acc }
		result := ChannelInt8{c}.ScanInt16(0, sum)

		// convert chan int8 to chan int16
		c2 := make(chan int8, 1)
		go func() {
			for el := range result {
				c2 <- int8(el)
			}
			close(c2)
		}()

		actual := ChannelInt8{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, []int8{})
	f([]int8{1}, []int8{1})
	f([]int8{1, 2}, []int8{1, 3})
	f([]int8{1, 2, 3, 4, 5}, []int8{1, 3, 6, 10, 15})
}

func TestChannelScanInt8Int32(t *testing.T) {
	f := func(given []int8, expected []int8) {
		c := make(chan int8, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int8, acc int32) int32 { return int32(el) + acc }
		result := ChannelInt8{c}.ScanInt32(0, sum)

		// convert chan int8 to chan int32
		c2 := make(chan int8, 1)
		go func() {
			for el := range result {
				c2 <- int8(el)
			}
			close(c2)
		}()

		actual := ChannelInt8{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, []int8{})
	f([]int8{1}, []int8{1})
	f([]int8{1, 2}, []int8{1, 3})
	f([]int8{1, 2, 3, 4, 5}, []int8{1, 3, 6, 10, 15})
}

func TestChannelScanInt8Int64(t *testing.T) {
	f := func(given []int8, expected []int8) {
		c := make(chan int8, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int8, acc int64) int64 { return int64(el) + acc }
		result := ChannelInt8{c}.ScanInt64(0, sum)

		// convert chan int8 to chan int64
		c2 := make(chan int8, 1)
		go func() {
			for el := range result {
				c2 <- int8(el)
			}
			close(c2)
		}()

		actual := ChannelInt8{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, []int8{})
	f([]int8{1}, []int8{1})
	f([]int8{1, 2}, []int8{1, 3})
	f([]int8{1, 2, 3, 4, 5}, []int8{1, 3, 6, 10, 15})
}

func TestChannelSumInt8(t *testing.T) {
	f := func(given []int8, expected int8) {
		c := make(chan int8, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := ChannelInt8{c}.Sum()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int8{}, 0)
	f([]int8{1}, 1)
	f([]int8{1, 2}, 3)
	f([]int8{1, 2, 3, 4, 5}, 15)
}

func TestChannelTakeInt8(t *testing.T) {
	f := func(count int, given int8, expected []int8) {
		ctx, cancel := context.WithCancel(context.Background())
		s := SequenceInt8{ctx: ctx}
		seq := s.Repeat(given)
		seq2 := ChannelInt8{seq}.Take(count)
		actual := ChannelInt8{seq2}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, 1, []int8{})
	f(1, 1, []int8{1})
	f(2, 1, []int8{1, 1})
}

func TestChannelTeeInt8(t *testing.T) {
	f := func(count int, given []int8) {
		c := make(chan int8, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		channels := ChannelInt8{c}.Tee(count)
		for _, ch := range channels {
			go func(ch chan int8) {
				actual := ChannelInt8{ch}.ToSlice()
				assert.Equal(t, given, actual, "they should be equal")
			}(ch)
		}
	}
	f(1, []int8{})
	f(1, []int8{1})
	f(1, []int8{1, 2})
	f(1, []int8{1, 2, 3})
	f(1, []int8{1, 2, 3, 1, 2})

	f(2, []int8{})
	f(2, []int8{1})
	f(2, []int8{1, 2})
	f(2, []int8{1, 2, 3})
	f(2, []int8{1, 2, 3, 1, 2})

	f(10, []int8{1, 2, 3, 1, 2})
}

func TestAsyncSliceAnyInt8(t *testing.T) {
	f := func(check func(t int8) bool, given []int8, expected bool) {
		s := AsyncSliceInt8{Data: given, Workers: 2}
		actual := s.Any(check)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	isEven := func(t int8) bool { return (t % 2) == 0 }

	f(isEven, []int8{}, false)
	f(isEven, []int8{1}, false)
	f(isEven, []int8{1, 3}, false)
	f(isEven, []int8{2}, true)
	f(isEven, []int8{1, 2}, true)
	f(isEven, []int8{1, 3, 5, 7, 9, 11}, false)
	f(isEven, []int8{1, 3, 5, 7, 9, 12}, true)
}

func TestAsyncSliceAllInt8(t *testing.T) {
	f := func(check func(t int8) bool, given []int8, expected bool) {
		s := AsyncSliceInt8{Data: given, Workers: 2}
		actual := s.All(check)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	isEven := func(t int8) bool { return (t % 2) == 0 }

	f(isEven, []int8{}, true)
	f(isEven, []int8{1}, false)
	f(isEven, []int8{1, 3}, false)
	f(isEven, []int8{2}, true)
	f(isEven, []int8{2, 4}, true)
	f(isEven, []int8{2, 3}, false)
	f(isEven, []int8{2, 4, 6, 8, 10, 12}, true)
	f(isEven, []int8{2, 4, 6, 8, 10, 11}, false)
}

func TestAsyncSliceEachInt8(t *testing.T) {
	f := func(given []int8) {
		s := AsyncSliceInt8{Data: given, Workers: 2}
		result := make(chan int8, len(given))
		mapper := func(t int8) { result <- t }
		s.Each(mapper)
		close(result)
		actual := ChannelInt8{result}.ToSlice()
		sorted := SliceInt8{actual}.Sort()
		assert.Equal(t, given, sorted, "they should be equal")
	}

	f([]int8{})
	f([]int8{1})
	f([]int8{1, 2, 3})
	f([]int8{1, 2, 3, 4, 5, 6, 7})
}

func TestAsyncSliceFilterInt8(t *testing.T) {
	f := func(given []int8, expected []int8) {
		filter := func(t int8) bool { return t > 10 }
		s := AsyncSliceInt8{Data: given, Workers: 2}
		actual := s.Filter(filter)
		assert.Equal(t, expected, actual, "they should be equal")
	}

	f([]int8{}, []int8{})
	f([]int8{5}, []int8{})
	f([]int8{15}, []int8{15})
	f([]int8{9, 11, 12, 13, 6}, []int8{11, 12, 13})
}

func TestAsyncSliceMapInt8Int(t *testing.T) {
	f := func(mapper func(t int8) int, given []int8, expected []int) {
		s := AsyncSliceInt8{Data: given, Workers: 2}
		actual := s.MapInt(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int8) int { return int((t * 2)) }

	f(double, []int8{}, []int{})
	f(double, []int8{1}, []int{2})
	f(double, []int8{1, 2, 3}, []int{2, 4, 6})
}

func TestAsyncSliceMapInt8Int8(t *testing.T) {
	f := func(mapper func(t int8) int8, given []int8, expected []int8) {
		s := AsyncSliceInt8{Data: given, Workers: 2}
		actual := s.MapInt8(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int8) int8 { return int8((t * 2)) }

	f(double, []int8{}, []int8{})
	f(double, []int8{1}, []int8{2})
	f(double, []int8{1, 2, 3}, []int8{2, 4, 6})
}

func TestAsyncSliceMapInt8Int16(t *testing.T) {
	f := func(mapper func(t int8) int16, given []int8, expected []int16) {
		s := AsyncSliceInt8{Data: given, Workers: 2}
		actual := s.MapInt16(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int8) int16 { return int16((t * 2)) }

	f(double, []int8{}, []int16{})
	f(double, []int8{1}, []int16{2})
	f(double, []int8{1, 2, 3}, []int16{2, 4, 6})
}

func TestAsyncSliceMapInt8Int32(t *testing.T) {
	f := func(mapper func(t int8) int32, given []int8, expected []int32) {
		s := AsyncSliceInt8{Data: given, Workers: 2}
		actual := s.MapInt32(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int8) int32 { return int32((t * 2)) }

	f(double, []int8{}, []int32{})
	f(double, []int8{1}, []int32{2})
	f(double, []int8{1, 2, 3}, []int32{2, 4, 6})
}

func TestAsyncSliceMapInt8Int64(t *testing.T) {
	f := func(mapper func(t int8) int64, given []int8, expected []int64) {
		s := AsyncSliceInt8{Data: given, Workers: 2}
		actual := s.MapInt64(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int8) int64 { return int64((t * 2)) }

	f(double, []int8{}, []int64{})
	f(double, []int8{1}, []int64{2})
	f(double, []int8{1, 2, 3}, []int64{2, 4, 6})
}

func TestAsyncSliceReduceInt8(t *testing.T) {
	f := func(reducer func(a int8, b int8) int8, given []int8, expected int8) {
		s := AsyncSliceInt8{Data: given, Workers: 4}
		actual := s.Reduce(reducer)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	sum := func(a int8, b int8) int8 { return a + b }

	f(sum, []int8{}, 0)
	f(sum, []int8{1}, 1)
	f(sum, []int8{1, 2}, 3)
	f(sum, []int8{1, 2, 3}, 6)
	f(sum, []int8{1, 2, 3, 4}, 10)
	f(sum, []int8{1, 2, 3, 4, 5}, 15)
}

func TestSlicesProductInt16(t *testing.T) {
	f := func(given [][]int16, expected [][]int16) {
		actual := make([][]int16, 0)
		i := 0
		s := SlicesInt16{given}
		for el := range s.Product() {
			actual = append(actual, el)
			i++
			if i > 50 {
				t.Fatal("infinite loop")
			}
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([][]int16{{1, 2}, {3, 4}}, [][]int16{{1, 3}, {1, 4}, {2, 3}, {2, 4}})
	f([][]int16{{1, 2}, {3}, {4, 5}}, [][]int16{{1, 3, 4}, {1, 3, 5}, {2, 3, 4}, {2, 3, 5}})
}

func TestSequenceCountInt16(t *testing.T) {
	f := func(start int16, step int16, count int, expected []int16) {
		ctx, cancel := context.WithCancel(context.Background())
		s := SequenceInt16{ctx: ctx}
		seq := s.Count(start, step)
		seq2 := ChannelInt16{seq}.Take(count)
		actual := ChannelInt16{seq2}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 2, 4, []int16{1, 3, 5, 7})
}

func TestSequenceExponentialInt16(t *testing.T) {
	f := func(start int16, factor int16, count int, expected []int16) {
		ctx, cancel := context.WithCancel(context.Background())
		s := SequenceInt16{ctx: ctx}
		seq := s.Exponential(start, factor)
		seq2 := ChannelInt16{seq}.Take(count)
		actual := ChannelInt16{seq2}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 1, 4, []int16{1, 1, 1, 1})
	f(1, 2, 4, []int16{1, 2, 4, 8})
}

func TestSequenceIterateInt16(t *testing.T) {
	f := func(start int16, count int, expected []int16) {
		ctx, cancel := context.WithCancel(context.Background())
		s := SequenceInt16{ctx: ctx}
		double := func(val int16) int16 { return val * 2 }
		seq := s.Iterate(start, double)
		seq2 := ChannelInt16{seq}.Take(count)
		actual := ChannelInt16{seq2}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 4, []int16{1, 2, 4, 8})
}

func TestSequenceRangeInt16(t *testing.T) {
	f := func(start int16, stop int16, step int16, expected []int16) {
		ctx, cancel := context.WithCancel(context.Background())
		s := SequenceInt16{ctx: ctx}
		seq := s.Range(start, stop, step)
		actual := ChannelInt16{seq}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 4, 1, []int16{1, 2, 3})
	f(3, 0, -1, []int16{3, 2, 1})
	f(1, 1, 1, []int16{})
	f(1, 2, 1, []int16{1})
}

func TestSequenceRepeatInt16(t *testing.T) {
	f := func(count int, given int16, expected []int16) {
		ctx, cancel := context.WithCancel(context.Background())
		s := SequenceInt16{ctx: ctx}
		seq := s.Repeat(given)
		seq2 := ChannelInt16{seq}.Take(count)
		actual := ChannelInt16{seq2}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, 1, []int16{1, 1})
}

func TestSequenceReplicateInt16(t *testing.T) {
	f := func(count int, given int16, expected []int16) {
		ctx, cancel := context.WithCancel(context.Background())
		s := SequenceInt16{ctx: ctx}
		seq := s.Replicate(given, count)
		actual := ChannelInt16{seq}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, 1, []int16{})
	f(1, 1, []int16{1})
	f(5, 1, []int16{1, 1, 1, 1, 1})
}

func TestSliceAnyInt16(t *testing.T) {
	f := func(given []int16, expected bool) {
		even := func(t int16) bool { return (t % 2) == 0 }
		actual := SliceInt16{given}.Any(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, false)
	f([]int16{1, 3}, false)
	f([]int16{2}, true)
	f([]int16{1, 2}, true)
}

func TestSliceAllInt16(t *testing.T) {
	f := func(given []int16, expected bool) {
		even := func(t int16) bool { return (t % 2) == 0 }
		actual := SliceInt16{given}.All(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, true)
	f([]int16{2}, true)
	f([]int16{1}, false)
	f([]int16{2, 4}, true)
	f([]int16{2, 4, 1}, false)
	f([]int16{1, 2, 4}, false)
}

func TestSliceChunkByInt16Int(t *testing.T) {
	f := func(given []int16, expected [][]int16) {
		reminder := func(t int16) int { return int((t % 2)) }
		actual := SliceInt16{given}.ChunkByInt(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, [][]int16{})
	f([]int16{1}, [][]int16{{1}})
	f([]int16{1, 2, 3}, [][]int16{{1}, {2}, {3}})
	f([]int16{1, 3, 2, 4, 5}, [][]int16{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkByInt16Int8(t *testing.T) {
	f := func(given []int16, expected [][]int16) {
		reminder := func(t int16) int8 { return int8((t % 2)) }
		actual := SliceInt16{given}.ChunkByInt8(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, [][]int16{})
	f([]int16{1}, [][]int16{{1}})
	f([]int16{1, 2, 3}, [][]int16{{1}, {2}, {3}})
	f([]int16{1, 3, 2, 4, 5}, [][]int16{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkByInt16Int16(t *testing.T) {
	f := func(given []int16, expected [][]int16) {
		reminder := func(t int16) int16 { return int16((t % 2)) }
		actual := SliceInt16{given}.ChunkByInt16(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, [][]int16{})
	f([]int16{1}, [][]int16{{1}})
	f([]int16{1, 2, 3}, [][]int16{{1}, {2}, {3}})
	f([]int16{1, 3, 2, 4, 5}, [][]int16{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkByInt16Int32(t *testing.T) {
	f := func(given []int16, expected [][]int16) {
		reminder := func(t int16) int32 { return int32((t % 2)) }
		actual := SliceInt16{given}.ChunkByInt32(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, [][]int16{})
	f([]int16{1}, [][]int16{{1}})
	f([]int16{1, 2, 3}, [][]int16{{1}, {2}, {3}})
	f([]int16{1, 3, 2, 4, 5}, [][]int16{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkByInt16Int64(t *testing.T) {
	f := func(given []int16, expected [][]int16) {
		reminder := func(t int16) int64 { return int64((t % 2)) }
		actual := SliceInt16{given}.ChunkByInt64(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, [][]int16{})
	f([]int16{1}, [][]int16{{1}})
	f([]int16{1, 2, 3}, [][]int16{{1}, {2}, {3}})
	f([]int16{1, 3, 2, 4, 5}, [][]int16{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkEveryInt16(t *testing.T) {
	f := func(count int, given []int16, expected [][]int16) {
		actual, _ := SliceInt16{given}.ChunkEvery(count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, []int16{}, [][]int16{})
	f(2, []int16{1}, [][]int16{{1}})
	f(2, []int16{1, 2, 3, 4}, [][]int16{{1, 2}, {3, 4}})
	f(2, []int16{1, 2, 3, 4, 5}, [][]int16{{1, 2}, {3, 4}, {5}})
}

func TestSliceContainsInt16(t *testing.T) {
	f := func(el int16, given []int16, expected bool) {
		actual := SliceInt16{given}.Contains(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, []int16{}, false)
	f(1, []int16{1}, true)
	f(1, []int16{2}, false)
	f(1, []int16{2, 3, 4, 5}, false)
	f(1, []int16{2, 3, 1, 4, 5}, true)
	f(1, []int16{2, 3, 1, 1, 4, 5}, true)
}

func TestSliceCountInt16(t *testing.T) {
	f := func(el int16, given []int16, expected int) {
		actual := SliceInt16{given}.Count(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, []int16{}, 0)
	f(1, []int16{1}, 1)
	f(1, []int16{2}, 0)
	f(1, []int16{2, 3, 4, 5}, 0)
	f(1, []int16{2, 3, 1, 4, 5}, 1)
	f(1, []int16{2, 3, 1, 1, 4, 5}, 2)
	f(1, []int16{1, 1, 1, 1, 1}, 5)
}

func TestSliceCountByInt16(t *testing.T) {
	f := func(given []int16, expected int) {
		even := func(t int16) bool { return (t % 2) == 0 }
		actual := SliceInt16{given}.CountBy(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, 0)
	f([]int16{1}, 0)
	f([]int16{2}, 1)
	f([]int16{1, 2, 3, 4, 5}, 2)
	f([]int16{1, 2, 3, 4, 5, 6}, 3)
}

func TestSliceCycleInt16(t *testing.T) {
	f := func(count int, given []int16, expected []int16) {
		c := SliceInt16{given}.Cycle()
		seq := ChannelInt16{c}.Take(count)
		actual := ChannelInt16{seq}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(5, []int16{}, []int16{})
	f(5, []int16{1}, []int16{1, 1, 1, 1, 1})
	f(5, []int16{1, 2}, []int16{1, 2, 1, 2, 1})
}

func TestSliceDedupInt16(t *testing.T) {
	f := func(given []int16, expected []int16) {
		actual := SliceInt16{given}.Dedup()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, []int16{})
	f([]int16{1}, []int16{1})
	f([]int16{1, 1}, []int16{1})
	f([]int16{1, 2}, []int16{1, 2})
	f([]int16{1, 2, 3}, []int16{1, 2, 3})
	f([]int16{1, 2, 2, 3}, []int16{1, 2, 3})
	f([]int16{1, 2, 2, 3, 3, 3, 2, 1, 1}, []int16{1, 2, 3, 2, 1})
}

func TestSliceDedupByInt16Int(t *testing.T) {
	f := func(given []int16, expected []int16) {
		even := func(el int16) int { return int(el % 2) }
		actual := SliceInt16{given}.DedupByInt(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, []int16{})
	f([]int16{1}, []int16{1})
	f([]int16{1, 1}, []int16{1})
	f([]int16{1, 2}, []int16{1, 2})
	f([]int16{1, 2, 3}, []int16{1, 2, 3})
	f([]int16{1, 2, 2, 3}, []int16{1, 2, 3})
	f([]int16{1, 2, 4, 3, 5, 7, 10}, []int16{1, 2, 3, 10})
}

func TestSliceDedupByInt16Int8(t *testing.T) {
	f := func(given []int16, expected []int16) {
		even := func(el int16) int8 { return int8(el % 2) }
		actual := SliceInt16{given}.DedupByInt8(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, []int16{})
	f([]int16{1}, []int16{1})
	f([]int16{1, 1}, []int16{1})
	f([]int16{1, 2}, []int16{1, 2})
	f([]int16{1, 2, 3}, []int16{1, 2, 3})
	f([]int16{1, 2, 2, 3}, []int16{1, 2, 3})
	f([]int16{1, 2, 4, 3, 5, 7, 10}, []int16{1, 2, 3, 10})
}

func TestSliceDedupByInt16Int16(t *testing.T) {
	f := func(given []int16, expected []int16) {
		even := func(el int16) int16 { return int16(el % 2) }
		actual := SliceInt16{given}.DedupByInt16(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, []int16{})
	f([]int16{1}, []int16{1})
	f([]int16{1, 1}, []int16{1})
	f([]int16{1, 2}, []int16{1, 2})
	f([]int16{1, 2, 3}, []int16{1, 2, 3})
	f([]int16{1, 2, 2, 3}, []int16{1, 2, 3})
	f([]int16{1, 2, 4, 3, 5, 7, 10}, []int16{1, 2, 3, 10})
}

func TestSliceDedupByInt16Int32(t *testing.T) {
	f := func(given []int16, expected []int16) {
		even := func(el int16) int32 { return int32(el % 2) }
		actual := SliceInt16{given}.DedupByInt32(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, []int16{})
	f([]int16{1}, []int16{1})
	f([]int16{1, 1}, []int16{1})
	f([]int16{1, 2}, []int16{1, 2})
	f([]int16{1, 2, 3}, []int16{1, 2, 3})
	f([]int16{1, 2, 2, 3}, []int16{1, 2, 3})
	f([]int16{1, 2, 4, 3, 5, 7, 10}, []int16{1, 2, 3, 10})
}

func TestSliceDedupByInt16Int64(t *testing.T) {
	f := func(given []int16, expected []int16) {
		even := func(el int16) int64 { return int64(el % 2) }
		actual := SliceInt16{given}.DedupByInt64(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, []int16{})
	f([]int16{1}, []int16{1})
	f([]int16{1, 1}, []int16{1})
	f([]int16{1, 2}, []int16{1, 2})
	f([]int16{1, 2, 3}, []int16{1, 2, 3})
	f([]int16{1, 2, 2, 3}, []int16{1, 2, 3})
	f([]int16{1, 2, 4, 3, 5, 7, 10}, []int16{1, 2, 3, 10})
}

func TestSliceDeleteInt16(t *testing.T) {
	f := func(given []int16, el int16, expected []int16) {
		actual := SliceInt16{given}.Delete(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, 1, []int16{})
	f([]int16{1}, 1, []int16{})
	f([]int16{2}, 1, []int16{2})
	f([]int16{1, 2}, 1, []int16{2})
	f([]int16{1, 2, 3}, 2, []int16{1, 3})
	f([]int16{1, 2, 2, 3, 2}, 2, []int16{1, 2, 3, 2})
}

func TestSliceDeleteAllInt16(t *testing.T) {
	f := func(given []int16, el int16, expected []int16) {
		actual := SliceInt16{given}.DeleteAll(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, 1, []int16{})
	f([]int16{1}, 1, []int16{})
	f([]int16{2}, 1, []int16{2})
	f([]int16{1, 2}, 1, []int16{2})
	f([]int16{1, 2, 3}, 2, []int16{1, 3})
	f([]int16{1, 2, 2, 3, 2}, 2, []int16{1, 3})
}

func TestSliceDeleteAtInt16(t *testing.T) {
	f := func(given []int16, indices []int, expected []int16) {
		actual, _ := SliceInt16{given}.DeleteAt(indices...)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, []int{}, []int16{})
	f([]int16{1}, []int{0}, []int16{})
	f([]int16{1, 2}, []int{0}, []int16{2})

	f([]int16{1, 2, 3}, []int{0}, []int16{2, 3})
	f([]int16{1, 2, 3}, []int{1}, []int16{1, 3})
	f([]int16{1, 2, 3}, []int{2}, []int16{1, 2})

	f([]int16{1, 2, 3}, []int{0, 1}, []int16{3})
	f([]int16{1, 2, 3}, []int{0, 2}, []int16{2})
	f([]int16{1, 2, 3}, []int{1, 2}, []int16{1})
}

func TestSliceDropEveryInt16(t *testing.T) {
	f := func(given []int16, nth int, from int, expected []int16) {
		actual, _ := SliceInt16{given}.DropEvery(nth, from)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, 1, 1, []int16{})
	f([]int16{1, 2, 3}, 1, 1, []int16{})

	f([]int16{1, 2, 3, 4}, 2, 1, []int16{1, 3})
	f([]int16{1, 2, 3, 4, 5}, 2, 1, []int16{1, 3, 5})

	f([]int16{1, 2, 3, 4}, 2, 0, []int16{2, 4})
	f([]int16{1, 2, 3, 4, 5}, 2, 0, []int16{2, 4})

	f([]int16{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2, 1, []int16{1, 3, 5, 7, 9})
	f([]int16{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, 1, []int16{1, 2, 4, 5, 7, 8, 10})
}

func TestSliceDropWhileInt16(t *testing.T) {
	f := func(given []int16, expected []int16) {
		even := func(el int16) bool { return el%2 == 0 }
		actual := SliceInt16{given}.DropWhile(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, []int16{})
	f([]int16{2}, []int16{})
	f([]int16{1}, []int16{1})
	f([]int16{2, 1}, []int16{1})
	f([]int16{2, 1, 2}, []int16{1, 2})
	f([]int16{1, 2}, []int16{1, 2})
	f([]int16{2, 4, 6, 1, 8}, []int16{1, 8})
}

func TestSliceEndsWithInt16(t *testing.T) {
	f := func(given []int16, suffix []int16, expected bool) {
		actual := SliceInt16{given}.EndsWith(suffix)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, []int16{}, true)
	f([]int16{1}, []int16{1}, true)
	f([]int16{1}, []int16{2}, false)
	f([]int16{2, 3}, []int16{1, 2, 3}, false)

	f([]int16{1, 2, 3}, []int16{3}, true)
	f([]int16{1, 2, 3}, []int16{2, 3}, true)
	f([]int16{1, 2, 3}, []int16{1, 2, 3}, true)

	f([]int16{1, 2, 3}, []int16{1}, false)
	f([]int16{1, 2, 3}, []int16{2}, false)
	f([]int16{1, 2, 3}, []int16{1, 2}, false)
	f([]int16{1, 2, 3}, []int16{3, 2}, false)
}

func TestSliceEqualInt16(t *testing.T) {
	f := func(left []int16, right []int16, expected bool) {
		actual := SliceInt16{left}.Equal(right)
		assert.Equal(t, expected, actual, "they should be equal")

		actual = SliceInt16{right}.Equal(left)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, []int16{}, true)
	f([]int16{1}, []int16{1}, true)
	f([]int16{1}, []int16{2}, false)
	f([]int16{1, 2, 3, 3}, []int16{1, 2, 3, 3}, true)
	f([]int16{1, 2, 3, 3}, []int16{1, 2, 2, 3}, false)
	f([]int16{1, 2, 3, 3}, []int16{1, 2, 4, 3}, false)

	// different len
	f([]int16{1, 2, 3}, []int16{1, 2}, false)
	f([]int16{1, 2}, []int16{1, 2, 3}, false)
	f([]int16{}, []int16{1, 2, 3}, false)
	f([]int16{1, 2, 3}, []int16{}, false)
}

func TestSliceFilterInt16(t *testing.T) {
	f := func(given []int16, expected []int16) {
		even := func(t int16) bool { return (t % 2) == 0 }
		actual := SliceInt16{given}.Filter(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, []int16{})
	f([]int16{1, 2, 3, 4}, []int16{2, 4})
	f([]int16{1, 3}, []int16{})
	f([]int16{2, 4}, []int16{2, 4})
}

func TestSliceFindInt16(t *testing.T) {
	f := func(given []int16, expectedEl int16, expectedErr error) {
		even := func(t int16) bool { return (t % 2) == 0 }
		el, err := SliceInt16{given}.Find(even)
		assert.Equal(t, expectedEl, el, "they should be equal")
		assert.Equal(t, expectedErr, err, "they should be equal")
	}
	f([]int16{}, 0, ErrNotFound)
	f([]int16{1}, 0, ErrNotFound)
	f([]int16{1}, 0, ErrNotFound)
	f([]int16{2}, 2, nil)
	f([]int16{1, 2}, 2, nil)
	f([]int16{1, 2, 3}, 2, nil)
	f([]int16{1, 3, 5}, 0, ErrNotFound)
}

func TestSliceFindIndexInt16(t *testing.T) {
	f := func(given []int16, expectedInd int, expectedErr error) {
		even := func(t int16) bool { return (t % 2) == 0 }
		index, err := SliceInt16{given}.FindIndex(even)
		assert.Equal(t, expectedInd, index, "they should be equal")
		assert.Equal(t, expectedErr, err, "they should be equal")
	}
	f([]int16{}, 0, ErrNotFound)
	f([]int16{1}, 0, ErrNotFound)
	f([]int16{1}, 0, ErrNotFound)
	f([]int16{2}, 0, nil)
	f([]int16{1, 2}, 1, nil)
	f([]int16{1, 2, 3}, 1, nil)
	f([]int16{1, 3, 5, 7, 9, 2}, 5, nil)
	f([]int16{1, 3, 5}, 0, ErrNotFound)
}

func TestSliceJoinInt16(t *testing.T) {
	f := func(given []int16, sep string, expected string) {
		actual := SliceInt16{given}.Join(sep)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, "", "")
	f([]int16{}, "|", "")

	f([]int16{1}, "", "1")
	f([]int16{1}, "|", "1")

	f([]int16{1, 2, 3}, "", "123")
	f([]int16{1, 2, 3}, "|", "1|2|3")
	f([]int16{1, 2, 3}, "<int16>", "1<int16>2<int16>3")
}

func TestSliceGroupByInt16Int(t *testing.T) {
	f := func(given []int16, expected map[int][]int16) {
		reminder := func(t int16) int { return int((t % 2)) }
		actual := SliceInt16{given}.GroupByInt(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, map[int][]int16{})
	f([]int16{1}, map[int][]int16{1: {1}})
	f([]int16{1, 3, 2, 4, 5}, map[int][]int16{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceGroupByInt16Int8(t *testing.T) {
	f := func(given []int16, expected map[int8][]int16) {
		reminder := func(t int16) int8 { return int8((t % 2)) }
		actual := SliceInt16{given}.GroupByInt8(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, map[int8][]int16{})
	f([]int16{1}, map[int8][]int16{1: {1}})
	f([]int16{1, 3, 2, 4, 5}, map[int8][]int16{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceGroupByInt16Int16(t *testing.T) {
	f := func(given []int16, expected map[int16][]int16) {
		reminder := func(t int16) int16 { return int16((t % 2)) }
		actual := SliceInt16{given}.GroupByInt16(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, map[int16][]int16{})
	f([]int16{1}, map[int16][]int16{1: {1}})
	f([]int16{1, 3, 2, 4, 5}, map[int16][]int16{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceGroupByInt16Int32(t *testing.T) {
	f := func(given []int16, expected map[int32][]int16) {
		reminder := func(t int16) int32 { return int32((t % 2)) }
		actual := SliceInt16{given}.GroupByInt32(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, map[int32][]int16{})
	f([]int16{1}, map[int32][]int16{1: {1}})
	f([]int16{1, 3, 2, 4, 5}, map[int32][]int16{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceGroupByInt16Int64(t *testing.T) {
	f := func(given []int16, expected map[int64][]int16) {
		reminder := func(t int16) int64 { return int64((t % 2)) }
		actual := SliceInt16{given}.GroupByInt64(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, map[int64][]int16{})
	f([]int16{1}, map[int64][]int16{1: {1}})
	f([]int16{1, 3, 2, 4, 5}, map[int64][]int16{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceInsertAtInt16(t *testing.T) {
	f := func(given []int16, index int, expected []int16, expectedErr error) {
		actual, err := SliceInt16{given}.InsertAt(index, 10)
		assert.Equal(t, expected, actual, "they should be equal")
		assert.Equal(t, expectedErr, err, "they should be equal")
	}
	f([]int16{}, -1, []int16{}, ErrNegativeValue)
	f([]int16{}, 0, []int16{10}, nil)
	f([]int16{}, 1, []int16{}, ErrOutOfRange)

	f([]int16{1, 2, 3}, -1, []int16{1, 2, 3}, ErrNegativeValue)
	f([]int16{1, 2, 3}, 0, []int16{10, 1, 2, 3}, nil)
	f([]int16{1, 2, 3}, 1, []int16{1, 10, 2, 3}, nil)
	f([]int16{1, 2, 3}, 3, []int16{1, 2, 3, 10}, nil)
	f([]int16{1, 2, 3}, 4, []int16{1, 2, 3}, ErrOutOfRange)
}

func TestSliceIntersperseInt16(t *testing.T) {
	f := func(el int16, given []int16, expected []int16) {
		actual := SliceInt16{given}.Intersperse(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, []int16{}, []int16{})
	f(0, []int16{1}, []int16{1})
	f(0, []int16{1, 2}, []int16{1, 0, 2})
	f(0, []int16{1, 2, 3}, []int16{1, 0, 2, 0, 3})
}

func TestSliceLastInt16(t *testing.T) {
	f := func(given []int16, expectedEl int16, expectedErr error) {
		el, err := SliceInt16{given}.Last()
		assert.Equal(t, expectedEl, el, "they should be equal")
		assert.Equal(t, expectedErr, err, "they should be equal")
	}
	f([]int16{}, 0, ErrEmpty)
	f([]int16{1}, 1, nil)
	f([]int16{1, 2, 3}, 3, nil)
}

func TestSliceMapInt16Int(t *testing.T) {
	f := func(given []int16, expected []int) {
		double := func(t int16) int { return int((t * 2)) }
		actual := SliceInt16{given}.MapInt(double)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, []int{})
	f([]int16{1}, []int{2})
	f([]int16{1, 2, 3}, []int{2, 4, 6})
}

func TestSliceMapInt16Int8(t *testing.T) {
	f := func(given []int16, expected []int8) {
		double := func(t int16) int8 { return int8((t * 2)) }
		actual := SliceInt16{given}.MapInt8(double)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, []int8{})
	f([]int16{1}, []int8{2})
	f([]int16{1, 2, 3}, []int8{2, 4, 6})
}

func TestSliceMapInt16Int16(t *testing.T) {
	f := func(given []int16, expected []int16) {
		double := func(t int16) int16 { return int16((t * 2)) }
		actual := SliceInt16{given}.MapInt16(double)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, []int16{})
	f([]int16{1}, []int16{2})
	f([]int16{1, 2, 3}, []int16{2, 4, 6})
}

func TestSliceMapInt16Int32(t *testing.T) {
	f := func(given []int16, expected []int32) {
		double := func(t int16) int32 { return int32((t * 2)) }
		actual := SliceInt16{given}.MapInt32(double)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, []int32{})
	f([]int16{1}, []int32{2})
	f([]int16{1, 2, 3}, []int32{2, 4, 6})
}

func TestSliceMapInt16Int64(t *testing.T) {
	f := func(given []int16, expected []int64) {
		double := func(t int16) int64 { return int64((t * 2)) }
		actual := SliceInt16{given}.MapInt64(double)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, []int64{})
	f([]int16{1}, []int64{2})
	f([]int16{1, 2, 3}, []int64{2, 4, 6})
}

func TestSliceMaxInt16(t *testing.T) {
	f := func(given []int16, expectedEl int16, expectedErr error) {
		el, err := SliceInt16{given}.Max()
		assert.Equal(t, expectedEl, el, "they should be equal")
		assert.Equal(t, expectedErr, err, "they should be equal")
	}
	f([]int16{}, 0, ErrEmpty)
	f([]int16{1}, 1, nil)
	f([]int16{1, 2, 3}, 3, nil)
	f([]int16{1, 3, 2}, 3, nil)
	f([]int16{3, 2, 1}, 3, nil)
}

func TestSliceMinInt16(t *testing.T) {
	f := func(given []int16, expectedEl int16, expectedErr error) {
		el, err := SliceInt16{given}.Min()
		assert.Equal(t, expectedEl, el, "they should be equal")
		assert.Equal(t, expectedErr, err, "they should be equal")
	}
	f([]int16{}, 0, ErrEmpty)
	f([]int16{1}, 1, nil)
	f([]int16{1, 2, 3}, 1, nil)
	f([]int16{2, 1, 3}, 1, nil)
	f([]int16{3, 2, 1}, 1, nil)
}

func TestSlicesPermutationsInt16(t *testing.T) {
	f := func(size int, given []int16, expected [][]int16) {
		actual := make([][]int16, 0)
		i := 0
		s := SliceInt16{given}
		for el := range s.Permutations(size) {
			actual = append(actual, el)
			i++
			if i > 50 {
				t.Fatal("infinite loop")
			}
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, []int16{}, [][]int16{})
	f(2, []int16{1}, [][]int16{{1}})
	f(2, []int16{1, 2, 3}, [][]int16{{1, 2}, {1, 3}, {2, 1}, {2, 3}, {3, 1}, {3, 2}})
}

func TestSliceProductInt16(t *testing.T) {
	f := func(given []int16, repeat int, expected [][]int16) {
		actual := make([][]int16, 0)
		i := 0
		s := SliceInt16{given}
		for el := range s.Product(repeat) {
			actual = append(actual, el)
			i++
			if i > 50 {
				t.Fatal("infinite loop")
			}
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}

	f([]int16{1, 2}, 0, [][]int16{})
	f([]int16{}, 2, [][]int16{})
	f([]int16{1}, 2, [][]int16{{1, 1}})

	f([]int16{1, 2}, 1, [][]int16{{1}, {2}})
	f([]int16{1, 2}, 2, [][]int16{{1, 1}, {1, 2}, {2, 1}, {2, 2}})
	f([]int16{1, 2}, 3, [][]int16{
		{1, 1, 1}, {1, 1, 2}, {1, 2, 1}, {1, 2, 2},
		{2, 1, 1}, {2, 1, 2}, {2, 2, 1}, {2, 2, 2},
	})
}

func TestSliceReduceInt16Int(t *testing.T) {
	f := func(given []int16, expected int) {
		sum := func(el int16, acc int) int { return int(el) + acc }
		actual := SliceInt16{given}.ReduceInt(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, 0)
	f([]int16{1}, 1)
	f([]int16{1, 2}, 3)
	f([]int16{1, 2, 3}, 6)
}

func TestSliceReduceInt16Int8(t *testing.T) {
	f := func(given []int16, expected int8) {
		sum := func(el int16, acc int8) int8 { return int8(el) + acc }
		actual := SliceInt16{given}.ReduceInt8(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, 0)
	f([]int16{1}, 1)
	f([]int16{1, 2}, 3)
	f([]int16{1, 2, 3}, 6)
}

func TestSliceReduceInt16Int16(t *testing.T) {
	f := func(given []int16, expected int16) {
		sum := func(el int16, acc int16) int16 { return int16(el) + acc }
		actual := SliceInt16{given}.ReduceInt16(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, 0)
	f([]int16{1}, 1)
	f([]int16{1, 2}, 3)
	f([]int16{1, 2, 3}, 6)
}

func TestSliceReduceInt16Int32(t *testing.T) {
	f := func(given []int16, expected int32) {
		sum := func(el int16, acc int32) int32 { return int32(el) + acc }
		actual := SliceInt16{given}.ReduceInt32(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, 0)
	f([]int16{1}, 1)
	f([]int16{1, 2}, 3)
	f([]int16{1, 2, 3}, 6)
}

func TestSliceReduceInt16Int64(t *testing.T) {
	f := func(given []int16, expected int64) {
		sum := func(el int16, acc int64) int64 { return int64(el) + acc }
		actual := SliceInt16{given}.ReduceInt64(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, 0)
	f([]int16{1}, 1)
	f([]int16{1, 2}, 3)
	f([]int16{1, 2, 3}, 6)
}

func TestSliceReduceWhileInt16Int(t *testing.T) {
	f := func(given []int16, expected int) {
		sum := func(el int16, acc int) (int, error) {
			if el == 0 {
				return acc, ErrEmpty
			}
			return int(el) + acc, nil
		}
		actual, _ := SliceInt16{given}.ReduceWhileInt(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, 0)
	f([]int16{1}, 1)
	f([]int16{1, 2}, 3)
	f([]int16{1, 2, 3}, 6)
	f([]int16{1, 2, 0, 3}, 3)
}

func TestSliceReduceWhileInt16Int8(t *testing.T) {
	f := func(given []int16, expected int8) {
		sum := func(el int16, acc int8) (int8, error) {
			if el == 0 {
				return acc, ErrEmpty
			}
			return int8(el) + acc, nil
		}
		actual, _ := SliceInt16{given}.ReduceWhileInt8(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, 0)
	f([]int16{1}, 1)
	f([]int16{1, 2}, 3)
	f([]int16{1, 2, 3}, 6)
	f([]int16{1, 2, 0, 3}, 3)
}

func TestSliceReduceWhileInt16Int16(t *testing.T) {
	f := func(given []int16, expected int16) {
		sum := func(el int16, acc int16) (int16, error) {
			if el == 0 {
				return acc, ErrEmpty
			}
			return int16(el) + acc, nil
		}
		actual, _ := SliceInt16{given}.ReduceWhileInt16(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, 0)
	f([]int16{1}, 1)
	f([]int16{1, 2}, 3)
	f([]int16{1, 2, 3}, 6)
	f([]int16{1, 2, 0, 3}, 3)
}

func TestSliceReduceWhileInt16Int32(t *testing.T) {
	f := func(given []int16, expected int32) {
		sum := func(el int16, acc int32) (int32, error) {
			if el == 0 {
				return acc, ErrEmpty
			}
			return int32(el) + acc, nil
		}
		actual, _ := SliceInt16{given}.ReduceWhileInt32(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, 0)
	f([]int16{1}, 1)
	f([]int16{1, 2}, 3)
	f([]int16{1, 2, 3}, 6)
	f([]int16{1, 2, 0, 3}, 3)
}

func TestSliceReduceWhileInt16Int64(t *testing.T) {
	f := func(given []int16, expected int64) {
		sum := func(el int16, acc int64) (int64, error) {
			if el == 0 {
				return acc, ErrEmpty
			}
			return int64(el) + acc, nil
		}
		actual, _ := SliceInt16{given}.ReduceWhileInt64(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, 0)
	f([]int16{1}, 1)
	f([]int16{1, 2}, 3)
	f([]int16{1, 2, 3}, 6)
	f([]int16{1, 2, 0, 3}, 3)
}

func TestChannelToSliceInt16(t *testing.T) {
	f := func(given []int16) {
		c := make(chan int16, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := ChannelInt16{c}.ToSlice()
		assert.Equal(t, given, actual, "they should be equal")
	}
	f([]int16{})
	f([]int16{1})
	f([]int16{1, 2, 3, 1, 2})
}

func TestChannelAnyInt16(t *testing.T) {
	f := func(given []int16, expected bool) {
		even := func(t int16) bool { return t%2 == 0 }
		c := make(chan int16, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := ChannelInt16{c}.Any(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, false)
	f([]int16{1}, false)
	f([]int16{2}, true)
	f([]int16{1, 2}, true)
	f([]int16{1, 2, 3}, true)
	f([]int16{1, 3, 5}, false)
	f([]int16{1, 3, 5, 7, 9, 11}, false)
	f([]int16{1, 3, 5, 7, 10, 11}, true)
}

func TestChannelAllInt16(t *testing.T) {
	f := func(given []int16, expected bool) {
		even := func(t int16) bool { return t%2 == 0 }
		c := make(chan int16, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := ChannelInt16{c}.All(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, true)
	f([]int16{1}, false)
	f([]int16{2}, true)
	f([]int16{1, 2}, false)
	f([]int16{2, 4}, true)
	f([]int16{2, 4, 6, 8, 10, 12}, true)
	f([]int16{2, 4, 6, 8, 11, 12}, false)
}

func TestChannelEachInt16(t *testing.T) {
	f := func(given []int16) {
		c := make(chan int16, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := make(chan int16, len(given))
		mapper := func(t int16) { result <- t }
		ChannelInt16{c}.Each(mapper)
		close(result)
		actual := ChannelInt16{result}.ToSlice()
		assert.Equal(t, given, actual, "they should be equal")
	}

	f([]int16{})
	f([]int16{1})
	f([]int16{1, 2, 3})
	f([]int16{1, 2, 3, 4, 5, 6, 7})
}

func TestChannelChunkEveryInt16(t *testing.T) {
	f := func(size int, given []int16, expected [][]int16) {
		c := make(chan int16, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := ChannelInt16{c}.ChunkEvery(size)
		actual := make([][]int16, 0)
		for el := range result {
			actual = append(actual, el)
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, []int16{}, [][]int16{})
	f(2, []int16{1}, [][]int16{{1}})
	f(2, []int16{1, 2}, [][]int16{{1, 2}})
	f(2, []int16{1, 2, 3, 4}, [][]int16{{1, 2}, {3, 4}})
	f(2, []int16{1, 2, 3, 4, 5}, [][]int16{{1, 2}, {3, 4}, {5}})
}

func TestChannelCountInt16(t *testing.T) {
	f := func(element int16, given []int16, expected int) {
		c := make(chan int16, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := ChannelInt16{c}.Count(element)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, []int16{}, 0)
	f(1, []int16{1}, 1)
	f(1, []int16{2}, 0)
	f(1, []int16{1, 2, 3, 1, 4}, 2)
}

func TestChannelDropInt16(t *testing.T) {
	f := func(count int, given []int16, expected []int16) {
		c := make(chan int16, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := ChannelInt16{c}.Drop(count)
		actual := make([]int16, 0)
		for el := range result {
			actual = append(actual, el)
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, []int16{}, []int16{})
	f(1, []int16{2}, []int16{})
	f(1, []int16{2, 3}, []int16{3})
	f(1, []int16{1, 2, 3}, []int16{2, 3})
	f(0, []int16{1, 2, 3}, []int16{1, 2, 3})
	f(3, []int16{1, 2, 3, 4, 5, 6}, []int16{4, 5, 6})
	f(1, []int16{1, 2, 3, 4, 5, 6}, []int16{2, 3, 4, 5, 6})
}

func TestChannelFilterInt16(t *testing.T) {
	f := func(given []int16, expected []int16) {
		even := func(t int16) bool { return t%2 == 0 }
		c := make(chan int16, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := ChannelInt16{c}.Filter(even)
		actual := ChannelInt16{result}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, []int16{})
	f([]int16{1}, []int16{})
	f([]int16{2}, []int16{2})
	f([]int16{1, 2, 3, 4}, []int16{2, 4})
}

func TestChannelMapInt16Int(t *testing.T) {
	f := func(given []int16, expected []int16) {
		double := func(el int16) int { return int(el * 2) }
		c := make(chan int16, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := ChannelInt16{c}.MapInt(double)

		// convert chan int16 to chan int
		c2 := make(chan int16, 1)
		go func() {
			for el := range result {
				c2 <- int16(el)
			}
			close(c2)
		}()

		actual := ChannelInt16{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, []int16{})
	f([]int16{1}, []int16{2})
	f([]int16{1, 2, 3}, []int16{2, 4, 6})
}

func TestChannelMapInt16Int8(t *testing.T) {
	f := func(given []int16, expected []int16) {
		double := func(el int16) int8 { return int8(el * 2) }
		c := make(chan int16, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := ChannelInt16{c}.MapInt8(double)

		// convert chan int16 to chan int8
		c2 := make(chan int16, 1)
		go func() {
			for el := range result {
				c2 <- int16(el)
			}
			close(c2)
		}()

		actual := ChannelInt16{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, []int16{})
	f([]int16{1}, []int16{2})
	f([]int16{1, 2, 3}, []int16{2, 4, 6})
}

func TestChannelMapInt16Int16(t *testing.T) {
	f := func(given []int16, expected []int16) {
		double := func(el int16) int16 { return int16(el * 2) }
		c := make(chan int16, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := ChannelInt16{c}.MapInt16(double)

		// convert chan int16 to chan int16
		c2 := make(chan int16, 1)
		go func() {
			for el := range result {
				c2 <- int16(el)
			}
			close(c2)
		}()

		actual := ChannelInt16{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, []int16{})
	f([]int16{1}, []int16{2})
	f([]int16{1, 2, 3}, []int16{2, 4, 6})
}

func TestChannelMapInt16Int32(t *testing.T) {
	f := func(given []int16, expected []int16) {
		double := func(el int16) int32 { return int32(el * 2) }
		c := make(chan int16, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := ChannelInt16{c}.MapInt32(double)

		// convert chan int16 to chan int32
		c2 := make(chan int16, 1)
		go func() {
			for el := range result {
				c2 <- int16(el)
			}
			close(c2)
		}()

		actual := ChannelInt16{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, []int16{})
	f([]int16{1}, []int16{2})
	f([]int16{1, 2, 3}, []int16{2, 4, 6})
}

func TestChannelMapInt16Int64(t *testing.T) {
	f := func(given []int16, expected []int16) {
		double := func(el int16) int64 { return int64(el * 2) }
		c := make(chan int16, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := ChannelInt16{c}.MapInt64(double)

		// convert chan int16 to chan int64
		c2 := make(chan int16, 1)
		go func() {
			for el := range result {
				c2 <- int16(el)
			}
			close(c2)
		}()

		actual := ChannelInt16{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, []int16{})
	f([]int16{1}, []int16{2})
	f([]int16{1, 2, 3}, []int16{2, 4, 6})
}

func TestChannelMaxInt16(t *testing.T) {
	f := func(given []int16, expected int16, expectedErr error) {
		c := make(chan int16, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual, actualErr := ChannelInt16{c}.Max()
		assert.Equal(t, expected, actual, "they should be equal")
		assert.Equal(t, expectedErr, actualErr, "they should be equal")
	}
	f([]int16{}, 0, ErrEmpty)
	f([]int16{1, 4, 2}, 4, nil)
	f([]int16{1, 2, 4}, 4, nil)
	f([]int16{4, 2, 1}, 4, nil)
}

func TestChannelMinInt16(t *testing.T) {
	f := func(given []int16, expected int16, expectedErr error) {
		c := make(chan int16, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual, actualErr := ChannelInt16{c}.Min()
		assert.Equal(t, expected, actual, "they should be equal")
		assert.Equal(t, expectedErr, actualErr, "they should be equal")
	}
	f([]int16{}, 0, ErrEmpty)
	f([]int16{4, 1, 2}, 1, nil)
	f([]int16{1, 2, 4}, 1, nil)
	f([]int16{4, 2, 1}, 1, nil)
}

func TestChannelReduceInt16Int(t *testing.T) {
	f := func(given []int16, expected int) {
		c := make(chan int16, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int16, acc int) int { return int(el) + acc }
		actual := ChannelInt16{c}.ReduceInt(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, 0)
	f([]int16{1}, 1)
	f([]int16{1, 2}, 3)
	f([]int16{1, 2, 3, 4, 5}, 15)
}

func TestChannelReduceInt16Int8(t *testing.T) {
	f := func(given []int16, expected int8) {
		c := make(chan int16, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int16, acc int8) int8 { return int8(el) + acc }
		actual := ChannelInt16{c}.ReduceInt8(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, 0)
	f([]int16{1}, 1)
	f([]int16{1, 2}, 3)
	f([]int16{1, 2, 3, 4, 5}, 15)
}

func TestChannelReduceInt16Int16(t *testing.T) {
	f := func(given []int16, expected int16) {
		c := make(chan int16, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int16, acc int16) int16 { return int16(el) + acc }
		actual := ChannelInt16{c}.ReduceInt16(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, 0)
	f([]int16{1}, 1)
	f([]int16{1, 2}, 3)
	f([]int16{1, 2, 3, 4, 5}, 15)
}

func TestChannelReduceInt16Int32(t *testing.T) {
	f := func(given []int16, expected int32) {
		c := make(chan int16, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int16, acc int32) int32 { return int32(el) + acc }
		actual := ChannelInt16{c}.ReduceInt32(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, 0)
	f([]int16{1}, 1)
	f([]int16{1, 2}, 3)
	f([]int16{1, 2, 3, 4, 5}, 15)
}

func TestChannelReduceInt16Int64(t *testing.T) {
	f := func(given []int16, expected int64) {
		c := make(chan int16, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int16, acc int64) int64 { return int64(el) + acc }
		actual := ChannelInt16{c}.ReduceInt64(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, 0)
	f([]int16{1}, 1)
	f([]int16{1, 2}, 3)
	f([]int16{1, 2, 3, 4, 5}, 15)
}

func TestChannelScanInt16Int(t *testing.T) {
	f := func(given []int16, expected []int16) {
		c := make(chan int16, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int16, acc int) int { return int(el) + acc }
		result := ChannelInt16{c}.ScanInt(0, sum)

		// convert chan int16 to chan int
		c2 := make(chan int16, 1)
		go func() {
			for el := range result {
				c2 <- int16(el)
			}
			close(c2)
		}()

		actual := ChannelInt16{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, []int16{})
	f([]int16{1}, []int16{1})
	f([]int16{1, 2}, []int16{1, 3})
	f([]int16{1, 2, 3, 4, 5}, []int16{1, 3, 6, 10, 15})
}

func TestChannelScanInt16Int8(t *testing.T) {
	f := func(given []int16, expected []int16) {
		c := make(chan int16, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int16, acc int8) int8 { return int8(el) + acc }
		result := ChannelInt16{c}.ScanInt8(0, sum)

		// convert chan int16 to chan int8
		c2 := make(chan int16, 1)
		go func() {
			for el := range result {
				c2 <- int16(el)
			}
			close(c2)
		}()

		actual := ChannelInt16{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, []int16{})
	f([]int16{1}, []int16{1})
	f([]int16{1, 2}, []int16{1, 3})
	f([]int16{1, 2, 3, 4, 5}, []int16{1, 3, 6, 10, 15})
}

func TestChannelScanInt16Int16(t *testing.T) {
	f := func(given []int16, expected []int16) {
		c := make(chan int16, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int16, acc int16) int16 { return int16(el) + acc }
		result := ChannelInt16{c}.ScanInt16(0, sum)

		// convert chan int16 to chan int16
		c2 := make(chan int16, 1)
		go func() {
			for el := range result {
				c2 <- int16(el)
			}
			close(c2)
		}()

		actual := ChannelInt16{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, []int16{})
	f([]int16{1}, []int16{1})
	f([]int16{1, 2}, []int16{1, 3})
	f([]int16{1, 2, 3, 4, 5}, []int16{1, 3, 6, 10, 15})
}

func TestChannelScanInt16Int32(t *testing.T) {
	f := func(given []int16, expected []int16) {
		c := make(chan int16, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int16, acc int32) int32 { return int32(el) + acc }
		result := ChannelInt16{c}.ScanInt32(0, sum)

		// convert chan int16 to chan int32
		c2 := make(chan int16, 1)
		go func() {
			for el := range result {
				c2 <- int16(el)
			}
			close(c2)
		}()

		actual := ChannelInt16{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, []int16{})
	f([]int16{1}, []int16{1})
	f([]int16{1, 2}, []int16{1, 3})
	f([]int16{1, 2, 3, 4, 5}, []int16{1, 3, 6, 10, 15})
}

func TestChannelScanInt16Int64(t *testing.T) {
	f := func(given []int16, expected []int16) {
		c := make(chan int16, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int16, acc int64) int64 { return int64(el) + acc }
		result := ChannelInt16{c}.ScanInt64(0, sum)

		// convert chan int16 to chan int64
		c2 := make(chan int16, 1)
		go func() {
			for el := range result {
				c2 <- int16(el)
			}
			close(c2)
		}()

		actual := ChannelInt16{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, []int16{})
	f([]int16{1}, []int16{1})
	f([]int16{1, 2}, []int16{1, 3})
	f([]int16{1, 2, 3, 4, 5}, []int16{1, 3, 6, 10, 15})
}

func TestChannelSumInt16(t *testing.T) {
	f := func(given []int16, expected int16) {
		c := make(chan int16, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := ChannelInt16{c}.Sum()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int16{}, 0)
	f([]int16{1}, 1)
	f([]int16{1, 2}, 3)
	f([]int16{1, 2, 3, 4, 5}, 15)
}

func TestChannelTakeInt16(t *testing.T) {
	f := func(count int, given int16, expected []int16) {
		ctx, cancel := context.WithCancel(context.Background())
		s := SequenceInt16{ctx: ctx}
		seq := s.Repeat(given)
		seq2 := ChannelInt16{seq}.Take(count)
		actual := ChannelInt16{seq2}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, 1, []int16{})
	f(1, 1, []int16{1})
	f(2, 1, []int16{1, 1})
}

func TestChannelTeeInt16(t *testing.T) {
	f := func(count int, given []int16) {
		c := make(chan int16, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		channels := ChannelInt16{c}.Tee(count)
		for _, ch := range channels {
			go func(ch chan int16) {
				actual := ChannelInt16{ch}.ToSlice()
				assert.Equal(t, given, actual, "they should be equal")
			}(ch)
		}
	}
	f(1, []int16{})
	f(1, []int16{1})
	f(1, []int16{1, 2})
	f(1, []int16{1, 2, 3})
	f(1, []int16{1, 2, 3, 1, 2})

	f(2, []int16{})
	f(2, []int16{1})
	f(2, []int16{1, 2})
	f(2, []int16{1, 2, 3})
	f(2, []int16{1, 2, 3, 1, 2})

	f(10, []int16{1, 2, 3, 1, 2})
}

func TestAsyncSliceAnyInt16(t *testing.T) {
	f := func(check func(t int16) bool, given []int16, expected bool) {
		s := AsyncSliceInt16{Data: given, Workers: 2}
		actual := s.Any(check)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	isEven := func(t int16) bool { return (t % 2) == 0 }

	f(isEven, []int16{}, false)
	f(isEven, []int16{1}, false)
	f(isEven, []int16{1, 3}, false)
	f(isEven, []int16{2}, true)
	f(isEven, []int16{1, 2}, true)
	f(isEven, []int16{1, 3, 5, 7, 9, 11}, false)
	f(isEven, []int16{1, 3, 5, 7, 9, 12}, true)
}

func TestAsyncSliceAllInt16(t *testing.T) {
	f := func(check func(t int16) bool, given []int16, expected bool) {
		s := AsyncSliceInt16{Data: given, Workers: 2}
		actual := s.All(check)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	isEven := func(t int16) bool { return (t % 2) == 0 }

	f(isEven, []int16{}, true)
	f(isEven, []int16{1}, false)
	f(isEven, []int16{1, 3}, false)
	f(isEven, []int16{2}, true)
	f(isEven, []int16{2, 4}, true)
	f(isEven, []int16{2, 3}, false)
	f(isEven, []int16{2, 4, 6, 8, 10, 12}, true)
	f(isEven, []int16{2, 4, 6, 8, 10, 11}, false)
}

func TestAsyncSliceEachInt16(t *testing.T) {
	f := func(given []int16) {
		s := AsyncSliceInt16{Data: given, Workers: 2}
		result := make(chan int16, len(given))
		mapper := func(t int16) { result <- t }
		s.Each(mapper)
		close(result)
		actual := ChannelInt16{result}.ToSlice()
		sorted := SliceInt16{actual}.Sort()
		assert.Equal(t, given, sorted, "they should be equal")
	}

	f([]int16{})
	f([]int16{1})
	f([]int16{1, 2, 3})
	f([]int16{1, 2, 3, 4, 5, 6, 7})
}

func TestAsyncSliceFilterInt16(t *testing.T) {
	f := func(given []int16, expected []int16) {
		filter := func(t int16) bool { return t > 10 }
		s := AsyncSliceInt16{Data: given, Workers: 2}
		actual := s.Filter(filter)
		assert.Equal(t, expected, actual, "they should be equal")
	}

	f([]int16{}, []int16{})
	f([]int16{5}, []int16{})
	f([]int16{15}, []int16{15})
	f([]int16{9, 11, 12, 13, 6}, []int16{11, 12, 13})
}

func TestAsyncSliceMapInt16Int(t *testing.T) {
	f := func(mapper func(t int16) int, given []int16, expected []int) {
		s := AsyncSliceInt16{Data: given, Workers: 2}
		actual := s.MapInt(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int16) int { return int((t * 2)) }

	f(double, []int16{}, []int{})
	f(double, []int16{1}, []int{2})
	f(double, []int16{1, 2, 3}, []int{2, 4, 6})
}

func TestAsyncSliceMapInt16Int8(t *testing.T) {
	f := func(mapper func(t int16) int8, given []int16, expected []int8) {
		s := AsyncSliceInt16{Data: given, Workers: 2}
		actual := s.MapInt8(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int16) int8 { return int8((t * 2)) }

	f(double, []int16{}, []int8{})
	f(double, []int16{1}, []int8{2})
	f(double, []int16{1, 2, 3}, []int8{2, 4, 6})
}

func TestAsyncSliceMapInt16Int16(t *testing.T) {
	f := func(mapper func(t int16) int16, given []int16, expected []int16) {
		s := AsyncSliceInt16{Data: given, Workers: 2}
		actual := s.MapInt16(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int16) int16 { return int16((t * 2)) }

	f(double, []int16{}, []int16{})
	f(double, []int16{1}, []int16{2})
	f(double, []int16{1, 2, 3}, []int16{2, 4, 6})
}

func TestAsyncSliceMapInt16Int32(t *testing.T) {
	f := func(mapper func(t int16) int32, given []int16, expected []int32) {
		s := AsyncSliceInt16{Data: given, Workers: 2}
		actual := s.MapInt32(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int16) int32 { return int32((t * 2)) }

	f(double, []int16{}, []int32{})
	f(double, []int16{1}, []int32{2})
	f(double, []int16{1, 2, 3}, []int32{2, 4, 6})
}

func TestAsyncSliceMapInt16Int64(t *testing.T) {
	f := func(mapper func(t int16) int64, given []int16, expected []int64) {
		s := AsyncSliceInt16{Data: given, Workers: 2}
		actual := s.MapInt64(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int16) int64 { return int64((t * 2)) }

	f(double, []int16{}, []int64{})
	f(double, []int16{1}, []int64{2})
	f(double, []int16{1, 2, 3}, []int64{2, 4, 6})
}

func TestAsyncSliceReduceInt16(t *testing.T) {
	f := func(reducer func(a int16, b int16) int16, given []int16, expected int16) {
		s := AsyncSliceInt16{Data: given, Workers: 4}
		actual := s.Reduce(reducer)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	sum := func(a int16, b int16) int16 { return a + b }

	f(sum, []int16{}, 0)
	f(sum, []int16{1}, 1)
	f(sum, []int16{1, 2}, 3)
	f(sum, []int16{1, 2, 3}, 6)
	f(sum, []int16{1, 2, 3, 4}, 10)
	f(sum, []int16{1, 2, 3, 4, 5}, 15)
}

func TestSlicesProductInt32(t *testing.T) {
	f := func(given [][]int32, expected [][]int32) {
		actual := make([][]int32, 0)
		i := 0
		s := SlicesInt32{given}
		for el := range s.Product() {
			actual = append(actual, el)
			i++
			if i > 50 {
				t.Fatal("infinite loop")
			}
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([][]int32{{1, 2}, {3, 4}}, [][]int32{{1, 3}, {1, 4}, {2, 3}, {2, 4}})
	f([][]int32{{1, 2}, {3}, {4, 5}}, [][]int32{{1, 3, 4}, {1, 3, 5}, {2, 3, 4}, {2, 3, 5}})
}

func TestSequenceCountInt32(t *testing.T) {
	f := func(start int32, step int32, count int, expected []int32) {
		ctx, cancel := context.WithCancel(context.Background())
		s := SequenceInt32{ctx: ctx}
		seq := s.Count(start, step)
		seq2 := ChannelInt32{seq}.Take(count)
		actual := ChannelInt32{seq2}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 2, 4, []int32{1, 3, 5, 7})
}

func TestSequenceExponentialInt32(t *testing.T) {
	f := func(start int32, factor int32, count int, expected []int32) {
		ctx, cancel := context.WithCancel(context.Background())
		s := SequenceInt32{ctx: ctx}
		seq := s.Exponential(start, factor)
		seq2 := ChannelInt32{seq}.Take(count)
		actual := ChannelInt32{seq2}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 1, 4, []int32{1, 1, 1, 1})
	f(1, 2, 4, []int32{1, 2, 4, 8})
}

func TestSequenceIterateInt32(t *testing.T) {
	f := func(start int32, count int, expected []int32) {
		ctx, cancel := context.WithCancel(context.Background())
		s := SequenceInt32{ctx: ctx}
		double := func(val int32) int32 { return val * 2 }
		seq := s.Iterate(start, double)
		seq2 := ChannelInt32{seq}.Take(count)
		actual := ChannelInt32{seq2}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 4, []int32{1, 2, 4, 8})
}

func TestSequenceRangeInt32(t *testing.T) {
	f := func(start int32, stop int32, step int32, expected []int32) {
		ctx, cancel := context.WithCancel(context.Background())
		s := SequenceInt32{ctx: ctx}
		seq := s.Range(start, stop, step)
		actual := ChannelInt32{seq}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 4, 1, []int32{1, 2, 3})
	f(3, 0, -1, []int32{3, 2, 1})
	f(1, 1, 1, []int32{})
	f(1, 2, 1, []int32{1})
}

func TestSequenceRepeatInt32(t *testing.T) {
	f := func(count int, given int32, expected []int32) {
		ctx, cancel := context.WithCancel(context.Background())
		s := SequenceInt32{ctx: ctx}
		seq := s.Repeat(given)
		seq2 := ChannelInt32{seq}.Take(count)
		actual := ChannelInt32{seq2}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, 1, []int32{1, 1})
}

func TestSequenceReplicateInt32(t *testing.T) {
	f := func(count int, given int32, expected []int32) {
		ctx, cancel := context.WithCancel(context.Background())
		s := SequenceInt32{ctx: ctx}
		seq := s.Replicate(given, count)
		actual := ChannelInt32{seq}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, 1, []int32{})
	f(1, 1, []int32{1})
	f(5, 1, []int32{1, 1, 1, 1, 1})
}

func TestSliceAnyInt32(t *testing.T) {
	f := func(given []int32, expected bool) {
		even := func(t int32) bool { return (t % 2) == 0 }
		actual := SliceInt32{given}.Any(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, false)
	f([]int32{1, 3}, false)
	f([]int32{2}, true)
	f([]int32{1, 2}, true)
}

func TestSliceAllInt32(t *testing.T) {
	f := func(given []int32, expected bool) {
		even := func(t int32) bool { return (t % 2) == 0 }
		actual := SliceInt32{given}.All(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, true)
	f([]int32{2}, true)
	f([]int32{1}, false)
	f([]int32{2, 4}, true)
	f([]int32{2, 4, 1}, false)
	f([]int32{1, 2, 4}, false)
}

func TestSliceChunkByInt32Int(t *testing.T) {
	f := func(given []int32, expected [][]int32) {
		reminder := func(t int32) int { return int((t % 2)) }
		actual := SliceInt32{given}.ChunkByInt(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, [][]int32{})
	f([]int32{1}, [][]int32{{1}})
	f([]int32{1, 2, 3}, [][]int32{{1}, {2}, {3}})
	f([]int32{1, 3, 2, 4, 5}, [][]int32{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkByInt32Int8(t *testing.T) {
	f := func(given []int32, expected [][]int32) {
		reminder := func(t int32) int8 { return int8((t % 2)) }
		actual := SliceInt32{given}.ChunkByInt8(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, [][]int32{})
	f([]int32{1}, [][]int32{{1}})
	f([]int32{1, 2, 3}, [][]int32{{1}, {2}, {3}})
	f([]int32{1, 3, 2, 4, 5}, [][]int32{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkByInt32Int16(t *testing.T) {
	f := func(given []int32, expected [][]int32) {
		reminder := func(t int32) int16 { return int16((t % 2)) }
		actual := SliceInt32{given}.ChunkByInt16(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, [][]int32{})
	f([]int32{1}, [][]int32{{1}})
	f([]int32{1, 2, 3}, [][]int32{{1}, {2}, {3}})
	f([]int32{1, 3, 2, 4, 5}, [][]int32{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkByInt32Int32(t *testing.T) {
	f := func(given []int32, expected [][]int32) {
		reminder := func(t int32) int32 { return int32((t % 2)) }
		actual := SliceInt32{given}.ChunkByInt32(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, [][]int32{})
	f([]int32{1}, [][]int32{{1}})
	f([]int32{1, 2, 3}, [][]int32{{1}, {2}, {3}})
	f([]int32{1, 3, 2, 4, 5}, [][]int32{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkByInt32Int64(t *testing.T) {
	f := func(given []int32, expected [][]int32) {
		reminder := func(t int32) int64 { return int64((t % 2)) }
		actual := SliceInt32{given}.ChunkByInt64(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, [][]int32{})
	f([]int32{1}, [][]int32{{1}})
	f([]int32{1, 2, 3}, [][]int32{{1}, {2}, {3}})
	f([]int32{1, 3, 2, 4, 5}, [][]int32{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkEveryInt32(t *testing.T) {
	f := func(count int, given []int32, expected [][]int32) {
		actual, _ := SliceInt32{given}.ChunkEvery(count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, []int32{}, [][]int32{})
	f(2, []int32{1}, [][]int32{{1}})
	f(2, []int32{1, 2, 3, 4}, [][]int32{{1, 2}, {3, 4}})
	f(2, []int32{1, 2, 3, 4, 5}, [][]int32{{1, 2}, {3, 4}, {5}})
}

func TestSliceContainsInt32(t *testing.T) {
	f := func(el int32, given []int32, expected bool) {
		actual := SliceInt32{given}.Contains(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, []int32{}, false)
	f(1, []int32{1}, true)
	f(1, []int32{2}, false)
	f(1, []int32{2, 3, 4, 5}, false)
	f(1, []int32{2, 3, 1, 4, 5}, true)
	f(1, []int32{2, 3, 1, 1, 4, 5}, true)
}

func TestSliceCountInt32(t *testing.T) {
	f := func(el int32, given []int32, expected int) {
		actual := SliceInt32{given}.Count(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, []int32{}, 0)
	f(1, []int32{1}, 1)
	f(1, []int32{2}, 0)
	f(1, []int32{2, 3, 4, 5}, 0)
	f(1, []int32{2, 3, 1, 4, 5}, 1)
	f(1, []int32{2, 3, 1, 1, 4, 5}, 2)
	f(1, []int32{1, 1, 1, 1, 1}, 5)
}

func TestSliceCountByInt32(t *testing.T) {
	f := func(given []int32, expected int) {
		even := func(t int32) bool { return (t % 2) == 0 }
		actual := SliceInt32{given}.CountBy(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, 0)
	f([]int32{1}, 0)
	f([]int32{2}, 1)
	f([]int32{1, 2, 3, 4, 5}, 2)
	f([]int32{1, 2, 3, 4, 5, 6}, 3)
}

func TestSliceCycleInt32(t *testing.T) {
	f := func(count int, given []int32, expected []int32) {
		c := SliceInt32{given}.Cycle()
		seq := ChannelInt32{c}.Take(count)
		actual := ChannelInt32{seq}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(5, []int32{}, []int32{})
	f(5, []int32{1}, []int32{1, 1, 1, 1, 1})
	f(5, []int32{1, 2}, []int32{1, 2, 1, 2, 1})
}

func TestSliceDedupInt32(t *testing.T) {
	f := func(given []int32, expected []int32) {
		actual := SliceInt32{given}.Dedup()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, []int32{})
	f([]int32{1}, []int32{1})
	f([]int32{1, 1}, []int32{1})
	f([]int32{1, 2}, []int32{1, 2})
	f([]int32{1, 2, 3}, []int32{1, 2, 3})
	f([]int32{1, 2, 2, 3}, []int32{1, 2, 3})
	f([]int32{1, 2, 2, 3, 3, 3, 2, 1, 1}, []int32{1, 2, 3, 2, 1})
}

func TestSliceDedupByInt32Int(t *testing.T) {
	f := func(given []int32, expected []int32) {
		even := func(el int32) int { return int(el % 2) }
		actual := SliceInt32{given}.DedupByInt(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, []int32{})
	f([]int32{1}, []int32{1})
	f([]int32{1, 1}, []int32{1})
	f([]int32{1, 2}, []int32{1, 2})
	f([]int32{1, 2, 3}, []int32{1, 2, 3})
	f([]int32{1, 2, 2, 3}, []int32{1, 2, 3})
	f([]int32{1, 2, 4, 3, 5, 7, 10}, []int32{1, 2, 3, 10})
}

func TestSliceDedupByInt32Int8(t *testing.T) {
	f := func(given []int32, expected []int32) {
		even := func(el int32) int8 { return int8(el % 2) }
		actual := SliceInt32{given}.DedupByInt8(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, []int32{})
	f([]int32{1}, []int32{1})
	f([]int32{1, 1}, []int32{1})
	f([]int32{1, 2}, []int32{1, 2})
	f([]int32{1, 2, 3}, []int32{1, 2, 3})
	f([]int32{1, 2, 2, 3}, []int32{1, 2, 3})
	f([]int32{1, 2, 4, 3, 5, 7, 10}, []int32{1, 2, 3, 10})
}

func TestSliceDedupByInt32Int16(t *testing.T) {
	f := func(given []int32, expected []int32) {
		even := func(el int32) int16 { return int16(el % 2) }
		actual := SliceInt32{given}.DedupByInt16(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, []int32{})
	f([]int32{1}, []int32{1})
	f([]int32{1, 1}, []int32{1})
	f([]int32{1, 2}, []int32{1, 2})
	f([]int32{1, 2, 3}, []int32{1, 2, 3})
	f([]int32{1, 2, 2, 3}, []int32{1, 2, 3})
	f([]int32{1, 2, 4, 3, 5, 7, 10}, []int32{1, 2, 3, 10})
}

func TestSliceDedupByInt32Int32(t *testing.T) {
	f := func(given []int32, expected []int32) {
		even := func(el int32) int32 { return int32(el % 2) }
		actual := SliceInt32{given}.DedupByInt32(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, []int32{})
	f([]int32{1}, []int32{1})
	f([]int32{1, 1}, []int32{1})
	f([]int32{1, 2}, []int32{1, 2})
	f([]int32{1, 2, 3}, []int32{1, 2, 3})
	f([]int32{1, 2, 2, 3}, []int32{1, 2, 3})
	f([]int32{1, 2, 4, 3, 5, 7, 10}, []int32{1, 2, 3, 10})
}

func TestSliceDedupByInt32Int64(t *testing.T) {
	f := func(given []int32, expected []int32) {
		even := func(el int32) int64 { return int64(el % 2) }
		actual := SliceInt32{given}.DedupByInt64(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, []int32{})
	f([]int32{1}, []int32{1})
	f([]int32{1, 1}, []int32{1})
	f([]int32{1, 2}, []int32{1, 2})
	f([]int32{1, 2, 3}, []int32{1, 2, 3})
	f([]int32{1, 2, 2, 3}, []int32{1, 2, 3})
	f([]int32{1, 2, 4, 3, 5, 7, 10}, []int32{1, 2, 3, 10})
}

func TestSliceDeleteInt32(t *testing.T) {
	f := func(given []int32, el int32, expected []int32) {
		actual := SliceInt32{given}.Delete(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, 1, []int32{})
	f([]int32{1}, 1, []int32{})
	f([]int32{2}, 1, []int32{2})
	f([]int32{1, 2}, 1, []int32{2})
	f([]int32{1, 2, 3}, 2, []int32{1, 3})
	f([]int32{1, 2, 2, 3, 2}, 2, []int32{1, 2, 3, 2})
}

func TestSliceDeleteAllInt32(t *testing.T) {
	f := func(given []int32, el int32, expected []int32) {
		actual := SliceInt32{given}.DeleteAll(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, 1, []int32{})
	f([]int32{1}, 1, []int32{})
	f([]int32{2}, 1, []int32{2})
	f([]int32{1, 2}, 1, []int32{2})
	f([]int32{1, 2, 3}, 2, []int32{1, 3})
	f([]int32{1, 2, 2, 3, 2}, 2, []int32{1, 3})
}

func TestSliceDeleteAtInt32(t *testing.T) {
	f := func(given []int32, indices []int, expected []int32) {
		actual, _ := SliceInt32{given}.DeleteAt(indices...)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, []int{}, []int32{})
	f([]int32{1}, []int{0}, []int32{})
	f([]int32{1, 2}, []int{0}, []int32{2})

	f([]int32{1, 2, 3}, []int{0}, []int32{2, 3})
	f([]int32{1, 2, 3}, []int{1}, []int32{1, 3})
	f([]int32{1, 2, 3}, []int{2}, []int32{1, 2})

	f([]int32{1, 2, 3}, []int{0, 1}, []int32{3})
	f([]int32{1, 2, 3}, []int{0, 2}, []int32{2})
	f([]int32{1, 2, 3}, []int{1, 2}, []int32{1})
}

func TestSliceDropEveryInt32(t *testing.T) {
	f := func(given []int32, nth int, from int, expected []int32) {
		actual, _ := SliceInt32{given}.DropEvery(nth, from)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, 1, 1, []int32{})
	f([]int32{1, 2, 3}, 1, 1, []int32{})

	f([]int32{1, 2, 3, 4}, 2, 1, []int32{1, 3})
	f([]int32{1, 2, 3, 4, 5}, 2, 1, []int32{1, 3, 5})

	f([]int32{1, 2, 3, 4}, 2, 0, []int32{2, 4})
	f([]int32{1, 2, 3, 4, 5}, 2, 0, []int32{2, 4})

	f([]int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2, 1, []int32{1, 3, 5, 7, 9})
	f([]int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, 1, []int32{1, 2, 4, 5, 7, 8, 10})
}

func TestSliceDropWhileInt32(t *testing.T) {
	f := func(given []int32, expected []int32) {
		even := func(el int32) bool { return el%2 == 0 }
		actual := SliceInt32{given}.DropWhile(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, []int32{})
	f([]int32{2}, []int32{})
	f([]int32{1}, []int32{1})
	f([]int32{2, 1}, []int32{1})
	f([]int32{2, 1, 2}, []int32{1, 2})
	f([]int32{1, 2}, []int32{1, 2})
	f([]int32{2, 4, 6, 1, 8}, []int32{1, 8})
}

func TestSliceEndsWithInt32(t *testing.T) {
	f := func(given []int32, suffix []int32, expected bool) {
		actual := SliceInt32{given}.EndsWith(suffix)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, []int32{}, true)
	f([]int32{1}, []int32{1}, true)
	f([]int32{1}, []int32{2}, false)
	f([]int32{2, 3}, []int32{1, 2, 3}, false)

	f([]int32{1, 2, 3}, []int32{3}, true)
	f([]int32{1, 2, 3}, []int32{2, 3}, true)
	f([]int32{1, 2, 3}, []int32{1, 2, 3}, true)

	f([]int32{1, 2, 3}, []int32{1}, false)
	f([]int32{1, 2, 3}, []int32{2}, false)
	f([]int32{1, 2, 3}, []int32{1, 2}, false)
	f([]int32{1, 2, 3}, []int32{3, 2}, false)
}

func TestSliceEqualInt32(t *testing.T) {
	f := func(left []int32, right []int32, expected bool) {
		actual := SliceInt32{left}.Equal(right)
		assert.Equal(t, expected, actual, "they should be equal")

		actual = SliceInt32{right}.Equal(left)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, []int32{}, true)
	f([]int32{1}, []int32{1}, true)
	f([]int32{1}, []int32{2}, false)
	f([]int32{1, 2, 3, 3}, []int32{1, 2, 3, 3}, true)
	f([]int32{1, 2, 3, 3}, []int32{1, 2, 2, 3}, false)
	f([]int32{1, 2, 3, 3}, []int32{1, 2, 4, 3}, false)

	// different len
	f([]int32{1, 2, 3}, []int32{1, 2}, false)
	f([]int32{1, 2}, []int32{1, 2, 3}, false)
	f([]int32{}, []int32{1, 2, 3}, false)
	f([]int32{1, 2, 3}, []int32{}, false)
}

func TestSliceFilterInt32(t *testing.T) {
	f := func(given []int32, expected []int32) {
		even := func(t int32) bool { return (t % 2) == 0 }
		actual := SliceInt32{given}.Filter(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, []int32{})
	f([]int32{1, 2, 3, 4}, []int32{2, 4})
	f([]int32{1, 3}, []int32{})
	f([]int32{2, 4}, []int32{2, 4})
}

func TestSliceFindInt32(t *testing.T) {
	f := func(given []int32, expectedEl int32, expectedErr error) {
		even := func(t int32) bool { return (t % 2) == 0 }
		el, err := SliceInt32{given}.Find(even)
		assert.Equal(t, expectedEl, el, "they should be equal")
		assert.Equal(t, expectedErr, err, "they should be equal")
	}
	f([]int32{}, 0, ErrNotFound)
	f([]int32{1}, 0, ErrNotFound)
	f([]int32{1}, 0, ErrNotFound)
	f([]int32{2}, 2, nil)
	f([]int32{1, 2}, 2, nil)
	f([]int32{1, 2, 3}, 2, nil)
	f([]int32{1, 3, 5}, 0, ErrNotFound)
}

func TestSliceFindIndexInt32(t *testing.T) {
	f := func(given []int32, expectedInd int, expectedErr error) {
		even := func(t int32) bool { return (t % 2) == 0 }
		index, err := SliceInt32{given}.FindIndex(even)
		assert.Equal(t, expectedInd, index, "they should be equal")
		assert.Equal(t, expectedErr, err, "they should be equal")
	}
	f([]int32{}, 0, ErrNotFound)
	f([]int32{1}, 0, ErrNotFound)
	f([]int32{1}, 0, ErrNotFound)
	f([]int32{2}, 0, nil)
	f([]int32{1, 2}, 1, nil)
	f([]int32{1, 2, 3}, 1, nil)
	f([]int32{1, 3, 5, 7, 9, 2}, 5, nil)
	f([]int32{1, 3, 5}, 0, ErrNotFound)
}

func TestSliceJoinInt32(t *testing.T) {
	f := func(given []int32, sep string, expected string) {
		actual := SliceInt32{given}.Join(sep)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, "", "")
	f([]int32{}, "|", "")

	f([]int32{1}, "", "1")
	f([]int32{1}, "|", "1")

	f([]int32{1, 2, 3}, "", "123")
	f([]int32{1, 2, 3}, "|", "1|2|3")
	f([]int32{1, 2, 3}, "<int32>", "1<int32>2<int32>3")
}

func TestSliceGroupByInt32Int(t *testing.T) {
	f := func(given []int32, expected map[int][]int32) {
		reminder := func(t int32) int { return int((t % 2)) }
		actual := SliceInt32{given}.GroupByInt(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, map[int][]int32{})
	f([]int32{1}, map[int][]int32{1: {1}})
	f([]int32{1, 3, 2, 4, 5}, map[int][]int32{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceGroupByInt32Int8(t *testing.T) {
	f := func(given []int32, expected map[int8][]int32) {
		reminder := func(t int32) int8 { return int8((t % 2)) }
		actual := SliceInt32{given}.GroupByInt8(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, map[int8][]int32{})
	f([]int32{1}, map[int8][]int32{1: {1}})
	f([]int32{1, 3, 2, 4, 5}, map[int8][]int32{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceGroupByInt32Int16(t *testing.T) {
	f := func(given []int32, expected map[int16][]int32) {
		reminder := func(t int32) int16 { return int16((t % 2)) }
		actual := SliceInt32{given}.GroupByInt16(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, map[int16][]int32{})
	f([]int32{1}, map[int16][]int32{1: {1}})
	f([]int32{1, 3, 2, 4, 5}, map[int16][]int32{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceGroupByInt32Int32(t *testing.T) {
	f := func(given []int32, expected map[int32][]int32) {
		reminder := func(t int32) int32 { return int32((t % 2)) }
		actual := SliceInt32{given}.GroupByInt32(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, map[int32][]int32{})
	f([]int32{1}, map[int32][]int32{1: {1}})
	f([]int32{1, 3, 2, 4, 5}, map[int32][]int32{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceGroupByInt32Int64(t *testing.T) {
	f := func(given []int32, expected map[int64][]int32) {
		reminder := func(t int32) int64 { return int64((t % 2)) }
		actual := SliceInt32{given}.GroupByInt64(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, map[int64][]int32{})
	f([]int32{1}, map[int64][]int32{1: {1}})
	f([]int32{1, 3, 2, 4, 5}, map[int64][]int32{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceInsertAtInt32(t *testing.T) {
	f := func(given []int32, index int, expected []int32, expectedErr error) {
		actual, err := SliceInt32{given}.InsertAt(index, 10)
		assert.Equal(t, expected, actual, "they should be equal")
		assert.Equal(t, expectedErr, err, "they should be equal")
	}
	f([]int32{}, -1, []int32{}, ErrNegativeValue)
	f([]int32{}, 0, []int32{10}, nil)
	f([]int32{}, 1, []int32{}, ErrOutOfRange)

	f([]int32{1, 2, 3}, -1, []int32{1, 2, 3}, ErrNegativeValue)
	f([]int32{1, 2, 3}, 0, []int32{10, 1, 2, 3}, nil)
	f([]int32{1, 2, 3}, 1, []int32{1, 10, 2, 3}, nil)
	f([]int32{1, 2, 3}, 3, []int32{1, 2, 3, 10}, nil)
	f([]int32{1, 2, 3}, 4, []int32{1, 2, 3}, ErrOutOfRange)
}

func TestSliceIntersperseInt32(t *testing.T) {
	f := func(el int32, given []int32, expected []int32) {
		actual := SliceInt32{given}.Intersperse(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, []int32{}, []int32{})
	f(0, []int32{1}, []int32{1})
	f(0, []int32{1, 2}, []int32{1, 0, 2})
	f(0, []int32{1, 2, 3}, []int32{1, 0, 2, 0, 3})
}

func TestSliceLastInt32(t *testing.T) {
	f := func(given []int32, expectedEl int32, expectedErr error) {
		el, err := SliceInt32{given}.Last()
		assert.Equal(t, expectedEl, el, "they should be equal")
		assert.Equal(t, expectedErr, err, "they should be equal")
	}
	f([]int32{}, 0, ErrEmpty)
	f([]int32{1}, 1, nil)
	f([]int32{1, 2, 3}, 3, nil)
}

func TestSliceMapInt32Int(t *testing.T) {
	f := func(given []int32, expected []int) {
		double := func(t int32) int { return int((t * 2)) }
		actual := SliceInt32{given}.MapInt(double)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, []int{})
	f([]int32{1}, []int{2})
	f([]int32{1, 2, 3}, []int{2, 4, 6})
}

func TestSliceMapInt32Int8(t *testing.T) {
	f := func(given []int32, expected []int8) {
		double := func(t int32) int8 { return int8((t * 2)) }
		actual := SliceInt32{given}.MapInt8(double)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, []int8{})
	f([]int32{1}, []int8{2})
	f([]int32{1, 2, 3}, []int8{2, 4, 6})
}

func TestSliceMapInt32Int16(t *testing.T) {
	f := func(given []int32, expected []int16) {
		double := func(t int32) int16 { return int16((t * 2)) }
		actual := SliceInt32{given}.MapInt16(double)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, []int16{})
	f([]int32{1}, []int16{2})
	f([]int32{1, 2, 3}, []int16{2, 4, 6})
}

func TestSliceMapInt32Int32(t *testing.T) {
	f := func(given []int32, expected []int32) {
		double := func(t int32) int32 { return int32((t * 2)) }
		actual := SliceInt32{given}.MapInt32(double)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, []int32{})
	f([]int32{1}, []int32{2})
	f([]int32{1, 2, 3}, []int32{2, 4, 6})
}

func TestSliceMapInt32Int64(t *testing.T) {
	f := func(given []int32, expected []int64) {
		double := func(t int32) int64 { return int64((t * 2)) }
		actual := SliceInt32{given}.MapInt64(double)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, []int64{})
	f([]int32{1}, []int64{2})
	f([]int32{1, 2, 3}, []int64{2, 4, 6})
}

func TestSliceMaxInt32(t *testing.T) {
	f := func(given []int32, expectedEl int32, expectedErr error) {
		el, err := SliceInt32{given}.Max()
		assert.Equal(t, expectedEl, el, "they should be equal")
		assert.Equal(t, expectedErr, err, "they should be equal")
	}
	f([]int32{}, 0, ErrEmpty)
	f([]int32{1}, 1, nil)
	f([]int32{1, 2, 3}, 3, nil)
	f([]int32{1, 3, 2}, 3, nil)
	f([]int32{3, 2, 1}, 3, nil)
}

func TestSliceMinInt32(t *testing.T) {
	f := func(given []int32, expectedEl int32, expectedErr error) {
		el, err := SliceInt32{given}.Min()
		assert.Equal(t, expectedEl, el, "they should be equal")
		assert.Equal(t, expectedErr, err, "they should be equal")
	}
	f([]int32{}, 0, ErrEmpty)
	f([]int32{1}, 1, nil)
	f([]int32{1, 2, 3}, 1, nil)
	f([]int32{2, 1, 3}, 1, nil)
	f([]int32{3, 2, 1}, 1, nil)
}

func TestSlicesPermutationsInt32(t *testing.T) {
	f := func(size int, given []int32, expected [][]int32) {
		actual := make([][]int32, 0)
		i := 0
		s := SliceInt32{given}
		for el := range s.Permutations(size) {
			actual = append(actual, el)
			i++
			if i > 50 {
				t.Fatal("infinite loop")
			}
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, []int32{}, [][]int32{})
	f(2, []int32{1}, [][]int32{{1}})
	f(2, []int32{1, 2, 3}, [][]int32{{1, 2}, {1, 3}, {2, 1}, {2, 3}, {3, 1}, {3, 2}})
}

func TestSliceProductInt32(t *testing.T) {
	f := func(given []int32, repeat int, expected [][]int32) {
		actual := make([][]int32, 0)
		i := 0
		s := SliceInt32{given}
		for el := range s.Product(repeat) {
			actual = append(actual, el)
			i++
			if i > 50 {
				t.Fatal("infinite loop")
			}
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}

	f([]int32{1, 2}, 0, [][]int32{})
	f([]int32{}, 2, [][]int32{})
	f([]int32{1}, 2, [][]int32{{1, 1}})

	f([]int32{1, 2}, 1, [][]int32{{1}, {2}})
	f([]int32{1, 2}, 2, [][]int32{{1, 1}, {1, 2}, {2, 1}, {2, 2}})
	f([]int32{1, 2}, 3, [][]int32{
		{1, 1, 1}, {1, 1, 2}, {1, 2, 1}, {1, 2, 2},
		{2, 1, 1}, {2, 1, 2}, {2, 2, 1}, {2, 2, 2},
	})
}

func TestSliceReduceInt32Int(t *testing.T) {
	f := func(given []int32, expected int) {
		sum := func(el int32, acc int) int { return int(el) + acc }
		actual := SliceInt32{given}.ReduceInt(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, 0)
	f([]int32{1}, 1)
	f([]int32{1, 2}, 3)
	f([]int32{1, 2, 3}, 6)
}

func TestSliceReduceInt32Int8(t *testing.T) {
	f := func(given []int32, expected int8) {
		sum := func(el int32, acc int8) int8 { return int8(el) + acc }
		actual := SliceInt32{given}.ReduceInt8(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, 0)
	f([]int32{1}, 1)
	f([]int32{1, 2}, 3)
	f([]int32{1, 2, 3}, 6)
}

func TestSliceReduceInt32Int16(t *testing.T) {
	f := func(given []int32, expected int16) {
		sum := func(el int32, acc int16) int16 { return int16(el) + acc }
		actual := SliceInt32{given}.ReduceInt16(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, 0)
	f([]int32{1}, 1)
	f([]int32{1, 2}, 3)
	f([]int32{1, 2, 3}, 6)
}

func TestSliceReduceInt32Int32(t *testing.T) {
	f := func(given []int32, expected int32) {
		sum := func(el int32, acc int32) int32 { return int32(el) + acc }
		actual := SliceInt32{given}.ReduceInt32(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, 0)
	f([]int32{1}, 1)
	f([]int32{1, 2}, 3)
	f([]int32{1, 2, 3}, 6)
}

func TestSliceReduceInt32Int64(t *testing.T) {
	f := func(given []int32, expected int64) {
		sum := func(el int32, acc int64) int64 { return int64(el) + acc }
		actual := SliceInt32{given}.ReduceInt64(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, 0)
	f([]int32{1}, 1)
	f([]int32{1, 2}, 3)
	f([]int32{1, 2, 3}, 6)
}

func TestSliceReduceWhileInt32Int(t *testing.T) {
	f := func(given []int32, expected int) {
		sum := func(el int32, acc int) (int, error) {
			if el == 0 {
				return acc, ErrEmpty
			}
			return int(el) + acc, nil
		}
		actual, _ := SliceInt32{given}.ReduceWhileInt(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, 0)
	f([]int32{1}, 1)
	f([]int32{1, 2}, 3)
	f([]int32{1, 2, 3}, 6)
	f([]int32{1, 2, 0, 3}, 3)
}

func TestSliceReduceWhileInt32Int8(t *testing.T) {
	f := func(given []int32, expected int8) {
		sum := func(el int32, acc int8) (int8, error) {
			if el == 0 {
				return acc, ErrEmpty
			}
			return int8(el) + acc, nil
		}
		actual, _ := SliceInt32{given}.ReduceWhileInt8(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, 0)
	f([]int32{1}, 1)
	f([]int32{1, 2}, 3)
	f([]int32{1, 2, 3}, 6)
	f([]int32{1, 2, 0, 3}, 3)
}

func TestSliceReduceWhileInt32Int16(t *testing.T) {
	f := func(given []int32, expected int16) {
		sum := func(el int32, acc int16) (int16, error) {
			if el == 0 {
				return acc, ErrEmpty
			}
			return int16(el) + acc, nil
		}
		actual, _ := SliceInt32{given}.ReduceWhileInt16(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, 0)
	f([]int32{1}, 1)
	f([]int32{1, 2}, 3)
	f([]int32{1, 2, 3}, 6)
	f([]int32{1, 2, 0, 3}, 3)
}

func TestSliceReduceWhileInt32Int32(t *testing.T) {
	f := func(given []int32, expected int32) {
		sum := func(el int32, acc int32) (int32, error) {
			if el == 0 {
				return acc, ErrEmpty
			}
			return int32(el) + acc, nil
		}
		actual, _ := SliceInt32{given}.ReduceWhileInt32(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, 0)
	f([]int32{1}, 1)
	f([]int32{1, 2}, 3)
	f([]int32{1, 2, 3}, 6)
	f([]int32{1, 2, 0, 3}, 3)
}

func TestSliceReduceWhileInt32Int64(t *testing.T) {
	f := func(given []int32, expected int64) {
		sum := func(el int32, acc int64) (int64, error) {
			if el == 0 {
				return acc, ErrEmpty
			}
			return int64(el) + acc, nil
		}
		actual, _ := SliceInt32{given}.ReduceWhileInt64(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, 0)
	f([]int32{1}, 1)
	f([]int32{1, 2}, 3)
	f([]int32{1, 2, 3}, 6)
	f([]int32{1, 2, 0, 3}, 3)
}

func TestChannelToSliceInt32(t *testing.T) {
	f := func(given []int32) {
		c := make(chan int32, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := ChannelInt32{c}.ToSlice()
		assert.Equal(t, given, actual, "they should be equal")
	}
	f([]int32{})
	f([]int32{1})
	f([]int32{1, 2, 3, 1, 2})
}

func TestChannelAnyInt32(t *testing.T) {
	f := func(given []int32, expected bool) {
		even := func(t int32) bool { return t%2 == 0 }
		c := make(chan int32, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := ChannelInt32{c}.Any(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, false)
	f([]int32{1}, false)
	f([]int32{2}, true)
	f([]int32{1, 2}, true)
	f([]int32{1, 2, 3}, true)
	f([]int32{1, 3, 5}, false)
	f([]int32{1, 3, 5, 7, 9, 11}, false)
	f([]int32{1, 3, 5, 7, 10, 11}, true)
}

func TestChannelAllInt32(t *testing.T) {
	f := func(given []int32, expected bool) {
		even := func(t int32) bool { return t%2 == 0 }
		c := make(chan int32, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := ChannelInt32{c}.All(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, true)
	f([]int32{1}, false)
	f([]int32{2}, true)
	f([]int32{1, 2}, false)
	f([]int32{2, 4}, true)
	f([]int32{2, 4, 6, 8, 10, 12}, true)
	f([]int32{2, 4, 6, 8, 11, 12}, false)
}

func TestChannelEachInt32(t *testing.T) {
	f := func(given []int32) {
		c := make(chan int32, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := make(chan int32, len(given))
		mapper := func(t int32) { result <- t }
		ChannelInt32{c}.Each(mapper)
		close(result)
		actual := ChannelInt32{result}.ToSlice()
		assert.Equal(t, given, actual, "they should be equal")
	}

	f([]int32{})
	f([]int32{1})
	f([]int32{1, 2, 3})
	f([]int32{1, 2, 3, 4, 5, 6, 7})
}

func TestChannelChunkEveryInt32(t *testing.T) {
	f := func(size int, given []int32, expected [][]int32) {
		c := make(chan int32, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := ChannelInt32{c}.ChunkEvery(size)
		actual := make([][]int32, 0)
		for el := range result {
			actual = append(actual, el)
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, []int32{}, [][]int32{})
	f(2, []int32{1}, [][]int32{{1}})
	f(2, []int32{1, 2}, [][]int32{{1, 2}})
	f(2, []int32{1, 2, 3, 4}, [][]int32{{1, 2}, {3, 4}})
	f(2, []int32{1, 2, 3, 4, 5}, [][]int32{{1, 2}, {3, 4}, {5}})
}

func TestChannelCountInt32(t *testing.T) {
	f := func(element int32, given []int32, expected int) {
		c := make(chan int32, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := ChannelInt32{c}.Count(element)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, []int32{}, 0)
	f(1, []int32{1}, 1)
	f(1, []int32{2}, 0)
	f(1, []int32{1, 2, 3, 1, 4}, 2)
}

func TestChannelDropInt32(t *testing.T) {
	f := func(count int, given []int32, expected []int32) {
		c := make(chan int32, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := ChannelInt32{c}.Drop(count)
		actual := make([]int32, 0)
		for el := range result {
			actual = append(actual, el)
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, []int32{}, []int32{})
	f(1, []int32{2}, []int32{})
	f(1, []int32{2, 3}, []int32{3})
	f(1, []int32{1, 2, 3}, []int32{2, 3})
	f(0, []int32{1, 2, 3}, []int32{1, 2, 3})
	f(3, []int32{1, 2, 3, 4, 5, 6}, []int32{4, 5, 6})
	f(1, []int32{1, 2, 3, 4, 5, 6}, []int32{2, 3, 4, 5, 6})
}

func TestChannelFilterInt32(t *testing.T) {
	f := func(given []int32, expected []int32) {
		even := func(t int32) bool { return t%2 == 0 }
		c := make(chan int32, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := ChannelInt32{c}.Filter(even)
		actual := ChannelInt32{result}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, []int32{})
	f([]int32{1}, []int32{})
	f([]int32{2}, []int32{2})
	f([]int32{1, 2, 3, 4}, []int32{2, 4})
}

func TestChannelMapInt32Int(t *testing.T) {
	f := func(given []int32, expected []int32) {
		double := func(el int32) int { return int(el * 2) }
		c := make(chan int32, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := ChannelInt32{c}.MapInt(double)

		// convert chan int32 to chan int
		c2 := make(chan int32, 1)
		go func() {
			for el := range result {
				c2 <- int32(el)
			}
			close(c2)
		}()

		actual := ChannelInt32{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, []int32{})
	f([]int32{1}, []int32{2})
	f([]int32{1, 2, 3}, []int32{2, 4, 6})
}

func TestChannelMapInt32Int8(t *testing.T) {
	f := func(given []int32, expected []int32) {
		double := func(el int32) int8 { return int8(el * 2) }
		c := make(chan int32, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := ChannelInt32{c}.MapInt8(double)

		// convert chan int32 to chan int8
		c2 := make(chan int32, 1)
		go func() {
			for el := range result {
				c2 <- int32(el)
			}
			close(c2)
		}()

		actual := ChannelInt32{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, []int32{})
	f([]int32{1}, []int32{2})
	f([]int32{1, 2, 3}, []int32{2, 4, 6})
}

func TestChannelMapInt32Int16(t *testing.T) {
	f := func(given []int32, expected []int32) {
		double := func(el int32) int16 { return int16(el * 2) }
		c := make(chan int32, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := ChannelInt32{c}.MapInt16(double)

		// convert chan int32 to chan int16
		c2 := make(chan int32, 1)
		go func() {
			for el := range result {
				c2 <- int32(el)
			}
			close(c2)
		}()

		actual := ChannelInt32{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, []int32{})
	f([]int32{1}, []int32{2})
	f([]int32{1, 2, 3}, []int32{2, 4, 6})
}

func TestChannelMapInt32Int32(t *testing.T) {
	f := func(given []int32, expected []int32) {
		double := func(el int32) int32 { return int32(el * 2) }
		c := make(chan int32, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := ChannelInt32{c}.MapInt32(double)

		// convert chan int32 to chan int32
		c2 := make(chan int32, 1)
		go func() {
			for el := range result {
				c2 <- int32(el)
			}
			close(c2)
		}()

		actual := ChannelInt32{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, []int32{})
	f([]int32{1}, []int32{2})
	f([]int32{1, 2, 3}, []int32{2, 4, 6})
}

func TestChannelMapInt32Int64(t *testing.T) {
	f := func(given []int32, expected []int32) {
		double := func(el int32) int64 { return int64(el * 2) }
		c := make(chan int32, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := ChannelInt32{c}.MapInt64(double)

		// convert chan int32 to chan int64
		c2 := make(chan int32, 1)
		go func() {
			for el := range result {
				c2 <- int32(el)
			}
			close(c2)
		}()

		actual := ChannelInt32{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, []int32{})
	f([]int32{1}, []int32{2})
	f([]int32{1, 2, 3}, []int32{2, 4, 6})
}

func TestChannelMaxInt32(t *testing.T) {
	f := func(given []int32, expected int32, expectedErr error) {
		c := make(chan int32, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual, actualErr := ChannelInt32{c}.Max()
		assert.Equal(t, expected, actual, "they should be equal")
		assert.Equal(t, expectedErr, actualErr, "they should be equal")
	}
	f([]int32{}, 0, ErrEmpty)
	f([]int32{1, 4, 2}, 4, nil)
	f([]int32{1, 2, 4}, 4, nil)
	f([]int32{4, 2, 1}, 4, nil)
}

func TestChannelMinInt32(t *testing.T) {
	f := func(given []int32, expected int32, expectedErr error) {
		c := make(chan int32, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual, actualErr := ChannelInt32{c}.Min()
		assert.Equal(t, expected, actual, "they should be equal")
		assert.Equal(t, expectedErr, actualErr, "they should be equal")
	}
	f([]int32{}, 0, ErrEmpty)
	f([]int32{4, 1, 2}, 1, nil)
	f([]int32{1, 2, 4}, 1, nil)
	f([]int32{4, 2, 1}, 1, nil)
}

func TestChannelReduceInt32Int(t *testing.T) {
	f := func(given []int32, expected int) {
		c := make(chan int32, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int32, acc int) int { return int(el) + acc }
		actual := ChannelInt32{c}.ReduceInt(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, 0)
	f([]int32{1}, 1)
	f([]int32{1, 2}, 3)
	f([]int32{1, 2, 3, 4, 5}, 15)
}

func TestChannelReduceInt32Int8(t *testing.T) {
	f := func(given []int32, expected int8) {
		c := make(chan int32, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int32, acc int8) int8 { return int8(el) + acc }
		actual := ChannelInt32{c}.ReduceInt8(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, 0)
	f([]int32{1}, 1)
	f([]int32{1, 2}, 3)
	f([]int32{1, 2, 3, 4, 5}, 15)
}

func TestChannelReduceInt32Int16(t *testing.T) {
	f := func(given []int32, expected int16) {
		c := make(chan int32, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int32, acc int16) int16 { return int16(el) + acc }
		actual := ChannelInt32{c}.ReduceInt16(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, 0)
	f([]int32{1}, 1)
	f([]int32{1, 2}, 3)
	f([]int32{1, 2, 3, 4, 5}, 15)
}

func TestChannelReduceInt32Int32(t *testing.T) {
	f := func(given []int32, expected int32) {
		c := make(chan int32, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int32, acc int32) int32 { return int32(el) + acc }
		actual := ChannelInt32{c}.ReduceInt32(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, 0)
	f([]int32{1}, 1)
	f([]int32{1, 2}, 3)
	f([]int32{1, 2, 3, 4, 5}, 15)
}

func TestChannelReduceInt32Int64(t *testing.T) {
	f := func(given []int32, expected int64) {
		c := make(chan int32, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int32, acc int64) int64 { return int64(el) + acc }
		actual := ChannelInt32{c}.ReduceInt64(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, 0)
	f([]int32{1}, 1)
	f([]int32{1, 2}, 3)
	f([]int32{1, 2, 3, 4, 5}, 15)
}

func TestChannelScanInt32Int(t *testing.T) {
	f := func(given []int32, expected []int32) {
		c := make(chan int32, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int32, acc int) int { return int(el) + acc }
		result := ChannelInt32{c}.ScanInt(0, sum)

		// convert chan int32 to chan int
		c2 := make(chan int32, 1)
		go func() {
			for el := range result {
				c2 <- int32(el)
			}
			close(c2)
		}()

		actual := ChannelInt32{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, []int32{})
	f([]int32{1}, []int32{1})
	f([]int32{1, 2}, []int32{1, 3})
	f([]int32{1, 2, 3, 4, 5}, []int32{1, 3, 6, 10, 15})
}

func TestChannelScanInt32Int8(t *testing.T) {
	f := func(given []int32, expected []int32) {
		c := make(chan int32, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int32, acc int8) int8 { return int8(el) + acc }
		result := ChannelInt32{c}.ScanInt8(0, sum)

		// convert chan int32 to chan int8
		c2 := make(chan int32, 1)
		go func() {
			for el := range result {
				c2 <- int32(el)
			}
			close(c2)
		}()

		actual := ChannelInt32{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, []int32{})
	f([]int32{1}, []int32{1})
	f([]int32{1, 2}, []int32{1, 3})
	f([]int32{1, 2, 3, 4, 5}, []int32{1, 3, 6, 10, 15})
}

func TestChannelScanInt32Int16(t *testing.T) {
	f := func(given []int32, expected []int32) {
		c := make(chan int32, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int32, acc int16) int16 { return int16(el) + acc }
		result := ChannelInt32{c}.ScanInt16(0, sum)

		// convert chan int32 to chan int16
		c2 := make(chan int32, 1)
		go func() {
			for el := range result {
				c2 <- int32(el)
			}
			close(c2)
		}()

		actual := ChannelInt32{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, []int32{})
	f([]int32{1}, []int32{1})
	f([]int32{1, 2}, []int32{1, 3})
	f([]int32{1, 2, 3, 4, 5}, []int32{1, 3, 6, 10, 15})
}

func TestChannelScanInt32Int32(t *testing.T) {
	f := func(given []int32, expected []int32) {
		c := make(chan int32, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int32, acc int32) int32 { return int32(el) + acc }
		result := ChannelInt32{c}.ScanInt32(0, sum)

		// convert chan int32 to chan int32
		c2 := make(chan int32, 1)
		go func() {
			for el := range result {
				c2 <- int32(el)
			}
			close(c2)
		}()

		actual := ChannelInt32{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, []int32{})
	f([]int32{1}, []int32{1})
	f([]int32{1, 2}, []int32{1, 3})
	f([]int32{1, 2, 3, 4, 5}, []int32{1, 3, 6, 10, 15})
}

func TestChannelScanInt32Int64(t *testing.T) {
	f := func(given []int32, expected []int32) {
		c := make(chan int32, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int32, acc int64) int64 { return int64(el) + acc }
		result := ChannelInt32{c}.ScanInt64(0, sum)

		// convert chan int32 to chan int64
		c2 := make(chan int32, 1)
		go func() {
			for el := range result {
				c2 <- int32(el)
			}
			close(c2)
		}()

		actual := ChannelInt32{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, []int32{})
	f([]int32{1}, []int32{1})
	f([]int32{1, 2}, []int32{1, 3})
	f([]int32{1, 2, 3, 4, 5}, []int32{1, 3, 6, 10, 15})
}

func TestChannelSumInt32(t *testing.T) {
	f := func(given []int32, expected int32) {
		c := make(chan int32, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := ChannelInt32{c}.Sum()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int32{}, 0)
	f([]int32{1}, 1)
	f([]int32{1, 2}, 3)
	f([]int32{1, 2, 3, 4, 5}, 15)
}

func TestChannelTakeInt32(t *testing.T) {
	f := func(count int, given int32, expected []int32) {
		ctx, cancel := context.WithCancel(context.Background())
		s := SequenceInt32{ctx: ctx}
		seq := s.Repeat(given)
		seq2 := ChannelInt32{seq}.Take(count)
		actual := ChannelInt32{seq2}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, 1, []int32{})
	f(1, 1, []int32{1})
	f(2, 1, []int32{1, 1})
}

func TestChannelTeeInt32(t *testing.T) {
	f := func(count int, given []int32) {
		c := make(chan int32, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		channels := ChannelInt32{c}.Tee(count)
		for _, ch := range channels {
			go func(ch chan int32) {
				actual := ChannelInt32{ch}.ToSlice()
				assert.Equal(t, given, actual, "they should be equal")
			}(ch)
		}
	}
	f(1, []int32{})
	f(1, []int32{1})
	f(1, []int32{1, 2})
	f(1, []int32{1, 2, 3})
	f(1, []int32{1, 2, 3, 1, 2})

	f(2, []int32{})
	f(2, []int32{1})
	f(2, []int32{1, 2})
	f(2, []int32{1, 2, 3})
	f(2, []int32{1, 2, 3, 1, 2})

	f(10, []int32{1, 2, 3, 1, 2})
}

func TestAsyncSliceAnyInt32(t *testing.T) {
	f := func(check func(t int32) bool, given []int32, expected bool) {
		s := AsyncSliceInt32{Data: given, Workers: 2}
		actual := s.Any(check)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	isEven := func(t int32) bool { return (t % 2) == 0 }

	f(isEven, []int32{}, false)
	f(isEven, []int32{1}, false)
	f(isEven, []int32{1, 3}, false)
	f(isEven, []int32{2}, true)
	f(isEven, []int32{1, 2}, true)
	f(isEven, []int32{1, 3, 5, 7, 9, 11}, false)
	f(isEven, []int32{1, 3, 5, 7, 9, 12}, true)
}

func TestAsyncSliceAllInt32(t *testing.T) {
	f := func(check func(t int32) bool, given []int32, expected bool) {
		s := AsyncSliceInt32{Data: given, Workers: 2}
		actual := s.All(check)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	isEven := func(t int32) bool { return (t % 2) == 0 }

	f(isEven, []int32{}, true)
	f(isEven, []int32{1}, false)
	f(isEven, []int32{1, 3}, false)
	f(isEven, []int32{2}, true)
	f(isEven, []int32{2, 4}, true)
	f(isEven, []int32{2, 3}, false)
	f(isEven, []int32{2, 4, 6, 8, 10, 12}, true)
	f(isEven, []int32{2, 4, 6, 8, 10, 11}, false)
}

func TestAsyncSliceEachInt32(t *testing.T) {
	f := func(given []int32) {
		s := AsyncSliceInt32{Data: given, Workers: 2}
		result := make(chan int32, len(given))
		mapper := func(t int32) { result <- t }
		s.Each(mapper)
		close(result)
		actual := ChannelInt32{result}.ToSlice()
		sorted := SliceInt32{actual}.Sort()
		assert.Equal(t, given, sorted, "they should be equal")
	}

	f([]int32{})
	f([]int32{1})
	f([]int32{1, 2, 3})
	f([]int32{1, 2, 3, 4, 5, 6, 7})
}

func TestAsyncSliceFilterInt32(t *testing.T) {
	f := func(given []int32, expected []int32) {
		filter := func(t int32) bool { return t > 10 }
		s := AsyncSliceInt32{Data: given, Workers: 2}
		actual := s.Filter(filter)
		assert.Equal(t, expected, actual, "they should be equal")
	}

	f([]int32{}, []int32{})
	f([]int32{5}, []int32{})
	f([]int32{15}, []int32{15})
	f([]int32{9, 11, 12, 13, 6}, []int32{11, 12, 13})
}

func TestAsyncSliceMapInt32Int(t *testing.T) {
	f := func(mapper func(t int32) int, given []int32, expected []int) {
		s := AsyncSliceInt32{Data: given, Workers: 2}
		actual := s.MapInt(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int32) int { return int((t * 2)) }

	f(double, []int32{}, []int{})
	f(double, []int32{1}, []int{2})
	f(double, []int32{1, 2, 3}, []int{2, 4, 6})
}

func TestAsyncSliceMapInt32Int8(t *testing.T) {
	f := func(mapper func(t int32) int8, given []int32, expected []int8) {
		s := AsyncSliceInt32{Data: given, Workers: 2}
		actual := s.MapInt8(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int32) int8 { return int8((t * 2)) }

	f(double, []int32{}, []int8{})
	f(double, []int32{1}, []int8{2})
	f(double, []int32{1, 2, 3}, []int8{2, 4, 6})
}

func TestAsyncSliceMapInt32Int16(t *testing.T) {
	f := func(mapper func(t int32) int16, given []int32, expected []int16) {
		s := AsyncSliceInt32{Data: given, Workers: 2}
		actual := s.MapInt16(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int32) int16 { return int16((t * 2)) }

	f(double, []int32{}, []int16{})
	f(double, []int32{1}, []int16{2})
	f(double, []int32{1, 2, 3}, []int16{2, 4, 6})
}

func TestAsyncSliceMapInt32Int32(t *testing.T) {
	f := func(mapper func(t int32) int32, given []int32, expected []int32) {
		s := AsyncSliceInt32{Data: given, Workers: 2}
		actual := s.MapInt32(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int32) int32 { return int32((t * 2)) }

	f(double, []int32{}, []int32{})
	f(double, []int32{1}, []int32{2})
	f(double, []int32{1, 2, 3}, []int32{2, 4, 6})
}

func TestAsyncSliceMapInt32Int64(t *testing.T) {
	f := func(mapper func(t int32) int64, given []int32, expected []int64) {
		s := AsyncSliceInt32{Data: given, Workers: 2}
		actual := s.MapInt64(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int32) int64 { return int64((t * 2)) }

	f(double, []int32{}, []int64{})
	f(double, []int32{1}, []int64{2})
	f(double, []int32{1, 2, 3}, []int64{2, 4, 6})
}

func TestAsyncSliceReduceInt32(t *testing.T) {
	f := func(reducer func(a int32, b int32) int32, given []int32, expected int32) {
		s := AsyncSliceInt32{Data: given, Workers: 4}
		actual := s.Reduce(reducer)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	sum := func(a int32, b int32) int32 { return a + b }

	f(sum, []int32{}, 0)
	f(sum, []int32{1}, 1)
	f(sum, []int32{1, 2}, 3)
	f(sum, []int32{1, 2, 3}, 6)
	f(sum, []int32{1, 2, 3, 4}, 10)
	f(sum, []int32{1, 2, 3, 4, 5}, 15)
}

func TestSlicesProductInt64(t *testing.T) {
	f := func(given [][]int64, expected [][]int64) {
		actual := make([][]int64, 0)
		i := 0
		s := SlicesInt64{given}
		for el := range s.Product() {
			actual = append(actual, el)
			i++
			if i > 50 {
				t.Fatal("infinite loop")
			}
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([][]int64{{1, 2}, {3, 4}}, [][]int64{{1, 3}, {1, 4}, {2, 3}, {2, 4}})
	f([][]int64{{1, 2}, {3}, {4, 5}}, [][]int64{{1, 3, 4}, {1, 3, 5}, {2, 3, 4}, {2, 3, 5}})
}

func TestSequenceCountInt64(t *testing.T) {
	f := func(start int64, step int64, count int, expected []int64) {
		ctx, cancel := context.WithCancel(context.Background())
		s := SequenceInt64{ctx: ctx}
		seq := s.Count(start, step)
		seq2 := ChannelInt64{seq}.Take(count)
		actual := ChannelInt64{seq2}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 2, 4, []int64{1, 3, 5, 7})
}

func TestSequenceExponentialInt64(t *testing.T) {
	f := func(start int64, factor int64, count int, expected []int64) {
		ctx, cancel := context.WithCancel(context.Background())
		s := SequenceInt64{ctx: ctx}
		seq := s.Exponential(start, factor)
		seq2 := ChannelInt64{seq}.Take(count)
		actual := ChannelInt64{seq2}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 1, 4, []int64{1, 1, 1, 1})
	f(1, 2, 4, []int64{1, 2, 4, 8})
}

func TestSequenceIterateInt64(t *testing.T) {
	f := func(start int64, count int, expected []int64) {
		ctx, cancel := context.WithCancel(context.Background())
		s := SequenceInt64{ctx: ctx}
		double := func(val int64) int64 { return val * 2 }
		seq := s.Iterate(start, double)
		seq2 := ChannelInt64{seq}.Take(count)
		actual := ChannelInt64{seq2}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 4, []int64{1, 2, 4, 8})
}

func TestSequenceRangeInt64(t *testing.T) {
	f := func(start int64, stop int64, step int64, expected []int64) {
		ctx, cancel := context.WithCancel(context.Background())
		s := SequenceInt64{ctx: ctx}
		seq := s.Range(start, stop, step)
		actual := ChannelInt64{seq}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, 4, 1, []int64{1, 2, 3})
	f(3, 0, -1, []int64{3, 2, 1})
	f(1, 1, 1, []int64{})
	f(1, 2, 1, []int64{1})
}

func TestSequenceRepeatInt64(t *testing.T) {
	f := func(count int, given int64, expected []int64) {
		ctx, cancel := context.WithCancel(context.Background())
		s := SequenceInt64{ctx: ctx}
		seq := s.Repeat(given)
		seq2 := ChannelInt64{seq}.Take(count)
		actual := ChannelInt64{seq2}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, 1, []int64{1, 1})
}

func TestSequenceReplicateInt64(t *testing.T) {
	f := func(count int, given int64, expected []int64) {
		ctx, cancel := context.WithCancel(context.Background())
		s := SequenceInt64{ctx: ctx}
		seq := s.Replicate(given, count)
		actual := ChannelInt64{seq}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, 1, []int64{})
	f(1, 1, []int64{1})
	f(5, 1, []int64{1, 1, 1, 1, 1})
}

func TestSliceAnyInt64(t *testing.T) {
	f := func(given []int64, expected bool) {
		even := func(t int64) bool { return (t % 2) == 0 }
		actual := SliceInt64{given}.Any(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, false)
	f([]int64{1, 3}, false)
	f([]int64{2}, true)
	f([]int64{1, 2}, true)
}

func TestSliceAllInt64(t *testing.T) {
	f := func(given []int64, expected bool) {
		even := func(t int64) bool { return (t % 2) == 0 }
		actual := SliceInt64{given}.All(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, true)
	f([]int64{2}, true)
	f([]int64{1}, false)
	f([]int64{2, 4}, true)
	f([]int64{2, 4, 1}, false)
	f([]int64{1, 2, 4}, false)
}

func TestSliceChunkByInt64Int(t *testing.T) {
	f := func(given []int64, expected [][]int64) {
		reminder := func(t int64) int { return int((t % 2)) }
		actual := SliceInt64{given}.ChunkByInt(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, [][]int64{})
	f([]int64{1}, [][]int64{{1}})
	f([]int64{1, 2, 3}, [][]int64{{1}, {2}, {3}})
	f([]int64{1, 3, 2, 4, 5}, [][]int64{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkByInt64Int8(t *testing.T) {
	f := func(given []int64, expected [][]int64) {
		reminder := func(t int64) int8 { return int8((t % 2)) }
		actual := SliceInt64{given}.ChunkByInt8(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, [][]int64{})
	f([]int64{1}, [][]int64{{1}})
	f([]int64{1, 2, 3}, [][]int64{{1}, {2}, {3}})
	f([]int64{1, 3, 2, 4, 5}, [][]int64{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkByInt64Int16(t *testing.T) {
	f := func(given []int64, expected [][]int64) {
		reminder := func(t int64) int16 { return int16((t % 2)) }
		actual := SliceInt64{given}.ChunkByInt16(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, [][]int64{})
	f([]int64{1}, [][]int64{{1}})
	f([]int64{1, 2, 3}, [][]int64{{1}, {2}, {3}})
	f([]int64{1, 3, 2, 4, 5}, [][]int64{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkByInt64Int32(t *testing.T) {
	f := func(given []int64, expected [][]int64) {
		reminder := func(t int64) int32 { return int32((t % 2)) }
		actual := SliceInt64{given}.ChunkByInt32(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, [][]int64{})
	f([]int64{1}, [][]int64{{1}})
	f([]int64{1, 2, 3}, [][]int64{{1}, {2}, {3}})
	f([]int64{1, 3, 2, 4, 5}, [][]int64{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkByInt64Int64(t *testing.T) {
	f := func(given []int64, expected [][]int64) {
		reminder := func(t int64) int64 { return int64((t % 2)) }
		actual := SliceInt64{given}.ChunkByInt64(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, [][]int64{})
	f([]int64{1}, [][]int64{{1}})
	f([]int64{1, 2, 3}, [][]int64{{1}, {2}, {3}})
	f([]int64{1, 3, 2, 4, 5}, [][]int64{{1, 3}, {2, 4}, {5}})
}

func TestSliceChunkEveryInt64(t *testing.T) {
	f := func(count int, given []int64, expected [][]int64) {
		actual, _ := SliceInt64{given}.ChunkEvery(count)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, []int64{}, [][]int64{})
	f(2, []int64{1}, [][]int64{{1}})
	f(2, []int64{1, 2, 3, 4}, [][]int64{{1, 2}, {3, 4}})
	f(2, []int64{1, 2, 3, 4, 5}, [][]int64{{1, 2}, {3, 4}, {5}})
}

func TestSliceContainsInt64(t *testing.T) {
	f := func(el int64, given []int64, expected bool) {
		actual := SliceInt64{given}.Contains(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, []int64{}, false)
	f(1, []int64{1}, true)
	f(1, []int64{2}, false)
	f(1, []int64{2, 3, 4, 5}, false)
	f(1, []int64{2, 3, 1, 4, 5}, true)
	f(1, []int64{2, 3, 1, 1, 4, 5}, true)
}

func TestSliceCountInt64(t *testing.T) {
	f := func(el int64, given []int64, expected int) {
		actual := SliceInt64{given}.Count(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, []int64{}, 0)
	f(1, []int64{1}, 1)
	f(1, []int64{2}, 0)
	f(1, []int64{2, 3, 4, 5}, 0)
	f(1, []int64{2, 3, 1, 4, 5}, 1)
	f(1, []int64{2, 3, 1, 1, 4, 5}, 2)
	f(1, []int64{1, 1, 1, 1, 1}, 5)
}

func TestSliceCountByInt64(t *testing.T) {
	f := func(given []int64, expected int) {
		even := func(t int64) bool { return (t % 2) == 0 }
		actual := SliceInt64{given}.CountBy(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, 0)
	f([]int64{1}, 0)
	f([]int64{2}, 1)
	f([]int64{1, 2, 3, 4, 5}, 2)
	f([]int64{1, 2, 3, 4, 5, 6}, 3)
}

func TestSliceCycleInt64(t *testing.T) {
	f := func(count int, given []int64, expected []int64) {
		c := SliceInt64{given}.Cycle()
		seq := ChannelInt64{c}.Take(count)
		actual := ChannelInt64{seq}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(5, []int64{}, []int64{})
	f(5, []int64{1}, []int64{1, 1, 1, 1, 1})
	f(5, []int64{1, 2}, []int64{1, 2, 1, 2, 1})
}

func TestSliceDedupInt64(t *testing.T) {
	f := func(given []int64, expected []int64) {
		actual := SliceInt64{given}.Dedup()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, []int64{})
	f([]int64{1}, []int64{1})
	f([]int64{1, 1}, []int64{1})
	f([]int64{1, 2}, []int64{1, 2})
	f([]int64{1, 2, 3}, []int64{1, 2, 3})
	f([]int64{1, 2, 2, 3}, []int64{1, 2, 3})
	f([]int64{1, 2, 2, 3, 3, 3, 2, 1, 1}, []int64{1, 2, 3, 2, 1})
}

func TestSliceDedupByInt64Int(t *testing.T) {
	f := func(given []int64, expected []int64) {
		even := func(el int64) int { return int(el % 2) }
		actual := SliceInt64{given}.DedupByInt(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, []int64{})
	f([]int64{1}, []int64{1})
	f([]int64{1, 1}, []int64{1})
	f([]int64{1, 2}, []int64{1, 2})
	f([]int64{1, 2, 3}, []int64{1, 2, 3})
	f([]int64{1, 2, 2, 3}, []int64{1, 2, 3})
	f([]int64{1, 2, 4, 3, 5, 7, 10}, []int64{1, 2, 3, 10})
}

func TestSliceDedupByInt64Int8(t *testing.T) {
	f := func(given []int64, expected []int64) {
		even := func(el int64) int8 { return int8(el % 2) }
		actual := SliceInt64{given}.DedupByInt8(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, []int64{})
	f([]int64{1}, []int64{1})
	f([]int64{1, 1}, []int64{1})
	f([]int64{1, 2}, []int64{1, 2})
	f([]int64{1, 2, 3}, []int64{1, 2, 3})
	f([]int64{1, 2, 2, 3}, []int64{1, 2, 3})
	f([]int64{1, 2, 4, 3, 5, 7, 10}, []int64{1, 2, 3, 10})
}

func TestSliceDedupByInt64Int16(t *testing.T) {
	f := func(given []int64, expected []int64) {
		even := func(el int64) int16 { return int16(el % 2) }
		actual := SliceInt64{given}.DedupByInt16(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, []int64{})
	f([]int64{1}, []int64{1})
	f([]int64{1, 1}, []int64{1})
	f([]int64{1, 2}, []int64{1, 2})
	f([]int64{1, 2, 3}, []int64{1, 2, 3})
	f([]int64{1, 2, 2, 3}, []int64{1, 2, 3})
	f([]int64{1, 2, 4, 3, 5, 7, 10}, []int64{1, 2, 3, 10})
}

func TestSliceDedupByInt64Int32(t *testing.T) {
	f := func(given []int64, expected []int64) {
		even := func(el int64) int32 { return int32(el % 2) }
		actual := SliceInt64{given}.DedupByInt32(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, []int64{})
	f([]int64{1}, []int64{1})
	f([]int64{1, 1}, []int64{1})
	f([]int64{1, 2}, []int64{1, 2})
	f([]int64{1, 2, 3}, []int64{1, 2, 3})
	f([]int64{1, 2, 2, 3}, []int64{1, 2, 3})
	f([]int64{1, 2, 4, 3, 5, 7, 10}, []int64{1, 2, 3, 10})
}

func TestSliceDedupByInt64Int64(t *testing.T) {
	f := func(given []int64, expected []int64) {
		even := func(el int64) int64 { return int64(el % 2) }
		actual := SliceInt64{given}.DedupByInt64(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, []int64{})
	f([]int64{1}, []int64{1})
	f([]int64{1, 1}, []int64{1})
	f([]int64{1, 2}, []int64{1, 2})
	f([]int64{1, 2, 3}, []int64{1, 2, 3})
	f([]int64{1, 2, 2, 3}, []int64{1, 2, 3})
	f([]int64{1, 2, 4, 3, 5, 7, 10}, []int64{1, 2, 3, 10})
}

func TestSliceDeleteInt64(t *testing.T) {
	f := func(given []int64, el int64, expected []int64) {
		actual := SliceInt64{given}.Delete(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, 1, []int64{})
	f([]int64{1}, 1, []int64{})
	f([]int64{2}, 1, []int64{2})
	f([]int64{1, 2}, 1, []int64{2})
	f([]int64{1, 2, 3}, 2, []int64{1, 3})
	f([]int64{1, 2, 2, 3, 2}, 2, []int64{1, 2, 3, 2})
}

func TestSliceDeleteAllInt64(t *testing.T) {
	f := func(given []int64, el int64, expected []int64) {
		actual := SliceInt64{given}.DeleteAll(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, 1, []int64{})
	f([]int64{1}, 1, []int64{})
	f([]int64{2}, 1, []int64{2})
	f([]int64{1, 2}, 1, []int64{2})
	f([]int64{1, 2, 3}, 2, []int64{1, 3})
	f([]int64{1, 2, 2, 3, 2}, 2, []int64{1, 3})
}

func TestSliceDeleteAtInt64(t *testing.T) {
	f := func(given []int64, indices []int, expected []int64) {
		actual, _ := SliceInt64{given}.DeleteAt(indices...)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, []int{}, []int64{})
	f([]int64{1}, []int{0}, []int64{})
	f([]int64{1, 2}, []int{0}, []int64{2})

	f([]int64{1, 2, 3}, []int{0}, []int64{2, 3})
	f([]int64{1, 2, 3}, []int{1}, []int64{1, 3})
	f([]int64{1, 2, 3}, []int{2}, []int64{1, 2})

	f([]int64{1, 2, 3}, []int{0, 1}, []int64{3})
	f([]int64{1, 2, 3}, []int{0, 2}, []int64{2})
	f([]int64{1, 2, 3}, []int{1, 2}, []int64{1})
}

func TestSliceDropEveryInt64(t *testing.T) {
	f := func(given []int64, nth int, from int, expected []int64) {
		actual, _ := SliceInt64{given}.DropEvery(nth, from)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, 1, 1, []int64{})
	f([]int64{1, 2, 3}, 1, 1, []int64{})

	f([]int64{1, 2, 3, 4}, 2, 1, []int64{1, 3})
	f([]int64{1, 2, 3, 4, 5}, 2, 1, []int64{1, 3, 5})

	f([]int64{1, 2, 3, 4}, 2, 0, []int64{2, 4})
	f([]int64{1, 2, 3, 4, 5}, 2, 0, []int64{2, 4})

	f([]int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2, 1, []int64{1, 3, 5, 7, 9})
	f([]int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, 1, []int64{1, 2, 4, 5, 7, 8, 10})
}

func TestSliceDropWhileInt64(t *testing.T) {
	f := func(given []int64, expected []int64) {
		even := func(el int64) bool { return el%2 == 0 }
		actual := SliceInt64{given}.DropWhile(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, []int64{})
	f([]int64{2}, []int64{})
	f([]int64{1}, []int64{1})
	f([]int64{2, 1}, []int64{1})
	f([]int64{2, 1, 2}, []int64{1, 2})
	f([]int64{1, 2}, []int64{1, 2})
	f([]int64{2, 4, 6, 1, 8}, []int64{1, 8})
}

func TestSliceEndsWithInt64(t *testing.T) {
	f := func(given []int64, suffix []int64, expected bool) {
		actual := SliceInt64{given}.EndsWith(suffix)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, []int64{}, true)
	f([]int64{1}, []int64{1}, true)
	f([]int64{1}, []int64{2}, false)
	f([]int64{2, 3}, []int64{1, 2, 3}, false)

	f([]int64{1, 2, 3}, []int64{3}, true)
	f([]int64{1, 2, 3}, []int64{2, 3}, true)
	f([]int64{1, 2, 3}, []int64{1, 2, 3}, true)

	f([]int64{1, 2, 3}, []int64{1}, false)
	f([]int64{1, 2, 3}, []int64{2}, false)
	f([]int64{1, 2, 3}, []int64{1, 2}, false)
	f([]int64{1, 2, 3}, []int64{3, 2}, false)
}

func TestSliceEqualInt64(t *testing.T) {
	f := func(left []int64, right []int64, expected bool) {
		actual := SliceInt64{left}.Equal(right)
		assert.Equal(t, expected, actual, "they should be equal")

		actual = SliceInt64{right}.Equal(left)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, []int64{}, true)
	f([]int64{1}, []int64{1}, true)
	f([]int64{1}, []int64{2}, false)
	f([]int64{1, 2, 3, 3}, []int64{1, 2, 3, 3}, true)
	f([]int64{1, 2, 3, 3}, []int64{1, 2, 2, 3}, false)
	f([]int64{1, 2, 3, 3}, []int64{1, 2, 4, 3}, false)

	// different len
	f([]int64{1, 2, 3}, []int64{1, 2}, false)
	f([]int64{1, 2}, []int64{1, 2, 3}, false)
	f([]int64{}, []int64{1, 2, 3}, false)
	f([]int64{1, 2, 3}, []int64{}, false)
}

func TestSliceFilterInt64(t *testing.T) {
	f := func(given []int64, expected []int64) {
		even := func(t int64) bool { return (t % 2) == 0 }
		actual := SliceInt64{given}.Filter(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, []int64{})
	f([]int64{1, 2, 3, 4}, []int64{2, 4})
	f([]int64{1, 3}, []int64{})
	f([]int64{2, 4}, []int64{2, 4})
}

func TestSliceFindInt64(t *testing.T) {
	f := func(given []int64, expectedEl int64, expectedErr error) {
		even := func(t int64) bool { return (t % 2) == 0 }
		el, err := SliceInt64{given}.Find(even)
		assert.Equal(t, expectedEl, el, "they should be equal")
		assert.Equal(t, expectedErr, err, "they should be equal")
	}
	f([]int64{}, 0, ErrNotFound)
	f([]int64{1}, 0, ErrNotFound)
	f([]int64{1}, 0, ErrNotFound)
	f([]int64{2}, 2, nil)
	f([]int64{1, 2}, 2, nil)
	f([]int64{1, 2, 3}, 2, nil)
	f([]int64{1, 3, 5}, 0, ErrNotFound)
}

func TestSliceFindIndexInt64(t *testing.T) {
	f := func(given []int64, expectedInd int, expectedErr error) {
		even := func(t int64) bool { return (t % 2) == 0 }
		index, err := SliceInt64{given}.FindIndex(even)
		assert.Equal(t, expectedInd, index, "they should be equal")
		assert.Equal(t, expectedErr, err, "they should be equal")
	}
	f([]int64{}, 0, ErrNotFound)
	f([]int64{1}, 0, ErrNotFound)
	f([]int64{1}, 0, ErrNotFound)
	f([]int64{2}, 0, nil)
	f([]int64{1, 2}, 1, nil)
	f([]int64{1, 2, 3}, 1, nil)
	f([]int64{1, 3, 5, 7, 9, 2}, 5, nil)
	f([]int64{1, 3, 5}, 0, ErrNotFound)
}

func TestSliceJoinInt64(t *testing.T) {
	f := func(given []int64, sep string, expected string) {
		actual := SliceInt64{given}.Join(sep)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, "", "")
	f([]int64{}, "|", "")

	f([]int64{1}, "", "1")
	f([]int64{1}, "|", "1")

	f([]int64{1, 2, 3}, "", "123")
	f([]int64{1, 2, 3}, "|", "1|2|3")
	f([]int64{1, 2, 3}, "<int64>", "1<int64>2<int64>3")
}

func TestSliceGroupByInt64Int(t *testing.T) {
	f := func(given []int64, expected map[int][]int64) {
		reminder := func(t int64) int { return int((t % 2)) }
		actual := SliceInt64{given}.GroupByInt(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, map[int][]int64{})
	f([]int64{1}, map[int][]int64{1: {1}})
	f([]int64{1, 3, 2, 4, 5}, map[int][]int64{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceGroupByInt64Int8(t *testing.T) {
	f := func(given []int64, expected map[int8][]int64) {
		reminder := func(t int64) int8 { return int8((t % 2)) }
		actual := SliceInt64{given}.GroupByInt8(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, map[int8][]int64{})
	f([]int64{1}, map[int8][]int64{1: {1}})
	f([]int64{1, 3, 2, 4, 5}, map[int8][]int64{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceGroupByInt64Int16(t *testing.T) {
	f := func(given []int64, expected map[int16][]int64) {
		reminder := func(t int64) int16 { return int16((t % 2)) }
		actual := SliceInt64{given}.GroupByInt16(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, map[int16][]int64{})
	f([]int64{1}, map[int16][]int64{1: {1}})
	f([]int64{1, 3, 2, 4, 5}, map[int16][]int64{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceGroupByInt64Int32(t *testing.T) {
	f := func(given []int64, expected map[int32][]int64) {
		reminder := func(t int64) int32 { return int32((t % 2)) }
		actual := SliceInt64{given}.GroupByInt32(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, map[int32][]int64{})
	f([]int64{1}, map[int32][]int64{1: {1}})
	f([]int64{1, 3, 2, 4, 5}, map[int32][]int64{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceGroupByInt64Int64(t *testing.T) {
	f := func(given []int64, expected map[int64][]int64) {
		reminder := func(t int64) int64 { return int64((t % 2)) }
		actual := SliceInt64{given}.GroupByInt64(reminder)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, map[int64][]int64{})
	f([]int64{1}, map[int64][]int64{1: {1}})
	f([]int64{1, 3, 2, 4, 5}, map[int64][]int64{0: {2, 4}, 1: {1, 3, 5}})
}

func TestSliceInsertAtInt64(t *testing.T) {
	f := func(given []int64, index int, expected []int64, expectedErr error) {
		actual, err := SliceInt64{given}.InsertAt(index, 10)
		assert.Equal(t, expected, actual, "they should be equal")
		assert.Equal(t, expectedErr, err, "they should be equal")
	}
	f([]int64{}, -1, []int64{}, ErrNegativeValue)
	f([]int64{}, 0, []int64{10}, nil)
	f([]int64{}, 1, []int64{}, ErrOutOfRange)

	f([]int64{1, 2, 3}, -1, []int64{1, 2, 3}, ErrNegativeValue)
	f([]int64{1, 2, 3}, 0, []int64{10, 1, 2, 3}, nil)
	f([]int64{1, 2, 3}, 1, []int64{1, 10, 2, 3}, nil)
	f([]int64{1, 2, 3}, 3, []int64{1, 2, 3, 10}, nil)
	f([]int64{1, 2, 3}, 4, []int64{1, 2, 3}, ErrOutOfRange)
}

func TestSliceIntersperseInt64(t *testing.T) {
	f := func(el int64, given []int64, expected []int64) {
		actual := SliceInt64{given}.Intersperse(el)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, []int64{}, []int64{})
	f(0, []int64{1}, []int64{1})
	f(0, []int64{1, 2}, []int64{1, 0, 2})
	f(0, []int64{1, 2, 3}, []int64{1, 0, 2, 0, 3})
}

func TestSliceLastInt64(t *testing.T) {
	f := func(given []int64, expectedEl int64, expectedErr error) {
		el, err := SliceInt64{given}.Last()
		assert.Equal(t, expectedEl, el, "they should be equal")
		assert.Equal(t, expectedErr, err, "they should be equal")
	}
	f([]int64{}, 0, ErrEmpty)
	f([]int64{1}, 1, nil)
	f([]int64{1, 2, 3}, 3, nil)
}

func TestSliceMapInt64Int(t *testing.T) {
	f := func(given []int64, expected []int) {
		double := func(t int64) int { return int((t * 2)) }
		actual := SliceInt64{given}.MapInt(double)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, []int{})
	f([]int64{1}, []int{2})
	f([]int64{1, 2, 3}, []int{2, 4, 6})
}

func TestSliceMapInt64Int8(t *testing.T) {
	f := func(given []int64, expected []int8) {
		double := func(t int64) int8 { return int8((t * 2)) }
		actual := SliceInt64{given}.MapInt8(double)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, []int8{})
	f([]int64{1}, []int8{2})
	f([]int64{1, 2, 3}, []int8{2, 4, 6})
}

func TestSliceMapInt64Int16(t *testing.T) {
	f := func(given []int64, expected []int16) {
		double := func(t int64) int16 { return int16((t * 2)) }
		actual := SliceInt64{given}.MapInt16(double)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, []int16{})
	f([]int64{1}, []int16{2})
	f([]int64{1, 2, 3}, []int16{2, 4, 6})
}

func TestSliceMapInt64Int32(t *testing.T) {
	f := func(given []int64, expected []int32) {
		double := func(t int64) int32 { return int32((t * 2)) }
		actual := SliceInt64{given}.MapInt32(double)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, []int32{})
	f([]int64{1}, []int32{2})
	f([]int64{1, 2, 3}, []int32{2, 4, 6})
}

func TestSliceMapInt64Int64(t *testing.T) {
	f := func(given []int64, expected []int64) {
		double := func(t int64) int64 { return int64((t * 2)) }
		actual := SliceInt64{given}.MapInt64(double)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, []int64{})
	f([]int64{1}, []int64{2})
	f([]int64{1, 2, 3}, []int64{2, 4, 6})
}

func TestSliceMaxInt64(t *testing.T) {
	f := func(given []int64, expectedEl int64, expectedErr error) {
		el, err := SliceInt64{given}.Max()
		assert.Equal(t, expectedEl, el, "they should be equal")
		assert.Equal(t, expectedErr, err, "they should be equal")
	}
	f([]int64{}, 0, ErrEmpty)
	f([]int64{1}, 1, nil)
	f([]int64{1, 2, 3}, 3, nil)
	f([]int64{1, 3, 2}, 3, nil)
	f([]int64{3, 2, 1}, 3, nil)
}

func TestSliceMinInt64(t *testing.T) {
	f := func(given []int64, expectedEl int64, expectedErr error) {
		el, err := SliceInt64{given}.Min()
		assert.Equal(t, expectedEl, el, "they should be equal")
		assert.Equal(t, expectedErr, err, "they should be equal")
	}
	f([]int64{}, 0, ErrEmpty)
	f([]int64{1}, 1, nil)
	f([]int64{1, 2, 3}, 1, nil)
	f([]int64{2, 1, 3}, 1, nil)
	f([]int64{3, 2, 1}, 1, nil)
}

func TestSlicesPermutationsInt64(t *testing.T) {
	f := func(size int, given []int64, expected [][]int64) {
		actual := make([][]int64, 0)
		i := 0
		s := SliceInt64{given}
		for el := range s.Permutations(size) {
			actual = append(actual, el)
			i++
			if i > 50 {
				t.Fatal("infinite loop")
			}
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, []int64{}, [][]int64{})
	f(2, []int64{1}, [][]int64{{1}})
	f(2, []int64{1, 2, 3}, [][]int64{{1, 2}, {1, 3}, {2, 1}, {2, 3}, {3, 1}, {3, 2}})
}

func TestSliceProductInt64(t *testing.T) {
	f := func(given []int64, repeat int, expected [][]int64) {
		actual := make([][]int64, 0)
		i := 0
		s := SliceInt64{given}
		for el := range s.Product(repeat) {
			actual = append(actual, el)
			i++
			if i > 50 {
				t.Fatal("infinite loop")
			}
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}

	f([]int64{1, 2}, 0, [][]int64{})
	f([]int64{}, 2, [][]int64{})
	f([]int64{1}, 2, [][]int64{{1, 1}})

	f([]int64{1, 2}, 1, [][]int64{{1}, {2}})
	f([]int64{1, 2}, 2, [][]int64{{1, 1}, {1, 2}, {2, 1}, {2, 2}})
	f([]int64{1, 2}, 3, [][]int64{
		{1, 1, 1}, {1, 1, 2}, {1, 2, 1}, {1, 2, 2},
		{2, 1, 1}, {2, 1, 2}, {2, 2, 1}, {2, 2, 2},
	})
}

func TestSliceReduceInt64Int(t *testing.T) {
	f := func(given []int64, expected int) {
		sum := func(el int64, acc int) int { return int(el) + acc }
		actual := SliceInt64{given}.ReduceInt(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, 0)
	f([]int64{1}, 1)
	f([]int64{1, 2}, 3)
	f([]int64{1, 2, 3}, 6)
}

func TestSliceReduceInt64Int8(t *testing.T) {
	f := func(given []int64, expected int8) {
		sum := func(el int64, acc int8) int8 { return int8(el) + acc }
		actual := SliceInt64{given}.ReduceInt8(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, 0)
	f([]int64{1}, 1)
	f([]int64{1, 2}, 3)
	f([]int64{1, 2, 3}, 6)
}

func TestSliceReduceInt64Int16(t *testing.T) {
	f := func(given []int64, expected int16) {
		sum := func(el int64, acc int16) int16 { return int16(el) + acc }
		actual := SliceInt64{given}.ReduceInt16(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, 0)
	f([]int64{1}, 1)
	f([]int64{1, 2}, 3)
	f([]int64{1, 2, 3}, 6)
}

func TestSliceReduceInt64Int32(t *testing.T) {
	f := func(given []int64, expected int32) {
		sum := func(el int64, acc int32) int32 { return int32(el) + acc }
		actual := SliceInt64{given}.ReduceInt32(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, 0)
	f([]int64{1}, 1)
	f([]int64{1, 2}, 3)
	f([]int64{1, 2, 3}, 6)
}

func TestSliceReduceInt64Int64(t *testing.T) {
	f := func(given []int64, expected int64) {
		sum := func(el int64, acc int64) int64 { return int64(el) + acc }
		actual := SliceInt64{given}.ReduceInt64(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, 0)
	f([]int64{1}, 1)
	f([]int64{1, 2}, 3)
	f([]int64{1, 2, 3}, 6)
}

func TestSliceReduceWhileInt64Int(t *testing.T) {
	f := func(given []int64, expected int) {
		sum := func(el int64, acc int) (int, error) {
			if el == 0 {
				return acc, ErrEmpty
			}
			return int(el) + acc, nil
		}
		actual, _ := SliceInt64{given}.ReduceWhileInt(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, 0)
	f([]int64{1}, 1)
	f([]int64{1, 2}, 3)
	f([]int64{1, 2, 3}, 6)
	f([]int64{1, 2, 0, 3}, 3)
}

func TestSliceReduceWhileInt64Int8(t *testing.T) {
	f := func(given []int64, expected int8) {
		sum := func(el int64, acc int8) (int8, error) {
			if el == 0 {
				return acc, ErrEmpty
			}
			return int8(el) + acc, nil
		}
		actual, _ := SliceInt64{given}.ReduceWhileInt8(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, 0)
	f([]int64{1}, 1)
	f([]int64{1, 2}, 3)
	f([]int64{1, 2, 3}, 6)
	f([]int64{1, 2, 0, 3}, 3)
}

func TestSliceReduceWhileInt64Int16(t *testing.T) {
	f := func(given []int64, expected int16) {
		sum := func(el int64, acc int16) (int16, error) {
			if el == 0 {
				return acc, ErrEmpty
			}
			return int16(el) + acc, nil
		}
		actual, _ := SliceInt64{given}.ReduceWhileInt16(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, 0)
	f([]int64{1}, 1)
	f([]int64{1, 2}, 3)
	f([]int64{1, 2, 3}, 6)
	f([]int64{1, 2, 0, 3}, 3)
}

func TestSliceReduceWhileInt64Int32(t *testing.T) {
	f := func(given []int64, expected int32) {
		sum := func(el int64, acc int32) (int32, error) {
			if el == 0 {
				return acc, ErrEmpty
			}
			return int32(el) + acc, nil
		}
		actual, _ := SliceInt64{given}.ReduceWhileInt32(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, 0)
	f([]int64{1}, 1)
	f([]int64{1, 2}, 3)
	f([]int64{1, 2, 3}, 6)
	f([]int64{1, 2, 0, 3}, 3)
}

func TestSliceReduceWhileInt64Int64(t *testing.T) {
	f := func(given []int64, expected int64) {
		sum := func(el int64, acc int64) (int64, error) {
			if el == 0 {
				return acc, ErrEmpty
			}
			return int64(el) + acc, nil
		}
		actual, _ := SliceInt64{given}.ReduceWhileInt64(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, 0)
	f([]int64{1}, 1)
	f([]int64{1, 2}, 3)
	f([]int64{1, 2, 3}, 6)
	f([]int64{1, 2, 0, 3}, 3)
}

func TestChannelToSliceInt64(t *testing.T) {
	f := func(given []int64) {
		c := make(chan int64, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := ChannelInt64{c}.ToSlice()
		assert.Equal(t, given, actual, "they should be equal")
	}
	f([]int64{})
	f([]int64{1})
	f([]int64{1, 2, 3, 1, 2})
}

func TestChannelAnyInt64(t *testing.T) {
	f := func(given []int64, expected bool) {
		even := func(t int64) bool { return t%2 == 0 }
		c := make(chan int64, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := ChannelInt64{c}.Any(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, false)
	f([]int64{1}, false)
	f([]int64{2}, true)
	f([]int64{1, 2}, true)
	f([]int64{1, 2, 3}, true)
	f([]int64{1, 3, 5}, false)
	f([]int64{1, 3, 5, 7, 9, 11}, false)
	f([]int64{1, 3, 5, 7, 10, 11}, true)
}

func TestChannelAllInt64(t *testing.T) {
	f := func(given []int64, expected bool) {
		even := func(t int64) bool { return t%2 == 0 }
		c := make(chan int64, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := ChannelInt64{c}.All(even)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, true)
	f([]int64{1}, false)
	f([]int64{2}, true)
	f([]int64{1, 2}, false)
	f([]int64{2, 4}, true)
	f([]int64{2, 4, 6, 8, 10, 12}, true)
	f([]int64{2, 4, 6, 8, 11, 12}, false)
}

func TestChannelEachInt64(t *testing.T) {
	f := func(given []int64) {
		c := make(chan int64, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := make(chan int64, len(given))
		mapper := func(t int64) { result <- t }
		ChannelInt64{c}.Each(mapper)
		close(result)
		actual := ChannelInt64{result}.ToSlice()
		assert.Equal(t, given, actual, "they should be equal")
	}

	f([]int64{})
	f([]int64{1})
	f([]int64{1, 2, 3})
	f([]int64{1, 2, 3, 4, 5, 6, 7})
}

func TestChannelChunkEveryInt64(t *testing.T) {
	f := func(size int, given []int64, expected [][]int64) {
		c := make(chan int64, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := ChannelInt64{c}.ChunkEvery(size)
		actual := make([][]int64, 0)
		for el := range result {
			actual = append(actual, el)
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(2, []int64{}, [][]int64{})
	f(2, []int64{1}, [][]int64{{1}})
	f(2, []int64{1, 2}, [][]int64{{1, 2}})
	f(2, []int64{1, 2, 3, 4}, [][]int64{{1, 2}, {3, 4}})
	f(2, []int64{1, 2, 3, 4, 5}, [][]int64{{1, 2}, {3, 4}, {5}})
}

func TestChannelCountInt64(t *testing.T) {
	f := func(element int64, given []int64, expected int) {
		c := make(chan int64, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := ChannelInt64{c}.Count(element)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, []int64{}, 0)
	f(1, []int64{1}, 1)
	f(1, []int64{2}, 0)
	f(1, []int64{1, 2, 3, 1, 4}, 2)
}

func TestChannelDropInt64(t *testing.T) {
	f := func(count int, given []int64, expected []int64) {
		c := make(chan int64, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := ChannelInt64{c}.Drop(count)
		actual := make([]int64, 0)
		for el := range result {
			actual = append(actual, el)
		}
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(1, []int64{}, []int64{})
	f(1, []int64{2}, []int64{})
	f(1, []int64{2, 3}, []int64{3})
	f(1, []int64{1, 2, 3}, []int64{2, 3})
	f(0, []int64{1, 2, 3}, []int64{1, 2, 3})
	f(3, []int64{1, 2, 3, 4, 5, 6}, []int64{4, 5, 6})
	f(1, []int64{1, 2, 3, 4, 5, 6}, []int64{2, 3, 4, 5, 6})
}

func TestChannelFilterInt64(t *testing.T) {
	f := func(given []int64, expected []int64) {
		even := func(t int64) bool { return t%2 == 0 }
		c := make(chan int64, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := ChannelInt64{c}.Filter(even)
		actual := ChannelInt64{result}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, []int64{})
	f([]int64{1}, []int64{})
	f([]int64{2}, []int64{2})
	f([]int64{1, 2, 3, 4}, []int64{2, 4})
}

func TestChannelMapInt64Int(t *testing.T) {
	f := func(given []int64, expected []int64) {
		double := func(el int64) int { return int(el * 2) }
		c := make(chan int64, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := ChannelInt64{c}.MapInt(double)

		// convert chan int64 to chan int
		c2 := make(chan int64, 1)
		go func() {
			for el := range result {
				c2 <- int64(el)
			}
			close(c2)
		}()

		actual := ChannelInt64{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, []int64{})
	f([]int64{1}, []int64{2})
	f([]int64{1, 2, 3}, []int64{2, 4, 6})
}

func TestChannelMapInt64Int8(t *testing.T) {
	f := func(given []int64, expected []int64) {
		double := func(el int64) int8 { return int8(el * 2) }
		c := make(chan int64, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := ChannelInt64{c}.MapInt8(double)

		// convert chan int64 to chan int8
		c2 := make(chan int64, 1)
		go func() {
			for el := range result {
				c2 <- int64(el)
			}
			close(c2)
		}()

		actual := ChannelInt64{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, []int64{})
	f([]int64{1}, []int64{2})
	f([]int64{1, 2, 3}, []int64{2, 4, 6})
}

func TestChannelMapInt64Int16(t *testing.T) {
	f := func(given []int64, expected []int64) {
		double := func(el int64) int16 { return int16(el * 2) }
		c := make(chan int64, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := ChannelInt64{c}.MapInt16(double)

		// convert chan int64 to chan int16
		c2 := make(chan int64, 1)
		go func() {
			for el := range result {
				c2 <- int64(el)
			}
			close(c2)
		}()

		actual := ChannelInt64{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, []int64{})
	f([]int64{1}, []int64{2})
	f([]int64{1, 2, 3}, []int64{2, 4, 6})
}

func TestChannelMapInt64Int32(t *testing.T) {
	f := func(given []int64, expected []int64) {
		double := func(el int64) int32 { return int32(el * 2) }
		c := make(chan int64, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := ChannelInt64{c}.MapInt32(double)

		// convert chan int64 to chan int32
		c2 := make(chan int64, 1)
		go func() {
			for el := range result {
				c2 <- int64(el)
			}
			close(c2)
		}()

		actual := ChannelInt64{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, []int64{})
	f([]int64{1}, []int64{2})
	f([]int64{1, 2, 3}, []int64{2, 4, 6})
}

func TestChannelMapInt64Int64(t *testing.T) {
	f := func(given []int64, expected []int64) {
		double := func(el int64) int64 { return int64(el * 2) }
		c := make(chan int64, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		result := ChannelInt64{c}.MapInt64(double)

		// convert chan int64 to chan int64
		c2 := make(chan int64, 1)
		go func() {
			for el := range result {
				c2 <- int64(el)
			}
			close(c2)
		}()

		actual := ChannelInt64{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, []int64{})
	f([]int64{1}, []int64{2})
	f([]int64{1, 2, 3}, []int64{2, 4, 6})
}

func TestChannelMaxInt64(t *testing.T) {
	f := func(given []int64, expected int64, expectedErr error) {
		c := make(chan int64, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual, actualErr := ChannelInt64{c}.Max()
		assert.Equal(t, expected, actual, "they should be equal")
		assert.Equal(t, expectedErr, actualErr, "they should be equal")
	}
	f([]int64{}, 0, ErrEmpty)
	f([]int64{1, 4, 2}, 4, nil)
	f([]int64{1, 2, 4}, 4, nil)
	f([]int64{4, 2, 1}, 4, nil)
}

func TestChannelMinInt64(t *testing.T) {
	f := func(given []int64, expected int64, expectedErr error) {
		c := make(chan int64, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual, actualErr := ChannelInt64{c}.Min()
		assert.Equal(t, expected, actual, "they should be equal")
		assert.Equal(t, expectedErr, actualErr, "they should be equal")
	}
	f([]int64{}, 0, ErrEmpty)
	f([]int64{4, 1, 2}, 1, nil)
	f([]int64{1, 2, 4}, 1, nil)
	f([]int64{4, 2, 1}, 1, nil)
}

func TestChannelReduceInt64Int(t *testing.T) {
	f := func(given []int64, expected int) {
		c := make(chan int64, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int64, acc int) int { return int(el) + acc }
		actual := ChannelInt64{c}.ReduceInt(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, 0)
	f([]int64{1}, 1)
	f([]int64{1, 2}, 3)
	f([]int64{1, 2, 3, 4, 5}, 15)
}

func TestChannelReduceInt64Int8(t *testing.T) {
	f := func(given []int64, expected int8) {
		c := make(chan int64, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int64, acc int8) int8 { return int8(el) + acc }
		actual := ChannelInt64{c}.ReduceInt8(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, 0)
	f([]int64{1}, 1)
	f([]int64{1, 2}, 3)
	f([]int64{1, 2, 3, 4, 5}, 15)
}

func TestChannelReduceInt64Int16(t *testing.T) {
	f := func(given []int64, expected int16) {
		c := make(chan int64, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int64, acc int16) int16 { return int16(el) + acc }
		actual := ChannelInt64{c}.ReduceInt16(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, 0)
	f([]int64{1}, 1)
	f([]int64{1, 2}, 3)
	f([]int64{1, 2, 3, 4, 5}, 15)
}

func TestChannelReduceInt64Int32(t *testing.T) {
	f := func(given []int64, expected int32) {
		c := make(chan int64, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int64, acc int32) int32 { return int32(el) + acc }
		actual := ChannelInt64{c}.ReduceInt32(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, 0)
	f([]int64{1}, 1)
	f([]int64{1, 2}, 3)
	f([]int64{1, 2, 3, 4, 5}, 15)
}

func TestChannelReduceInt64Int64(t *testing.T) {
	f := func(given []int64, expected int64) {
		c := make(chan int64, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int64, acc int64) int64 { return int64(el) + acc }
		actual := ChannelInt64{c}.ReduceInt64(0, sum)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, 0)
	f([]int64{1}, 1)
	f([]int64{1, 2}, 3)
	f([]int64{1, 2, 3, 4, 5}, 15)
}

func TestChannelScanInt64Int(t *testing.T) {
	f := func(given []int64, expected []int64) {
		c := make(chan int64, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int64, acc int) int { return int(el) + acc }
		result := ChannelInt64{c}.ScanInt(0, sum)

		// convert chan int64 to chan int
		c2 := make(chan int64, 1)
		go func() {
			for el := range result {
				c2 <- int64(el)
			}
			close(c2)
		}()

		actual := ChannelInt64{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, []int64{})
	f([]int64{1}, []int64{1})
	f([]int64{1, 2}, []int64{1, 3})
	f([]int64{1, 2, 3, 4, 5}, []int64{1, 3, 6, 10, 15})
}

func TestChannelScanInt64Int8(t *testing.T) {
	f := func(given []int64, expected []int64) {
		c := make(chan int64, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int64, acc int8) int8 { return int8(el) + acc }
		result := ChannelInt64{c}.ScanInt8(0, sum)

		// convert chan int64 to chan int8
		c2 := make(chan int64, 1)
		go func() {
			for el := range result {
				c2 <- int64(el)
			}
			close(c2)
		}()

		actual := ChannelInt64{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, []int64{})
	f([]int64{1}, []int64{1})
	f([]int64{1, 2}, []int64{1, 3})
	f([]int64{1, 2, 3, 4, 5}, []int64{1, 3, 6, 10, 15})
}

func TestChannelScanInt64Int16(t *testing.T) {
	f := func(given []int64, expected []int64) {
		c := make(chan int64, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int64, acc int16) int16 { return int16(el) + acc }
		result := ChannelInt64{c}.ScanInt16(0, sum)

		// convert chan int64 to chan int16
		c2 := make(chan int64, 1)
		go func() {
			for el := range result {
				c2 <- int64(el)
			}
			close(c2)
		}()

		actual := ChannelInt64{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, []int64{})
	f([]int64{1}, []int64{1})
	f([]int64{1, 2}, []int64{1, 3})
	f([]int64{1, 2, 3, 4, 5}, []int64{1, 3, 6, 10, 15})
}

func TestChannelScanInt64Int32(t *testing.T) {
	f := func(given []int64, expected []int64) {
		c := make(chan int64, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int64, acc int32) int32 { return int32(el) + acc }
		result := ChannelInt64{c}.ScanInt32(0, sum)

		// convert chan int64 to chan int32
		c2 := make(chan int64, 1)
		go func() {
			for el := range result {
				c2 <- int64(el)
			}
			close(c2)
		}()

		actual := ChannelInt64{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, []int64{})
	f([]int64{1}, []int64{1})
	f([]int64{1, 2}, []int64{1, 3})
	f([]int64{1, 2, 3, 4, 5}, []int64{1, 3, 6, 10, 15})
}

func TestChannelScanInt64Int64(t *testing.T) {
	f := func(given []int64, expected []int64) {
		c := make(chan int64, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		sum := func(el int64, acc int64) int64 { return int64(el) + acc }
		result := ChannelInt64{c}.ScanInt64(0, sum)

		// convert chan int64 to chan int64
		c2 := make(chan int64, 1)
		go func() {
			for el := range result {
				c2 <- int64(el)
			}
			close(c2)
		}()

		actual := ChannelInt64{c2}.ToSlice()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, []int64{})
	f([]int64{1}, []int64{1})
	f([]int64{1, 2}, []int64{1, 3})
	f([]int64{1, 2, 3, 4, 5}, []int64{1, 3, 6, 10, 15})
}

func TestChannelSumInt64(t *testing.T) {
	f := func(given []int64, expected int64) {
		c := make(chan int64, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		actual := ChannelInt64{c}.Sum()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f([]int64{}, 0)
	f([]int64{1}, 1)
	f([]int64{1, 2}, 3)
	f([]int64{1, 2, 3, 4, 5}, 15)
}

func TestChannelTakeInt64(t *testing.T) {
	f := func(count int, given int64, expected []int64) {
		ctx, cancel := context.WithCancel(context.Background())
		s := SequenceInt64{ctx: ctx}
		seq := s.Repeat(given)
		seq2 := ChannelInt64{seq}.Take(count)
		actual := ChannelInt64{seq2}.ToSlice()
		cancel()
		assert.Equal(t, expected, actual, "they should be equal")
	}
	f(0, 1, []int64{})
	f(1, 1, []int64{1})
	f(2, 1, []int64{1, 1})
}

func TestChannelTeeInt64(t *testing.T) {
	f := func(count int, given []int64) {
		c := make(chan int64, 1)
		go func() {
			for _, el := range given {
				c <- el
			}
			close(c)
		}()
		channels := ChannelInt64{c}.Tee(count)
		for _, ch := range channels {
			go func(ch chan int64) {
				actual := ChannelInt64{ch}.ToSlice()
				assert.Equal(t, given, actual, "they should be equal")
			}(ch)
		}
	}
	f(1, []int64{})
	f(1, []int64{1})
	f(1, []int64{1, 2})
	f(1, []int64{1, 2, 3})
	f(1, []int64{1, 2, 3, 1, 2})

	f(2, []int64{})
	f(2, []int64{1})
	f(2, []int64{1, 2})
	f(2, []int64{1, 2, 3})
	f(2, []int64{1, 2, 3, 1, 2})

	f(10, []int64{1, 2, 3, 1, 2})
}

func TestAsyncSliceAnyInt64(t *testing.T) {
	f := func(check func(t int64) bool, given []int64, expected bool) {
		s := AsyncSliceInt64{Data: given, Workers: 2}
		actual := s.Any(check)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	isEven := func(t int64) bool { return (t % 2) == 0 }

	f(isEven, []int64{}, false)
	f(isEven, []int64{1}, false)
	f(isEven, []int64{1, 3}, false)
	f(isEven, []int64{2}, true)
	f(isEven, []int64{1, 2}, true)
	f(isEven, []int64{1, 3, 5, 7, 9, 11}, false)
	f(isEven, []int64{1, 3, 5, 7, 9, 12}, true)
}

func TestAsyncSliceAllInt64(t *testing.T) {
	f := func(check func(t int64) bool, given []int64, expected bool) {
		s := AsyncSliceInt64{Data: given, Workers: 2}
		actual := s.All(check)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	isEven := func(t int64) bool { return (t % 2) == 0 }

	f(isEven, []int64{}, true)
	f(isEven, []int64{1}, false)
	f(isEven, []int64{1, 3}, false)
	f(isEven, []int64{2}, true)
	f(isEven, []int64{2, 4}, true)
	f(isEven, []int64{2, 3}, false)
	f(isEven, []int64{2, 4, 6, 8, 10, 12}, true)
	f(isEven, []int64{2, 4, 6, 8, 10, 11}, false)
}

func TestAsyncSliceEachInt64(t *testing.T) {
	f := func(given []int64) {
		s := AsyncSliceInt64{Data: given, Workers: 2}
		result := make(chan int64, len(given))
		mapper := func(t int64) { result <- t }
		s.Each(mapper)
		close(result)
		actual := ChannelInt64{result}.ToSlice()
		sorted := SliceInt64{actual}.Sort()
		assert.Equal(t, given, sorted, "they should be equal")
	}

	f([]int64{})
	f([]int64{1})
	f([]int64{1, 2, 3})
	f([]int64{1, 2, 3, 4, 5, 6, 7})
}

func TestAsyncSliceFilterInt64(t *testing.T) {
	f := func(given []int64, expected []int64) {
		filter := func(t int64) bool { return t > 10 }
		s := AsyncSliceInt64{Data: given, Workers: 2}
		actual := s.Filter(filter)
		assert.Equal(t, expected, actual, "they should be equal")
	}

	f([]int64{}, []int64{})
	f([]int64{5}, []int64{})
	f([]int64{15}, []int64{15})
	f([]int64{9, 11, 12, 13, 6}, []int64{11, 12, 13})
}

func TestAsyncSliceMapInt64Int(t *testing.T) {
	f := func(mapper func(t int64) int, given []int64, expected []int) {
		s := AsyncSliceInt64{Data: given, Workers: 2}
		actual := s.MapInt(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int64) int { return int((t * 2)) }

	f(double, []int64{}, []int{})
	f(double, []int64{1}, []int{2})
	f(double, []int64{1, 2, 3}, []int{2, 4, 6})
}

func TestAsyncSliceMapInt64Int8(t *testing.T) {
	f := func(mapper func(t int64) int8, given []int64, expected []int8) {
		s := AsyncSliceInt64{Data: given, Workers: 2}
		actual := s.MapInt8(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int64) int8 { return int8((t * 2)) }

	f(double, []int64{}, []int8{})
	f(double, []int64{1}, []int8{2})
	f(double, []int64{1, 2, 3}, []int8{2, 4, 6})
}

func TestAsyncSliceMapInt64Int16(t *testing.T) {
	f := func(mapper func(t int64) int16, given []int64, expected []int16) {
		s := AsyncSliceInt64{Data: given, Workers: 2}
		actual := s.MapInt16(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int64) int16 { return int16((t * 2)) }

	f(double, []int64{}, []int16{})
	f(double, []int64{1}, []int16{2})
	f(double, []int64{1, 2, 3}, []int16{2, 4, 6})
}

func TestAsyncSliceMapInt64Int32(t *testing.T) {
	f := func(mapper func(t int64) int32, given []int64, expected []int32) {
		s := AsyncSliceInt64{Data: given, Workers: 2}
		actual := s.MapInt32(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int64) int32 { return int32((t * 2)) }

	f(double, []int64{}, []int32{})
	f(double, []int64{1}, []int32{2})
	f(double, []int64{1, 2, 3}, []int32{2, 4, 6})
}

func TestAsyncSliceMapInt64Int64(t *testing.T) {
	f := func(mapper func(t int64) int64, given []int64, expected []int64) {
		s := AsyncSliceInt64{Data: given, Workers: 2}
		actual := s.MapInt64(mapper)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	double := func(t int64) int64 { return int64((t * 2)) }

	f(double, []int64{}, []int64{})
	f(double, []int64{1}, []int64{2})
	f(double, []int64{1, 2, 3}, []int64{2, 4, 6})
}

func TestAsyncSliceReduceInt64(t *testing.T) {
	f := func(reducer func(a int64, b int64) int64, given []int64, expected int64) {
		s := AsyncSliceInt64{Data: given, Workers: 4}
		actual := s.Reduce(reducer)
		assert.Equal(t, expected, actual, "they should be equal")
	}
	sum := func(a int64, b int64) int64 { return a + b }

	f(sum, []int64{}, 0)
	f(sum, []int64{1}, 1)
	f(sum, []int64{1, 2}, 3)
	f(sum, []int64{1, 2, 3}, 6)
	f(sum, []int64{1, 2, 3, 4}, 10)
	f(sum, []int64{1, 2, 3, 4, 5}, 15)
}
