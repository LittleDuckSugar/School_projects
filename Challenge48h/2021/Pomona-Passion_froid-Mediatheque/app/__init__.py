from flask import Flask
from app import config
from app.models import db
from app.controllers import image


app = Flask(__name__)
app.secret_key = b'_5#y2L"F4Q8z\n\xec]/'
app.config.from_object(config.Config)

db.init_app(app)

# permet de creer la base de données, à mettre en commentaire quand la creation de la bdd est terminer
# with app.app_context():
#     db.drop_all()
#     db.create_all()

app.register_blueprint(image.image)
