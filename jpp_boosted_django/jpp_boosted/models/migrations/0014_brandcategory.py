# Generated by Django 2.2.4 on 2019-08-25 18:30

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('models', '0013_project_last_modified'),
    ]

    operations = [
        migrations.CreateModel(
            name='BrandCategory',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('created', models.DateTimeField(auto_now_add=True)),
                ('name', models.CharField(max_length=50)),
            ],
        ),
    ]
