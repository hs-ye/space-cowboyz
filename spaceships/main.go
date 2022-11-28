package main

import (
	"fmt"
	"reflect"
	b "spaceships/battle"
) // imports are folder paths, will can only have one package per path

// run a test battle
func testbattle(nf1, nf2 int) {
	f1 := b.NewFleet("f1")
	f1.AddShip(nf1)
	f2 := b.NewFleet("f2")
	f2.AddShip(nf2)
	fmt.Printf("two fleets are %p vs %p\n ", &f1, &f2)

	btl := b.NewBattle(&[2]b.Fleet{f1, f2})
	fmt.Printf("two fleets in the battle are %p vs %p\n ", btl.Fleets(), btl.Fleets())
	fmt.Println(*btl.Fleets())
	btl.StartBattle()
}

// ok so what is the deal with pointers getting and setting
func debug() {
	f1 := b.NewFleet("f1")
	f1.AddShip(1)
	f2 := b.NewFleet("f2")
	f2.AddShip(1)

	f2.Lastship.Findtarget(&f1)
	f1.Lastship.Findtarget(&f2)
	fmt.Printf("ship %s target found: %s at location %p\n", f1.Lastship.S_id, f1.Lastship.Arms().Target().S_id, f1.Lastship.Arms().Target())
	fmt.Printf("ship %s target found: %s at location %p\n", f2.Lastship.S_id, f2.Lastship.Arms().Target().S_id, f2.Lastship.Arms().Target())
	fmt.Printf("f1 ship: %p - f2 ship %p\n", &f1.Ships[0], &f2.Ships[0])
	fmt.Printf("f2 target: %p - f1 target %p\n", f1.Ships[0].Arms().Target(), f2.Ships[0].Arms().Target())

	s1 := f1.Ships[0]
	s2 := f2.Ships[0]
	s1.Arms().Firewpn()
	fmt.Println("HP left after attack:", s2.Hull())

	// manual call modify - target is wrong here
	f1.Ships[0].Arms().Target().ModifyHull(-100) // this is the line that's failing to resolve pointer, modifying on a copy here
	fmt.Println("HP left after manual change #1:", s2.Hull())
	fmt.Println(reflect.TypeOf(f1.Ships[0].Arms().Target()))
	fmt.Printf("s1 %p s2 target %p\n", s1, f2.Ships[0].Arms().Target())
	if s1 == f1.Ships[0].Arms().Target() {
		fmt.Println("Target is the same!")
	}
	// manually set target  -- works
	f1.Ships[0].Arms().SetTarget(s2)
	f1.Ships[0].Arms().Firewpn()
	fmt.Println("HP left after manual target set:", s2.Hull())

	// manually modify hull hp - works
	s2.ModifyHull(-100)
	fmt.Println("HP left after manual direct update:", s2.Hull())

}

func main() {
	// println("hello test")

	// how to append to a slice, vs how to append to pointer to slice?

	// scr.ScratchTest()
	testbattle(8, 10)
	// debug()
	// scr.Slicingtest()
}
