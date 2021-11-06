package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
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
	rand.Seed(time.Now().Unix())
	t := time.Now()
	wg := &sync.WaitGroup{}
	users := generateUsers(1000)
	for _, user := range users {
		wg.Add(1)
		go saveUserInfo(user, wg)
	}
	wg.Wait() //wait go routing is done
	fmt.Println("TIME ELAPSED", time.Since(t))
}
func saveUserInfo(user User, wg *sync.WaitGroup) error {
	fmt.Printf("Writing file for user id: %d\n", user.id)
	time.Sleep(time.Millisecond * 10) //wait like a network
	fileName := fmt.Sprintf("mod2ConcurrencyInGo/les2/logs/uid%d.txt", user.id)
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	_, err = file.WriteString(user.getActivityInfo())
	if err != nil {
		return err
	}
	wg.Done()
	return nil
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
