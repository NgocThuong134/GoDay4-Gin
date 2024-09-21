package main

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)
type Student struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Gender string `json:"gender"`
	Email string `json:"email"`
	Age int `json:"age"`
	Math float32 `json:"math"`
	English float32 `json:"english"`
}
var students = []Student{
	{ID: 1, Name: "Nguyen Ngoc Thuong", Gender: "Nam",Email: "thuong123@gmail.com", Age: 20, Math: 9.8, English: 8.0},
	{ID: 2, Name: "Pham Ngoc Anh", Gender: "Nu",Email: "anh1234@gmail.com", Age: 21, Math: 7.8, English: 5 },
	{ID: 3, Name: "Tan Thieu Ta", Gender: "Nam", Email: "thieuta@gmail.com", Age: 22, Math: 4, English: 6.8},
}
func getStudents(c *gin.Context){
	c.JSON(http.StatusOK, students)
}

func getStudentDetail(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, student := range students {
		if student.ID == id {
			c.JSON(http.StatusOK, student)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Student not found"})
}

func addStudent(c *gin.Context){
	var newStudent Student
	if err := c.ShouldBindJSON(&newStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newStudent.ID = len(students) + 1
	students = append(students, newStudent)
	c.JSON(http.StatusCreated, newStudent)
}

func updateStudent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedStudent Student
	if err := c.ShouldBindJSON(&updatedStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, student := range students {
		if student.ID == id {
			students[i].Name = updatedStudent.Name
			students[i].Age = updatedStudent.Age
			students[i].Gender = updatedStudent.Gender
			students[i].Email = updatedStudent.Email
			students[i].Math = updatedStudent.Math
			students[i].English = updatedStudent.English
			c.JSON(http.StatusOK, students[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Student not found"})
}

func deleteStudent(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	for i, student := range students {
		if student.ID == id {
			students = append(students[:i], students[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Student deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Student not found"})
}
func main() {
  r := gin.Default()
    r.GET("/get-students", getStudents)
	r.GET("/get-student-detail/:id", getStudentDetail)
	r.POST("/add-student", addStudent)
	r.PUT("/update-student/:id", updateStudent)
	r.DELETE("/delete-student/:id", deleteStudent)

	r.Run() 
}