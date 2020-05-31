package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

type str2func map[string]func(http.ResponseWriter, *http.Request)

var public str2func

func ParsePrefix() {
	public = make(str2func)
	public["/"] = index
	public["/sc"] = sc
	public["/student"] = student
	public["/course"] = course

	public["/stu_info"] = studentInfo
	public["/owned_list"] = getOwnedList
	public["/unowned_list"] = getUnownedList
	public["/choose_course"] = chooseCourse
	public["/undo_course"] = undoCourse

	public["/post_stu"] = addStudent
	public["/update_stu"] = updateStudent
	public["/delete_stu"] = deleteStudent
	public["/stu_list"] = getStudentList

	public["/post_course"] = addCourse
	public["/update_course"] = updateCourse
	public["/delete_course"] = deleteCourse
	public["/course_list"] = getCourseList
}

type MyHandler struct {}

func (*MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Public
	if h, ok := public[r.URL.Path]; ok {
		h(w, r)
		return
	}

	if ok, _ := regexp.MatchString("/favicon.ico", r.URL.Path); ok {
		download(w, r, "./favicon.ico")
	}

}

func download(w http.ResponseWriter, r *http.Request, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		_, _ = fmt.Fprintln(w, "File Not Found")
		return
	}
	defer func() { _ = file.Close() }()
	data := make([]byte, 1024)
	for {
		n, err1 := file.Read(data)
		if err1 != nil && err1 != io.EOF {
			_, _ = fmt.Fprintln(w, "File Read Error")
			return
		}
		nn, err2 := w.Write(data[:n])
		if err2 != nil || nn != n {
			_, _ = fmt.Fprintln(w, "File Write Error")
			return
		}
		if err1 == io.EOF {
			return
		}
	}
}

func Fprint(w http.ResponseWriter, a ...interface{}) {
	_, _ = fmt.Fprint(w, a...)
}

func Fprintf(w http.ResponseWriter, format string, a ...interface{}) {
	_, _ = fmt.Fprintf(w, format, a...)
}
func Fprintln(w http.ResponseWriter, a ...interface{}) {
	_, _ = fmt.Fprintln(w, a...)
}