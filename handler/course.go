package handler

import (
	"dbwork/da"
	"dbwork/tpl"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func course(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//phrase := map[string]interface{}{}
		if err := tpl.Course.Execute(w, nil); err != nil {
			fmt.Println(err)
			_, _ = fmt.Fprintf(w, "%v", "Error")
		}
	}
}

func addCourse(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		ch := r.FormValue("ch")
		cm := r.FormValue("cm")
		xf, err := strconv.Atoi(r.FormValue("xf"))
		if err != nil {
			Fprint(w, "Failure")
			return
		}

		if da.AddCourse(ch, cm, xf) {
			Fprint(w, "Successful")
		} else {
			Fprint(w, "Failure")
		}
	}
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		chs := r.FormValue("ch_s")
		ch := r.FormValue("ch")
		cm := r.FormValue("cm")
		xf, err := strconv.Atoi(r.FormValue("xf"))
		if err != nil {
			Fprint(w, "Failure")
			return
		}
		if da.UpdateCourse(chs, da.Course{
			Ch: ch,
			Cm: cm,
			Xf: xf,
		}) {
			Fprint(w, "Successful")
		} else {
			Fprint(w, "Failure")
		}
	}
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		chs := r.FormValue("ch")

		ch := make([]string, 0)

		err := json.Unmarshal([]byte(chs), &ch)
		if err != nil {
			Fprint(w, "Failure")
			return
		}

		if da.DeleteCourses(ch) {
			Fprint(w, "Successful")
		} else {
			Fprint(w, "Failure")
		}
	}
}

func getCourseList(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		stu, ok := da.FindCourses()
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

func getUnownedList(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		sc, ok := da.FindUnownedCourseByStudent(r.FormValue("xh"))
		if !ok {
			Fprint(w, "")
			return
		}
		data, err := json.Marshal(sc)
		if err != nil {
			Fprint(w, "")
			return
		}
		Fprint(w, string(data))
	}
}