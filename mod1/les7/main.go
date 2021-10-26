package main

import "fmt"

type User struct {
	name   string
	age    int
	sex    string
	weight int
	height int
}

func NewUser(name string, age int, sex string, weight int, height int) User {
	return User{
		name:   name,
		age:    age,
		sex:    "Male",
		weight: 75,
		height: 185,
	}
}
func (u *User) changNameUser(name string) {
	u.name = name
	return
}
func (u User) printUserInfo() {
	fmt.Println(u.name, u.age, u.sex, u.weight, u.height)
}

//type DumpDatabase struct {
//	m map[string]string
//}
//
//func NewDumpDatabase() *DumpDatabase {
//	return &DumpDatabase{
//		m: make(map[string]string),
// 	}
//}

func main() {
	user1 := NewUser("Vasay", 23, "Male", 75, 185)
	user2 := User{"Pety", 33, "Male", 85, 190}
	user1.printUserInfo()
	user2.printUserInfo()
	user1.changNameUser("Oleg")
	user1.printUserInfo()

	//fmt.Println(user1, user2)
	//fmt.Println(user1.name, user2.name)
	//fmt.Printf("%+v \n %+v \n", user1, user2)

	//printUserInfo(user1)
	//printUserInfo(user2)
}

//func printUserInfo(user User)  {
//fmt.Println(user.name,user.age,user.sex,user.weight,user.height)
//}
