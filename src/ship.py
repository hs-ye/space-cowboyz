import json, uuid
from src.weapons import Weapon
from projectile import Projectile

# all classes related to the ship



class ShipFactory:
    """Creates ships """

    @staticmethod
    def mk_ship(shipdata):
        """Makes a single ship. TODO validate the data file"""
        return Ship(shipdata)
    
    @staticmethod
    def mk_ships_from_file(filepath):
        """Makes ships from a file"""
        with open(filepath) as f:
            ship_types = json.load(f)

class ShipState:
    """ holds the state of the ship, once a ship is created """
    def __init__(self, ship_data) -> None:
        self.current_armour = ship_data.armour  # starts at max
        self.current_shield = ship_data.shield  # starts at max
        self.current_engine = ship_data.engine  # starts at max
        self.target = None  # no target
        self.battle_ready = True  # if false, ship cannot fight


class BridgeTactics:
    """ Bridge sets the tactical combat commands
    # handles things like positioning
    # how to select targets
    # method to select the next target in enemy fleet
    """
    def __init__(self, bridge_data) -> None:
        self.row = bridge_data.row
        self.target_priority = bridge_data.target_priority

class Ship:
    """ 
    Ships have stuff mounted on them, this is the base stats of the ship.
    E.g. the state should check against this for HP
    """
    def __init__(self, ship_data):
        self.weapon = ship_data.weapon  # could be multiple weapons, later
        self.shipclass = ship_data.shipclass
        self.max_armour = ship_data.armour
        self.max_shield = ship_data.shield
        self.max_engine = ship_data.engine  # engine combined with mass determine speed
        self.mass = ship_data.mass  # mass also used as HP
        self.position = None  # depends on the initial positions on the map
        self.name = uuid.uuid4()
        # state object: Holds things for the actual battle
        self.state = ShipState(ship_data)


    def on_hit(self, proj: Projectile):
        """Calculate damage due to projectile hit, implements a damage model"""
        # first, check if weapon hits
        # next, check armour / shield reduction
        # finally, calculate amount of damage using projectile properties
        pass

    def attack_round(self):
        # TODO need to go through all weapons here
        self.find_target(self.weapon)
        if self.weapon != None:  # not none
            self.fire_weapons()  # possible that weapons lock onto different targets?

    def find_target(self, weapon):
        # uses battle computer to find suitable targets for each weapon
        # selects a target based on firing strategy

        # if no targets remaining, skip
        weapon.target = None

    def fire_weapons(self):
        """ Fires the main weapon system"""
        # applies damage against target ship in target fleet
        pass

    def movement(self):
        """ Parses a movement command input"""
        pass
    
    def repair_ship(self):

        if not self.state.battle_ready:
            self.state.battle_ready = True
        else:
            print(f"ship {self.name} doesn't need repairing")