package battle

import (
	"fmt"
	"strconv"
)

// for now just gives it default arms
// TODO: set custom arms from template
func Setarms(s *Ship) {
	if s.weapon != nil { // must check it's not nill,
		fmt.Println(*s.weapon) // before attemping to deference & print
	}
	fmt.Println("setting default guns on ship", s.S_id)
	s.weapon = &Arms{"testgun", "normal", 100, nil} // get the pointer to new arms

	// this below won't work, as you are trying to set an object into a supposed pointer field
	// newgun := Arms{"testgun", "normal", 100, nil}
	// *s.weapon = newgun

}

// todo: Construct a ship from template
func Makeship(id string) Ship {
	s := Ship{"dd", id, 1000, 10, nil} //  note you can't ever access/evaluate a nil field - will result in panic
	fmt.Println("setting guns on ship", s.S_id)
	Setarms(&s) // this breaks if the Arms pointer is Nil, works if blank arms initialised.
	return s
}

func NewFleet(fname string) (f Fleet) {
	f = Fleet{fname, []Ship{}, nil}
	return f
}

// Adds a ship to a fleet
// TODO adds specified ship
func (f *Fleet) AddShip(n int) {
	if n <= 0 {
		fmt.Println("Must add 1 or more!")
		return
	}
	var newship Ship
	var _sid string
	for i := 1; i <= n; i++ {
		_sid = f.f_id + " " + strconv.Itoa(i)
		newship = Makeship(_sid)
		// not sure if this loop is optimised? Maybe preallocate the looped slice?
		// it's fine as long as not adding huge amount and causing lots of re-allocations
		f.Ships = append(f.Ships, newship)
	}
	fmt.Printf("fleet %s added ships %d\n", f.f_id, len(f.Ships))
	f.Lastship = &newship
}
