package sortedmap_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/matsuyoshi30/sortedmap"
)

func TestNewSortedMap(t *testing.T) {
	sm := sortedmap.NewSortedMap[int, string]()
	if !reflect.DeepEqual(reflect.TypeOf(&sortedmap.SortedMap[int, string]{}), reflect.TypeOf(sm)) {
		t.Fatalf("want %v but got %v\n", reflect.TypeOf(&sortedmap.SortedMap[int, string]{}), reflect.TypeOf(sm))
	}
}

func TestSortedMap_Get(t *testing.T) {
	tests := []struct {
		desc    string
		vals    []string
		key     int
		wantVal string
		wantOK  bool
	}{
		{
			desc:    "empty",
			wantVal: "",
			wantOK:  false,
		},
		{
			desc:    "single value",
			vals:    []string{"foo"},
			key:     1,
			wantVal: "foo",
			wantOK:  true,
		},
		{
			desc:    "multiple values",
			vals:    []string{"foo", "bar", "baz"},
			key:     2,
			wantVal: "bar",
			wantOK:  true,
		},
		{
			desc:    "multiple values but not found",
			vals:    []string{"foo", "bar", "baz"},
			key:     4,
			wantVal: "",
			wantOK:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			sm := sortedmap.NewSortedMap[int, string]()
			for i, val := range tt.vals {
				sm.Put(i+1, val)
			}
			got, ok := sm.Get(tt.key)
			if tt.wantVal != got {
				t.Fatalf("want %v but got %v\n", tt.wantVal, got)
			}
			if tt.wantOK != ok {
				t.Fatalf("want %v but got %v\n", tt.wantOK, ok)
			}
		})
	}
}

func TestSortedMap_Put(t *testing.T) {
	tests := []struct {
		desc     string
		initVals []string
		key      int
		val      string
		want     string
	}{
		{
			desc: "empty",
			key:  1,
			val:  "foo",
		},
		{
			desc:     "single value",
			initVals: []string{"foo"},
			key:      2,
			val:      "bar",
		},
		{
			desc:     "multiple values",
			initVals: []string{"foo", "bar", "baz"},
			key:      4,
			val:      "qux",
		},
		{
			desc:     "has previous value",
			initVals: []string{"foo", "bar", "baz"},
			key:      1,
			val:      "qux",
			want:     "foo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			sm := sortedmap.NewSortedMap[int, string]()
			for i, val := range tt.initVals {
				sm.Put(i+1, val)
			}
			got := sm.Put(tt.key, tt.val)
			if tt.want != got {
				t.Fatalf("want %v but got %v\n", tt.want, got)
			}
			got2, ok := sm.Get(tt.key)
			if !ok {
				t.Fatal("want true but got false")
			}
			if tt.val != got2 {
				t.Fatalf("want %v but got %v\n", tt.val, got2)
			}
		})
	}
}

func TestSortedMap_Remove(t *testing.T) {
	tests := []struct {
		desc string
		vals []string
		key  int
		want string
	}{
		{
			desc: "empty",
			key:  1,
			want: "",
		},
		{
			desc: "single value",
			vals: []string{"foo"},
			key:  1,
			want: "foo",
		},
		{
			desc: "multiple values",
			vals: []string{"foo", "bar", "baz"},
			key:  2,
			want: "bar",
		},
		{
			desc: "multiple values but not found",
			vals: []string{"foo", "bar", "baz"},
			key:  4,
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			sm := sortedmap.NewSortedMap[int, string]()
			for i, val := range tt.vals {
				sm.Put(i+1, val)
			}
			got := sm.Remove(tt.key)
			if tt.want != got {
				t.Fatalf("want %v but got %v\n", tt.want, got)
			}
			_, ok := sm.Get(tt.key)
			if ok {
				t.Fatal("want true but got false")
			}
		})
	}
}

func TestSortedMap_FirstKey(t *testing.T) {
	tests := []struct {
		desc   string
		keys   []int
		want   int
		wantOK bool
	}{
		{
			desc:   "empty",
			want:   0,
			wantOK: false,
		},
		{
			desc:   "single value",
			keys:   []int{1},
			want:   1,
			wantOK: true,
		},
		{
			desc:   "multiple values",
			keys:   []int{1, 2, 3},
			want:   1,
			wantOK: true,
		},
		{
			desc:   "multiple values unordered",
			keys:   []int{3, 2, 1},
			want:   1,
			wantOK: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			sm := sortedmap.NewSortedMap[int, string]()
			for _, key := range tt.keys {
				sm.Put(key, fmt.Sprintf("%d", key))
			}
			got, ok := sm.FirstKey()
			if tt.wantOK != ok {
				t.Fatalf("want %v but got %v\n", tt.wantOK, ok)
			}
			if tt.want != got {
				t.Fatalf("want %v but got %v\n", tt.want, got)
			}
		})
	}
}

func TestSortedMap_LastKey(t *testing.T) {
	tests := []struct {
		desc   string
		keys   []int
		want   int
		wantOK bool
	}{
		{
			desc:   "empty",
			want:   0,
			wantOK: false,
		},
		{
			desc:   "single value",
			keys:   []int{1},
			want:   1,
			wantOK: true,
		},
		{
			desc:   "multiple values",
			keys:   []int{1, 2, 3},
			want:   3,
			wantOK: true,
		},
		{
			desc:   "multiple values unordered",
			keys:   []int{3, 2, 1},
			want:   3,
			wantOK: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			sm := sortedmap.NewSortedMap[int, string]()
			for _, key := range tt.keys {
				sm.Put(key, fmt.Sprintf("%d", key))
			}
			got, ok := sm.LastKey()
			if tt.wantOK != ok {
				t.Fatalf("want %v but got %v\n", tt.wantOK, ok)
			}
			if tt.want != got {
				t.Fatalf("want %v but got %v\n", tt.want, got)
			}
		})
	}
}

func TestSortedMap_SubMap(t *testing.T) {
	tests := []struct {
		desc     string
		vals     []string
		from, to int
		wantNil  bool
	}{
		{
			desc:    "empty",
			from:    2,
			to:      4,
			wantNil: true,
		},
		{
			desc: "submap",
			vals: []string{"a", "b", "c", "d", "e"},
			from: 2,
			to:   4,
		},
		// TODO: more tests
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			sm := sortedmap.NewSortedMap[int, string]()
			for i, val := range tt.vals {
				sm.Put(i+1, val)
			}
			sm2 := sm.SubMap(tt.from, tt.to)
			if tt.wantNil {
				if sm2 != nil {
					t.Fatal("want sm2 nil but not nil")
				}
				return
			}
			if sm2 == nil {
				t.Fatal("want sm2 not nil but nil")
			}
			new := "z"
			sm2.Put(tt.from, new)
			got, ok := sm.Get(tt.from)
			if !ok {
				t.Fatal("want true but got false")
			}
			if got != "z" {
				t.Fatalf("want %v but got %v\n", new, got)
			}
		})
	}
}

func TestSortedMap_HeadMap(t *testing.T) {
	tests := []struct {
		desc    string
		vals    []string
		to      int
		wantNil bool
	}{
		{
			desc:    "empty",
			to:      3,
			wantNil: true,
		},
		{
			desc: "headmap",
			vals: []string{"a", "b", "c", "d", "e"},
			to:   3,
		},
		// TODO: more tests
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			sm := sortedmap.NewSortedMap[int, string]()
			for i, val := range tt.vals {
				sm.Put(i+1, val)
			}
			sm2 := sm.HeadMap(tt.to)
			if tt.wantNil {
				if sm2 != nil {
					t.Fatal("want sm2 nil but not nil")
				}
				return
			}
			if sm2 == nil {
				t.Fatal("want sm2 not nil but nil")
			}
			new := "z"
			sm2.Put(tt.to-1, new)
			got, ok := sm.Get(tt.to - 1)
			if !ok {
				t.Fatal("want true but got false")
			}
			if got != "z" {
				t.Fatalf("want %v but got %v\n", new, got)
			}
		})
	}
}

func TestSortedMap_TailMap(t *testing.T) {
	tests := []struct {
		desc    string
		vals    []string
		from    int
		wantNil bool
	}{
		{
			desc:    "empty",
			from:    3,
			wantNil: true,
		},
		{
			desc: "tailmap",
			vals: []string{"a", "b", "c", "d", "e"},
			from: 3,
		},
		// TODO: more tests
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			sm := sortedmap.NewSortedMap[int, string]()
			for i, val := range tt.vals {
				sm.Put(i+1, val)
			}
			sm2 := sm.TailMap(tt.from)
			if tt.wantNil {
				if sm2 != nil {
					t.Fatal("want sm2 nil but not nil")
				}
				return
			}
			if sm2 == nil {
				t.Fatal("want sm2 not nil but nil")
			}
			new := "z"
			sm2.Put(tt.from, new)
			got, ok := sm.Get(tt.from)
			if !ok {
				t.Fatal("want true but got false")
			}
			if got != "z" {
				t.Fatalf("want %v but got %v\n", new, got)
			}
		})
	}
}
