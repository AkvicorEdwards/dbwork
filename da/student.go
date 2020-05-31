package da

func AddStudent(xh, xm string, age int, xb, sdept string) bool {
	if !Connected {
		Connect()
	}
	err := studentsT.Create(Students{
		Xh:    xh,
		Xm:    xm,
		Age:   age,
		Xb:    xb,
		Sdept: sdept,
	}).Error

	return err == nil
}

func UpdateStudent(xh string, student Students) bool {
	if !Connected {
		Connect()
	}

	err := studentsT.Where("xh = ?", xh).Update(student).Error

	return err == nil
}

func DeleteStudent(xh string) bool {
	if !Connected {
		Connect()
	}

	if !DeleteStudentChoose(xh) {
		return false
	}

	err := studentsT.Where("xh = ?", xh).Delete(Students{}).Error

	return err == nil
}

func DeleteStudents(xh []string) bool {
	if !Connected {
		Connect()
	}

	if !DeleteStudentsChoose(xh) {
		return false
	}

	err := studentsT.Where("xh IN (?)", xh).Delete(Students{}).Error

	return err == nil
}

func FindStudent(xh string) (Students, bool) {
	if !Connected {
		Connect()
	}

	result := Students{}
	err := studentsT.Where("xh = ?", xh).First(&result).Error

	return result, err == nil
}

func FindStudents() ([]Students, bool) {
	if !Connected {
		Connect()
	}

	result := make([]Students, 0)
	err := studentsT.Find(&result).Error

	return result, err == nil
}

func FindStudentsBy(condition string, args ...interface{}) ([]Students, bool) {
	if !Connected {
		Connect()
	}

	result := make([]Students, 0)
	err := studentsT.Where(condition, args...).Find(&result).Error

	return result, err == nil
}