package handler

import (
	"dbwork/da"
	"dbwork/tpl"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func student(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if err := tpl.Student.Execute(w, nil); err != nil {
			fmt.Println(err)
			_, _ = fmt.Fprintf(w, "%v", "Error")
		}
	}
}

func addStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		xh := r.FormValue("xh")
		xm := r.FormValue("xm")
		age, err := strconv.Atoi(r.FormValue("age"))
		if err != nil {
			Fprint(w, "Failure")
			return
		}
		xb := r.FormValue("xb")
		sdept := r.FormValue("sdept")

		if da.AddStudent(xh, xm, age, xb, sdept) {
			Fprint(w, "Successful")
		} else {
			Fprint(w, "Failure")
		}
	}
}

func updateStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		xhs := r.FormValue("xh_s")
		xh := r.FormValue("xh")
		xm := r.FormValue("xm")
		age, err := strconv.Atoi(r.FormValue("age"))
		if err != nil {
			Fprint(w, "Failure")
			return
		}
		xb := r.FormValue("xb")
		sdept := r.FormValue("sdept")

		if da.UpdateStudent(xhs, da.Students{
			Xh:    xh,
			Xm:    xm,
			Age:   age,
			Xb:    xb,
			Sdept: sdept,
		}) {
			Fprint(w, "Successful")
		} else {
			Fprint(w, "Failure")
		}
	}
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		xhs := r.FormValue("xh")

		xh := make([]string, 0)

		err := json.Unmarshal([]byte(xhs), &xh)
		if err != nil {
			Fprint(w, "Failure")
			return
		}

		if da.DeleteStudents(xh) {
			Fprint(w, "Successful")
		} else {
			Fprint(w, "Failure")
		}
	}
}

func getStudentList(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		stu, ok := da.FindStudents()
		if !ok {
			Fprint(w, "")
			return
		}
		data, err := json.Marshal(stu)
		if err != nil {
			Fprint(w, "")
			return
		}
		Fprint(w, string(data))
	}
}

func studentInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		stu, ok := da.FindStudent(r.FormValue("xh"))
		if !ok {
			Fprint(w, "")
			return
		}
		data, err := json.Marshal(stu)
		if err != nil {
			Fprint(w, "")
			return
		}
		Fprint(w, string(data))
	}
}