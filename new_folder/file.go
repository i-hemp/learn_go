package main

import (
	// "pack1/d1.go"
	"fmt"
	"new_folder/pack1"
	"strconv"
	"sync"
	"time"
	// "strings"
)

type UserData struct {
	firstName   string
	lastName    string
	email       string
	noOfTickets uint
}

var details = make([]UserData, 0)
var as = UserData{
	firstName: string("232ff"),
}

func sleeps() {
	fmt.Println("before waited for 3sec")

	time.Sleep(3 * time.Second)

	fmt.Println("waited for 3sec")
	wg1.Done()
}

// func sleeps(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println("before waited for 10sec")
// 	time.Sleep(3 * time.Second)
// 	fmt.Println("waited for 10sec")
// }

var wg sync.WaitGroup
var wg1 = sync.WaitGroup{}

func main() {
	fmt.Println(pack1.D1())

	// fmt.Println(pack1.D1())
	wg1.Add(1)
	go sleeps()

	// go sleeps()
	fmt.Println("after sleeps")
	var map1 = make(map[string]string) // {"a":"efib"}
	map1["tickets"] = strconv.FormatInt(200, 10)
	fmt.Println(map1)
	wg1.Wait()
	// fmt.Println(consts)
	// const tickets = 20
	// fmt.Println(tickets)
	// var remaining = tickets - 1
	// fmt.Print(remaining, "%")
	// fmt.Printf("%T", remaining)
	// a := []int{}
	// a = append(a, 12)
	// fmt.Println(a)
	// fmt.Println(a)
	// fmt.Println(len(a))
	// a, b, c := ins()
	// fmt.Println(a, b, c)
	// names := []string{"hema driti"}
	// fmt.Println("outsid/e", names)
	// nums := []int{1, 23}

	// fmt.Println(add(nums))
	// // var count int = 0
	// // red := count == 0
	// // var reds bool = count == 0
	// for _, name := range names {
	// 	s := strings.Fields(name)
	// 	fmt.Println(s)

	// 	fmt.Println(strings.Split(name, " "))
	// }
}

func add(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
