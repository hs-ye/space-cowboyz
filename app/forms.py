
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

    def __init__(self, money=1000, MarketGoods=None):
        self.money = money
        self.MarketGoods = MarketGoods
        if self.MarketGoods==None:
            self.MarketGoods = {
            "Spice": 200,
            "Robots": 150,
            "Crystals": 100
        }
        # TODO: needs to be wrapped in a WTF.form_fields and rendered altogether
        for key, value in self.MarketGoods.items():
            # goes through the dict and creates each item as a field
            exec('Item_{0}=IntegerField("{0}", description={1})'.format(key, value))

        # hardcoded version - to be removed
        # item1_buy = IntegerField("Spice")
        # item2_buy = IntegerField("Robots")
        # item3_buy = IntegerField("Crystals")
        submit = SubmitField('Buy')
        