package cases

import "fmt"

type Database interface {
	AddData(value string)
}

type PostgresDB struct{}

func (db *PostgresDB) AddData(value string) {
	fmt.Printf("Add data to PostgresDB: %s", value)
}

type MockDB struct{}

func (db *MockDB) AddData(value string) {
	fmt.Printf("Add data to MockDB: %s", value)
}

type UserService struct {
	db Database
}

func (s *UserService) AddUser(name string) {
	s.db.AddData(name)
}

func NewUserService(db Database) *UserService {
	return &UserService{db: db}
}
