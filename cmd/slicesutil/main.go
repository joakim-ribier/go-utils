package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/joakim-ribier/go-utils/pkg/slicesutil"
)

type User struct {
	Id   int
	Age  int
	Name string
}

func getUsers() []User {
	return []User{
		{Id: 1, Name: "Name 1", Age: 1},
		{Id: 2, Name: "Name 2", Age: 2},
		{Id: 3, Name: "Name 3", Age: 3}}
}

func main() {
	filterByNonEmpty()
	findNextEl()

	filterT()
	findT()
	transformT()
	sortT()
	toString()
	println("")
}

func filterByNonEmpty() {
	in := []string{"1", "", "3", "  "}
	fmt.Println("FilterByNonEmpty")

	var noneEmpty []string
	for _, s := range in {
		if strings.TrimSpace(s) != "" {
			noneEmpty = append(noneEmpty, s)
		}
	}
	fmt.Printf("// R0: %v", noneEmpty)

	// new way to do with the lib
	noneEmpty = slicesutil.FilterByNonEmpty(in)
	fmt.Printf("\n// R1: %v", noneEmpty)
}

func findNextEl() {
	in := []string{"1", "2", "3", "4"}
	fmt.Println("\nFindNextEl")

	var el string
	nextEl := false
	for _, s := range in {
		if nextEl {
			el = s
			break
		}
		if s == "3" {
			nextEl = true
		}
	}
	fmt.Printf("// R0: %v", el)

	// new way to do with the lib
	el = slicesutil.FindNextEl(in, "3")
	fmt.Printf("\n// R1: %v", el)
}

func filterT() {
	users := getUsers()
	fmt.Println("\nFilterT")

	var filteredUsers []User
	for _, user := range users {
		if user.Age > 1 {
			filteredUsers = append(filteredUsers, user)
		}
	}
	fmt.Printf("// R0: %v", filteredUsers)

	// new way to do with the lib
	filteredUsers = slicesutil.FilterT[User](users, func(u User) bool {
		return u.Age > 1
	})
	fmt.Printf("\n// R1: %v", filteredUsers)
}

func findT() {
	users := getUsers()
	fmt.Println("\nFindT")

	var foundUser *User
	for _, user := range users {
		if user.Age == 2 {
			foundUser = &user
			break
		}
	}
	fmt.Printf("// R0: %v", foundUser)

	// new way to do with the lib
	foundUser = slicesutil.FindT[User](users, func(u User) bool {
		return u.Age == 2
	})
	fmt.Printf("\n// R1: %v", foundUser)
}

func transformT() {
	users := getUsers()
	fmt.Println("\nTransformT")

	var names []string
	for _, user := range users {
		names = append(names, user.Name)
	}
	fmt.Printf("// R0: %v", names)

	// new way to do with the lib
	names = slicesutil.TransformT[User, string](users, func(u User) (*string, error) {
		return &u.Name, nil
	})
	fmt.Printf("\n// R1: %v", names)
}

func sortT() {
	users := getUsers()
	fmt.Println("\nSortT")

	sort.SliceStable(users, func(i, j int) bool {
		return users[i].Age < users[j].Age
	})
	fmt.Printf("// R0: %v", users)

	// new way to do with the lib -> clone and sort collections
	users = slicesutil.SortT[User](getUsers(), func(u1, u2 User) int {
		switch {
		case u1.Age < u2.Age:
			return -1
		case u1.Age > u2.Age:
			return +1
		}
		return 0
	})
	fmt.Printf("\n// R1: %v", users)
}

func toString() {
	users := getUsers()
	fmt.Println("\nToStringT")

	var names []string
	for _, user := range users {
		names = append(names, user.Name)
	}
	str := strings.Join(names, ", ")
	fmt.Printf("// R0: %s", str)

	// new way to do with the lib
	str = slicesutil.ToStringT[User](users, func(u User) *string {
		return &u.Name
	}, ", ")
	fmt.Printf("\n// R1: %s", str)
}
