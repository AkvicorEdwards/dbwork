package da

import "testing"

// 测试前的初始化：
// 需要在数据库内无任何数据的情况下成功运行下列两个函数
// TestAddCourse
// TestAddStudent
func TestChooseCourse(t *testing.T) {
	if !ChooseCourse("0001", "1", 46.0) {
		t.Error("0001 1")
	}
	if !ChooseCourse("0001", "2", 78.0) {
		t.Error("0001 2")
	}
	if !ChooseCourse("0001", "3", 69.0) {
		t.Error("0001 3")
	}
	if !ChooseCourse("0002", "1", 57.0) {
		t.Error("0002 1")
	}
	if !ChooseCourse("0002", "2", 56.0) {
		t.Error("0002 2")
	}
	if !ChooseCourse("0003", "1", 88.0) {
		t.Error("0003 1")
	}
}

func TestUpdateChoose(t *testing.T) {
	if !UpdateChoose("0001", "2", Sc{
		Xh: "0001",
		Ch: "2",
		G:  76.4,
	}) {
		t.Error("0001 2")
	}
}

func TestDeleteChoose(t *testing.T) {
	if !DeleteChoose("0001", "1") {
		t.Error("0001")
	}
}

func TestDeleteStudentChoose(t *testing.T) {
	if !DeleteStudentChoose("0002") {
		t.Error("Delete 0002")
	}
}

func TestDeleteStudentsChoose(t *testing.T) {
	if !DeleteStudentsChoose([]string{
		"0003",
	}) {
		t.Error("Delete 0003")
	}
}

func TestDeleteCourseChoose(t *testing.T) {
	if !DeleteCourseChoose("1") {
		t.Error("Delete 1")
	}
}

func TestDeleteCoursesChoose(t *testing.T) {
	if !DeleteCoursesChoose([]string{
		"2",
	}) {
		t.Error("Delete 2")
	}
}