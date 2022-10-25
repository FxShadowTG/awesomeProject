package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 1.定义一个全局对象db
var db *sql.DB

// 2.创建初始化数据库方法
func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True"
	//不会校验账号密码是否正确
	//注意，这里不要用:=，因为我们是给上面的全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	//尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {

		return err
	}
	return nil
}

type User struct {
	id       int
	username string
	password string
}

/*
插入方法
*/
func insert(un string, pw string) {
	s := "insert into user_tb1 (username,password) values(?,?)"
	r, err := db.Exec(s, un, pw)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		i, _ := r.LastInsertId()
		fmt.Printf("i: %v\n", i)
	}
}

/*
删除方法
*/
func delete() {

	s := "delete from user_tb1 where id = ?"
	r, err := db.Exec(s, 10)
	if err != nil {
		fmt.Println(err)
	} else {
		i, err := r.RowsAffected()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(i)
	}
}

/*
修改方法
*/
func update() {
	s := "update user_tb1 set username = ?, password = ? where id = ?"
	r, err := db.Exec(s, "wwm", "88888888", 7)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		//返回的i是影响的行数
		i, _ := r.RowsAffected()
		fmt.Printf("i: %v\n", i)
	}
}

/*
查询方法
*/
// 查询单条数据
func queryOneRow(id int) {
	//创建sql语句，占位符用?代替
	s := "select * from user_tb1 where id = ?"
	//创建user对象
	var u User
	//调用QueryRow()查询单条数据
	err := db.QueryRow(s, id).Scan(&u.id, &u.username, &u.password)
	if err != nil {
		//如果错就打印err
		fmt.Printf("err: %v\n", err)
	} else {
		//如果没错就打印那条数据
		fmt.Printf("u: %v\n", u)
	}
}

// 查询多条数据
func queryManyRow() {
	//创建sql语句，占位符用?代替
	s := "select * from user_tb1"
	//调用Query()查询多条数据
	rows, err := db.Query(s)
	//创建user对象
	var u User
	//注意最后要关闭rows释放持有的数据库连接
	defer rows.Close()
	if err != nil {
		//如果错就打印err
		fmt.Printf("err: %v\n", err)
	} else {
		////如果没错就遍历数据并加入user对象
		for rows.Next() {
			rows.Scan(&u.id, &u.username, &u.password)
			fmt.Printf("u: %v\n", u)
		}
	}
}

func main() {
	// 3.调用
	err := initDB()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("数据库连接成功!\n")
	}

	//insert()
	//insert("Aaron", "666")

	//queryOneRow(2)
	//queryManyRow()

	//update()
	delete()
}
