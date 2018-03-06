package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	//"time"
)

func main() {
	db, err := sql.Open("mysql", "root:guddqs@tcp(192.168.1.92:3306)/test?charset=utf8")
	checkErr(err)
	
	stmt, err := db.Prepare("insert userinfo set username=?,departname=?,created=?")
	checkErr(err)	

	res, err := stmt.Exec("wq","SCOTT","2012-12-23")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)
	
	fmt.Println(id)

	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("wqguddqs", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	rows, err := db.Query("select * from userinfo")
	checkErr(err)
	
	for rows.Next() {
		var (
			uid int
			username string
			department string
			created string
		)
		err := rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid, username, department, created)
	}
	
//	stmt, err = db.Prepare("delete from userinfo where uid=?")
//	checkErr(err)

//	res, err = stmt.Exec(id)
	
//	affect, err = res.RowsAffected()
//	checkErr(err)

//	fmt.Println(affect)

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
