# Generated by Django 2.2.4 on 2019-08-16 12:40

from django.db import migrations, models
import django.db.models.deletion


class Migration(migrations.Migration):

    dependencies = [
        ('models', '0002_auto_20190816_1045'),
    ]

    operations = [
        migrations.CreateModel(
            name='Brand',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('created', models.DateTimeField(auto_now_add=True)),
                ('name', models.CharField(max_length=100)),
                ('url', models.URLField()),
            ],
        ),
        migrations.CreateModel(
            name='TimePerformanceMeasurementDevice',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('created', models.DateTimeField(auto_now_add=True)),
                ('name', models.CharField(max_length=100)),
            ],
        ),
        migrations.CreateModel(
            name='TimePerformanceMeasurement',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('created', models.DateTimeField(auto_now_add=True)),
                ('speed_start', models.PositiveIntegerField()),
                ('speed_end', models.PositiveIntegerField()),
                ('duration', models.TimeField()),
                ('measurement_device', models.ForeignKey(on_delete=django.db.models.deletion.DO_NOTHING, to='models.TimePerformanceMeasurementDevice')),
            ],
        ),
        migrations.CreateModel(
            name='State',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('created', models.DateTimeField(auto_now_add=True)),
                ('title', models.CharField(max_length=100)),
                ('description', models.TextField()),
                ('modified_at', models.DateField()),
                ('youtubevideo', models.ManyToManyField(to='models.YouTubeVideo')),
            ],
        ),
        migrations.CreateModel(
            name='Project',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('created', models.DateTimeField(auto_now_add=True)),
                ('title', models.CharField(max_length=100)),
                ('code', models.CharField(max_length=50)),
                ('series', models.CharField(max_length=50)),
                ('model', models.CharField(max_length=50)),
                ('fuel', models.CharField(max_length=50)),
                ('production_date', models.DateField()),
                ('purchase_date', models.DateField()),
                ('brand', models.ForeignKey(on_delete=django.db.models.deletion.DO_NOTHING, to='models.Brand')),
            ],
        ),
        migrations.CreateModel(
            name='Part',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('created', models.DateTimeField(auto_now_add=True)),
                ('category', models.CharField(max_length=100)),
                ('name', models.CharField(max_length=100)),
                ('url', models.URLField()),
                ('weight', models.DecimalField(decimal_places=2, max_digits=5)),
                ('weight_reduction', models.DecimalField(decimal_places=2, max_digits=5)),
                ('brand', models.ForeignKey(on_delete=django.db.models.deletion.DO_NOTHING, to='models.Brand')),
            ],
        ),
    ]
