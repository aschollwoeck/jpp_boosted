from django.urls import include, path, re_path

from . import views

urlpatterns = [
    path('projects/<int:pk>', views.ProjectDetailView.as_view()),
    path('projects/', views.ProjectsView.as_view()),
    path('tracktimes/', views.TrackTimesView.as_view()),
    path('dynos/', views.DynoMeasurementsView.as_view()),
    path('times/', views.TimesView.as_view()),
    path('about/', views.about),
    path('privacy/', views.privacy),
    #re_path(r'^search/(?P<q>[a-zA-Z0-9]+)/$', views.search),
    path('', views.IndexView.as_view()),
]