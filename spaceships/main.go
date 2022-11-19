package main

import (
	"fmt"
	b "spaceships/battle"
) // folder name, will grab all packages in this path

func main() {
	// println("hello test")

	s := b.Makeship("testshipid")
	b.Setarms(&s) // pass it the pointer
	fmt.Println(s)

	fmt.Println(s.Weapon()) // prints the weapon
	// how to append to a slice, vs how to append to pointer to slice?

	// ships := []b.Ship{}
	// s_p := &ships
	// ships = append(ships, b.Makeship())
	// ships = append(*s_p, b.Makeship())

	f1 := b.NewFleet("f1")
	f1.AddShip(5)
	f2 := b.NewFleet("f2")
	f2.AddShip(3)

	f2.Lastship.Findtarget(&f1)
	f1.LastShip.Findtarget(&f1)
	fmt.Printf("ship %s target found: %s", f2.Lastship.S_id, f2.Lastship.Arms().Target().S_id)

	// *append(f.ships, Makeship())
}
