from src.ship import Ship

class Fleet():
    """ Fleets are composed of ships"""

    def __init__(self, fleetname: str, max_ships = 100) -> None:
        self.fleetname = fleetname
        self.fleet_ships = {}  # dict of ships, by name
        self.n_battle_ready_ships = 0
        self.max_ships = max_ships


    def add_ship(self, name: str, ship: Ship):
        """ Instantiates a ship, adds there"""
        if len(self.fleet_ships) >= self.max_ships:
            self.fleet_ships[name] = ship
            self.n_battle_ready_ships += 1  # increments number of battle ready ships
        else:
            print(f"Max ships in fleet reached: {self.max_ships}, cannot add any more")
    
    def repair_ship(self, name:str):
        self.fleet_ships[name].repair_ship()


        

    def remove_ship(self, name):
        """ Should be a way to remove a ship by name also """
        del self.fleet_ships[name]
        # otherwise do nothing
        