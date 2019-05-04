# use a class to store all config options & settings

import os

class Config(object):
    # store config variables inside a dict in here
    # prefer an environment var, but if not available, use a placeholder
    SECRET_KEY = os.environ.get('SECRET_KEY') or 'placeholder_secret_key'

