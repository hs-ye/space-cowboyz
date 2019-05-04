from flask import Flask  # gets Flask object from flask package
from config import Config # gets the Config object from config.py

app_instance = Flask(__name__)  # this is the 'app' object that gets imported 
# form app import app, because it is in the app folder (from app)
# renamed the instance to app_instance to make it less confusing
app_instance.config.from_object(Config)  # gets the info from the Config object

from app import routes