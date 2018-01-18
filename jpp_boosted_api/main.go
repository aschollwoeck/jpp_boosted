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

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/carmanufacturers", GetCarManufacturers).Methods("GET")
	router.HandleFunc("/api/carmanufacturers", PostCarManufacturer).Methods("POST")
	router.HandleFunc("/api/carmanufacturers/{id}/models", GetCarManufacturerModels).Methods("GET")
	router.HandleFunc("/api/carmanufacturers/{id}", GetCarManufacturer).Methods("GET")
	router.HandleFunc("/api/carmodels/{id}", GetCarModel).Methods("GET")
	router.HandleFunc("/api/carmodels", PostCarModel).Methods("POST")
	router.HandleFunc("/api/projects/include", GetProjectsInclude).Methods("GET")
	router.HandleFunc("/api/projects/{id}/include", GetProjectInclude).Methods("GET")
	router.HandleFunc("/api/projects", GetProjects).Methods("GET")
	router.HandleFunc("/api/projects/{id}", GetProject).Methods("GET")
	router.HandleFunc("/api/projects", PostProject).Methods("POST")
	router.HandleFunc("/api/tunings", PostTuning).Methods("POST")
	router.HandleFunc("/api/parts", PostPart).Methods("POST")
	router.HandleFunc("/api/times", PostTime).Methods("POST")
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

// GetCarModel liest ein Modell
func GetCarModel(w http.ResponseWriter, r *http.Request) {
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

// QueryProjects übernimmt die Suche in der Datenbank
// Parameter werden ebenfalls behandelt
func QueryProjects(r *http.Request) (*sql.Rows, error) {
	db, err := GetDatabaseConnection()
	checkErr(err)
	defer db.Close()

	limit := r.URL.Query().Get("limit")
	if limit != "" {
		limit = " LIMIT " + limit
	}

	title := r.URL.Query().Get("title")
	if title != "" {
		title := "%" + title + "%"
		return db.Query("SELECT DISTINCT ID, CarModelID, Title, BuildYear, ImageUrl FROM Project WHERE title like ? ORDER BY Title ASC"+limit, title)
	}

	return db.Query("SELECT ID, CarModelID, Title, BuildYear, ImageUrl FROM Project ORDER BY Title ASC" + limit)
}

// QueryProjectsInclude übernimmt die Suche in der Datenbank
// Parameter werden ebenfalls behandelt
func QueryProjectsInclude(r *http.Request) (*sql.Rows, error) {
	db, err := GetDatabaseConnection()
	checkErr(err)
	defer db.Close()

	search := r.URL.Query().Get("search")
	if search != "" {
		search := "%" + search + "%"
		return db.Query(`
				select 
				p.id as projectid, p.carmodelid, p.title, p.buildYear, p.imageUrl as projimgurl, 
				c.id as carmodelid, c.buildSeries, c.imageUrl, c.name, c.seriescode, c.buildstart, c.buildend, c.seriesbuildstart, c.seriesbuildend, c.type, c.Cylinder, c.KW, c.PS, c.Torque, c.Tare, c.Wheelsize, c.Acceleration, c.VMax, 
				m.name, m.url,
				t.id as tuningid, t.stage, t.description, t.horsepower, t.torque, t.date, t.youtubeurl, 
				coalesce(ti.id, 0) as timeid, coalesce(ti.speedRange, "") as speedRange, coalesce(ti.time, 0) as time, 
				coalesce(part.id, 0) as partid, coalesce(part.Name, "") as Name, coalesce(part.Url, "") as Url, coalesce(part.Manufacturer, "") as Manufacturer, coalesce(part.ManufacturerUrl, "") as ManufacturerUrl
				from Project p
				left join CarModel c on c.id = p.carmodelid
				left join CarManufacturer m on m.id = c.manufacturerid
				left join Tuning t on t.projectid = p.id
				left join Time ti on t.id = ti.TuningID
				left join Part part on t.id = part.TuningID
				where p.title like ? or m.name like ? or c.buildSeries like ? 
				order by
				p.id, c.id, t.id, ti.id, part.id
				`, search, search, search)
	}

	params := mux.Vars(r)
	i, _ := strconv.Atoi(params["id"])
	if err == nil && i > 0 {
		return db.Query(`
				select 
				p.id as projectid, p.carmodelid, p.title, p.buildYear, p.imageUrl as projimgurl, 
				c.id as carmodelid, c.buildSeries, c.imageUrl, c.name, c.seriescode, c.buildstart, c.buildend, c.seriesbuildstart, c.seriesbuildend, c.type, c.Cylinder, c.KW, c.PS, c.Torque, c.Tare, c.Wheelsize, c.Acceleration, c.VMax, 
				m.name, m.url,
				t.id as tuningid, t.stage, t.description, t.horsepower, t.torque, t.date, t.youtubeurl, 
				coalesce(ti.id, 0) as timeid, coalesce(ti.speedRange, "") as speedRange, coalesce(ti.time, 0) as time, 
				coalesce(part.id, 0) as partid, coalesce(part.Name, "") as Name, coalesce(part.Url, "") as Url, coalesce(part.Manufacturer, "") as Manufacturer, coalesce(part.ManufacturerUrl, "") as ManufacturerUrl
				from Project p
				left join CarModel c on c.id = p.carmodelid
				left join CarManufacturer m on m.id = c.manufacturerid
				left join Tuning t on t.projectid = p.id
				left join Time ti on t.id = ti.TuningID
				left join Part part on t.id = part.TuningID
				where p.id = ?
				order by
				p.id, c.id, t.id, ti.id, part.id
				`, i)
	}

	log.Println("Standard Query")
	return db.Query(`
		select 
		p.id as projectid, p.carmodelid, p.title, p.buildYear, p.imageUrl as projimgurl, 
		c.id as carmodelid, c.buildSeries, c.imageUrl, c.name, c.seriescode, c.buildstart, c.buildend, c.seriesbuildstart, c.seriesbuildend, c.type, c.Cylinder, c.KW, c.PS, c.Torque, c.Tare, c.Wheelsize, c.Acceleration, c.VMax, 
		m.name, m.url,
		t.id as tuningid, t.stage, t.description, t.horsepower, t.torque, t.date, t.youtubeurl, 
		coalesce(ti.id, 0) as timeid, coalesce(ti.speedRange, "") as speedRange, coalesce(ti.time, 0) as time, 
		coalesce(part.id, 0) as partid, coalesce(part.Name, "") as Name, coalesce(part.Url, "") as Url, coalesce(part.Manufacturer, "") as Manufacturer, coalesce(part.ManufacturerUrl, "") as ManufacturerUrl
		from Project p
		left join CarModel c on c.id = p.carmodelid
		left join CarManufacturer m on m.id = c.manufacturerid
		left join Tuning t on t.projectid = p.id
		left join Time ti on t.id = ti.TuningID
		left join Part part on t.id = part.TuningID
		order by
		p.id, c.id, t.id, ti.id, part.id
		`)
}

// GetProjects gibt alle vorhandenen Projekte zurück
// Es werden alle zugehörigen Objekte (Tunings, Automodel, etc.) mit zurückgegeben
func GetProjects(w http.ResponseWriter, r *http.Request) {
	rows, err := QueryProjects(r)
	checkErr(err)
	defer rows.Close()

	var projects []Project
	for rows.Next() {
		proj := Project{}
		err = rows.Scan(&proj.ID, &proj.CarModelID, &proj.Title, &proj.BuildYear, &proj.ProjImageURL)

		projects = append(projects, proj)
	}

	SetJSONContentHeader(w)
	json.NewEncoder(w).Encode(projects)
}

// GetProjectsInclude gibt alle vorhandenen Projekte zurück
// Es werden alle zugehörigen Objekte (Tunings, Automodel, etc.) mit zurückgegeben
// func GetProjectsInclude(w http.ResponseWriter, r *http.Request) {
// 	rows, err := QueryProjectsInclude(r)
// 	checkErr(err)
// 	defer rows.Close()

// 	var projects Projects
// 	for rows.Next() {
// 		proj := Project{}
// 		tuning := Tuning{}
// 		tuningPart := TuningPart{}
// 		tuningTime := TuningTime{}

// 		err = rows.Scan(&proj.ID, &proj.CarModelID, &proj.Title, &proj.CarModel.ID, &proj.CarModel.BuildSeries, &proj.CarModel.ImageURL, &proj.CarModel.Manufacturer.Name, &proj.CarModel.Manufacturer.URL, &tuning.ID, &tuning.Stage, &tuning.Description, &tuning.HorsePower, &tuning.Torque, &tuning.Date, &tuning.YoutubeURL, &tuningTime.ID, &tuningTime.SpeedRange, &tuningTime.Time, &tuningPart.ID, &tuningPart.Name, &tuningPart.URL, &tuningPart.Manufacturer, &tuningPart.ManufacturerURL)

// 		if tuningPart.ID > 0 {
// 			tuning.Parts = append(tuning.Parts, tuningPart)
// 		}
// 		if tuningTime.ID > 0 {
// 			tuning.Times = append(tuning.Times, tuningTime)
// 		}
// 		if tuning.ID > 0 {
// 			proj.Tunings = append(proj.Tunings, tuning)
// 		}

// 		// Wenn Projekt noch nicht existiert, dann nehmen
// 		// wir es im Array auf und gehen zur nächsten Zeile
// 		tempProj := projects.GetProjectFromSlice(proj.ID)
// 		if tempProj == nil {
// 			projects.AddProject(proj)
// 			continue
// 		}

// 		// Wenn Tuning noch nicht existiert, dann
// 		// fügen wir es dem Projekt hinzu
// 		tempTuning := tempProj.GetTuning(tuning.ID)
// 		if tempTuning == nil {
// 			tempProj.AddTuning(tuning)
// 			continue
// 		}

// 		// Tuning existiert ebenfalls schon, also machen wir mit
// 		// den Parts weiter
// 		tempPart := tempTuning.GetPart(tuningPart.ID)
// 		if tempPart == nil {
// 			tempTuning.AddPart(tuningPart)
// 		}

// 		// Und jetzt noch für Times
// 		tempTime := tempTuning.GetTime(tuningTime.ID)
// 		if tempTime == nil {
// 			log.Println(tempProj.Tunings[0].Times)
// 			log.Println(tempTuning.Times)
// 			tempTuning.AddTime(tuningTime)
// 			log.Println(tempProj.Tunings[0].Times)
// 			log.Println(tempTuning.Times)
// 		}
// 	}

// 	SetJSONContentHeader(w)
// 	json.NewEncoder(w).Encode(projects.Projects)
// }

func GetProjectsInclude(w http.ResponseWriter, r *http.Request) {
	rows, err := QueryProjectsInclude(r)
	checkErr(err)
	defer rows.Close()

	var projects Projects
	for rows.Next() {
		proj := Project{}
		tuning := Tuning{}
		tuningPart := TuningPart{}
		tuningTime := TuningTime{}
		err = rows.Scan(&proj.ID, &proj.CarModelID, &proj.Title, &proj.BuildYear, &proj.ProjImageURL,
			&proj.CarModel.ID, &proj.CarModel.BuildSeries, &proj.CarModel.ImageURL, &proj.CarModel.Name, &proj.CarModel.SeriesCode, &proj.CarModel.BuildStart, &proj.CarModel.BuildEnd, &proj.CarModel.SeriesBuildStart, &proj.CarModel.SeriesBuildEnd, &proj.CarModel.Type, &proj.CarModel.Cylinder, &proj.CarModel.KW, &proj.CarModel.PS, &proj.CarModel.Torque, &proj.CarModel.Tare, &proj.CarModel.WheelSize, &proj.CarModel.Acceleration, &proj.CarModel.VMax,
			&proj.CarModel.Manufacturer.Name, &proj.CarModel.Manufacturer.URL,
			&tuning.ID, &tuning.Stage, &tuning.Description, &tuning.HorsePower, &tuning.Torque, &tuning.Date, &tuning.YoutubeURL,
			&tuningTime.ID, &tuningTime.SpeedRange, &tuningTime.Time,
			&tuningPart.ID, &tuningPart.Name, &tuningPart.URL, &tuningPart.Manufacturer, &tuningPart.ManufacturerURL)

		log.Println("TuningPart")
		log.Println(tuningPart)

		if tuningPart.ID > 0 {
			tuning.Parts = append(tuning.Parts, tuningPart)
		}
		if tuningTime.ID > 0 {
			tuning.Times = append(tuning.Times, tuningTime)
		}
		if tuning.ID > 0 {
			proj.Tunings = append(proj.Tunings, tuning)
		}

		// Wenn Projekt noch nicht existiert, dann nehmen
		// wir es im Array auf und gehen zur nächsten Zeile
		tempProjIndex := -1
		for i, p := range projects.Projects {
			if p.ID == proj.ID {
				tempProjIndex = i
				break
			}
		}

		if tempProjIndex == -1 {
			projects.AddProject(proj)
			continue
		}

		// Wenn Tuning noch nicht existiert, dann
		// fügen wir es dem Projekt hinzu
		tempTuningIndex := -1
		for i, t := range projects.Projects[tempProjIndex].Tunings {
			if t.ID == tuning.ID {
				tempTuningIndex = i
				break
			}
		}
		if tempTuningIndex == -1 {
			projects.Projects[tempProjIndex].AddTuning(tuning)
			continue
		}

		// Tuning existiert ebenfalls schon, also machen wir mit
		// den Parts weiter
		if tuningPart.ID > 0 {
			tempPart := projects.Projects[tempProjIndex].Tunings[tempTuningIndex].GetPart(tuningPart.ID)
			if tempPart == nil {
				projects.Projects[tempProjIndex].Tunings[tempTuningIndex].AddPart(tuningPart)
			}
		}

		if tuningTime.ID > 0 {
			// Und jetzt noch für Times
			tempTime := projects.Projects[tempProjIndex].Tunings[tempTuningIndex].GetTime(tuningTime.ID)
			if tempTime == nil {
				projects.Projects[tempProjIndex].Tunings[tempTuningIndex].AddTime(tuningTime)
			}
		}
	}

	SetJSONContentHeader(w)
	json.NewEncoder(w).Encode(projects.Projects)
}

// GetProjectInclude gibt genau 1 Projekt zurück
func GetProjectInclude(w http.ResponseWriter, r *http.Request) {
	rows, err := QueryProjectsInclude(r)
	checkErr(err)
	defer rows.Close()

	var projects Projects
	for rows.Next() {
		proj := Project{}
		tuning := Tuning{}
		tuningPart := TuningPart{}
		tuningTime := TuningTime{}
		err = rows.Scan(&proj.ID, &proj.CarModelID, &proj.Title, &proj.BuildYear, &proj.ProjImageURL,
			&proj.CarModel.ID, &proj.CarModel.BuildSeries, &proj.CarModel.ImageURL, &proj.CarModel.Name, &proj.CarModel.SeriesCode, &proj.CarModel.BuildStart, &proj.CarModel.BuildEnd, &proj.CarModel.SeriesBuildStart, &proj.CarModel.SeriesBuildEnd, &proj.CarModel.Type, &proj.CarModel.Cylinder, &proj.CarModel.KW, &proj.CarModel.PS, &proj.CarModel.Torque, &proj.CarModel.Tare, &proj.CarModel.WheelSize, &proj.CarModel.Acceleration, &proj.CarModel.VMax,
			&proj.CarModel.Manufacturer.Name, &proj.CarModel.Manufacturer.URL,
			&tuning.ID, &tuning.Stage, &tuning.Description, &tuning.HorsePower, &tuning.Torque, &tuning.Date, &tuning.YoutubeURL,
			&tuningTime.ID, &tuningTime.SpeedRange, &tuningTime.Time,
			&tuningPart.ID, &tuningPart.Name, &tuningPart.URL, &tuningPart.Manufacturer, &tuningPart.ManufacturerURL)

		if tuningPart.ID > 0 {
			tuning.Parts = append(tuning.Parts, tuningPart)
		}
		if tuningTime.ID > 0 {
			tuning.Times = append(tuning.Times, tuningTime)
		}
		if tuning.ID > 0 {
			proj.Tunings = append(proj.Tunings, tuning)
		}

		// Wenn Projekt noch nicht existiert, dann nehmen
		// wir es im Array auf und gehen zur nächsten Zeile
		tempProjIndex := -1
		for i, p := range projects.Projects {
			if p.ID == proj.ID {
				tempProjIndex = i
				break
			}
		}

		if tempProjIndex == -1 {
			projects.AddProject(proj)
			continue
		}

		// Wenn Tuning noch nicht existiert, dann
		// fügen wir es dem Projekt hinzu
		tempTuningIndex := -1
		for i, t := range projects.Projects[tempProjIndex].Tunings {
			if t.ID == tuning.ID {
				tempTuningIndex = i
				break
			}
		}
		if tempTuningIndex == -1 {
			projects.Projects[tempProjIndex].AddTuning(tuning)
			continue
		}

		// Tuning existiert ebenfalls schon, also machen wir mit
		// den Parts weiter
		tempPart := projects.Projects[tempProjIndex].Tunings[tempTuningIndex].GetPart(tuningPart.ID)
		if tempPart == nil {
			projects.Projects[tempProjIndex].Tunings[tempTuningIndex].AddPart(tuningPart)
		}

		// Und jetzt noch für Times
		tempTime := projects.Projects[tempProjIndex].Tunings[tempTuningIndex].GetTime(tuningTime.ID)
		if tempTime == nil {
			projects.Projects[tempProjIndex].Tunings[tempTuningIndex].AddTime(tuningTime)
		}
	}

	SetJSONContentHeader(w)

	if len(projects.Projects) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(projects.Projects[0])
}

// GetProject gibt ein vorhandenes Projekt zurück
// Es werden alle zugehörigen Objekte (Tunings, Automodel, etc.) mit zurückgegeben
func GetProject(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	i, _ := strconv.Atoi(params["id"])

	db, err := GetDatabaseConnection()
	checkErr(err)
	defer db.Close()

	stmt, err := db.Prepare("SELECT ID, CarModelID, Title, BuildYear, ImageUrl FROM Project WHERE ID = ?")
	checkErr(err)
	defer stmt.Close()

	var proj Project
	err = stmt.QueryRow(i).Scan(&proj.ID, &proj.CarModelID, &proj.Title, &proj.BuildYear, &proj.ProjImageURL)
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

	stmt, err := db.Prepare("INSERT INTO Project(carmodelid, title, rating, buildyear, imageUrl) VALUES (?,?,?,?,?)")
	checkErr(err)
	defer stmt.Close()
	log.Println(proj)
	return stmt.Exec(proj.CarModelID, proj.Title, proj.Rating, proj.BuildYear, proj.ProjImageURL)
}

// UpdateProject aktualisiert einen Datenbankeintrage
// ID muss im Objekt gepflegt sein
func UpdateProject(proj *Project) (sql.Result, error) {
	db, err := GetDatabaseConnection()
	checkErr(err)
	defer db.Close()

	stmt, err := db.Prepare("UPDATE Project SET carmodelid=?, title=?, rating=?, buildyear=?, imageUrl=? WHERE ID=?")
	checkErr(err)
	defer stmt.Close()
	log.Println(proj)
	return stmt.Exec(proj.CarModelID, proj.Title, proj.Rating, proj.BuildYear, proj.ProjImageURL, proj.ID)
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

	fmt.Printf("%+v\n", proj)

	SetJSONContentHeader(w)

	var id int64
	if proj.ID > 0 {
		id = proj.ID
		_, err := UpdateProject(&proj)
		checkErr(err)
	} else {
		res, err := InsertProject(&proj)
		checkErr(err)

		id, err = res.LastInsertId()
		log.Println("Created with ID ", id)
	}

	for _, tun := range proj.Tunings {
		log.Println(tun)
		tun.ProjectID = id

		var tunID int64
		if tun.ID > 0 {
			tunID = tun.ID
			_, err := UpdateTuning(&tun)
			checkErr(err)
		} else {
			tunRes, err := InsertTuning(&tun)
			checkErr(err)
			tunID, _ = tunRes.LastInsertId()
		}

		for _, part := range tun.Parts {
			// Wir fügen nur Parts ein, die nicht leer sind
			if part.Name == "" && part.URL == "" && part.Manufacturer == "" && part.ManufacturerURL == "" {
				continue
			}
			part.TuningID = tunID
			if part.ID > 0 {
				_, err := UpdatePart(&part)
				checkErr(err)
			} else {
				_, err := InsertPart(&part)
				checkErr(err)
			}
		}

		for _, time := range tun.Times {
			// Wir fügen nur Zeiten ein, die nicht leer sind
			if time.SpeedRange == "" {
				continue
			}
			log.Println(time)
			time.TuningID = tunID
			if time.ID > 0 {
				_, err := UpdateTime(&time)
				checkErr(err)
			} else {
				_, err := InsertTime(&time)
				checkErr(err)
			}
		}
	}

	w.WriteHeader(http.StatusCreated)
}

// InsertTuning erstellt einen neuen Datenbankeintrag
func InsertTuning(tun *Tuning) (sql.Result, error) {
	db, err := GetDatabaseConnection()
	checkErr(err)
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO Tuning(projectid, stage, description, horsepower, torque, date, youtubeurl) VALUES (?,?,?,?,?,?,?)")
	checkErr(err)
	defer stmt.Close()
	log.Println(tun)
	return stmt.Exec(tun.ProjectID, tun.Stage, tun.Description, tun.HorsePower, tun.Torque, tun.Date, tun.YoutubeURL)
}

// UpdateTuning aktualisiert einen Datenbankeintrag
// ID muss im Objekt gepflegt sein
func UpdateTuning(tun *Tuning) (sql.Result, error) {
	db, err := GetDatabaseConnection()
	checkErr(err)
	defer db.Close()

	stmt, err := db.Prepare("UPDATE Tuning SET projectid=?, stage=?, description=?, horsepower=?, torque=?, date=?, youtubeurl=? WHERE ID=?")
	checkErr(err)
	defer stmt.Close()
	log.Println(tun)
	return stmt.Exec(tun.ProjectID, tun.Stage, tun.Description, tun.HorsePower, tun.Torque, tun.Date, tun.YoutubeURL, tun.ID)
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

// UpdatePart aktualisiert einen Datenbankeintrag
// Objekt muss ID enthalten
func UpdatePart(part *TuningPart) (sql.Result, error) {
	db, err := GetDatabaseConnection()
	checkErr(err)
	defer db.Close()

	stmt, err := db.Prepare("UPDATE Part SET tuningid=?, name=?, url=?, youtubeurl=?, manufacturer=?, manufacturerurl=? WHERE ID=?")
	checkErr(err)
	defer stmt.Close()
	log.Println(part)
	return stmt.Exec(part.TuningID, part.Name, part.URL, part.YouTubeURL, part.Manufacturer, part.ManufacturerURL, part.ID)
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

// UpdateTime aktualisiert einen Datenbankeintrag
// Objekt muss ID enthalten
func UpdateTime(ti *TuningTime) (sql.Result, error) {
	db, err := GetDatabaseConnection()
	checkErr(err)
	defer db.Close()

	stmt, err := db.Prepare("UPDATE Time SET tuningid=?, speedrange=?, time=? WHERE ID=?")
	checkErr(err)
	defer stmt.Close()
	log.Println(ti)
	return stmt.Exec(ti.TuningID, ti.SpeedRange, ti.Time, ti.ID)
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
