package domain

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type courseDaoInterface interface {
	Create(name string, professor string, description string) error
}

type courseDao struct {

}

var (
	client *sqlx.DB
	CourseDao courseDaoInterface
)

// initializeClient establishes a database connection for MYSQL
func initializeClient() *sqlx.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbDatabase := os.Getenv("DB_DATABASE")
	dbPort := os.Getenv("DB_PORT")

	connString := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":"+ dbPort +")/" + dbDatabase + "?parseTime=true"
	client, err := sqlx.Open("mysql", connString)
	if err != nil {
		log.Println(err)
		panic("Something went wrong with your database connection.")
	}

	maxConnections, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	client.SetMaxOpenConns(maxConnections)

	return client
}

func init() {
	CourseDao = &courseDao{}
	client = initializeClient()

	fmt.Println("=============== Started ===============")
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

	_, err := client.Exec(sql, name, professor, description)

	return err
}