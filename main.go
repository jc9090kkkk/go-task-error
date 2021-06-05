package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	pkgErrors "github.com/pkg/errors"
	"log"
	"runtime"
)

var db *sql.DB

// 数据库连接初始化
func init() {
	var err error

	// mysql配置
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal("database connect failed：", err)
		return
	}

	// 懒加载连接
	err = db.Ping()
	if err != nil {
		log.Fatal("database connect failed：", err)
		return
	}
}

func QueryOne() (int, string, error) {
	var id int
	var name string
	var err error

	// 获取当前文件的文件名 行号
	_, file, line, _ := runtime.Caller(1)

	// 查询记录
	querySql := "select id, name from user where id = ?"
	err = db.QueryRow(querySql, 0).Scan(&id, &name)

	// 记录为空
	if err == sql.ErrNoRows {
		notFound := errors.New("row data not found")
		return id, name, pkgErrors.Wrapf(notFound, fmt.Sprintf("[SQL: %s], [file:%s], [line:%d], error: %v", querySql, file, line, err))
	}

	return id, name, nil
}

func main() {
	// 查询db
	id, name, err := QueryOne()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(id)
	fmt.Println(name)
}