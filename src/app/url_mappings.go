package app

import "github.com/dembygenesis/student_enrollment_exam/src/api/controllers"

func mapUrls() {
	// Student
	router.GET("/student", controllers.StudentController.GetStudents)
	router.POST("/student", controllers.StudentController.Create)
	router.POST("/student/enroll", controllers.StudentController.Enroll)

	// Course
	router.POST("/course", controllers.CourseController.Create)
	router.DELETE("/course/:course_id", controllers.CourseController.DeleteCourse)
}