from flask import Flask
from app.celery_task import make_celery

app = Flask(__name__)
app.config.update(
    CELERY_BROKER_URL='redis://localhost:6379',
    CELERY_BACKEND_URL='redis://localhost:6379'
)
celery = make_celery(app)