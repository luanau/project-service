package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/benc-uk/dapr-store/pkg/dapr"
	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func PostHomePage(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(200, gin.H{"message": string(value)})
}

func QueryString(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func PathParameters(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

type Project struct {
	ProjectId                 string  `json:"projectId" binding:"required"`
	Name                      string  `json:"name" binding:"required"`
	StartDate                 string  `json:"startDate"`
	EndDate                   string  `json:"endDate"`
	EstimatedDurationInDays   uint64  `json:"estimatedDurationInDays"`
	EstimatedCost             float64 `json:"estimatedCost"`
	AnnualInterestRateOffered float64 `json:"annualInterestRateOffered"`
	ProjectManager            string  `json:"projectManager"`
}

func SaveProject(c *gin.Context) {
	proj := Project{}
	if err := c.ShouldBind(&proj); err == nil {
		SaveProjectToStore(&proj)
		c.JSON(http.StatusCreated, proj)
	} else {
		c.String(http.StatusBadRequest, `the body should be a project structure`)
	}
}

func SaveProjectToStore(project *Project) error {
	// TODO - Check project.ProjectId has proper value
	prob := helper.SaveState("statestore", project.ProjectId, *project)
	if prob != nil {
		log.Printf("### Error!, unable to save project wid id '%s'", project.ProjectId)
		return prob
	}

	return nil
}

var store string = "statestore"
var helper *dapr.Helper

func main() {
	helper = dapr.NewHelper("project-service")
	if helper == nil {
		os.Exit(1)
	}

	r := gin.Default()
	r.Use(gin.Logger())
	r.GET("/", HomePage)
	r.POST("/", SaveProject)
	r.GET("/query", QueryString)
	r.GET("/path/:name/:age", PathParameters)
	r.Run()
}
