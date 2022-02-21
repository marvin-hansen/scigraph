package graph_types

import (
	"fmt"
	"testing"
)

func TestSetDifference(t *testing.T) {
	cases := []struct {
		Name     string
		A, B     []myint
		Expected []myint
	}{
		{
			"same",
			[]myint{1, 2, 3},
			[]myint{3, 1, 2},
			[]myint{},
		},

		{
			"A has extra elements",
			[]myint{1, 2, 3},
			[]myint{3, 2},
			[]myint{1},
		},

		{
			"B has extra elements",
			[]myint{1, 2, 3},
			[]myint{3, 2, 1, 4},
			[]myint{},
		},
		{
			"B is nil",
			[]myint{1, 2, 3},
			nil,
			[]myint{1, 2, 3},
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("%d-%s", i, tc.Name), func(t *testing.T) {
			one := make(Set[myint])
			two := make(Set[myint])
			expected := make(Set[myint])
			for _, v := range tc.A {
				one.Add(v)
			}
			for _, v := range tc.B {
				two.Add(v)
			}
			if tc.B == nil {
				two = nil
			}
			for _, v := range tc.Expected {
				expected.Add(v)
			}

			actual := one.Difference(two)
			match := actual.Intersection(expected)
			if match.Len() != expected.Len() {
				t.Fatalf("bad: %#v", actual.List())
			}
		})
	}
}

func TestSetFilter(t *testing.T) {
	cases := []struct {
		Input    []myint
		Expected []myint
	}{
		{
			[]myint{1, 2, 3},
			[]myint{1, 2, 3},
		},

		{
			[]myint{4, 5, 6},
			[]myint{4},
		},

		{
			[]myint{7, 8, 9},
			[]myint{},
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("%d-%#v", i, tc.Input), func(t *testing.T) {
			input := make(Set[myint])
			expected := make(Set[myint])
			for _, v := range tc.Input {
				input.Add(v)
			}
			for _, v := range tc.Expected {
				expected.Add(v)
			}

			actual := input.Filter(func(v myint) bool {
				return v < 5
			})
			match := actual.Intersection(expected)
			if match.Len() != expected.Len() {
				t.Fatalf("bad: %#v", actual.List())
			}
		})
	}
}

func TestSetCopy(t *testing.T) {
	a := make(Set[myint])
	a.Add(myint(1))
	a.Add(myint(2))

	b := a.Copy()
	b.Add(myint(3))

	diff := b.Difference(a)

	if diff.Len() != 1 {
		t.Fatalf("expected single diff value, got %#v", diff)
	}

	if !diff.Include(3) {
		t.Fatalf("diff does not contain 3, got %#v", diff)
	}

}

func makeSet(n int) Set[myint] {
	ret := make(Set[myint], n)
	for i := 0; i < n; i++ {
		ret.Add(myint(i))
	}
	return ret
}

func BenchmarkSetIntersection_100_100000(b *testing.B) {
	small := makeSet(100)
	large := makeSet(100000)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		small.Intersection(large)
	}
}

func BenchmarkSetIntersection_100000_100(b *testing.B) {
	small := makeSet(100)
	large := makeSet(100000)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		large.Intersection(small)
	}
}
