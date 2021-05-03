package controllers

import (
	"fmt"
	"github.com/dembygenesis/student_enrollment_exam/src/api/domain"
	"github.com/gin-gonic/gin"
)

type courseController struct {

}

var (
	CourseController *courseController
)

func init() {
	CourseController = &courseController{}
}

func (controller *courseController) Create(c *gin.Context) {

	err := domain.CourseDao.Create("1", "2", "3")

	if err != nil {
		fmt.Println("I have an error after inserting the course", err.Error())
	}

	// Run domain
	c.String(202, "hello course controller")
}