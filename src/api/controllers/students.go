package controllers

import (
	"github.com/dembygenesis/student_enrollment_exam/src/api/domain"
	"github.com/dembygenesis/student_enrollment_exam/src/api/services"
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

func (controller *studentController) Enroll(c *gin.Context) {
	// Validate params
	var body domain.EnrollStudent

	if err := c.ShouldBindJSON(&body); err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "body must conform to format",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		utils.RespondError(c, apiErr)
		return
	}

	// Perform enroll
	err := services.StudentService.Enroll(body.StudentId, body.CourseId)

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "Error when attempting to enroll the student to the course : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code:       "bad_request",
		}

		utils.RespondError(c, apiErr)
		return
	}

	utils.Respond(c, http.StatusOK, "Successfully enrolled the student to the course!")
	return
}

func (controller *studentController) GetStudents(c *gin.Context) {
	// Perform fetch
	students, err := services.StudentService.GetStudents()

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "Error when attempting to fetch the student data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code:       "bad_request",
		}

		utils.RespondError(c, apiErr)
		return
	}

	utils.Respond(c, http.StatusOK, students)
}

func (controller *studentController) Create(c *gin.Context) {
	// Validate params
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

	// Perform create
	err := services.StudentService.Create(body.Name, body.Email, body.Phone)

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