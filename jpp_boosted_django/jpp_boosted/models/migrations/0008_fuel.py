# Generated by Django 2.2.4 on 2019-08-18 15:58

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('models', '0007_auto_20190818_1520'),
    ]

    operations = [
        migrations.CreateModel(
            name='Fuel',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.CharField(max_length=50)),
            ],
        ),
    ]
