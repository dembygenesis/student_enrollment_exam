package services

import (
	"errors"
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

	// Quick validations to prevent hitting the database
	if studentId == 0 {
		return errors.New("student_id is invalid")
	}
	if courseId == 0 {
		return errors.New("course_id is invalid")
	}

	// Validate student_id
	isValidStudentId, err := domain.StudentDao.IsValidId(studentId)

	if err != nil {
		return err
	}

	if isValidStudentId == false {
		return errors.New("student_id is invalid")
	}

	// Validate course_id
	isValidCourseId, err := domain.CourseDao.IsValidId(courseId)

	if err != nil {
		return err
	}

	if isValidCourseId == false {
		return errors.New("course_id is invalid")
	}

	// Perform insert

	return nil

	// return domain.StudentDao.Create(name, email, phone)
}