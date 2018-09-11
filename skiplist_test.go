package skiplist

import "testing"

func TestNew(t *testing.T) {
	sl := New()
	if sl == nil {
		t.Fatalf("New should not return nil")
	}
	if sl.Len() != 0 {
		t.Fatalf("New SkipList should have Len = 0")
	}
	if sl.MaxLevel() != DefaultMaxLevel {
		t.Fatalf("bad max level")
	}
	//if sl.Level() != 1 {
	//	t.Fatalf("bad level")
	//}
}

func TestNewMaxLevel(t *testing.T) {
	maxLevel := 10
	sl := NewMaxLevel(maxLevel)
	if sl == nil {
		t.Fatalf("New should not return nil")
	}
	if sl.Len() != 0 {
		t.Fatalf("New SkipList should have Len = 0")
	}
	if sl.MaxLevel() != maxLevel {
		t.Fatalf("bad max level")
	}
}

func TestSkipList_Has(t *testing.T) {
	sl := New()
	sl.ReplaceOrInsert(Int(4))
	sl.ReplaceOrInsert(Int(2))
	sl.ReplaceOrInsert(Int(3))
	if sl.Len() != 3 {
		t.Fatalf("bad len")
	}
	if sl.Has(Int(1)) {
		t.Fatalf("should not has 1")
	}
	if !sl.Has(Int(2)) {
		t.Fatalf("should has 2")
	}
}

func TestSkipList_ReplaceOrInsert(t *testing.T) {
	sl := New()
	sl.ReplaceOrInsert(Int(5))
	sl.ReplaceOrInsert(Int(6))
	sl.ReplaceOrInsert(Int(3))
	assertHas(t, sl, 5)
	assertHas(t, sl, 6)
	assertHas(t, sl, 3)
	assertNotHas(t, sl, 4)
	assertNotHas(t, sl, 1)
	oitem := sl.ReplaceOrInsert(Int(5))
	if oitem == nil {
		t.Fatalf("should not be nil")
	}
}

func TestSkipList_Delete(t *testing.T) {
	sl := New()
	sl.ReplaceOrInsert(Int(5))
	sl.ReplaceOrInsert(Int(6))
	sl.ReplaceOrInsert(Int(3))
	assertHas(t, sl, 6)
	assertLen(t, sl, 3)
	if sl.Delete(Int(6)) == nil {
		t.Fatalf("should not be nil")
	}
	assertNotHas(t, sl, 6)
	assertLen(t, sl, 2)
	if sl.Delete(Int(2)) != nil {
		t.Fatalf("should be nil")
	}
	assertNotHas(t, sl, 2)
	assertLen(t, sl, 2)
}

func TestSkipList_DeleteMin(t *testing.T) {
	sl := New()
	item := sl.DeleteMin()
	if item != nil {
		t.Fatalf("should be nil")
	}
	sl.ReplaceOrInsert(Int(5))
	sl.ReplaceOrInsert(Int(6))
	sl.ReplaceOrInsert(Int(3))
	assertHas(t, sl, 3)
	assertLen(t, sl, 3)
	item = sl.DeleteMin()
	if item == nil {
		t.Fatalf("should not be nil")
	}
	if item != Int(3) {
		t.Fatalf("should delete 3")
	}
}

func TestSkipList_InsertNoReplace(t *testing.T) {
	sl := New()
	sl.InsertNoReplace(Int(3))
	sl.InsertNoReplace(Int(5))
	sl.InsertNoReplace(Int(4))
	assertLen(t, sl, 3)
	sl.InsertNoReplace(Int(4))
	assertLen(t, sl, 4)
	if sl.Delete(Int(4)) == nil {
		t.Fatalf("should not be nil")
	}
	assertLen(t, sl, 3)
	if sl.Delete(Int(4)) == nil {
		t.Fatalf("should not be nil")
	}
	assertLen(t, sl, 2)
	if sl.Delete(Int(4)) != nil {
		t.Fatalf("should be nil")
	}
	assertLen(t, sl, 2)
}

func TestSkipList_Traverse(t *testing.T) {
	sl := New()
	sl.ReplaceOrInsert(Int(3))
	sl.ReplaceOrInsert(Int(2))
	sl.ReplaceOrInsert(Int(1))
	sl.ReplaceOrInsert(Int(5))
	sl.ReplaceOrInsert(Int(4))
	sl.ReplaceOrInsert(Int(6))
	sl.ReplaceOrInsert(Int(7))
	sl.ReplaceOrInsert(Int(9))
	sl.ReplaceOrInsert(Int(8))

	order1 := []Int{}
	sl.Ascend(func(i Item) bool {
		order1 = append(order1, i.(Int))
		return true
	})
	assertInOrder(t, order1, 1, 9)
	order2 := []Int{}
	sl.AscendGreaterOrEqual(Int(5), func(i Item) bool {
		order2 = append(order2, i.(Int))
		return true
	})
	assertInOrder(t, order2, 5, 9)

	order3 := []Int{}
	sl.AscendLessThan(Int(5), func(i Item) bool {
		order3 = append(order3, i.(Int))
		return true
	})
	assertInOrder(t, order3, 1, 4)
	order4 := []Int{}
	sl.AscendRange(Int(2), Int(7), func(i Item) bool {
		order4 = append(order4, i.(Int))
		return true
	})
	assertInOrder(t, order4, 2, 6)
}

func assertInOrder(t *testing.T, order []Int, min, max Int) {
	if len(order) != int(max)-int(min)+1 {
		t.Fatalf("order is not %d~%d: %+v", min, max, order)
	}

	for i, v := range order {
		if v != min+Int(i) {
			t.Fatalf("order is not %d~%d: %+v", min, max, order)
		}
	}
}

func assertLen(t *testing.T, sl *SkipList, v int) {
	if sl.Len() != v {
		t.Fatalf("len should be %d, but is %d", v, sl.Len())
	}
}

func assertHas(t *testing.T, sl *SkipList, v int) {
	if !sl.Has(Int(v)) {
		t.Fatalf("should has %d", v)
	}
}

func assertNotHas(t *testing.T, sl *SkipList, v int) {
	if sl.Has(Int(v)) {
		t.Fatalf("should not has %d", v)
	}
}
