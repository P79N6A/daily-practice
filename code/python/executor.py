from app.celery_task import add_together

a = 11
b = 22
result = add_together.delay(11, 22)
print('add_together({0}, {1})= {2}'.format(a, b, result.wait()))