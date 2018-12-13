from celery import Celery
from flask import Flask

def make_celery(app):
    celery = Celery(app.import_name, backend=app.config['CELERY_BACKEND_URL'], 
                    broker=app.config['CELERY_BROKER_URL'])
    celery.conf.update(app.config)    
    TaskBase = celery.Task

    class ContextTask(TaskBase):
        abstract = True
        def __call__(self, *args, **kwargs):
            with app.app_context():
                return TaskBase.__call__(self, *args, **kwargs)
        def direct_call(self, *args, **kwargs):
            return task_base.__call__(self, *args, **kwargs)

    celery.Task = ContextTask
    return celery

app = Flask(__name__)
app.config.update(
    CELERY_BROKER_URL='redis://localhost:6379',
    CELERY_BACKEND_URL='redis://localhost:6379'
)
celery = make_celery(app)

@celery.task()
def add_together(a, b):
    return a + b