package battle

// todo:
// add in firing speed/cooldown time
type Arms struct {
	name     string
	dmg_type string
	dps      int
	target   *Ship // target to enemy ship, starts as not set.
}

type Ship struct {
	s_class string // to be replaced by enum
	S_id    string
	hull    int   // ship hitpoints
	armour  int   // damage reduction
	weapon  *Arms // can only have 1 weapon, for now. make array later
	// also, capitalisation determines if it's exported (acessible by methods outside the function)
}

// getter for the weapon, returns a copy.
func (s *Ship) Weapon() (w Arms) {
	return *s.weapon
}

type Fleet struct {
	f_id     string // fleet identifier
	Ships    []Ship
	Lastship *Ship // testing setting pointer
	canfight bool  // default to false, maybe change this to an int, as a count?
}

type Battle struct {
	fleets [2]Fleet // for now, only 2 fleets can fight
}

// Getter for Ship arms
func (s *Ship) Arms() (arms *Arms) {
	arms = s.weapon
	return arms
}

// Getter for Arm target
func (a *Arms) Target() (t *Ship) {
	return a.target
}

// modifies the hull points remaining
func (s *Ship) ModifyHull(m int) {
	s.hull += m
}
