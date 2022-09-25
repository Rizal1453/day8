package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	route := mux.NewRouter()

	route.PathPrefix("/public").Handler(http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))

	route.HandleFunc("/",home).Methods("GET")
	route.HandleFunc("/home", home).Methods("GET")
	route.HandleFunc("/contact",contact).Methods("GET")
	route.HandleFunc("/project",project).Methods("GET")
	route.HandleFunc("/blog-detail/{index}",blogDetail).Methods("GET")
	route.HandleFunc("/form-project",AddProject).Methods("POST")
	route.HandleFunc("/form-contact",AddContact).Methods("POST")
	route.HandleFunc("/delete-blog/{index}",deleteBlog).Methods("GET")
	route.HandleFunc("/project-edit/{index}",editProject).Methods("GET")
	route.HandleFunc("/submit-edit/{id}",submitEdit).Methods("POST")
	

	fmt.Println("server running port 7000")
	http.ListenAndServe("localhost:7000",route)
}

func helloWorld(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello World"))
}
func home(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","text/html; charset=utf8")
	var tmpl, err = template.ParseFiles("home.html")

	if err != nil{
		w.Write([]byte("web tidak tersedia" + err.Error()))
		return
	}
	response := map[string]interface{}{
		"Projects":dataProject,
	}
	tmpl.Execute(w,response)
}
func contact(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","text/html; charset=utf8")
	var tmpl, err = template.ParseFiles("contact.html")

	if err != nil{
		w.Write([]byte("web tidak tersedia" + err.Error()))
		return
	}
	tmpl.Execute(w,nil)
}
func project(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","text/html; charset=utf8")
	var tmpl, err = template.ParseFiles("project.html")

	if err != nil{
		w.Write([]byte("web tidak tersedia" + err.Error()))
	
	}
	tmpl.Execute(w,nil)
	
}
var dataProject=[] Struct{}

type Struct struct{
	NamaProject string
	StartDate string
	EndDate string
	Description string
	Nodejs string
	Golang string
	Reactjs string
	Vuejs string
	Duration string
	Id int
}
func AddProject(w http.ResponseWriter,r *http.Request){
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	
	var namaProject = r.PostForm.Get("input-project")
	var startDate = r.PostForm.Get("input-start")
	var endDate = r.PostForm.Get("input-end")
	var description =r.PostForm.Get("input-description")
	var nodejs = r.PostForm.Get("nodejs")
	var golang = r.PostForm.Get("golang")
	var reactjs = r.PostForm.Get("reactjs")
	var vuejs = r.PostForm.Get("vuejs")

	layout := "2006-01-02"
	startDateParse,_:= time.Parse(layout,startDate)
	endDateParse,_ := time.Parse(layout,endDate)

	hours := endDateParse.Sub(startDateParse).Hours()
	days := hours / 24
	weeks := math.Round(days / 7)
  	months := math.Round(days / 30)
 	years := math.Round(days / 365)

	var duration string
	

	if years > 0{
		duration = strconv.FormatFloat(years,'f',0,64) + "years"
	}else if months > 0 {
		duration = strconv.FormatFloat(months, 'f', 0, 64) + " Months"
	}else if weeks > 0 {
		duration = strconv.FormatFloat(weeks,'f',0,64) + "weeks"
	} else if days > 0 {
		duration = strconv.FormatFloat(days, 'f', 0, 64) + " Days"
	} else if hours > 0 {
		duration = strconv.FormatFloat(hours, 'f', 0, 64) + " Hours"
	} else {
		duration = "0 Days"
	}


	var newProject = Struct{
		NamaProject : namaProject,
		StartDate:startDate,
		EndDate:endDate,
		Description:description,
		Nodejs: nodejs,
		Golang: golang,
		Reactjs: reactjs,
		Vuejs: vuejs,
		Duration: duration,
		Id : len(dataProject),
	}
	dataProject = append(dataProject,newProject)
	fmt.Println(dataProject)
	

	http.Redirect(w,r,"/home",http.StatusMovedPermanently)
}
func AddContact(w http.ResponseWriter,r *http.Request){
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Nama : " + r.PostForm.Get("input-nama"))
	fmt.Println("email : " + r.PostForm.Get("input-email"))
	fmt.Println("phone Number : " + r.PostForm.Get("input-phone"))
	fmt.Println("subject : " + r.PostForm.Get("input-subject"))
	fmt.Println("Description : " + r.PostForm.Get("input-description"))
	http.Redirect(w,r,"/home",http.StatusMovedPermanently)
}
func blogDetail(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("blog-detail.html")

	if err != nil {
		w.Write([]byte("message :" + err.Error()))
		return
	}
	var BlogDetail = Struct{}
	index,_ := strconv.Atoi(mux.Vars(r)["index"])
	for i,data := range dataProject{
		if index ==i { 
			BlogDetail = Struct{
				NamaProject: data.NamaProject,
				Description: data.Description,
				StartDate: data.StartDate,
				EndDate: data.EndDate,
				Duration: data.Duration,
				Nodejs: data.Nodejs,
				Golang: data.Golang,
				Reactjs: data.Reactjs,
				Vuejs: data.Vuejs,
			}

		}
	}
	data := map[string]interface{}{
		"Blog": BlogDetail,
		
	}
	fmt.Println(data)
	tmpl.Execute(w,data)
}
func deleteBlog(w http.ResponseWriter,r *http.Request){
	index,_ := strconv.Atoi(mux.Vars(r)["index"])

	dataProject = append(dataProject[:index],dataProject[index+1:]...)
	http.Redirect(w,r,"/home",http.StatusFound)
}
func editProject(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("project-edit.html")

	if err != nil {
		w.Write([]byte("message :" + err.Error()))
		return
	}
	var BlogDetail = Struct{}
	index,_ := strconv.Atoi(mux.Vars(r)["index"])
	for i,data := range dataProject{
		if index ==i { 
			BlogDetail = Struct{
				NamaProject: data.NamaProject,
				Description: data.Description,
				StartDate: data.StartDate,
				EndDate: data.EndDate,
				Duration: data.Duration,
				Id : data.Id,
			}

		}
	}
	data := map[string]interface{}{
		"EditProject": BlogDetail,
	}
	tmpl.Execute(w,data)
}
func submitEdit(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	id,_ := strconv.Atoi(mux.Vars(r)["id"])
	
	
	var namaProject = r.PostForm.Get("input-project")
	var startDate = r.PostForm.Get("input-start")
	var endDate = r.PostForm.Get("input-end")
	var description =r.PostForm.Get("input-description")
	nodejs := r.PostForm.Get("nodejs")
	golang := r.PostForm.Get("golang")
	reactjs := r.PostForm.Get("reactjs")
	vuejs := r.PostForm.Get("vuejs")

	layout := "2006-01-02"
	startDateParse,_ := time.Parse(layout,startDate)
	endDateParse,_ := time.Parse(layout,endDate)

	hours := endDateParse.Sub(startDateParse).Hours()
	days := hours / 24
	weeks := math.Round(days / 7)
  	months := math.Round(days / 30)
 	years := math.Round(days / 365)

	var duration string
	

	if years > 0{
		duration = strconv.FormatFloat(years,'f',0,64) + "years"
	}else if months > 0 {
		duration = strconv.FormatFloat(months, 'f', 0, 64) + " Months"
	}else if weeks > 0 {
		duration = strconv.FormatFloat(weeks,'f',0,64) + "weeks"
	} else if days > 0 {
		duration = strconv.FormatFloat(days, 'f', 0, 64) + " Days"
	} else if hours > 0 {
		duration = strconv.FormatFloat(hours, 'f', 0, 64) + " Hours"
	} else {
		duration = "0 Days"
	}


	var newProject = Struct{
		NamaProject : namaProject,
		StartDate:startDate,
		EndDate:endDate,
		Description:description,
		Nodejs: nodejs,
		Golang: golang,
		Reactjs: reactjs,
		Vuejs: vuejs,
		Duration: duration,
		Id: id,
	}
	dataProject[id]=newProject

	http.Redirect(w,r,"/home",http.StatusMovedPermanently)

}