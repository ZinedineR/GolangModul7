package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

var menu int
var ID int = 0
var title, name, saver string
var hour, billing float64
var data = [][4]string{{"ID", "name", "hour", "billing (IDR)"}}

func max() string {
	f := math.MaxFloat64
	return strconv.FormatFloat(f, 'f', 1, 64)
}

func add() {
	fmt.Println("Input name : ")
	fmt.Scan(&name)
	ID++
	fmt.Println("Input hour : ")
	fmt.Scan(&hour)
	saver = strconv.FormatFloat(hour, 'f', 1, 64)
	billing = hour * 60 * 1000
	if valid, err := validate(saver); valid {
		newdata := [4]string{(strconv.Itoa(ID)), name, (strconv.FormatFloat(hour, 'f', 1, 64)), (strconv.FormatFloat(billing, 'f', 1, 64))}
		data = append(data, newdata)
		newdata = [4]string{}
	} else {
		fmt.Println(err.Error())
	}
}
func showAll() {
	fmt.Println("List of name available : \n", len(data))
	for i := 0; i < len(data); i++ {
		for j := 0; j < 4; j++ {
			fmt.Print(data[i][j], "  \t")
		}
		fmt.Println("")
	}
}

func searchID(ID int) int {
	var index int
	search := strconv.Itoa(ID)
	for i := 0; i < len(data); i++ {
		if data[i][0] == search {
			index = i
			break
		}
	}
	return index
}

func deleteID(ID int) {
	var before, after [][4]string
	index := searchID(ID)
	before = data[:index]
	after = data[index+1:]
	data = [][4]string{}
	for i := 0; i < len(before); i++ {
		data = append(data, before[i])
	}
	for i := 0; i < len(after); i++ {
		data = append(data, after[i])
	}
}

func searchname() {
	fmt.Println("Search result(s) : ")
	for i := 0; i < len(data); i++ {
		matched, _ := regexp.MatchString(`^a|^A`, data[i][2])
		if matched == true {
			fmt.Println("ID : ", data[i][0])
			fmt.Println("Title : ", data[i][1])
			fmt.Println("name : ", data[i][2])
			fmt.Println("Vote : ", data[i][3])
		}
	}
}
func stringtoFloat(input string) float64 {
	var save float64
	if s, err := strconv.ParseFloat(input, 64); err == nil {
		save = s
	}
	return save
}
func topthree() {
	if len(data) < 4 {
		fmt.Println("Sorry the data is too short, please add more")
	} else {
		var max string = max()
		// third := [4]string{"99999", "99999", "99999", "99999"}
		// second := [4]string{"99999", "99999", "99999", "99999"}
		// first := [4]string{"99999", "99999", "99999", "99999"}
		third := [4]string{max, max, max, max}
		second := [4]string{max, max, max, max}
		first := [4]string{max, max, max, max}
		for i := 1; i < len(data); i++ {
			if stringtoFloat(data[i][2]) < stringtoFloat(first[2]) {
				third = second
				second = first
				first = data[i]
			} else if stringtoFloat(data[i][2]) < stringtoFloat(second[2]) {
				third = second
				second = data[i]
			} else if stringtoFloat(data[i][2]) < stringtoFloat(third[2]) {
				third = data[i]
			}
		}
		fmt.Print("Top 3 least hours billing : \n")
		fmt.Print(data[0][0], "\t", data[0][1], "\t", data[0][2], "\t", data[0][3], "\n")
		fmt.Print(first[0], "\t", first[1], "\t", first[2], "\t", first[3], "\n")
		fmt.Print(second[0], "\t", second[1], "\t", second[2], "\t", second[3], "\n")
		fmt.Print(third[0], "\t", third[1], "\t", third[2], "\t", third[3], "\n")
	}
}
func countAverage() float64 {
	var count float64 = 0
	var sum float64 = 0
	for i := 1; i < len(data); i++ {
		sum = sum + stringtoFloat(data[i][2])
		count++
	}
	average := sum / count
	return average
}

func onlyfour() {
	average := countAverage()
	fmt.Println("List of billing hours < ", average)
	for i := 1; i < len(data); i++ {
		if stringtoFloat(data[i][2]) < average {
			fmt.Print(data[i][0], "\t", data[i][1], "\t", data[i][2], "\t", data[i][3], "\n")
		}
	}
}
func validate(input string) (bool, error) { //bug, don't know why doesn't work
	if input == "" {
		return false, errors.New("cannot be empty")
	}
	m := regexp.MustCompile("[0-9]")
	if m.MatchString(input) == false {
		return false, errors.New("please input numbers only")
	}
	return true, nil
}

func main() {
	fmt.Println("Welcome to billing system! Please select menu :\n1. Add new billing\n2. Delete a billing\n3. Show all billing\n4. Average hours i\n5. Top 3 least hours\n6. Billing less than average hours\n7. Exit")
	fmt.Scan(&menu)
	if menu == 1 {
		add()
		main()
	} else if menu == 2 {
		fmt.Println("What ID you want to delete? : ")
		fmt.Scan(&saver)
		var num, err = strconv.Atoi(saver)
		if err == nil && num > 0 {
			deleteID(num)
		}
		showAll()
		main()
	} else if menu == 3 {
		showAll()
		main()
	} else if menu == 4 {
		fmt.Println("Average hours used : ", countAverage())
		main()
	} else if menu == 5 {
		topthree()
		main()
	} else if menu == 6 {
		onlyfour()
		main()
	} else if menu == 7 {
		fmt.Println("Thanks")
		os.Exit(1)
	}

}
