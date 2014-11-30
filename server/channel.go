package server

type Channel struct {
	Name  string
	Users map[string]*User
}
