package da

import "encoding/json"

type Course struct {
	Ch string `json:"ch"` // 课程号
	Cm string `json:"cm"` // 课程名
	Xf int    `json:"xf"` // 学分
}

type Sc struct {
	Xh string  `json:"xh"` // 学号
	Ch string  `json:"ch"` // 课程号
	G  float32 `json:"g"`  // 成绩
}

type Students struct {
	Xh    string `json:"xh"`    // 学号
	Xm    string `json:"xm"`    // 姓名
	Age   int    `json:"age"`   // 年龄
	Xb    string `json:"xb"`    // 性别
	Sdept string `json:"sdept"` // 系别
}

func (c *Students) GetJson() string {
	data, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(data)
}
