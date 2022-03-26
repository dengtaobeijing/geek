package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

//1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
// 应该上抛给业务 做处理

type user struct {
	id         string
	name       string
	createTime string
	modifyTime string
}

func insert(db *sql.DB) (bool, error) {
	_, err := db.Exec("insert into user(id,name,createTime,modifyTime) values('a1','test','2022-03-26 10:43:06','2022-03-26 10:43:06')")
	if err != nil {

		return false, errors.Wrap(err, "insert error")
	}

	return true, nil

}

func getUserData(db *sql.DB) (user, error) {
	sqlStr := "select id,name,create_time,modify_time from user where id='a1'"
	var u user
	err := db.QueryRow(sqlStr).Scan(&u.id, &u.name, &u.createTime, &u.modifyTime)
	if err != nil {
		return u, errors.Wrap(err, "getData error")
	}
	return u, nil
}

func main() {
	var userData user
	db, err := sql.Open("mysql", "root:root@/geek")
	if err != nil {
		fmt.Println(err, errors.Wrap(err, "sql conn error"))
	}
	//insert(db)
	userData, err = getUserData(db)
	if err != nil {
		fmt.Printf("original error:%T\n%v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace:\n%+v\n", err)
		return
	}

	fmt.Println(userData)
}
