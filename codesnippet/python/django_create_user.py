#https://docs.djangoproject.com/en/1.10/topics/auth/default/

from django.contrib.auth.models import User

User.objects.create_user(username=username, email=email, password=password)
User.objects.create_superuser(username=username, email=email, password=password)

