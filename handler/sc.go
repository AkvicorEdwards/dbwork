package handler

import (
	"dbwork/da"
	"dbwork/tpl"
	"encoding/json"
	"fmt"
	"net/http"
)

func sc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if err := tpl.Sc.Execute(w, nil); err != nil {
			fmt.Println(err)
			_, _ = fmt.Fprintf(w, "%v", "Error")
		}
	}
}

func getOwnedList(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		sc, ok := da.FindOwnedScByStudent(r.FormValue("xh"))
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

func chooseCourse(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		xh := r.FormValue("xh")
		chs := r.FormValue("ch")
		ch := make([]string, 0)
		err := json.Unmarshal([]byte(chs), &ch)
		if err != nil {
			Fprint(w, "Failure")
			return
		}
		successful := 0
		failure := 0
		for _, v := range ch {
			if da.ChooseCourse(xh, v, 0) {
				successful++
			} else {
				failure++
			}
		}
		Fprintf(w, "Successful: %d\nFailure: %d", successful, failure)
	}
}

func undoCourse(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		xh := r.FormValue("xh")
		chs := r.FormValue("ch")
		ch := make([]string, 0)
		err := json.Unmarshal([]byte(chs), &ch)
		if err != nil {
			Fprint(w, "Failure")
			return
		}
		successful := 0
		failure := 0
		for _, v := range ch {
			if da.DeleteChoose(xh, v) {
				successful++
			} else {
				failure++
			}
		}
		Fprintf(w, "Successful: %d\nFailure: %d", successful, failure)
	}
}