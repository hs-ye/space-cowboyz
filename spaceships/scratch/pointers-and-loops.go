package scratch

import "fmt"

type Item struct {
	val int
}

type Testslice struct {
	Items []Item
}

type TestPSlice struct {
	Items []*Item // pointer to items
}

// Tests to check:
// - Make a list of items, loop over and then try to modify
//   - construct the list through copying the same item over and over
//   - creating a new item each time. both ways should create new objects
//
// - Make a list of pointers, loop over and try to modify
//   - construct through copying in the same pointer each time
//   - create a new pointer each time. This time, should only work when creating a new pointer each time
func Slicingtest() {
	t1 := Testslice{}
	a := Item{1}
	t1.Items = append(t1.Items, a, a, a) // copy being appended each time
	fmt.Println(t1)                      // 1,1,1
	t1.ModifySlice(4)
	fmt.Println(t1) // 1,1,1
	t1.ModifyPSlice(4)
	fmt.Println(t1)      // 1,1,1
	t1.Items[0].val += 1 // only the first copy should be affected here.
	fmt.Println(t1)      // 2,1,1

	t2 := Testslice{}
	t2.Items = append(t2.Items, t1.Items...)
	fmt.Println(t2) // 2,1,1
	t2.ModifySlice(4)
	fmt.Println(t2) // 2,1,1
	t2.ModifyPSlice(4)
	t2.Items[0].val += 1 // only the first copy should be affected here.
	fmt.Println(t2)      // 3,1,1
	fmt.Println(t1)      // 2,1,1

	t3 := TestPSlice{}
	t4 := t3                                // copy of empty struct, we will manually copy across the values
	t5 := t3                                // this one we will leave blank
	t3.Items = append(t3.Items, &a, &a, &a) // this should have 3 of the same items
	t6 := t3                                // this one we make a full copy, incl. the underlying pointers, so expect it to show all updates persisted through the items of t3
	fmt.Println("initial t3: ", t3)
	t3.PrintItems() // {1,1,1}
	fmt.Println("initial t4: ", t4)
	t4.Items = append(t4.Items, t3.Items...)
	fmt.Println(t4) // it's a copy now
	fmt.Println(t5) // still a copy
	t3.ModifySlice(2)
	t3.PrintItems() // {7,7,7}, Note 2 key things here:
	// 1) modification of the pointer classes is persistent
	// 2) the value is not as expcted: we are expecting {3,3,3}, but get {7,7,7}
	// this is because if we put in the exact same pointer, the +2 operation is applying 3x due to 3 elements, all pointing to the same object
	t4.PrintItems() // {7,7,7}  note item pointers to original object is persisted here
	t5.PrintItems() // {}  empty, as no items were ever added
	t6.PrintItems() // {7,7,7}

	// Now test if we make brand new slice objects:
	tp1 := TestPSlice{}
	for i := 0; i < 3; i++ {
		tp1.Items = append(tp1.Items, &t1.Items[i]) // all pointing to the individual items
	} // if we modify these, it should refect in the original as well
	tp1.PrintItems() // 2,1,1
	tp1.ModifyPSlice(1)
	tp1.PrintItems() // 3,2,2
	t1.PrintItems()  //3,2,2

	tp1.ModifySlice(1) // should this persist?
	tp1.PrintItems()   // 4,3,3 apparently it does. but why, if we are passing in a copy? Copy of pointer, when dereferenced, still point to the original memory location, so update is valid.
	t1.PrintItems()    // 4,3,3

}

func (t Testslice) PrintItems() {
	for _, _t := range t.Items {
		fmt.Print(_t)
	}
	fmt.Println("")

}

// this one should be modifying copies, leaving the original unchanged
func (t Testslice) ModifySlice(val int) {
	for _, i := range t.Items {
		i.val += val
	}
}
func (t *Testslice) ModifyPSlice(val int) {
	for _, i := range t.Items {
		i.val += val

	}
}

func (t TestPSlice) PrintItems() {
	for _, _t := range t.Items {
		fmt.Print(*_t) // it's a ref to a pointer now
	}
	fmt.Println("")

}

// this is modifying a slice of points in the struct
// so both the copy and pointer receiver should be modifying the original.
// because the pointer value to the object is the same whether it's a copy or not, both have the same result outcomes.
func (t *TestPSlice) ModifyPSlice(val int) {
	for _, i := range t.Items {
		i.val += val // i here is a pointer, so changing the value the pointer points to, persists changes

	}
}
func (t TestPSlice) ModifySlice(val int) {
	for _, i := range t.Items {
		i.val += val // even though t is a copy, because the type is a reference it becomes auto-dereferenced. So changes here are still persisted (?)
	}
}
