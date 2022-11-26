package battle

import (
	"fmt"
	"reflect"
)

// battle stuff goes here
const MaxBattleRounds = 3600 // max rounds

// Given a weapon, finds a target in enemy fleet, sets self weapon target
// todo: give specific target to synchronise fire upon
func (a *Arms) Findtarget(f *Fleet) {
	for _, _s := range f.Ships { // BUG: This is the problem!
		// using range gives copy of the values.
		// if you are needing original copy, store the data as pointers and pass the copies of the pointers around.
		if _s.hull > 0 {
			fmt.Printf("gun %s getting target: %p\n", a.name, &_s)
			a.target = &_s // set target
			fmt.Printf("gun %s target set: %p\n", a.name, a.target)
			return
		}
	}
}

// Ship level find target
// TODO multiple weapons, individual target logic, sync weapons
func (s *Ship) Findtarget(tf *Fleet) {
	fmt.Printf("finding target in fleet: %p\n ", tf)
	s.weapon.Findtarget(tf)
}

// arms fires at own target
// findtarget shuold be separate step
func (a *Arms) Firewpn() {
	// if target is dead, or no target don't fire
	// fmt.Printf("targetfleet: %p\n ", targetfleet)
	if a.target == nil {
		return
		// a.Findtarget(targetfleet)
	}
	if a.target.hull <= 0 { // turns out somehow ship is already destroyed
		a.target = nil // reset target
		return         // skip firing sequence
	}
	fmt.Printf("Firing at object: %p \n ", a.target)
	e_dmg := max(a.dps-a.target.armour, 0)
	a.target.ModifyHull(-1 * e_dmg)
	fmt.Printf("Weapon %s fired at %s, caused %d dmg, %d hull remains\n", a.name, a.target.S_id, e_dmg, a.target.hull)
	// a.target.ModifyHull(-1 * (a.dps - a.target.armour))
}

// Processes a single attack round for all ships in fleet, against enemy target fleet
func (f *Fleet) AttackRound(targetfleet *Fleet) {
	fmt.Printf("Attack round from %s: %p against target %s: %p\n ", f.f_id, f, targetfleet.f_id, targetfleet)
	fmt.Println(reflect.TypeOf(f), reflect.TypeOf(targetfleet))
	for _, s := range f.Ships {
		if s.hull > 0 {
			//TODO loop through different weapon systems to fire
			// Set targeting here against fleet
			// fmt.Printf("Ship %s firing weapons: ", s.S_id)
			s.weapon.Firewpn()
		}
	}
}

func (b *Battle) StartBattle() {
	f1 := &b.fleets[0]
	f2 := &b.fleets[1]
	timer := 0
	var f1s, f2s int
	fmt.Printf("Battle begins for fleets: %s vs %s\n", f1.f_id, f2.f_id)
	fmt.Printf("objects: %p vs %p\n", f1, f2)
	for timer = 0; f1.canfight && f2.canfight && timer < MaxBattleRounds; timer++ {
		fmt.Printf("Battle Round %d\n", timer)
		f1.AttackRound(f2)
		f2.AttackRound(f1)
		f1s = f1.CanFight()
		f2s = f2.CanFight()
		fmt.Printf("Ships remaining: Fleet 1 %d - Fleet 2 %d\n", f1s, f2s)
	}
	fmt.Printf("Battle over \n")
}
