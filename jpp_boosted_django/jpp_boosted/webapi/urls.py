from django.urls import include, path

from rest_framework import routers
from . import views

router = routers.DefaultRouter()
router.register(r'videos', views.YouTubeVideoViewSet)
router.register(r'brands', views.BrandViewSet)
router.register(r'projects', views.ProjectViewSet)
router.register(r'states', views.StateViewSet)
router.register(r'parts', views.PartViewSet)
router.register(r'dynos', views.DynoMeasurementDeviceViewSet)
router.register(r'dynomeasurements', views.DynoMeasurementViewSet)
router.register(r'timerdevices', views.TimePerformanceMeasurementDeviceViewSet)
router.register(r'timermeasurements', views.TimePerformanceMeasurementViewSet)
router.register(r'tracks', views.TrackViewSet)
router.register(r'trackparts', views.TrackPartViewSet)
router.register(r'trackmeasurements', views.TrackMeasurementViewSet)
router.register(r'trackpartmeasurements', views.TrackPartMeasurementViewSet)

urlpatterns = [
    path('', include(router.urls)),
]
