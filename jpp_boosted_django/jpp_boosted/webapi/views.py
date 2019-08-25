from django.shortcuts import render
from rest_framework import viewsets

from django.contrib.auth.models import User, Group
from models.models import YouTubeVideo, Brand, Project, State, Part, DynoMeasurementDevice, DynoMeasurement, TimePerformanceMeasurementDevice, TimePerformanceMeasurement, Track, TrackPart, TrackMeasurement, TrackPartMeasurement
from webapi.serializers import UserSerializer, GroupSerializer, YouTubeVideoSerializer, BrandSerializer, ProjectSerializer, StateSerializer, PartSerializer, DynoMeasurementDeviceSerializer, DynoMeasurementSerializer, TimePerformanceMeasurementDeviceSerializer, TimePerformanceMeasurementSerializer, TrackSerializer, TrackPartSerializer, TrackPartMeasurementSerializer, TrackMeasurementSerializer

# Create your views here.

class YouTubeVideoViewSet(viewsets.ModelViewSet):
    """
    API endpoint that allows YouTube videos to be viewed or edited.
    """
    queryset = YouTubeVideo.objects.all()
    serializer_class = YouTubeVideoSerializer

class BrandViewSet(viewsets.ModelViewSet):
    """
    API endpoint that allows YouTube videos to be viewed or edited.
    """
    queryset = Brand.objects.all()
    serializer_class = BrandSerializer

class ProjectViewSet(viewsets.ModelViewSet):
    """
    API endpoint that allows YouTube videos to be viewed or edited.
    """
    queryset = Project.objects.all()
    serializer_class = ProjectSerializer

class StateViewSet(viewsets.ModelViewSet):
    """
    API endpoint that allows YouTube videos to be viewed or edited.
    """
    queryset = State.objects.all()
    serializer_class = StateSerializer

class PartViewSet(viewsets.ModelViewSet):
    """
    API endpoint that allows YouTube videos to be viewed or edited.
    """
    queryset = Part.objects.all()
    serializer_class = PartSerializer

class DynoMeasurementDeviceViewSet(viewsets.ModelViewSet):
    """
    API endpoint that allows YouTube videos to be viewed or edited.
    """
    queryset = DynoMeasurementDevice.objects.all()
    serializer_class = DynoMeasurementDeviceSerializer

class DynoMeasurementViewSet(viewsets.ModelViewSet):
    """
    API endpoint that allows YouTube videos to be viewed or edited.
    """
    queryset = DynoMeasurement.objects.all()
    serializer_class = DynoMeasurementSerializer

class TimePerformanceMeasurementDeviceViewSet(viewsets.ModelViewSet):
    """
    API endpoint that allows YouTube videos to be viewed or edited.
    """
    queryset = TimePerformanceMeasurementDevice.objects.all()
    serializer_class = TimePerformanceMeasurementDeviceSerializer

class TimePerformanceMeasurementViewSet(viewsets.ModelViewSet):
    """
    API endpoint that allows YouTube videos to be viewed or edited.
    """
    queryset = TimePerformanceMeasurement.objects.all()
    serializer_class = TimePerformanceMeasurementSerializer

class TrackViewSet(viewsets.ModelViewSet):
    """
    API endpoint that allows YouTube videos to be viewed or edited.
    """
    queryset = Track.objects.all()
    serializer_class = TrackSerializer

class TrackPartViewSet(viewsets.ModelViewSet):
    """
    API endpoint that allows YouTube videos to be viewed or edited.
    """
    queryset = TrackPart.objects.all()
    serializer_class = TrackPartSerializer

class TrackMeasurementViewSet(viewsets.ModelViewSet):
    """
    API endpoint that allows YouTube videos to be viewed or edited.
    """
    queryset = TrackMeasurement.objects.all()
    serializer_class = TrackMeasurementSerializer

class TrackPartMeasurementViewSet(viewsets.ModelViewSet):
    """
    API endpoint that allows YouTube videos to be viewed or edited.
    """
    queryset = TrackPartMeasurement.objects.all()
    serializer_class = TrackPartMeasurementSerializer