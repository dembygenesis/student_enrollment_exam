package app

import "github.com/dembygenesis/student_enrollment_exam/src/api/controllers"

func mapUrls() {
	// Student
	router.POST("/student", controllers.StudentController.Create)
	router.POST("/student/enroll", controllers.StudentController.Enroll)

	// Course
	router.POST("/course", controllers.CourseController.Create)
}