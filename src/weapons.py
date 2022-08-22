# classes related to teh weapons
from enum import Enum
from projectile import Bullet


class weapon_type(Enum):
    KINETIC = 1
    ENERGY = 2
    MISSILE = 3

# weapon data
class Weapon():
    """ Abstract class for weapon. Subclass to use effectively"""
    def __init__(self, weapon_data):
        self.name = "weapon_name"
        self.damage_type = Enum()
        self.damage = 10
        self.rof = 1  # rate of fire, per sec
        self.reload_time = 10  # number of ticks to reload weapon
        self.current_cd = 0  # track timer against total. When it reaches cooldown, ready to fire
        self.target = None  # should be a pointer to an enemy ship
    
    def fire(self, cdtime = 1):
        """ needs a way to pass the timer/cooldown into the weapon"""
        if self.target == None:
            return  # skip, no action as no target locked
        if self.current_cd > 0:
            self.current_cd -= cdtime
            return  # weapon still reloading
        self.target.on_hit(self)
        self.current_cd = self.reload_time  # reset reload counter. TODO make this account for game ticks that is not 1




    