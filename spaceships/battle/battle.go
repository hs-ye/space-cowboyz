package battle

import "fmt"

// battle stuff goes here

// Given a weapon, finds a target in enemy fleet, sets self weapon target
// todo: give specific target to synchronise fire upon
func (a *Arms) Findtarget(f *Fleet) {
	for _, _s := range f.Ships {
		if _s.hull > 0 {
			a.target = &_s // set target
		}
	}
}

// Ship level find target
// TODO multiple weapons, individual target logic, sync weapons
func (s *Ship) Findtarget(tf *Fleet) {
	s.weapon.Findtarget(tf)
}

// arms fires at own target
func (a *Arms) Firewpn(targetfleet *Fleet) {
	// if target is dead, or no target don't fire
	if a.target == nil {
		a.Findtarget(targetfleet)
	}
	if a.target.hull <= 0 { // turns out somehow ship is already destroyed
		a.target = nil // reset target
		return         // skip firing sequence
	}
	a.target.ModifyHull(-1 * (a.dps - a.target.armour))
}

// Processes a single attack round for all ships in fleet, against enemy target fleet
func (f *Fleet) AttackRound(targetfleet *Fleet) {
	for _, s := range f.Ships {
		if s.hull > 0 {
			//TODO loop through different weapon systems to fire
			// fmt.Printf("Ship %s firing weapons: ", s.S_id)
			s.weapon.Firewpn(targetfleet)
		}
	}
}

func (b *Battle) StartBattle() {
	f1 := b.fleets[0]
	f2 := b.fleets[1]
	timer := 0
	var f1s, f2s int
	fmt.Printf("Battle begins for fleets: %s vs %s\n", f1.f_id, f2.f_id)
	for timer = 0; f1.canfight && f2.canfight; timer++ {
		fmt.Printf("Battle Round %d\n", timer)
		f1.AttackRound(&f2)
		f2.AttackRound(&f1)
		f1s = f1.CanFight()
		f2s = f2.CanFight()
		fmt.Printf("Ships remaining: Fleet 1 %d - Fleet 2 %d\n", f1s, f2s)
	}
	fmt.Printf("Battle over \n")
}
