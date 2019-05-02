from flask import render_template
from app import app

@app.route('/')
@app.route('/index')
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