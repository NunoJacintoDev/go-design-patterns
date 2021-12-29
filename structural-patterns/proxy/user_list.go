package proxy

import "fmt"

type UserList []User

func (t *UserList) FindUser(id int32) (User, error) {
	for i := 0; i < len(*t); i++ {
		if (*t)[i].ID == id {
			return (*t)[i], nil
		}
	}
	return User{}, fmt.Errorf("User %d could not be found\n", id)
}

func (t *UserList) addUser(newUser User) {
	*t = append(*t, newUser)
}
