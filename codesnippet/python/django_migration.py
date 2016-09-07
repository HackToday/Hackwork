# -*- coding: utf-8 -*-

## Before we coded migration business logic, we could first use tools:
#
#	python manage.py makemigrations
#
# That would include field changes, like alterfield, addfield etc.
# Then add your business logic there according to your needs.
#
## refer https://docs.djangoproject.com/en/1.10/ref/migration-operations/#runpython


from __future__ import unicode_literals

from django.db import migrations, models


def forwards_func(apps, schema_editor):
    # Note: This is important!
    # We get the model from the versioned app registry;
    # if we directly import it, it'll be the wrong version
    Country = apps.get_model("myapp", "Country")
    db_alias = schema_editor.connection.alias
    Country.objects.using(db_alias).bulk_create([
        Country(name="USA", code="us"),
        Country(name="France", code="fr"),
    ])

def reverse_func(apps, schema_editor):
    # forwards_func() creates two Country instances,
    # so reverse_func() should delete them.
    Country = apps.get_model("myapp", "Country")
    db_alias = schema_editor.connection.alias
    Country.objects.using(db_alias).filter(name="USA", code="us").delete()
    Country.objects.using(db_alias).filter(name="France", code="fr").delete()


class Migration(migrations.Migration):

    dependencies = []

    operations = [
        migrations.RunPython(forwards_func, reverse_func),
    ]
