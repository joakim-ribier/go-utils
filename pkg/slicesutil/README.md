# go-utils

## slices utils

A few examples on the "slices" package.

string

* [FilterT](#filtert)
* [FindT](#findt)

#### FilterByNonEmpty
```go
slicesutil.FilterByNonEmpty([]string{"1", "", "3", ""})
// Result: [1 3]
```

#### FindNextEl
```go
slicesutil.FindNextEl([]string{"1", "2", "3", "4"}, "2")
// Result: 3
```

T
* [FilterT](#filtert)
* [FindT](#findt)
* [TransformT](#transformt)
* [SortT](#sortt)
* [ToStringT](#tostringt)

```go
users := []User{
	{Id: 1, Name: "Name 1", Age: 1},
	{Id, 2, Name: "Name 2", Age: 2},
	{Id: 3, Name: "Name 3", Age: 3}}
```

#### FilterT
```go
var filteredUsers []User
for _, user := range users {
	if user.Age > 1 {
		filteredUsers = append(filteredUsers, user)
	}
}

// new way
filteredUsers = slicesutil.FilterT[User](users, func(u User) bool {
	return u.Age > 1
}) // Result: [{2 2 Name 2} {3 3 Name 3}]
```

#### FindT
```go
var foundUser *User
for _, user := range users {
	if user.Age == 2 {
		foundUser = &user
		break
	}
}

// new way
foundUser = slicesutil.FindT[User](users, func(u User) bool {
	return u.Age == 2
}) // Result: &{2 2 Name 2}
```

#### TransformT
```go
var names []string
for _, user := range users {
	names = append(names, user.Name)
}

// new way
names = slicesutil.TransformT[User, string](users, func(u User) (*string, error) {
	return &u.Name, nil
}) // Result: [Name 1 Name 2 Name 3]
```

#### SortT
```go
sort.SliceStable(users, func(i, j int) bool {
	return users[i].Age > users[j].Age
})

// new way (clone and sort values)
users = slicesutil.SortT[User](getUsers(), func(u1, u2 User) bool {
	return u1.Age > u2.Age
}) // Result: [{3 3 Name 3} {2 2 Name 2} {1 1 Name 1}]
```

#### ToStringT
```go
var names []string
for _, user := range users {
	names = append(names, user.Name)
}
str := strings.Join(names, ", ")

// new way
str = slicesutil.ToStringT[User](users, func(u User) *string {
	return &u.Name
}, ", ") // Result: Name 1, Name 2, Name 3
```