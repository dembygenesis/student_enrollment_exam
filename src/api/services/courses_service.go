package services

import (
	"errors"
	"github.com/dembygenesis/student_enrollment_exam/src/api/domain"
)

type courseService struct {

}

var (
	CourseService *courseService
)

func init() {
	CourseService = &courseService{}
}

func (s *courseService) Create(name string, professor string, description string) error {
	// Quick validations to prevent hitting the database
	if name == "" {
		return errors.New("name is invalid")
	}
	if professor == "" {
		return errors.New("professor is invalid")
	}
	if description == "" {
		return errors.New("description is invalid")
	}

	// Perform create
	return domain.CourseDao.Create(name, professor, description)
}

func (s *courseService) DeleteCourse(courseId int) error {
	// Quick validations to prevent hitting the database
	if courseId == 0 {
		return errors.New("course_id is invalid")
	}

	isValidCourseId, err := domain.CourseDao.IsValidId(courseId)

	if err != nil {
		return err
	}

	if isValidCourseId == false {
		return errors.New("course_id is invalid")
	}

	return domain.CourseDao.DeleteCourse(courseId)
}