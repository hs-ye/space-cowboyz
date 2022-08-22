# Testing scenario

# for setting up two basic fleets in a fight

# loads the configs for fleet 1, 2
# initiates initial positions, state etc.

# sets the timer, and runs attack
# game continues until one fleet runs out of battle ready ships

from ship import ShipFactory
from fleet import Fleet

ships = ShipFactory.mk_ships_from_file("data/ships.json")

fleet1 = Fleet("test-fleet-1")
for s in ships:
    fleet1.add_ship(s)

