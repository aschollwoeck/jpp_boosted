from django.contrib import admin
from models.models import YouTubeVideo, Brand, Project, State, Part, DynoMeasurementDevice, DynoMeasurement, TimePerformanceMeasurementDevice, TimePerformanceMeasurement, Track, TrackPart, TrackMeasurement, TrackPartMeasurement, Fuel

# Register your models here.
admin.site.register(YouTubeVideo)
admin.site.register(Brand)
admin.site.register(Project)
admin.site.register(State)
admin.site.register(Part)
admin.site.register(DynoMeasurementDevice)
admin.site.register(DynoMeasurement)
admin.site.register(TimePerformanceMeasurementDevice)
admin.site.register(TimePerformanceMeasurement)
admin.site.register(Track)
admin.site.register(TrackPart)
admin.site.register(TrackMeasurement)
admin.site.register(TrackPartMeasurement)
admin.site.register(Fuel)