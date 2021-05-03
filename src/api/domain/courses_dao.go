package domain

import (
	"github.com/dembygenesis/student_enrollment_exam/src/api/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type courseDaoInterface interface {
	Create(name string, professor string, description string) error
	SetClient()
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