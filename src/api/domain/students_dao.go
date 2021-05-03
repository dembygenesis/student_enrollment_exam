package domain

import (
	"github.com/dembygenesis/student_enrollment_exam/src/api/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type studentDaoInterface interface {
	Create(name string, email string, phone string) error
	Enroll(studentId int, courseId int) error
	IsValidId(studentId int) (bool, error)
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

// Enroll inserts a new entry to the students_enrolled table
func (s *studentDao) Enroll(studentId int, courseId int) error {
	sql := `
		INSERT INTO students_enrolled (                
		  student_ref_id,
		  course_ref_id
		)
		VALUES
		  (
			?,
			?
		  );
	`

	_, err := s.client.Exec(sql, studentId, courseId)

	return err
}


func (s *studentDao) IsValidId(id int) (bool, error) {
	var count int
	sql := `
		SELECT COUNT(*) AS countt
		FROM student 
		WHERE student_id = ?
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