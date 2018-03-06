package main

import f "fmt"

type UserDao interface {
	add(u User)
	update(u User)
	remove(u User)
	findAll() []User
}

type User struct {
	name string
	id int
}	

type UserDaoMysqlImpl struct {
	
}

type UserDaoOracleImpl struct {

}

func (dao UserDaoMysqlImpl) add(u User) {
	f.Println("add user[mysql] :",u)
}

func (dao UserDaoMysqlImpl) update(u User) {
	f.Println("update user[mysql] :",u)
}

func (dao UserDaoMysqlImpl) remove(u User) {
	f.Println("remove user[mysql] :",u.id)
}

func (dao UserDaoOracleImpl) add(u User) {
	f.Println("add user[oralce] :",u)
}

func (dao UserDaoOracleImpl) update(u User) {
	f.Println("update user[oracle] :",u)
}

func (dao UserDaoOracleImpl) remove(u User) {
	f.Println("remove user[oracle] :",u.id)
}

func (dao UserDaoMysqlImpl) findAll() []User {
	users := make([]User,3)
	users[0], users[1], users[2] = User{"wq",1}, User{"aa",2}, User{"gudqs",3}
	return users;
}

func (dao UserDaoOracleImpl) findAll() []User {
	users := make([]User,3)
	users[0] = User{"gg",007}
	return users
}

func main(){
	var dao UserDao
	u := User{"qq",88}
	dao =UserDaoOracleImpl{}
	dao.add(u)
	dao.update(u)
	dao.remove(u)
	for i,u := range dao.findAll() {
		f.Println(i,u)
	}
	dao =UserDaoMysqlImpl{}
	dao.add(u)
	dao.update(u)
	dao.remove(u)
	for j,u2 := range dao.findAll() {
		f.Println("mysql:",j,u2)
	}
	service := UserServiceImpl{}
	service.add(u,dao)
	dao = UserDaoOracleImpl{}
	service.add(u,dao)
}

type UserServiceImpl struct {}

func (service UserServiceImpl) add(u User,dao UserDao) {
	f.Print("service do:")
	dao.add(u)
}

