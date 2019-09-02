from django.db import models
import re

# Create your models here.

class YouTubeVideo(models.Model):
    created = models.DateTimeField(auto_now_add=True)
    
    url = models.URLField(max_length=100)
    title = models.CharField(max_length=100)
    city = models.CharField(max_length=100)
    country = models.CharField(max_length=100)

    def __str__(self):
        return self.title

    def get_watch_code(self):
        m = re.search("watch\?v=(.+)", self.url)
        if(m is None):
            return ""
        
        return m.group(1)

class BrandCategory(models.Model):
    created = models.DateTimeField(auto_now_add=True)

    name = models.CharField(max_length=50)

    def __str__(self):
        return self.name
    

class Brand(models.Model):
    created = models.DateTimeField(auto_now_add=True)

    name = models.CharField(max_length=100)
    category = models.ForeignKey(BrandCategory, on_delete=models.DO_NOTHING)
    url = models.URLField()

    def __str__(self):
        return self.name

class Fuel(models.Model):
    name = models.CharField(max_length=50)

    def __str__(self):
        return self.name
    

class Project(models.Model):
    created = models.DateTimeField(auto_now_add=True)

    title = models.CharField(max_length=100, unique=True)

    brand = models.ForeignKey(Brand, on_delete=models.DO_NOTHING)
    code = models.CharField(max_length=50)
    series = models.CharField(max_length=50)
    model = models.CharField(max_length=50)
    fuel = models.ForeignKey(Fuel, on_delete=models.DO_NOTHING)
    production_date = models.DateField()
    purchase_date = models.DateField()
    last_modified = models.DateField()
    image = models.ImageField(max_length=200, upload_to="static/")

    def __str__(self):
        return self.title

class Part(models.Model):
    created = models.DateTimeField(auto_now_add=True)

    category = models.CharField(max_length=100)
    name = models.CharField(max_length=100)
    url = models.URLField()

    weight = models.DecimalField(max_digits=5, decimal_places=2)
    weight_reduction = models.DecimalField(max_digits=5, decimal_places=2)

    brand = models.ForeignKey(Brand, on_delete=models.DO_NOTHING)

    def __str__(self):
        return self.name

class State(models.Model):
    created = models.DateTimeField(auto_now_add=True)

    title = models.CharField(max_length=100)
    description = models.TextField()
    modified_at = models.DateField()
    image = models.ImageField(max_length=200, upload_to="static/")

    project = models.ForeignKey(Project, on_delete=models.CASCADE)
    youtubevideo = models.ManyToManyField(YouTubeVideo)
    parts = models.ManyToManyField(Part, blank=True)

    def save(self, *args, **kwargs):
        if(self.project.last_modified < self.modified_at):
            self.project.last_modified = self.modified_at
            self.project.save()
        
        super(State, self).save(*args, **kwargs)

    def __str__(self):
        return self.project.title + " - " + self.title

    def get_time_0_to_100(self):
        return TimePerformanceMeasurement.objects.get(state__id = self.id, speed_start = 0, speed_end = 100)
    
    def get_time_100_to_200(self):
        return TimePerformanceMeasurement.objects.get(state__id = self.id, speed_start = 100, speed_end = 200)

    def get_time_0_to_200(self):
        return TimePerformanceMeasurement.objects.get(state__id = self.id, speed_start = 0, speed_end = 200)

class TimePerformanceMeasurementDevice(models.Model):
    created = models.DateTimeField(auto_now_add=True)

    name = models.CharField(max_length=100)

    def __str__(self):
        return self.name

class TimePerformanceMeasurement(models.Model):
    created = models.DateTimeField(auto_now_add=True)

    speed_start = models.PositiveIntegerField()
    speed_end = models.PositiveIntegerField()
    duration = models.DurationField()

    measurement_device = models.ForeignKey(TimePerformanceMeasurementDevice, on_delete=models.DO_NOTHING)

    state = models.ForeignKey(State, on_delete=models.DO_NOTHING)

    def __str__(self):
        return self.state.project.title + " - " + self.state.title
        #return self.speed_start.__str__ + " " + self.speed_end.__str__ + " = " + self.duration.__str__

class DynoMeasurementDevice(models.Model):
    created = models.DateTimeField(auto_now_add=True)

    name = models.CharField(max_length=100)
    owner = models.CharField(max_length=100)

    def __str__(self):
        return self.name + " (" + self.owner + ")"

class DynoMeasurement(models.Model):
    created = models.DateTimeField(auto_now_add=True)

    horse_power = models.PositiveIntegerField()
    newtonmeter = models.PositiveIntegerField()
    temperature = models.DecimalField(max_digits=5, decimal_places=2)    

    dyno_device = models.ForeignKey(DynoMeasurementDevice, on_delete=models.DO_NOTHING)

    state = models.ForeignKey(State, on_delete=models.DO_NOTHING)

    def __str__(self):
        return self.state.__str__() + ": " + self.dyno_device.__str__() + " (PS: " + self.horse_power.__str__() + " NM: " + self.newtonmeter.__str__() + ")"

    def get_state(self):
        return self.state

class Track(models.Model):
    created = models.DateTimeField(auto_now_add=True)

    name = models.CharField(max_length=100)

    def __str__(self):
        return self.name    

class TrackPart(models.Model):
    created = models.DateTimeField(auto_now_add=True)

    name = models.CharField(max_length=100)

    track = models.ForeignKey(Track, on_delete=models.CASCADE)

    def __str__(self):
        return self.name

class TrackMeasurement(models.Model):
    created = models.DateTimeField(auto_now_add=True)

    duration = models.DurationField()
    highest_speed = models.PositiveIntegerField()
    youtubevideo = models.ForeignKey(YouTubeVideo, on_delete=models.DO_NOTHING)

    track = models.ForeignKey(Track, on_delete=models.DO_NOTHING)

    state = models.ForeignKey(State, on_delete=models.DO_NOTHING)

    def __str__(self):
        return self.state.title + " - " + self.track.name + " - " + self.youtubevideo.title

class TrackPartMeasurement(models.Model):
    created = models.DateTimeField(auto_now_add=True)

    duration = models.DurationField()
    highest_speed = models.PositiveIntegerField()

    track_part = models.ForeignKey(TrackPart, on_delete=models.DO_NOTHING)

    track_measurement = models.ForeignKey(TrackMeasurement, on_delete=models.DO_NOTHING)

    def __str__(self):
        return self.track_part.track.__str__() + " - " + self.track_part.name + " - " + self.track_measurement.youtubevideo.title