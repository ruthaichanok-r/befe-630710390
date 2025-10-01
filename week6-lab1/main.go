package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type Student struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Email string  `json:"email"`
	Year  int     `json:"year"`
	GPA   float64 `json:"gpa"`
}

var students = []Student{
	{ID: "1", Name: "John Doe", Email: "john@example.com", Year: 3, GPA: 3.50},
	{ID: "2", Name: "Jane Smith", Email: "jane@example.com", Year: 2, GPA: 3.75},
}

func getStudents(c*gin.Context) {
	if yearQuery != "" {
		
		for _, Student := range students {
			if fmt.Sprint(student.Year) == yearQuery {

			}
		}
	}
}

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "healthy",
		})
	})

	api := r.Group("/api/v1")
	{
		api.GET("/students", getStudents)
		
		api.GET("/students/:id", getStudent)
		api.POST("/students", createStudent)
		api.PUT("/students/:id", updateStudent)
		api.DELETE("/students/:id", deleteStudent)
	}

	r.Run(":8080")
}*/