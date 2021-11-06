package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var action = []string{
	"logged in",
	"logged out",
	"create record",
	"delete record",
	"update record",
}

type logItem struct {
	action    string
	timestamp time.Time
}
type User struct {
	id    int
	email string
	logs  []logItem
}

func (u User) getActivityInfo() string {
	out := fmt.Sprintf("ID: %d | Email: %s\nActivity Log:\n", u.id, u.email)
	for i, item := range u.logs {
		out += fmt.Sprintf("%d. [%s] at %s\n", i+1, item.action, item.timestamp)
	}
	return out
}

func main() {

	//user := User{1, "max@ninja",
	//	[]logItem{
	//		{action[0], time.Now()},
	//		{action[2], time.Now()},
	//		{action[3], time.Now()},
	//		{action[1], time.Now()},
	//	}}
	//fmt.Println(user.getActivityInfo())

	rand.Seed(time.Now().Unix())
	users := generateUsers(1000)
	for _, user := range users {
		err := saveUserInfo(user)
		if err != nil {
			log.Fatalln(err)
		}
	}

}
func saveUserInfo(user User) error {
	fmt.Printf("Writing file for user id: %d\n", user.id)
	fileName := fmt.Sprintf("mod2ConcurrencyInGo/les2/logs/uid%d.txt", user.id)
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	_, err = file.WriteString(user.getActivityInfo())
	return err
}

//Count of users
func generateUsers(count int) []User {
	users := make([]User, count)
	for i := 0; i < count; i++ {
		users[i] = User{
			id:    i + 1,
			email: fmt.Sprintf("comeUser%d@gmail.com", i+1),
			logs:  generateLogs(rand.Intn(5)),
		}
	}
	return users
}
func generateLogs(count int) []logItem {
	logs := make([]logItem, count)
	for i := 0; i < count; i++ {
		logs[i] = logItem{
			action:    action[rand.Intn(len(action)-1)],
			timestamp: time.Now(),
		}
	}
	return logs
}
