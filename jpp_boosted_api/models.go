package main

// Manufacturer ist ein Autohersteller
type Manufacturer struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Model ist ein Automodell eines bestimmten Herstellers
// Jedes Projekt besitzt ein Model als Basis
type Model struct {
	ID                 int64        `json:"id"`
	ManufacturerID     int64        `json:"manufacturerId"`
	Name               string       `json:"name"`
	BuildStart         string       `json:"buildStart"`
	BuildEnd           string       `json:"buildEnd"`
	SeriesCode         string       `json:"seriesCode"`
	SeriesBuildStart   string       `json:"seriesBuildStart"`
	SeriesBuildEnd     string       `json:"seriesBuildEnd"`
	Type               string       `json:"type"`
	BuildSeries        string       `json:"buildSeries"`
	EngineType         string       `json:"engineType"`
	Fuel               string       `json:"fuel"`
	Fuel2              string       `json:"fuel2"`
	EngineConstruction string       `json:"engineConstruction"`
	Cylinder           int64        `json:"cylinder"`
	Capacity           float32      `json:"capacity"`
	KW                 float32      `json:"kw"`
	PS                 float32      `json:"ps"`
	Torque             float32      `json:"torque"`
	DriveType          string       `json:"driveType"`
	Gearbox            string       `json:"gearbox"`
	Tare               float32      `json:"tare"`
	WheelSize          string       `json:"wheelSize"`
	Acceleration       float32      `json:"acceleration"`
	VMax               float32      `json:"vmax"`
	ImageURL           string       `json:"imageUrl"`
	Manufacturer       Manufacturer `json:"manufacturer"`
}

// Project ist ein Projekt in dem ein Auto modifiziert wird
type Project struct {
	ID           int64    `json:"id"`
	CarModelID   int64    `json:"carModelId"`
	Title        string   `json:"title"`
	Rating       float32  `json:"rating"`
	BuildYear    int64    `json:"buildYear"`
	ProjImageURL string   `json:"projImageUrl"`
	Tunings      []Tuning `json:"tunings"`
	CarModel     Model    `json:"baseModel"`
}

// Projects stellt ein Slice für Projekte dar
type Projects struct {
	Projects []Project
}

// Tuning ist eine Stufe in einem Projekt
// z.B. Stage 1
type Tuning struct {
	ID          int64        `json:"id"`
	ProjectID   int64        `json:"projectId"`
	Stage       string       `json:"stage"`
	Description string       `json:"description"`
	HorsePower  float32      `json:"horsePower"`
	Torque      float32      `json:"torque"`
	Date        string       `json:"date"`
	YoutubeURL  string       `json:"youtubeUrl"`
	Times       []TuningTime `json:"times"`
	Parts       []TuningPart `json:"parts"`
}

// TuningTime ist eine Messung in einer bestimmten Stufe
type TuningTime struct {
	ID         int64   `json:"id"`
	TuningID   int64   `json:"tuningId"`
	SpeedRange string  `json:"speedRange"`
	Time       float32 `json:"time"`
}

// TuningPart ist ein Teil, welches in einer bestimmten Stufe verbaut wurde
type TuningPart struct {
	ID              int64  `json:"id"`
	TuningID        int64  `json:"tuningId"`
	Name            string `json:"name"`
	URL             string `json:"url"`
	YouTubeURL      string `json:"youtubeUrl"`
	Manufacturer    string `json:"manufacturer"`
	ManufacturerURL string `json:"manufacturerUrl"`
}

func (tuning *Tuning) AddTime(time TuningTime) []TuningTime {
	tuning.Times = append(tuning.Times, time)
	return tuning.Times
}

func (tuning *Tuning) AddPart(part TuningPart) []TuningPart {
	tuning.Parts = append(tuning.Parts, part)
	return tuning.Parts
}

func (project *Project) AddTuning(tuning Tuning) []Tuning {
	project.Tunings = append(project.Tunings, tuning)
	return project.Tunings
}

func (projects *Projects) AddProject(proj Project) []Project {
	projects.Projects = append(projects.Projects, proj)
	return projects.Projects
}

func (projects *Projects) GetProjectFromSlice(id int64) *Project {
	for i := range projects.Projects {
		if projects.Projects[i].ID == id {
			return &projects.Projects[i]
		}
	}

	return nil
}

func (project *Project) GetTuning(id int64) *Tuning {
	for i := range project.Tunings {
		if project.Tunings[i].ID == id {
			return &project.Tunings[i]
		}
	}

	return nil
}

func (tuning *Tuning) GetTime(id int64) *TuningTime {
	for i := range tuning.Times {
		if tuning.Times[i].ID == id {
			return &tuning.Times[i]
		}
	}

	return nil
}

func (tuning *Tuning) GetPart(id int64) *TuningPart {
	for i := range tuning.Parts {
		if tuning.Parts[i].ID == id {
			return &tuning.Parts[i]
		}
	}

	return nil
}
