package battle

// battle stuff goes here

// Given a ship, finds a target in enemy fleet, sets self weapon target
// todo: conditions under which a target switch occurs
func (s *Ship) Findtarget(f *Fleet) {
	for _, _s := range f.Ships {
		if _s.hull > 0 {
			s.weapon.target = &_s // set target
		}
	}
}

// arms fires at an enemy ship
func (a *Arms) Firewpn(s *Ship) {
	// if target is dead, or no target don't fire
	if a.target != nil && a.target.hull <= 0 {
		return // stop function
	}

	// &Ship.s_class = "test" // doesn't work

}

func (b *Battle) StartBattle(f1, f2 *Fleet) {

}
