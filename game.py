from app import app_instance
from app.routes import game, GameMarketBuyForm

# for debugging using 'flask shell' in CMD. need to 'set FLASK_APP=game.py' in cmd env variables
@app_instance.shell_context_processor
def make_shell_context():
    return {'game': game, 'GameMarketBuyForm': GameMarketBuyForm}

if __name__ == "__main__":
    app_instance.run(debug=True)