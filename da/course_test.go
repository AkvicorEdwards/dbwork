package da

import (
	"fmt"
	"testing"
)

func TestAddCourse(t *testing.T) {
	if !AddCourse("1", "数据库", 3) {
		t.Error("1")
	}
	if !AddCourse("2", "数据结构", 4) {
		t.Error("2")
	}
	if !AddCourse("3", "C语言", 4) {
		t.Error("3")
	}
}

func TestUpdateCourse(t *testing.T) {
	if !UpdateCourse("1", Course{
		Ch: "1",
		Cm: "数据库",
		Xf: 4,
	}) {
		t.Error("1")
	}
}

func TestDeleteCourse(t *testing.T) {
	if !DeleteCourse("1") {
		t.Error("1")
	}
}

func TestDeleteCourses(t *testing.T) {
	if !DeleteCourses([]string{
		"1",
		"2",
	}) {
		t.Error("1 2")
	}
}

func TestFindCourse(t *testing.T) {
	res, ok := FindCourse("2")
	if !ok {
		t.Error("2")
	} else {
		fmt.Println(res)
	}
}

func TestFindCourses(t *testing.T) {
	res, ok := FindCourses()
	if !ok {
		t.Error("ALL")
	} else {
		for _, v := range res{
			fmt.Println(v)
		}
	}
}

func TestFindCoursesBy(t *testing.T) {
	res, ok := FindCoursesBy("xf = ?", 4)

	if !ok {
		t.Error(4)
	} else {
		for _, v := range res{
			fmt.Println(v)
		}
	}
}