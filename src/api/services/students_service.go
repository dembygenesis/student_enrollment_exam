package services

import (
	"errors"
	"fmt"
	"github.com/dembygenesis/student_enrollment_exam/src/api/domain"
)


type studentService struct {

}

var (
	StudentService *studentService
)

func init() {
	StudentService = &studentService{}
}

func (s *studentService) Create(name string, email string, phone string) error {
	return domain.StudentDao.Create(name, email, phone)
}

func (s *studentService) Enroll(studentId int, courseId int) error {
	// Validate student_id
	isValidStudentId, err := domain.StudentDao.IsValidId(studentId)

	if err != nil {
		return err
	}

	if isValidStudentId == false {
		return errors.New("student_id is invalid")
	}

	fmt.Println("isValidStudentId", isValidStudentId)
	fmt.Println("err", err)

	// Validate course_id

	// Perform insert

	return nil

	// return domain.StudentDao.Create(name, email, phone)
}