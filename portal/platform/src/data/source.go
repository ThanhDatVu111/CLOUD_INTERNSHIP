package platform

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

// Structures for the JSON data
// Users ----------------------------------------------------------------
type User struct {
	ID                  int     `json:"id"`
	Name                string  `json:"name"`
	Password            string  `json:"password"`
	Mail                string  `json:"mail"`
	Phone               string  `json:"phone"`
	PFP                 string  `json:"pfp"`
	Title               string  `json:"title"`
	Department          string  `json:"department"`
	Status              string  `json:"status"`
	Calendar            []Event `json:"calendar"`
	OwnedProjects       []int   `json:"ownedProjects"`
	ContributedProjects []int   `json:"contributedProjects"`
	OwnedReports        []int   `json:"ownedReports"`
	ContributedReports  []int   `json:"contributedReports"`
}

// Event represents a calendar event
type Event struct {
	Date        string `json:"date"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

// Projects -------------------------------------------------------------
type Project struct {
	ID             int    `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Date           string `json:"date"`
	LastModified   string `json:"lastModified"`
	LastModifiedBy int    `json:"lastModifiedBy"`
	Thumbnail      string `json:"thumbnail"`
	Owner          []int  `json:"owner"`
	Contributors   []int  `json:"contributors"`
	Labs           []Lab  `json:"labs"`
}

// Lab represents a lab within a project
type Lab struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	Thumbnail    string `json:"thumbnail"`
	Date         string `json:"date"`
	Function     string `json:"function"`
	Locked       bool   `json:"locked"`
	Contributors []int  `json:"contributors"`
	Runs         []Run  `json:"runs"`
}

// Run represents a run of a lab experiment
type Run struct {
	Author  int    `json:"author"`
	Date    string `json:"date"`
	Status  string `json:"status"`
	Results string `json:"results"`
}

// Reports --------------------------------------------------------------
type Report struct {
	ID             int    `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Date           string `json:"date"`
	LastModified   string `json:"lastModified"`
	LastModifiedBy int    `json:"lastModifiedBy"`
	Thumbnail      string `json:"thumbnail"`
	Owner          []int  `json:"owner"`
	Contributors   []int  `json:"contributors"`
	Logs           []Logs `json:"logs"`
}

// Logs represents a log of a report
type Logs struct {
	Title             string       `json:"title"`
	Description       string       `json:"description"`
	Comment           string       `json:"comment"`
	Date              string       `json:"date"`
	LastModified      string       `json:"lastModified"`
	LastModifiedBy    int          `json:"lastModifiedBy"`
	Locked            bool         `json:"locked"`
	Contributors      []int        `json:"contributors"`
	Tags              []string     `json:"tags"`
	Files             []string     `json:"files"`
	LinkedProjectLabs []ProjectLab `json:"linkedProjectLabs"`
}

// Simple object for lab linking
type ProjectLab struct {
	ProjectID int `json:"projectID"`
	LabNumber int `json:"labNumber"`
}

// Routes ----------------------------------------------------------------
func Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/projects", func(r chi.Router) {
		r.Get("/", GetProjectsHandler)
	})

	r.Route("/reports", func(r chi.Router) {
		r.Get("/", GetReportsHandler)
	})

	r.Route("/users", func(r chi.Router) {
		r.Get("/", GetUsersHandler)
	})

	return r
}

func GetProjectsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("INTO THE PROJECTS")

	jsonFile, err := os.Open("platform/src/data/sampleProjects.json")
	if err != nil {
		http.Error(w, "Error opening JSON file", http.StatusInternalServerError)
		return
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		http.Error(w, "COULD NOT READ JSON FILE", http.StatusInternalServerError)
		return
	}

	var projects []Project
	err = json.Unmarshal(jsonData, &projects)
	if err != nil {
		http.Error(w, "COULD NOT UNAMRSHALL JSON DATA", http.StatusInternalServerError)
		return
	}

	resJSON, err := json.Marshal(projects)
	if err != nil {
		http.Error(w, "COULD NOT UNAMRSHALL JSON DATA", http.StatusInternalServerError)
		return
	}

	//log.Println(string(resJSON))
	w.Header().Set("Content-Type", "application/json")
	w.Write(resJSON)
}

func GetReportsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("INTO THE REPORTS")

	jsonFile, err := os.Open("platform/src/data/sampleReports.json")
	if err != nil {
		http.Error(w, "Error opening JSON file", http.StatusInternalServerError)
		return
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		http.Error(w, "COULD NOT READ JSON FILE", http.StatusInternalServerError)
		return
	}

	var reports []Report
	err = json.Unmarshal(jsonData, &reports)
	if err != nil {
		http.Error(w, "COULD NOT UNAMRSHALL JSON DATA", http.StatusInternalServerError)
		return
	}

	resJSON, err := json.Marshal(reports)
	if err != nil {
		http.Error(w, "COULD NOT UNAMRSHALL JSON DATA", http.StatusInternalServerError)
		return
	}

	//log.Println(string(resJSON))
	w.Header().Set("Content-Type", "application/json")
	w.Write(resJSON)
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("INTO THE USERS")

	jsonFile, err := os.Open("platform/src/data/sampleUsers.json")
	if err != nil {
		http.Error(w, "Error opening JSON file", http.StatusInternalServerError)
		return
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		http.Error(w, "COULD NOT READ JSON FILE", http.StatusInternalServerError)
		return
	}

	//log.Println(string(jsonData))

	var users []User
	err = json.Unmarshal(jsonData, &users)
	/* //Uncommenting this will cause an error
	if err != nil {
		http.Error(w, "COULD NOT UNAMRSHALL JSON DATA", http.StatusInternalServerError)
		return
	}
	*/

	resJSON, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "COULD NOT UNAMRSHALL JSON DATA", http.StatusInternalServerError)
		return
	}

	//log.Println(string(resJSON))
	w.Header().Set("Content-Type", "application/json")
	w.Write(resJSON)
}
