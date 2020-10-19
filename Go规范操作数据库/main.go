package main

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)

/**
TODO Golang和Java一样，定义了一套sql规范，在database/sql包下，需要具体的驱动才能操作数据库。不过原生的数据库驱动都是基于SQL语句操作的
	类似redigo的形式，对于复杂的业务有点繁琐。所以实际开发中建议使用数据库框架，如gorm。本文件下的代码基于原生数据库驱动开发。
	Mysql驱动：go get github.com/go-sql-driver/mysql
 */
const(
	DB_USERNAME = "root"
	DB_PASSWORD = "ABC123"
	DB_CONNECT_NETWORK = "tcp"
	DB_HOST = "localhost"
	DB_PORT = "3306"
	DB_NAME = "db4golang"
)

type student struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Class string `json:"class"`
	teacherName string `json:"teacher_name"`
}

func main() {
	/**
		定义连接信息(类似jdbcurl)
		用户名:密码@连接协议(主机:端口)/库名
	*/
	connParams := fmt.Sprintf("%s:%s@%s(%s:%s)/%s",DB_USERNAME,DB_PASSWORD,DB_CONNECT_NETWORK,DB_HOST,DB_PORT,DB_NAME)
	conn,error :=sql.Open("mysql",connParams)
	if error==nil{
		//insertData(conn)
		updateData(conn)
		selectData(conn)
	}else{
		fmt.Println("连接出错:",error)
	}

}
//核心API：DB.Exec(sql,params...) TODO sql语句中的参数占位符用?
func insertData(conn *sql.DB){
	var sql string = "insert into student(id,name,class,teacher_name)values (?,?,?,?)"
	id := 1
	name := "KouJyunGenn"
	class := "class4"
	teacherName := "miss.liu"
	result,error :=conn.Exec(sql,id,name,class,teacherName)
	if error == nil{
		fmt.Printf("查询的结果是%v\n",result)
	}else{
		fmt.Println("出现了错误",error)
	}
}

func updateData(conn *sql.DB){
	var sql string = "update student set Teacher_NAME = ? where id = ?"
	id := 1
	result,error :=conn.Exec(sql,"更新的老师名",id)
	if error == nil{
		fmt.Printf("更新的结果是%v\n",result)
	}else{
		fmt.Println("出现了错误",error)
	}
}

//查询的核心API:queryXXX()
func selectData(conn *sql.DB){
	var sql string = "select * from student where id = ?"
	id := 1
	result :=conn.QueryRow(sql,id)
	student := new(student)
	//TODO 转换的顺序必须和查询结果集顺序一致
	result.Scan(&student.Name,&student.Id,&student.Class,&student.teacherName)
	fmt.Printf("查询的结果是%v\n",student)

}