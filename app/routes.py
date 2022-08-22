from flask import render_template, flash, redirect, url_for
from app import app_instance
from app.forms import SignInForm, GameMarketBuyForm  # import neds to be from the specific py file

@app_instance.route('/')
@app_instance.route('/index')
def index():
    user = {'username': 'Test'}
    posts = [
        {
            'author': {'username': 'Wep'},
            'body': 'This is a test'
        }
        ,{
            'author': {'username': 'the second'},
            'body': '2nd line of text'
        }
    ]

    return render_template('index.html', title='game1', user=user, posts=posts)

@app_instance.route("/login", methods=["GET", "POST"])  # default only allows GET method
def login_page():
    form = SignInForm()  # object class obtained from the app.forms module
    # this contains the field defnitions (booleans, string & pw fields)
    # it is passed into the render_template, which expects these definitions

    # to handle user data being submitted (this function does the heavy lifting):
    if form.validate_on_submit():
        # the validate_on_submit method won't trigger during the initial GET request
        # from the browser to fetch the html of the page, hence this check
        # or if any of the field validations (e.g. pw) fails. Can add error message later
        flash("Login requested for user {}, remember_me={}".format(
            form.username.data, form.remember_login.data
        ))  # Flask method, flash shows a message to the user. Needs html (see base.html)
        # the login doesn't do anything yet, besides confirm data is recieved
        return redirect(url_for('game'))  # sends user somewhere else
    # the url_for calls the corresponding page handling function rather than the path directly
    return render_template("login.html", form=form, title="Login to Game")

@app_instance.route("/game", methods=["GET", "POST"])
def game():
    """The main game screen. WIP. Currently just has the market buy"""
    form = GameMarketBuyForm()
    if form.validate_on_submit():
        """TODO Updates variables for the market/inventory of the player
        get the variable changes from form fields, then pass through 
        to form object
        """
        pass
    return render_template("game.html", form=form, title="Game Market")
