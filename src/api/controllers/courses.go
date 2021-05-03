package controllers

import (
	"github.com/dembygenesis/student_enrollment_exam/src/api/domain"
	"github.com/dembygenesis/student_enrollment_exam/src/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
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
	var body domain.CreateCourse

	if err := c.ShouldBindJSON(&body); err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "body must conform to format",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		utils.RespondError(c, apiErr)
		return
	}

	err := domain.CourseDao.Create(body.Name, body.Professor, body.Description)

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "Error when attempting to insert course data : " + err.Error(),
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		utils.RespondError(c, apiErr)
		return
	}

	utils.Respond(c, http.StatusOK, "Successfully created the course!")
	return
}