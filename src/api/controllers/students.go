package controllers

import (
	"github.com/dembygenesis/student_enrollment_exam/src/api/domain"
	"github.com/dembygenesis/student_enrollment_exam/src/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type studentController struct {

}

var (
	StudentController *studentController
)

func init() {
	StudentController = &studentController{}
}

func (controller *studentController) Create(c *gin.Context) {
	var body domain.CreateStudent

	if err := c.ShouldBindJSON(&body); err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "body must conform to format",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		utils.RespondError(c, apiErr)
		return
	}

	err := domain.StudentDao.Create(body.Name, body.Email, body.Phone)

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "Error when attempting to insert student data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code:       "bad_request",
		}

		utils.RespondError(c, apiErr)
		return
	}

	utils.Respond(c, http.StatusOK, "Successfully created the student!")
	return
}