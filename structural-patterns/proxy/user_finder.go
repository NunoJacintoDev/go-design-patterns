package proxy

type User struct {
	ID int32
}

type UserFinder interface {
	FindUser(id int32) (User, error)
}
