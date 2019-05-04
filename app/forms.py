
from flask_wtf import FlaskForm
from wtforms import StringField, PasswordField, BooleanField, SubmitField, IntegerField
from wtforms.validators import DataRequired

class SignInForm(FlaskForm):
    """extends the base FlaskForm to allow for logins"""
    username = StringField('Username', validators=[DataRequired()])
    pw = PasswordField('Password', validators=[DataRequired()])
    remember_login = BooleanField('Remember Me')
    submit = SubmitField('Log in')

class GameMarketBuyForm(FlaskForm):
    """used in the game to buy from the current market
    requires as input an object containing all the goods & prices
    available for sale in the current market
    """

    # will need to replace above with a dynamic generated version, from market object
    # for good in MarketGoods:
    # TODO: i think if i wrap this in an array object, then access the same object via Jinja it should work?
    item1_buy = IntegerField("Spice")
    item2_buy = IntegerField("Robots")
    item3_buy = IntegerField("Crystals")
    submit = SubmitField('Buy')

    def __init__(self, money=1000, MarketGoods=None):
        self.money = money
        self.MarketGoods = MarketGoods
        if MarketGoods==None:
            self.MarketGoods = {
            "Spice": 200,
            "Robots": 150,
            "Crystals": 100
        }
        #if MarketGoods==None:  # need to find a way to pass a market data object in
        