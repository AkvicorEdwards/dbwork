此项目为大二下的数据库期末作业

使用go作为后端数据处理，HTML+CSS+JS来显示前端数据，MySQL8.0储存数据  
- HTML静态文件储存在tpl文件夹下，运行时会把html从文件加载到内存
- 网站图标为根目录下`favicon.ico`文件

# 使用指导

## 编译

```shell script
go mod download
go build
```

## 初始化数据库

```shell script
# dbwork init_database 数据库用户 用户密码 数据库名字[不能与原有数据库重名，
#                                          该命令将以此名字创建一个新的数据库]
./dbwork init_database root password dbwork
```

## 使用

```shell script
# port 监听的端口
# dbname 数据库名
# dbuser 数据库用户名
# dbpassword 用户密码
./dbwork port 8080 dbname dbwork dbuser root dbpassword password

# 详细参数及参数的默认值使用help查看
./dbwork help
```



# TODO

- [x] 学生基本信息表
    - [x] 新增
    - [x] 删除
    - [x] 修改
    - [x] 关闭 返回主页
- [x] 课程基本信息表
    - [x] 新增
    - [x] 删除
    - [X] 修改
    - [x] 关闭 返回主页
- [x] 学生选课信息表
    - [x] 查询 输入学生学号，查询学生详细的信息和选课的详细信息
    - [x] 选课 输入学生的学号和课程号，进行选课，可以显示学生选课的详细信息
    - [x] 退课 输入学生的学号和课程号可以退课
    - [x] 关闭 返回主页
    - [x] 清除 清除显示的全部信息

# Database struct

## students 学生信息

```mysql
create table students
(
	xh char(10) null comment 'Student id',
	xm char(8) null comment 'Student name',
	age integer(2) null comment 'Student age',
	xb char(2) null comment 'Student sex',
	sdept char(20) null comment 'Student department'
);

create unique index students_xh_uindex on students (xh);

alter table students add constraint students_pk primary key (xh);
```

## course 课程信息

```mysql
create table course
(
	ch char(2) null comment 'Curriculum id',
	cm char(20) null comment 'Curriculum name',
	xf integer(1) null comment 'Academic credit'
);

create unique index course_ch_uindex on course (ch);

alter table course add constraint course_pk primary key (ch);
```

## sc 学生的选课

```mysql
create table sc
(
	xh char(10) null comment 'Student id',
	ch char(2) null comment 'Curriculum ID',
	g DECIMAL(4,1) null comment 'Curriculum score',
	constraint sc_course_ch_fk foreign key (ch) references course (ch),
	constraint sc_students_xh_fk foreign key (xh) references students (xh)
);
```