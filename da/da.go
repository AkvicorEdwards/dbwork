package da

import (
	"dbwork/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var (
	Connected = false
	db        *gorm.DB
	studentsT *gorm.DB
	scT       *gorm.DB
	courseT   *gorm.DB
)

func Connect() {
	if Connected {
		Close()
	}
	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf(
		"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		config.DBUser, config.DBPassword, config.DBName))
	if err != nil {
		log.Println(err)
		return
	}
	studentsT = db.Table("students")
	scT = db.Table("sc")
	courseT = db.Table("course")
	Connected = true
}

func Close() {
	if !Connected {
		return
	}
	err := db.Close()
	if err != nil {
		log.Println(err)
	}
	Connected = false
}

func Init(user, password, dbname string) {
	db, err := gorm.Open("mysql", fmt.Sprintf(
		"%s:%s@/?charset=utf8&parseTime=True&loc=Local", user, password))
	if err != nil {
		log.Println(err)
		log.Println("连接数据库失败")
		return
	}
	defer func() {_=db.Close()}()

	err = db.Exec("CREATE DATABASE "+dbname).Error
	if err != nil {
		log.Println(err)
		log.Println("创建数据库失败")
		return
	}

	err = db.Exec("USE "+dbname).Error
	if err != nil {
		log.Println(err)
		log.Println("选择数据库失败")
		return
	}

	sqls := []string{
		`
create table students
(
	xh char(10) null comment 'Student id',
	xm char(8) null comment 'Student name',
	age integer(2) null comment 'Student age',
	xb char(2) null comment 'Student sex',
	sdept char(20) null comment 'Student department'
);
`,`
create unique index students_xh_uindex on students (xh);
`,`
alter table students add constraint students_pk primary key (xh);
`,`
create table course
(
	ch char(2) null comment 'Curriculum id',
	cm char(20) null comment 'Curriculum name',
	xf integer(1) null comment 'Academic credit'
);
`,`
create unique index course_ch_uindex on course (ch);
`,`
alter table course add constraint course_pk primary key (ch);
`,`
create table sc
(
	xh char(10) null comment 'Student id',
	ch char(2) null comment 'Curriculum ID',
	g DECIMAL(4,1) null comment 'Curriculum score',
	constraint sc_course_ch_fk foreign key (ch) references course (ch),
	constraint sc_students_xh_fk foreign key (xh) references students (xh)
);
`,
	}

	for k, v := range sqls {
		err = db.Exec(v).Error
		if err != nil {
			log.Println(err)
			log.Printf("%d 创建数据表失败", k)
			return
		}
	}

	log.Println("数据库初始化成功")
}
