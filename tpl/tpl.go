package tpl

import (
	"fmt"
	"html/template"
	"log"
)

var err error
var errCnt = 0

var Index *template.Template
var Sc *template.Template
var Student *template.Template
var Course *template.Template


func Parse() {
	err = nil
	errCnt = 0

	Index = addFromFile("./tpl/index.html")
	Sc = addFromFile("./tpl/sc.html")
	Student = addFromFile("./tpl/student.html")
	Course = addFromFile("./tpl/course.html")

	log.Printf("Parsing the html template was completed with %d errors\n", errCnt)
}

func add(name, tpl string) (t *template.Template) {
	t, err = template.New(name).Parse(tpl)
	if err != nil {
		errCnt++
		fmt.Println(err)
	}
	return
}

func addFromFile(file string) (t *template.Template) {
	t, err = template.ParseFiles(file)
	if err != nil {
		errCnt++
		fmt.Println(err)
	}
	return
}
