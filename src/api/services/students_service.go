package services

import (
	"errors"
	"github.com/dembygenesis/student_enrollment_exam/src/api/domain"
	"github.com/dembygenesis/student_enrollment_exam/src/api/utils"
)


type studentService struct {

}

var (
	StudentService *studentService
)

func init() {
	StudentService = &studentService{}
}

func (s *studentService) GetStudents() (*[]domain.Student, error) {
	return domain.StudentDao.GetStudents()
}

func (s *studentService) Create(name string, email string, phone string) error {
	// Quick validations to prevent hitting the database
	if name == "" {
		return errors.New("name is invalid")
	}
	if utils.IsValidEmail(email) == false {
		return errors.New("email is invalid")
	}
	if phone == "" {
		return errors.New("phone is invalid")
	}

	// Perform create
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

	// Perform enroll
	err = domain.StudentDao.Enroll(studentId, courseId)

	return err

}