package main

import (
	"fmt"
	b "spaceships/battle"
) // imports are folder paths, will can only have one package per path

// run a test battle
func testbattle() {
	f1 := b.NewFleet("f1")
	f1.AddShip(5)
	f2 := b.NewFleet("f2")
	f2.AddShip(3)

	f2.Lastship.Findtarget(&f1)
	f1.Lastship.Findtarget(&f1)
	fmt.Printf("ship %s target found: %s\n", f2.Lastship.S_id, f2.Lastship.Arms().Target().S_id)

	btl := b.NewBattle([2]b.Fleet{f1, f2})
	btl.StartBattle()
}

func main() {
	// println("hello test")

	// how to append to a slice, vs how to append to pointer to slice?

	// scr.ScratchTest()
	testbattle()
}
