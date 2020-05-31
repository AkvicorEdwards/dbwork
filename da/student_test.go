package da

import (
	"fmt"
	"testing"
)

func TestAddStudent(t *testing.T) {
	if !AddStudent("0001", "张三", 21, "男", "计算机") {
		t.Error("0001")
	}
	if !AddStudent("0002", "李四", 23, "女", "计算机") {
		t.Error("0002")
	}
	if !AddStudent("0003", "王五", 22, "男", "信息管理") {
		t.Error("0003")
	}
}

func TestUpdateStudent(t *testing.T) {
	if !UpdateStudent("0001", Students{
		Xh:    "0001",
		Xm:    "张三",
		Age:   22,
		Xb:    "男",
		Sdept: "计算机",
	}) {
		t.Error("0001")
	}
}

func TestFindStudent(t *testing.T) {
	res, ok := FindStudent("0002")
	if !ok {
		t.Error("0002")
	} else {
		fmt.Println(res)
	}
}

func TestFindStudents(t *testing.T) {
	res, ok := FindStudents()
	if !ok {
		t.Error("ALL")
	} else {
		for _, v := range res{
			fmt.Println(v)
		}
	}
}

func TestFindStudentsBy(t *testing.T) {
	res, ok := FindStudentsBy("xb = ?", "男")

	if !ok {
		t.Error("男")
	} else {
		for _, v := range res{
			fmt.Println(v)
		}
	}
}

func TestDeleteStudent(t *testing.T) {
	if !DeleteStudent("0001") {
		t.Error("0001")
	}
}

func TestDeleteStudents(t *testing.T) {
	if !AddStudent("0005", "张三1", 21, "男", "计算机") {
		t.Error("0005")
	}
	if !AddStudent("0006", "张三2", 21, "男", "计算机") {
		t.Error("0006")
	}

	all := []string{
		"0005",
		"0006",
	}

	if !DeleteStudents(all) {
		t.Error("Delete All")
	}

}