package controllers

import "github.com/gin-gonic/gin"

type studentController struct {

}

var (
	StudentController *studentController
)

func init() {
	StudentController = &studentController{}
}

func (controller *studentController) Create(c *gin.Context) {
	c.String(202, "hello student controller")
}