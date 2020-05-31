package da

func ChooseCourse(xh, ch string, g float32) bool {
	if !Connected {
		Connect()
	}

	err := scT.Create(Sc{
		Xh: xh,
		Ch: ch,
		G:  g,
	}).Error

	return err == nil
}

func DeleteChoose(xh, ch string) bool {
	if !Connected {
		Connect()
	}

	err := scT.Where("xh = ? AND ch = ?", xh, ch).Delete(Sc{}).Error

	return err == nil
}

func DeleteStudentChoose(xh string) bool {
	if !Connected {
		Connect()
	}

	err := scT.Where("xh = ?", xh).Delete(Sc{}).Error

	return err == nil
}

func DeleteStudentsChoose(xh []string) bool {
	if !Connected {
		Connect()
	}

	err := scT.Where("xh IN (?)", xh).Delete(Sc{}).Error

	return err == nil
}

func DeleteCourseChoose(ch string) bool {
	if !Connected {
		Connect()
	}

	err := scT.Where("ch = ?", ch).Delete(Sc{}).Error

	return err == nil
}

func DeleteCoursesChoose(ch []string) bool {
	if !Connected {
		Connect()
	}

	err := scT.Where("ch IN (?)", ch).Delete(Sc{}).Error

	return err == nil
}

func UpdateChoose(xh, ch string, sc Sc) bool {
	if !Connected {
		Connect()
	}

	err := scT.Where("xh = ? AND ch = ?", xh, ch).Update(sc).Error

	return err == nil
}

func FindOwnedScByStudent(xh string) ([]Sc, bool) {
	if !Connected {
		Connect()
	}

	result := make([]Sc, 0)
	err := scT.Where("xh = ?", xh).Find(&result).Error

	return result, err == nil
}

