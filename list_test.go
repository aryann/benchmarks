package benchmarks

import (
	"fmt"
	"reflect"
	"testing"
)

var sizes = []int{100_000, 1_000_000, 10_000_000, 100_000_000, 1_000_000_000}
var cutPoints = []float64{0.1, 0.5, 0.9}

// BenchmarkSliceRemoval measures the performance of removing an element from
// a Slice of int64 numbers. The benchmark runs over a variety of Slice sizes
// and removal points.
func BenchmarkSliceRemoval(b *testing.B) {
	for _, size := range sizes {
		slice := make([]int64, size)

		for _, cutPoint := range cutPoints {
			removeIdx := int(float64(size) * cutPoint)
			b.Run(fmt.Sprintf("Size=%d,CutPoint=%0.2f", size, cutPoint), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_ = append(slice[:removeIdx], slice[removeIdx+1:]...)
				}
			})
		}
	}
}

type int64Node struct {
	val  int64
	next *int64Node
}

func makeList(size int) *int64Node {
	var head *int64Node
	for size > 0 {
		head = &int64Node{
			val:  0,
			next: head,
		}
		size--
	}
	return head
}

func (i *int64Node) String() string {
	if i.next == nil {
		return fmt.Sprint(i.val)
	}
	return fmt.Sprintf("%d %s", i.val, i.next)
}

func removeAt(idx int, list *int64Node) (*int64Node, error) {
	dummy := &int64Node{next: list}
	curr := dummy
	originalIdx := idx
	for idx > 0 && curr.next != nil {
		curr = curr.next
		idx--
	}
	if curr.next == nil {
		return nil, fmt.Errorf("index %d is out of range for list", originalIdx)
	}
	curr.next = curr.next.next
	return dummy.next, nil
}

func TestRemove(t *testing.T) {
	testCases := []struct {
		in         *int64Node
		idx        int
		wantOutput *int64Node
		wantErr    string
	}{
		{in: nil, idx: 0, wantErr: "index 0 is out of range for list"},
		{in: nil, idx: 10, wantErr: "index 10 is out of range for list"},
		{in: &int64Node{val: 1}, idx: 0, wantOutput: nil},
		{in: &int64Node{val: 1}, idx: 1, wantErr: "index 1 is out of range for list"},
		{in: &int64Node{val: 1, next: &int64Node{val: 2, next: &int64Node{val: 3}}}, idx: 0,
			wantOutput: &int64Node{val: 2, next: &int64Node{val: 3}}},
		{in: &int64Node{val: 1, next: &int64Node{val: 2, next: &int64Node{val: 3}}}, idx: 1,
			wantOutput: &int64Node{val: 1, next: &int64Node{val: 3}}},
		{in: &int64Node{val: 1, next: &int64Node{val: 2, next: &int64Node{val: 3}}}, idx: 2,
			wantOutput: &int64Node{val: 1, next: &int64Node{val: 2}}},
		{in: &int64Node{val: 1, next: &int64Node{val: 2, next: &int64Node{val: 3}}}, idx: 100,
			wantErr: "index 100 is out of range for list"},
		{in: makeList(10), idx: 5, wantOutput: makeList(9)},
		{in: makeList(1000), idx: 900, wantOutput: makeList(999)},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got, err := removeAt(testCase.idx, testCase.in)

			if testCase.wantErr == "" {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if !reflect.DeepEqual(got, testCase.wantOutput) {
					t.Fatalf("want list %+v, got %+v", testCase.wantOutput, got)
				}

			} else {
				if err == nil {
					t.Fatalf("want error '%s', got none", testCase.wantErr)
				}
				if err.Error() != testCase.wantErr {
					t.Fatalf("want error '%s', got '%v'", testCase.wantErr, err.Error())
				}
			}
		})
	}
}

// BenchmarkListRemoval measures the performance of removing an element from
// a linked list of int64 numbers. The benchmark runs over a variety of list
// sizes and removal points.
func BenchmarkListRemoval(b *testing.B) {
	for _, size := range sizes {
		list := makeList(size - 1)

		for _, cutPoint := range cutPoints {
			removeIdx := int(float64(size) * cutPoint)
			b.Run(fmt.Sprintf("Size=%d,CutPoint=%0.2f", size, cutPoint), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					// Making a new list is too expensive to do on each iteration.
					// To speed up the set-up, we create a list that has size - 1 elements,
					// then we add one more element prior to the timing of each removal,
					// such that each iteration starts with a list of size elements.
					list = &int64Node{next: list}

					if _, err := removeAt(removeIdx, list); err != nil {
						b.Fatal(err)
					}
				}
			})
		}
	}
}
