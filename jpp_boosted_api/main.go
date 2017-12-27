package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gorilla/mux"
)

// Manufacturer ist ein Autohersteller
type Manufacturer struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Model ist ein Automodell eines bestimmten Herstellers
// Jedes Projekt besitzt ein Model als Basis
type Model struct {
	ID                 int64   `json:"id"`
	ManufacturerID     int64   `json:"manufacturerId"`
	Name               string  `json:"name"`
	BuildStart         string  `json:"buildStart"`
	BuildEnd           string  `json:"buildEnd"`
	SeriesCode         string  `json:"seriesCode"`
	SeriesBuildStart   string  `json:"seriesBuildStart"`
	SeriesBuildEnd     string  `json:"seriesBuildEnd"`
	Type               string  `json:"type"`
	BuildSeries        string  `json:"buildSeries"`
	EngineType         string  `json:"engineType"`
	Fuel               string  `json:"fuel"`
	Fuel2              string  `json:"fuel2"`
	EngineConstruction string  `json:"engineConstruction"`
	Cylinder           int64   `json:"cylinder"`
	Capacity           float32 `json:"capacity"`
	KW                 float32 `json:"kw"`
	PS                 float32 `json:"ps"`
	Torque             float32 `json:"torque"`
	DriveType          string  `json:"driveType"`
	Gearbox            string  `json:"gearbox"`
	Tare               float32 `json:"tare"`
	WheelSize          string  `json:"wheelSize"`
	Acceleration       float32 `json:"Acceleration"`
	VMax               float32 `json:"vmax"`
	ImageURL           string  `json:"imageUrl"`
}

// Project ist ein Projekt in dem ein Auto modifiziert wird
type Project struct {
	ID                int64    `json:"id"`
	CarModelID        int64    `json:"carModelId"`
	Title             string   `json:"title"`
	ModificationStart string   `json:"modificationStart"`
	ModificationEnd   string   `json:"modificationEnd"`
	Rating            float32  `json:"rating"`
	Tunings           []Tuning `json:"tunings"`
	CarModel          Model    `json:"baseModel"`
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

func main() {
	// manufacturers = append(manufacturers, Manufacturer{ID: 1, Name: "Mercedes-Benz"})
	// manufacturers = append(manufacturers, Manufacturer{ID: 2, Name: "BMW"})
	// manufacturers = append(manufacturers, Manufacturer{ID: 3, Name: "Audi"})

	router := mux.NewRouter()
	router.HandleFunc("/carmanufacturers", GetCarManufacturers).Methods("GET")
	router.HandleFunc("/carmanufacturers", PostCarManufacturer).Methods("POST")
	router.HandleFunc("/carmanufacturers/{id}/models", GetCarManufacturerModels).Methods("GET")
	router.HandleFunc("/carmanufacturers/{id}", GetCarManufacturer).Methods("GET")
	router.HandleFunc("/carmodels", PostCarModel).Methods("POST")
	router.HandleFunc("/projects", GetProjects).Methods("GET")
	router.HandleFunc("/projects/{id}", GetProject).Methods("GET")
	router.HandleFunc("/projects", PostProject).Methods("POST")
	router.HandleFunc("/tunings", PostTuning).Methods("POST")
	router.HandleFunc("/parts", PostPart).Methods("POST")
	router.HandleFunc("/times", PostTime).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// GetDatabaseConnection gibt die aktuelle Datenbankverbindung zurück
// Muss aber immer noch mit defer db.Close() behandelt werden!
func GetDatabaseConnection() (*sql.DB, error) {
	return sql.Open("sqlite3", "./jpp_boosted.db")
}

// QueryManufacturers sucht übernimmt die Suche in der Datenbank
// Parameter werden ebenfalls behandelt
func QueryManufacturers(r *http.Request) (*sql.Rows, error) {
	db, err := GetDatabaseConnection()
	checkErr(err)
	defer db.Close()

	name := r.URL.Query().Get("name")
	if name != "" {
		name := "%" + name + "%"
		return db.Query("SELECT ID, Name, Url FROM CarManufacturer WHERE name like ? ORDER BY Name ASC", name)
	}

	return db.Query("SELECT ID, Name, Url FROM CarManufacturer ORDER BY name ASC")
}

// QueryManufacturersCarModels sucht übernimmt die Suche in der Datenbank
// Parameter werden ebenfalls behandelt
func QueryManufacturersCarModels(manufacturerID int, r *http.Request) (*sql.Rows, error) {
	db, err := GetDatabaseConnection()
	checkErr(err)
	defer db.Close()

	name := r.URL.Query().Get("name")
	if name != "" {
		name := "%" + name + "%"
		return db.Query("SELECT ID, ManufacturerID, Name, BuildStart, BuildEnd, SeriesCode, SeriesBuildStart, SeriesBuildEnd, BuildSeries, Cylinder, Capacity, KW, PS, Torque, DriveType, Gearbox, Tare, Acceleration, VMax FROM CarModel WHERE ManufacturerID = ? AND BuildSeries LIKE ? ORDER BY Name ASC", manufacturerID, name)
	}

	return db.Query("SELECT ID, ManufacturerID, Name, BuildStart, BuildEnd, SeriesCode, SeriesBuildStart, SeriesBuildEnd, BuildSeries, Cylinder, Capacity, KW, PS, Torque, DriveType, Gearbox, Tare, Acceleration, VMax FROM CarModel WHERE ManufacturerID = ? ORDER BY name ASC", manufacturerID)
}

// GetCarManufacturers gibt alle Hersteller zurück
func GetCarManufacturers(w http.ResponseWriter, r *http.Request) {
	rows, err := QueryManufacturers(r)
	checkErr(err)
	defer rows.Close()

	var manufacturers []Manufacturer
	for rows.Next() {
		manu := Manufacturer{}
		err = rows.Scan(&manu.ID, &manu.Name, &manu.URL)

		manufacturers = append(manufacturers, manu)
	}

	SetJSONContentHeader(w)
	json.NewEncoder(w).Encode(manufacturers)
}

// GetCarManufacturer gibt einen bestimmten Hersteller anhand der ID zurück
func GetCarManufacturer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	i, _ := strconv.Atoi(params["id"])

	db, err := GetDatabaseConnection()
	checkErr(err)
	defer db.Close()

	stmt, err := db.Prepare("SELECT ID, Name, Url FROM CarManufacturer WHERE ID = ?")
	checkErr(err)
	defer stmt.Close()

	var manu Manufacturer
	err = stmt.QueryRow(i).Scan(&manu.ID, &manu.Name, &manu.URL)
	if err == sql.ErrNoRows {
		// log.Fatal(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	SetJSONContentHeader(w)
	json.NewEncoder(w).Encode(manu)
}

// PostCarManufacturer erzeugt einen neuen Hersteller
func PostCarManufacturer(w http.ResponseWriter, r *http.Request) {
	var manu Manufacturer

	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)

	err = r.Body.Close()
	checkErr(err)

	if err := json.Unmarshal(body, &manu); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			checkErr(err)
		}
	}

	SetJSONContentHeader(w)
	db, err := GetDatabaseConnection()
	checkErr(err)
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO CarManufacturer(name, url) VALUES (?,?)")
	checkErr(err)
	defer stmt.Close()
	log.Println(manu)
	res, err := stmt.Exec(manu.Name, manu.URL)
	checkErr(err)

	id, err := res.LastInsertId()
	log.Println("Created with ID ", id)

	w.WriteHeader(http.StatusCreated)
}

// SetJSONContentHeader setzt in der Antwort der HTTP-Response, den Content-Type application/json charset=UTF-8
func SetJSONContentHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	enableCors(&w)
}

func enableCors(w *http.ResponseWriter) {
	// (*w).Header().Set("Access-Control-Allow-Origin", "*")
	// (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	// (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// GetCarManufacturerModels liest alle Modelle von einem Hersteller
func GetCarManufacturerModels(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	i, _ := strconv.Atoi(params["id"])

	rows, err := QueryManufacturersCarModels(i, r)
	checkErr(err)
	defer rows.Close()

	var carModels []Model
	for rows.Next() {
		mod := Model{}
		err = rows.Scan(&mod.ID, &mod.ManufacturerID, &mod.Name, &mod.BuildStart, &mod.BuildEnd, &mod.SeriesCode, &mod.SeriesBuildStart, &mod.SeriesBuildEnd, &mod.BuildSeries, &mod.Cylinder, &mod.Capacity, &mod.KW, &mod.PS, &mod.Torque, &mod.DriveType, &mod.Gearbox, &mod.Tare, &mod.Acceleration, &mod.VMax)

		carModels = append(carModels, mod)
	}

	SetJSONContentHeader(w)
	json.NewEncoder(w).Encode(carModels)
}

// PostCarModel erzeugt einen neuen Wert in der Datenbank
func PostCarModel(w http.ResponseWriter, r *http.Request) {
	var mod Model

	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)

	err = r.Body.Close()
	checkErr(err)

	if err := json.Unmarshal(body, &mod); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			checkErr(err)
		}
	}

	SetJSONContentHeader(w)
	db, err := GetDatabaseConnection()
	checkErr(err)
	defer db.Close()

	stmt, err := db.Prepare(`INSERT INTO CarModel(
		manufacturerid, 
		name, 
		buildstart, 
		buildend, 
		seriescode, 
		seriesbuildstart, 
		seriesbuildend,
		type,
		buildSeries,
		engineType,
		fuel,
		fuel2,
		engineConstruction,
		cylinder,
		capacity,
		kw,
		ps,
		torque,
		driveType,
		gearbox,
		tare,
		wheelsize,
		Acceleration,
		vmax,
		imageurl) 
		VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`)

	checkErr(err)
	defer stmt.Close()
	log.Println(mod)
	res, err := stmt.Exec(mod.ManufacturerID,
		mod.Name,
		mod.BuildStart,
		mod.BuildEnd,
		mod.SeriesCode,
		mod.SeriesBuildStart,
		mod.SeriesBuildEnd,
		mod.Type,
		mod.BuildSeries,
		mod.EngineType,
		mod.Fuel,
		mod.Fuel2,
		mod.EngineConstruction,
		mod.Cylinder,
		mod.Capacity,
		mod.KW,
		mod.PS,
		mod.Torque,
		mod.DriveType,
		mod.Gearbox,
		mod.Tare,
		mod.WheelSize,
		mod.Acceleration,
		mod.VMax,
		mod.ImageURL)
	checkErr(err)

	id, err := res.LastInsertId()
	log.Println("Created with ID ", id)

	w.WriteHeader(http.StatusCreated)
}

// GetProjects gibt alle vorhandenen Projekte zurück
// Es werden alle zugehörigen Objekte (Tunings, Automodel, etc.) mit zurückgegeben
func GetProjects(w http.ResponseWriter, r *http.Request) {
	db, err := GetDatabaseConnection()
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT ID, CarModelID, Title, ModificationStart, ModificationEnd FROM Project")
	checkErr(err)
	defer rows.Close()

	var projects []Project
	for rows.Next() {
		proj := Project{}
		err = rows.Scan(&proj.ID, &proj.CarModelID, &proj.Title, &proj.ModificationStart, &proj.ModificationEnd)

		projects = append(projects, proj)
	}

	SetJSONContentHeader(w)
	json.NewEncoder(w).Encode(projects)
}

// GetProject gibt ein vorhandenes Projekt zurück
// Es werden alle zugehörigen Objekte (Tunings, Automodel, etc.) mit zurückgegeben
func GetProject(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	i, _ := strconv.Atoi(params["id"])

	db, err := GetDatabaseConnection()
	checkErr(err)
	defer db.Close()

	stmt, err := db.Prepare("SELECT ID, CarModelID, Title, ModificationStart, ModificationEnd FROM Project WHERE ID = ?")
	checkErr(err)
	defer stmt.Close()

	var proj Project
	err = stmt.QueryRow(i).Scan(&proj.ID, &proj.CarModelID, &proj.Title, &proj.ModificationStart, &proj.ModificationEnd)
	if err == sql.ErrNoRows {
		// log.Fatal(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	SetJSONContentHeader(w)
	json.NewEncoder(w).Encode(proj)
}

// InsertProject erstellt einen neuen Datenbankeintrag
func InsertProject(proj *Project) (sql.Result, error) {
	db, err := GetDatabaseConnection()
	checkErr(err)
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO Project(carmodelid, title, modificationstart, modificationend, rating) VALUES (?,?,?,?,?)")
	checkErr(err)
	defer stmt.Close()
	log.Println(proj)
	return stmt.Exec(proj.CarModelID, proj.Title, proj.ModificationStart, proj.ModificationEnd, proj.Rating)
}

// PostProject erzeugt ein neues Projekt
// Verknüpfungen (Messungen, Teile, etc.) werden mit angelegt
func PostProject(w http.ResponseWriter, r *http.Request) {
	var proj Project

	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)

	err = r.Body.Close()
	checkErr(err)

	if err := json.Unmarshal(body, &proj); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			checkErr(err)
		}
	}

	log.Println(string(body))
	fmt.Printf("%+v\n", proj)

	SetJSONContentHeader(w)

	res, err := InsertProject(&proj)
	checkErr(err)

	id, err := res.LastInsertId()
	log.Println("Created with ID ", id)

	for _, tun := range proj.Tunings {
		log.Println(tun)
		tun.ProjectID = id
		tunRes, err := InsertTuning(&tun)
		checkErr(err)

		tunID, _ := tunRes.LastInsertId()
		for _, part := range tun.Parts {
			log.Println(part)
			part.TuningID = tunID
			_, err := InsertPart(&part)
			checkErr(err)
		}

		for _, time := range tun.Times {
			log.Println(time)
			time.TuningID = tunID
			_, err := InsertTime(&time)
			checkErr(err)
		}
	}

	w.WriteHeader(http.StatusCreated)
}

// InsertTuning erstellt einen neuen Datenbankeintrag
func InsertTuning(tun *Tuning) (sql.Result, error) {
	db, err := GetDatabaseConnection()
	checkErr(err)
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO Tuning(projectid, stage, description, horsepower, torque, date) VALUES (?,?,?,?,?,?)")
	checkErr(err)
	defer stmt.Close()
	log.Println(tun)
	return stmt.Exec(tun.ProjectID, tun.Stage, tun.Description, tun.HorsePower, tun.Torque, tun.Date)
}

// PostTuning erstellt ein neues Tuning
// Verknüpfungen (Messungen, Teile, etc.) werden mit angelegt
func PostTuning(w http.ResponseWriter, r *http.Request) {
	var tun Tuning

	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)

	err = r.Body.Close()
	checkErr(err)

	if err := json.Unmarshal(body, &tun); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			checkErr(err)
		}
	}

	SetJSONContentHeader(w)

	res, err := InsertTuning(&tun)
	checkErr(err)

	id, err := res.LastInsertId()
	log.Println("Created with ID ", id)

	for _, part := range tun.Parts {
		log.Println(part)
		part.TuningID = id
		_, err := InsertPart(&part)
		checkErr(err)
	}

	for _, time := range tun.Times {
		log.Println(time)
		time.TuningID = id
		_, err := InsertTime(&time)
		checkErr(err)
	}

	w.WriteHeader(http.StatusCreated)
}

// InsertPart erstellt einen neuen Datenbankeintrag
func InsertPart(part *TuningPart) (sql.Result, error) {
	db, err := GetDatabaseConnection()
	checkErr(err)
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO Part(tuningid, name, url, youtubeurl, manufacturer, manufacturerurl) VALUES (?,?,?,?,?,?)")
	checkErr(err)
	defer stmt.Close()
	log.Println(part)
	return stmt.Exec(part.TuningID, part.Name, part.URL, part.YouTubeURL, part.Manufacturer, part.ManufacturerURL)
}

// PostPart erstellt ein neues Teil
func PostPart(w http.ResponseWriter, r *http.Request) {
	var part TuningPart

	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)

	err = r.Body.Close()
	checkErr(err)

	if err := json.Unmarshal(body, &part); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			checkErr(err)
		}
	}

	SetJSONContentHeader(w)

	res, err := InsertPart(&part)
	checkErr(err)

	id, err := res.LastInsertId()
	log.Println("Created with ID ", id)

	w.WriteHeader(http.StatusCreated)
}

// InsertTime erstellet einen neuen Datenbankeintrag
func InsertTime(ti *TuningTime) (sql.Result, error) {
	db, err := GetDatabaseConnection()
	checkErr(err)
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO Time(tuningid, speedrange, time) VALUES (?,?,?)")
	checkErr(err)
	defer stmt.Close()
	log.Println(ti)
	return stmt.Exec(ti.TuningID, ti.SpeedRange, ti.Time)
}

// PostTime erstellt eine neue Messung
func PostTime(w http.ResponseWriter, r *http.Request) {
	var ti TuningTime

	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)

	err = r.Body.Close()
	checkErr(err)

	if err := json.Unmarshal(body, &ti); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			checkErr(err)
		}
	}

	SetJSONContentHeader(w)

	res, err := InsertTime(&ti)
	checkErr(err)

	id, err := res.LastInsertId()
	log.Println("Created with ID ", id)

	w.WriteHeader(http.StatusCreated)
}

// Prüft ob es einen Fehler gab, wenn ja wird mit log.Fatal abgebrochen
func checkErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
