package storage

import "github.com/satish/golangdev/internal/types"

type Storage interface {
	CreateStudent(name, email string, age int) (int64, error)
	GetStudentByID(id int64) (types.Student, error)
	GetAllStudents() ([]types.Student, error)
	UpdateStudentByID(id int64, name, email string, age int) error
	DeleteStudentByID(id int64) error
}
