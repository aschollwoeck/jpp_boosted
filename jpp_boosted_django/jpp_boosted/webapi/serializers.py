from django.contrib.auth.models import User, Group
from models.models import YouTubeVideo, Brand, Project, State, Part, DynoMeasurementDevice, DynoMeasurement, TimePerformanceMeasurementDevice, TimePerformanceMeasurement, Track, TrackPart, TrackMeasurement, TrackPartMeasurement
from rest_framework import serializers

class UserSerializer(serializers.HyperlinkedModelSerializer):
    class Meta:
        model = User
        fields = ['url', 'username', 'email', 'groups']

class GroupSerializer(serializers.HyperlinkedModelSerializer):
    class Meta:
        model = Group
        fields = ['url', 'name']

class YouTubeVideoSerializer(serializers.HyperlinkedModelSerializer):
    class Meta:
        model = YouTubeVideo
        fields = ['url', 'title', 'city', 'country', 'created']

class BrandSerializer(serializers.HyperlinkedModelSerializer):
    class Meta:
        model = Brand
        fields = ['url', 'name']

class ProjectSerializer(serializers.HyperlinkedModelSerializer):
    class Meta:
        model = Project
        fields = ['url', 'title', 'brand', 'code', 'series', 'model', 'fuel', 'production_date', 'purchase_date']

class StateSerializer(serializers.HyperlinkedModelSerializer):
    class Meta:
        model = State
        fields = ['url', 'title', 'description', 'modified_at', 'youtubevideo', 'parts']

class PartSerializer(serializers.HyperlinkedModelSerializer):
    class Meta:
        model = Part
        fields = ['url', 'category', 'name', 'url', 'weight', 'weight_reduction', 'brand']

class DynoMeasurementDeviceSerializer(serializers.HyperlinkedModelSerializer):
    class Meta:
        model = DynoMeasurementDevice
        fields = ['url', 'name', 'owner']

class DynoMeasurementSerializer(serializers.HyperlinkedModelSerializer):
    class Meta:
        model = DynoMeasurement
        fields = ['url', 'horse_power', 'newtonmeter', 'temperature', 'dyno_device', 'state']

class TimePerformanceMeasurementDeviceSerializer(serializers.HyperlinkedModelSerializer):
    class Meta:
        model = TimePerformanceMeasurementDevice
        fields = ['url', 'name']

class TimePerformanceMeasurementSerializer(serializers.HyperlinkedModelSerializer):
    class Meta:
        model = TimePerformanceMeasurement
        fields = ['url', 'speed_start', 'speed_end', 'duration', 'measurement_device', 'state']

class TrackSerializer(serializers.HyperlinkedModelSerializer):
    class Meta:
        model = Track
        fields = ['url', 'name']

class TrackPartSerializer(serializers.HyperlinkedModelSerializer):
    class Meta:
        model = TrackPart
        fields = ['url', 'name', 'track']
class TrackMeasurementSerializer(serializers.HyperlinkedModelSerializer):
    class Meta:
        model = TrackMeasurement
        fields = ['url', 'duration', 'highest_speed', 'youtubevideo', 'track', 'state']

class TrackPartMeasurementSerializer(serializers.HyperlinkedModelSerializer):
    class Meta:
        model = TrackPartMeasurement
        fields = ['url', 'duration', 'highest_speed', 'track_part', 'track_measurement']