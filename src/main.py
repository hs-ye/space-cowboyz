from enum import Enum

from pip import main
from ship import ship, ship_state


game_config = []  # add in ship config


class SHIP_TAC_POS(Enum):
    FORWARD = 0
    MID = 1
    REAR = 2



class gamemap():
    """ Maps hold fleets, features/obstacles, have a size """
    def __init__(self) -> None:
        self.mapsize = None  # placeholder
        # self.grid =   # placeholder

class game():
    """ 
    Games have ships and maps, has a ticker that simulates the engine 
    """
    def __init__(self, ships) -> None:
        self.ships = ships  # of type ship
        self.map = gamemap
        self.game_state = 0  # 0 is not ended, 1 is ended
        self.active = ships[1]
    
    def setup_game(self):
        """ sets up the game """
        pass

    def start_game(self):
        """ Starts the game"""
        while self.game_state != 1:
            print("Game Simulation running!")
            # ask for player input here
            self.process_turn()  # runs turns until finished
        
        print("Game ended.")
            
    def ask_for_player_input(self):
        pass # takes input here

    def check_turn(self):
        """ Checks who's turn it is it move/shoot"""
        pass

    def process_turn(self):
        """ does a turn """
        pass

if __name__