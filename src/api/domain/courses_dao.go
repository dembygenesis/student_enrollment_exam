package domain

import (
	"github.com/dembygenesis/student_enrollment_exam/src/api/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type courseDaoInterface interface {
	Create(name string, professor string, description string) error
	SetClient()
	IsValidId(courseId int) (bool, error)
	DeleteCourse(courseId int) error
}

type courseDao struct {
	client *sqlx.DB
}

var (
	CourseDao courseDaoInterface
)

func init() {
	CourseDao = &courseDao{}
	CourseDao.SetClient()
}

func (s *courseDao) SetClient() {
	s.client = utils.GetMYSQLConnection()
}

// Create inserts a new entry to the course table
func (s *courseDao) Create(name string, professor string, description string) error {
	sql := `
		INSERT INTO course (
		  course_name,
		  course_professor,
		  course_description
		)
		VALUES
		  (
			?,
			?,
			?
		  );
	`

	_, err := s.client.Exec(sql, name, professor, description)

	return err
}

func (s *courseDao) IsValidId(id int) (bool, error) {
	var count int
	sql := `
		SELECT COUNT(*) AS countt
		FROM course 
		WHERE course_id = ?
	`

	err := s.client.Get(&count, sql, id)

	if err != nil {
		return false, err
	}

	if count == 0 {
		return false, nil
	}

	return true, nil
}

func (s *courseDao) DeleteCourse(id int) error {
	sql := `
		DELETE FROM course 
		WHERE course_id = ?
	`

	_, err := s.client.Exec(sql, id)
	return err
}

