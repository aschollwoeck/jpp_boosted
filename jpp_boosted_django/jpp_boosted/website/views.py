from django.shortcuts import render
from django.views import generic

from models.models import Project, State, TrackMeasurement, DynoMeasurement, TimePerformanceMeasurement, Brand, Fuel, Track

# Create your views here.
class IndexView(generic.ListView):
    template_name = 'index.html'
    context_object_name = 'project_list'

    def get_queryset(self):
        return Project.objects.order_by('-state__modified_at')[:5]

class FilterData():
    query = ""
    brand = ""
    brands = []
    model = ""
    models = []
    fuel = ""
    fuels = []
    ps_from = ""
    ps_to = ""
    nm_from = ""
    nm_to = ""
    order = ""

    def load_from_request(self, request):
        # Load the selected filter fields
        self.query = request.GET.get("q")
        self.brand = request.GET.get("brand")
        self.model = request.GET.get("model")
        self.fuel = request.GET.get("fuel")
        self.hp_from = request.GET.get("hp_from")
        self.hp_to = request.GET.get("hp_to")
        self.nm_from = request.GET.get("nm_from")
        self.nm_to = request.GET.get("nm_to")
        self.order = request.GET.get("order")

        # Load the possible filter lists for each field
        self.brands = Brand.objects.all().order_by("name").values_list("name", flat=True).distinct()
        if(not self.brand in self.brands):
            self.brand = ""

        if(self.brand != "" or self.brand is None):
            self.models = Project.objects.filter(brand__name__iexact=self.brand).order_by("model").values_list("model", flat=True).distinct()
        
        # Check if selected model is even available
        # If not, delete the current selected model
        if(not self.model in self.models):
            self.model = ""

        self.fuels = Fuel.objects.all()
        if(self.brand != "" and not self.brand is None):
            self.fuels = self.fuels.filter(project__brand__name__iexact=self.brand)
        if(self.model != "" and not self.model is None):
            self.fuels = self.fuels.filter(project__model__iexact=self.model)
        self.fuels = self.fuels.order_by("name").values_list("name", flat=True).distinct()
        
        if(not self.fuel in self.fuels):
            self.fuel = ""

class ProjectsView(generic.ListView):
    model = Project
    template_name = 'projects.html'
    context_object_name = 'project_list'

    def get_context_data(self, **kwargs):
        context = super().get_context_data(**kwargs)

        filterData = FilterData()
        filterData.load_from_request(self.request)
        context["filterData"] = filterData

        # First we specifiy the objects
        r = Project.objects.all()

        # Then we add the filters
        if(filterData.query != "" and not filterData.query is None):
            r = r.filter(title__icontains=filterData.query)
        
        if(filterData.brand != "" and not filterData.brand is None):
            r = r.filter(brand__name__iexact=filterData.brand)
        
        if(filterData.model != "" and not filterData.model is None):
            r = r.filter(model__iexact=filterData.model)
        
        if(filterData.fuel != "" and not filterData.fuel is None):
            r = r.filter(fuel__name__iexact=filterData.fuel)

        if(filterData.hp_from != "" and not filterData.hp_from is None):
            r = r.filter(state__dynomeasurement__horse_power__gte=filterData.hp_from)
        
        if(filterData.hp_to != "" and not filterData.hp_to is None):
            r = r.filter(state__dynomeasurement__horse_power__lte=filterData.hp_to)

        if(filterData.nm_from != "" and not filterData.nm_from is None):
            r = r.filter(state__dynomeasurement__newtonmeter__gte=filterData.nm_from)
        
        if(filterData.nm_to != "" and not filterData.nm_to is None):
            r = r.filter(state__dynomeasurement__newtonmeter__lte=filterData.nm_to)

        # Sort projects
        if(filterData.order is None or filterData.order == "" or filterData.order == "mod_desc"):
            r = r.order_by('-state__modified_at')

        if(filterData.order == "mod_asc"):
            r = r.order_by('state__modified_at')

        context[self.context_object_name] = r
        return context

class ProjectDetailView(generic.DetailView):
    model = Project
    template_name = 'project_detail.html'

class FilterDataTrack(FilterData):
    track = ""
    tracks = []

    def load_from_request(self, request):
        super().load_from_request(request)

        self.track = request.GET.get("track")

        self.tracks = Track.objects.all().order_by("name").values_list("name", flat=True).distinct()
        
        # Check if selected model is even available
        # If not, delete the current selected model
        if(not self.track in self.tracks):
            self.track = ""

class TrackTimesView(generic.ListView):
    model = TrackMeasurement
    template_name = 'track_times.html'
    context_object_name = 'track_times_list'
    
    def get_context_data(self, **kwargs):
        context = super().get_context_data(**kwargs)

        filterData = FilterDataTrack()
        filterData.load_from_request(self.request)
        context["filterData"] = filterData

        # First we specifiy the objects
        r = TrackMeasurement.objects.all()

        # Then we add the filters
        if(filterData.query != "" and not filterData.query is None):
            r = r.filter(state__title__icontains=filterData.query)
        
        if(filterData.brand != "" and not filterData.brand is None):
            r = r.filter(state__project__brand__name__iexact=filterData.brand)
        
        if(filterData.model != "" and not filterData.model is None):
            r = r.filter(state__project__model__iexact=filterData.model)
        
        if(filterData.fuel != "" and not filterData.fuel is None):
            r = r.filter(state__project__fuel__name__iexact=filterData.fuel)

        if(filterData.hp_from != "" and not filterData.hp_from is None):
            r = r.filter(state__dynomeasurement__horse_power__gte=filterData.hp_from)
        
        if(filterData.hp_to != "" and not filterData.hp_to is None):
            r = r.filter(state__dynomeasurement__horse_power__lte=filterData.hp_to)

        if(filterData.nm_from != "" and not filterData.nm_from is None):
            r = r.filter(state__dynomeasurement__newtonmeter__gte=filterData.nm_from)
        
        if(filterData.nm_to != "" and not filterData.nm_to is None):
            r = r.filter(state__dynomeasurement__newtonmeter__lte=filterData.nm_to)

        if(filterData.track != "" and not filterData.track is None):
            r = r.filter(track__name__iexact=filterData.track)

        # Sort
        if(filterData.order is None or filterData.order == "" or filterData.order == "time_asc"):
            r = r.order_by('duration')
        else:
            r = r.order_by('-duration')
        
        context[self.context_object_name] = r
        return context

class DynoMeasurementsView(generic.ListView):
    model = DynoMeasurement
    template_name = 'dyno_measurements.html'
    context_object_name = 'dyno_measurements_list'
    
    def get_context_data(self, **kwargs):
        context = super().get_context_data(**kwargs)

        filterData = FilterData()
        filterData.load_from_request(self.request)
        context["filterData"] = filterData

        # First we specifiy the objects
        r = DynoMeasurement.objects.all()

        # Then we add the filters
        if(filterData.query != "" and not filterData.query is None):
            r = r.filter(state__title__icontains=filterData.query)
        
        if(filterData.brand != "" and not filterData.brand is None):
            r = r.filter(state__project__brand__name__iexact=filterData.brand)
        
        if(filterData.model != "" and not filterData.model is None):
            r = r.filter(state__project__model__iexact=filterData.model)
        
        if(filterData.fuel != "" and not filterData.fuel is None):
            r = r.filter(state__project__fuel__name__iexact=filterData.fuel)

        if(filterData.hp_from != "" and not filterData.hp_from is None):
            r = r.filter(state__dynomeasurement__horse_power__gte=filterData.hp_from)
        
        if(filterData.hp_to != "" and not filterData.hp_to is None):
            r = r.filter(state__dynomeasurement__horse_power__lte=filterData.hp_to)

        if(filterData.nm_from != "" and not filterData.nm_from is None):
            r = r.filter(state__dynomeasurement__newtonmeter__gte=filterData.nm_from)
        
        if(filterData.nm_to != "" and not filterData.nm_to is None):
            r = r.filter(state__dynomeasurement__newtonmeter__lte=filterData.nm_to)

        # Sort
        if(filterData.order is None or filterData.order == "" or filterData.order == "hp_desc"):
            r = r.order_by('-horse_power')
        elif(filterData.order == "hp_asc"):
            r = r.order_by('horse_power')
        elif(filterData.order == "nm_desc"):
            r = r.order_by('-newtonmeter')
        elif(filterData.order == "nm_asc"):
            r = r.order_by('newtonmeter')
        
        context[self.context_object_name] = r
        return context

class TimesView(generic.ListView):
    model = State
    template_name = 'times.html'
    context_object_name = 'states_list'

    def get_context_data(self, **kwargs):
        context = super().get_context_data(**kwargs)

        filterData = FilterData()
        filterData.load_from_request(self.request)
        context["filterData"] = filterData

        # First we specifiy the objects
        r = State.objects.all()

        # Then we add the filters
        if(filterData.query != "" and not filterData.query is None):
            r = r.filter(title__icontains=filterData.query)
        
        if(filterData.brand != "" and not filterData.brand is None):
            r = r.filter(project__brand__name__iexact=filterData.brand)
        
        if(filterData.model != "" and not filterData.model is None):
            r = r.filter(project__model__iexact=filterData.model)
        
        if(filterData.fuel != "" and not filterData.fuel is None):
            r = r.filter(project__fuel__name__iexact=filterData.fuel)

        if(filterData.hp_from != "" and not filterData.hp_from is None):
            r = r.filter(dynomeasurement__horse_power__gte=filterData.hp_from)
        
        if(filterData.hp_to != "" and not filterData.hp_to is None):
            r = r.filter(dynomeasurement__horse_power__lte=filterData.hp_to)

        if(filterData.nm_from != "" and not filterData.nm_from is None):
            r = r.filter(dynomeasurement__newtonmeter__gte=filterData.nm_from)
        
        if(filterData.nm_to != "" and not filterData.nm_to is None):
            r = r.filter(dynomeasurement__newtonmeter__lte=filterData.nm_to)

        # Sort
        # if(filterData.order is None or filterData.order == "" or filterData.order == "100200_desc"):
        #      r = r.order_by('timeperformancemeasurement__speed_start', 'timeperformancemeasurement__speed_end', '-timeperformancemeasurement__duration')
        # elif(filterData.order == "100200_asc"):
        #     r = r.order_by('timeperformancemeasurement__speed_start', 'timeperformancemeasurement__speed_end', 'timeperformancemeasurement__duration')
        # elif(filterData.order == "0200_desc"):
        # elif(filterData.order == "0200_asc"):
        # elif(filterData.order == "0100_desc"):
        # elif(filterData.order == "0100_asc"):
        
        # Boost performance a bit and select all related objects
        # So there are not too much database queries
        # r = r.select_related()
        
        context[self.context_object_name] = r
        return context

def about(request):
    return render(request, 'about.html')

def privacy(request):
    return render(request, 'datenschutz.html')