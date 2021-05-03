package domain

import (
	"github.com/dembygenesis/student_enrollment_exam/src/api/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type studentDaoInterface interface {
	Create(name string, email string, phone string) error
	SetClient()
}

type studentDao struct {
	client *sqlx.DB
}

var (
	StudentDao studentDaoInterface
)

func init() {
	StudentDao = &studentDao{}
	StudentDao.SetClient()
}

func (s *studentDao) SetClient() {
	s.client = utils.GetMYSQLConnection()
}

// Create inserts a new entry to the student table
func (s *studentDao) Create(name string, email string, phone string) error {
	sql := `
		INSERT INTO student (
		  student_name,
		  student_email,
		  student_phone
		)
		VALUES
		  (
			?,
			?,
			?
		  );
	`

	_, err := s.client.Exec(sql, name, email, phone)

	return err
}