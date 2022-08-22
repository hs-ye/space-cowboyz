from 

class Battle:

    def __init__(self) -> None:
        self.fleets = []
        self.btimer = 0  # int for tracking battle time


    def add_fleet(self, fleet):
        self.fleets.append(fleet)

    def begin_battle(self):
        if len(self.fleets) != 2:
            
            return
