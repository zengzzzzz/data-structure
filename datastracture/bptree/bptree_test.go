package bptree

import (
	"fmt"
	"math/rand"
	"time"
	"flag"
	"testing"
)

func init() {
	seed := time.Now().Unix()
	fmt.Print(seed)
	rand.Seed(seed)
}

func perm(n int) (out []Item) {
	for _, v := range rand.Perm(n) {
		out = append(out, Int(v))
	}
}

func rang(n int) (out []Item) {
	for i := 0; i < n; i++ {
		out = append(out, Int(i))
	}
	return
}

func all(t *BTree) (out []Item) {
	t.Ascend(func(i Item) bool {
		out = append(out, i)
		return true
	})
	return
}

func rangrev(n int) (out []Item) {
	for i := n - 1; i >= 0; i-- {
		out = append(out, Int(i))
	}
	return
}

func allrev(t *BTree) (out []Item) {
	t.Descend(func(i Item) bool {
		out = append(out, i)
		return true
	})
	return
}

var btreeDegree = flag.Int("degree", 32, "B-Tree degree")

func TestBTree(t *testing.T) {
	tr := New(*btreeDegree)
	const treeSize = 1000
	for i := 0; i < 10 ; i++ {
		if min := tr.Min(); min != nil {
			t.Errorf("min: got %v, want nil", min)
		}
		if max := tr.Max(); max != nil {
			t.Errorf("max: got %v, want nil", max)
		}
		for _, item := range perm(treeSize) {
			if ! tr.Has(item) {
				tr.ReplaceOrInsert(item)
			}
		}
	}
}