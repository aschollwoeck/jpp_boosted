# Generated by Django 2.2.4 on 2019-08-25 17:45

from django.db import migrations, models
import django.utils.timezone


class Migration(migrations.Migration):

    dependencies = [
        ('models', '0012_auto_20190818_1608'),
    ]

    operations = [
        migrations.AddField(
            model_name='project',
            name='last_modified',
            field=models.DateField(default=django.utils.timezone.now),
            preserve_default=False,
        ),
    ]
