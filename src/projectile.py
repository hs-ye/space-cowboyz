from enum import Enum

from pyproj import Proj

class DMG_TYPE(Enum):
    KINETIC = 1
    EXPLOSION = 2
    EM = 3
    HEAT = 4


class Projectile:
    """ 
    abstract class for projectile, subclass with defualt enums to use correctly
    for implements damage model from a weapon vs ship armour etc. to determine damage

    OR, should this be a config that is read in?
    """

    def __init__(self, dmg_val: int, dtype: DMG_TYPE) -> None:
        self.dtype = dtype
        self.dmg = dmg_val
        # AP mechanic, implement later
        self.velocity = None  
        self.mass = None
    

class Bullet(Projectile):

    def __init__(self, dmg_val: int, dtype=DMG_TYPE.KINETIC) -> None:
        super().__init__(dmg_val, dtype)

class PulseBeam(Projectile):

    def __init__(self, dmg_val: int, dtype=DMG_TYPE.EM) -> None:
        super().__init__(dmg_val, dtype)

class IonBeam(Projectile):

    def __init__(self, dmg_val: int, dtype=DMG_TYPE.EM) -> None:
        super().__init__(dmg_val, dtype)

