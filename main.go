package main

/*
	@author: Sammidev
	@date: Tuesday, November 3 | 2020 | 9 : 46
*/

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/rs/cors"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Student struct {
	gorm.Model
	Name string
	NIM string
	Mores []More
}

type More struct {
	gorm.Model
	Class string
	Major string
	Email string
	PhoneNumber string
	University string
	Skills string
	ExperienceSkills int
	StudentID int
}

var (
	student = []Student{
		{Name: "Sammi Aldhi Yanto", NIM: "2003119xxx",},
		{Name: "Aditya Andika Putra", NIM: "20031231xxx",},
		{Name: "Gusnur", NIM: "231203119xxx",},
		{Name: "Ayatullah Ramadhan Jacoeb", NIM: "23121233119xxx",},
		{Name: "Abdul Rauf", NIM: "2871203119xxx",},
		{Name: "Aditya Fauzan Nul Haq", NIM: "2871203119xxx",},
		{Name: "Dandi Arnanda", NIM: "2871203119xxx",},
	}

	mores = []More {
		{
			Class:       "A",
			Major:       "Information System",
			Email:       "sammidev@gmai.com",
			PhoneNumber: "08123120931",
			University:  "Riau University",
			Skills:      "Java,Golang,Scala,Ruby",
			ExperienceSkills: 12,
			StudentID:  1,
		},
		{
			Class:       "B",
			Major:       "Information System",
			Email:       "adit@gmai.com",
			PhoneNumber: "08123120931",
			University:  "Riau University",
			Skills:      "Java,Golang,Scala,Ruby",
			ExperienceSkills: 10,
			StudentID:  2,
		},
		{
			Class:       "C",
			Major:       "Information System",
			Email:       "a@gmai.com",
			PhoneNumber: "08123120931",
			University:  "Riau University",
			Skills:      "Java,Golang,Scala,Ruby",
			ExperienceSkills: 12,
			StudentID:  3,
		},
		{
			Class:       "C",
			Major:       "Information System",
			Email:       "rahma@gmai.com",
			PhoneNumber: "b",
			University:  "Riau University",
			Skills:      "Java,Golang,Scala,Ruby",
			ExperienceSkills: 12,
			StudentID:  4,
		},
		{
			Class:       "D",
			Major:       "Information System",
			Email:       "c@gmai.com",
			PhoneNumber: "08123120931",
			University:  "Riau University",
			Skills:      "Java,Golang,Scala,Ruby",
			ExperienceSkills: 12,
			StudentID:  5,
		},
		{
			Class:       "E",
			Major:       "Information System",
			Email:       "d@gmai.com",
			PhoneNumber: "08123120931",
			University:  "Riau University",
			Skills:      "Java,Golang,Scala,Ruby",
			ExperienceSkills: 12,
			StudentID:  6,
		},
		{
			Class:       "F",
			Major:       "Information System",
			Email:       "e@gmai.com",
			PhoneNumber: "08123120931",
			University:  "Riau University",
			Skills:      "Java,Golang,Scala,Ruby",
			ExperienceSkills: 12,
			StudentID:  7,
		},
	}
)

var db *gorm.DB
var err error

const (
	dialect = "postgres"
	host = "host=localhost"
	user = " port=postgres"
	port = " port=5432"
	dbname = " user=postgres"
	sslmode = " sslmode=disable"
	password = " sslmode=sammidev"
	runningOnPort = ":9000"
)

func main() {

	router := mux.NewRouter()
	db, err = gorm.Open(
		dialect,
		host + user + port + dbname + sslmode + password);

	if err != nil {
		panic("FAILED TO CONNECT DATABASE")
	}

	defer db.Close()

	db.AutoMigrate(&Student{})
	db.AutoMigrate(&More{})

	for index := range student {
		db.Create(&student[index])
	}

	for index := range mores {
		db.Create(&mores[index])
	}

	router.HandleFunc("/mores/{id}", Getmore).Methods("GET")
	router.HandleFunc("/mores/{id}", DeleteMores).Methods("DELETE")
	router.HandleFunc("/mores", GetMores).Methods("GET")
	router.HandleFunc("/students/{id}", GetStudent).Methods("GET")

	handler := cors.Default().Handler(router)
	log.Println("RUN SMOOTHLY")
	log.Fatal(http.ListenAndServe(runningOnPort, handler))
}

func Getmore(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var more More
	db.First(&more, params["id"])
	json.NewEncoder(w).Encode(&more)
}

func GetMores(w http.ResponseWriter, r *http.Request) {
	var mores []More
	db.Find(&mores)
	json.NewEncoder(w).Encode(&mores)
}

func DeleteMores(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var more More
	db.First(more, params["id"])
	db.Delete(&more)
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var student Student
	var more []More
	db.First(&student, params["id"])
	db.Model(&student).Related(&more)
	student.Mores = more
	json.NewEncoder(w).Encode(&student)
}