package da

func AddCourse(ch, cm string, xf int) bool {
	if !Connected {
		Connect()
	}

	err := courseT.Create(Course{
		Ch: ch,
		Cm: cm,
		Xf: xf,
	}).Error

	return err == nil
}

func UpdateCourse(ch string, course Course) bool {
	if !Connected {
		Connect()
	}

	err := courseT.Where("ch = ?", ch).Update(course).Error

	return err == nil
}

func DeleteCourse(ch string) bool {
	if !Connected {
		Connect()
	}

	err := courseT.Where("ch = ?", ch).Delete(Course{}).Error

	return err == nil
}

func DeleteCourses(ch []string) bool {
	if !Connected {
		Connect()
	}

	if !DeleteCoursesChoose(ch) {
		return false
	}

	err := courseT.Where("ch IN (?)", ch).Delete(Course{}).Error

	return err == nil
}

func FindCourse(ch string) (Course, bool) {
	if !Connected {
		Connect()
	}

	result := Course{}
	err := courseT.Where("ch = ?", ch).First(&result).Error

	return result, err == nil
}

func FindCourses() ([]Course, bool) {
	if !Connected {
		Connect()
	}

	result := make([]Course, 0)
	err := courseT.Find(&result).Error

	return result, err == nil
}

func FindCoursesBy(condition string, args ...interface{}) ([]Course, bool) {
	if !Connected {
		Connect()
	}

	result := make([]Course, 0)
	err := courseT.Where(condition, args...).Find(&result).Error

	return result, err == nil
}

func FindUnownedCourseByStudent(xh string) ([]Course, bool) {
	if !Connected {
		Connect()
	}

	result := make([]Course, 0)


	owned, ok := FindOwnedScByStudent(xh)
	if !ok {
		return result, false
	}
	ch := make([]string, len(owned))
	for k, v := range owned {
		ch[k] = v.Ch
	}

	if len(ch) == 0 {
		return FindCourses()
	}

	err := courseT.Where("ch NOT IN (?)", ch).Find(&result).Error

	return result, err == nil
}
