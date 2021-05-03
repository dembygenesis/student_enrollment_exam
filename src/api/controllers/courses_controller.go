package controllers

import (
	"github.com/dembygenesis/student_enrollment_exam/src/api/domain"
	"github.com/dembygenesis/student_enrollment_exam/src/api/services"
	"github.com/dembygenesis/student_enrollment_exam/src/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type courseController struct {

}

var (
	CourseController *courseController
)

func init() {
	CourseController = &courseController{}
}

func (controller *courseController) DeleteCourse(c *gin.Context) {
	// Validate params
	courseId, err := strconv.ParseInt(c.Param("course_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "course_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	// Perform delete
	err = services.CourseService.DeleteCourse(int(courseId))

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "Error when attempting to delete the course: " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code:       "bad_request",
		}

		utils.RespondError(c, apiErr)
		return
	}

	utils.Respond(c, http.StatusOK, "Successfully deleted the course!")
	return
}

func (controller *courseController) GetEnrolledStudents(c *gin.Context) {
	// Validate params
	courseId, err := strconv.ParseInt(c.Param("course_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "course_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	// Perform fetch
	students, err := services.CourseService.GetEnrolledStudents(int(courseId))

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "Error when attempting to fetch the enrolled students: " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code:       "bad_request",
		}

		utils.RespondError(c, apiErr)
		return
	}

	utils.Respond(c, http.StatusOK, students)
	return
}

func (controller *courseController) Create(c *gin.Context) {
	var body domain.CreateCourse

	// Validate params
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
	err := services.CourseService.Create(body.Name, body.Professor, body.Description)

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "Error when attempting to insert course data : " + err.Error(),
			StatusCode: http.StatusInternalServerError,
			Code:       "bad_request",
		}

		utils.RespondError(c, apiErr)
		return
	}

	utils.Respond(c, http.StatusOK, "Successfully created the course!")
	return
}