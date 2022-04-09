package main

import "encoding/json"

type User struct {
	Userid   int
	Username string
}

func getUsers() ([]User, error) {
	var u User
	res, err := db.Query("SELECT userid, username FROM mersal.users")
	if err != nil {
		return nil, err
	}
	defer res.Close()

	items := make([]User, 0, 20)
	for res.Next() {
		res.Scan(&u.Userid, &u.Username)
		items = append(items, u)
	}
	return items, nil
}

func marchal(users []User) (result string, err error) {

	res, err := json.Marshal(users)
	result = string(res)
	return result, nil
}
