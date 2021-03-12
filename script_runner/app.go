package main

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

var port = "8080"

type TaskRequest struct {
	ImageName    string   `json:"image_name"`    // docker image (overiddes all other params)
	DockerParams []string `json:"docker_params"` // docekr parameters
	Version      string   `json:"version"`       // required python version
	Script       string   `json:"script"`        // script.py stuffed into a string
	Args         []string `json:"args"`          // script arguments
	Dependencies string   `json:"dependencies"`  // requirements.txt stuffed in a string
}

func runTask(task *TaskRequest) string {
	cmd := exec.Command("python", []string{"-c", task.Script}...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	return string(out)
}

func runDocker(image string, params []string) string {
	cmd := exec.Command("docker", []string{"pull", "-e", "X=123", image}...)
	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Infof("Failed to pull image: %s with error: %s\n", image, err.Error())
	}
	options := append([]string{"run", "--rm"}, params...)
	options = append(options, image)
	cmd = exec.Command("docker", options...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to run image: %s with error: %s\n", image, err.Error())
	}
	return string(out)
}

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	router := echo.New()

	router.GET("/docker-status", func(c echo.Context) error {
		cmd := exec.Command("/usr/local/bin/docker", []string{"ps", "-a"}...)
		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err)
		}
		log.Infof("combined out:\n%s\n", c)
		return c.String(http.StatusOK, string(out))
	})

	router.POST("/script/run", func(c echo.Context) error {
		log.Println("Request")
		u := new(TaskRequest)
		if err := c.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if u.ImageName != "" {
			out := runDocker(u.ImageName, u.DockerParams)
			return c.JSON(http.StatusOK, map[string]interface{}{
				"request": u,
				"output":  out,
			})
		} else {
			out := runTask(u)
			return c.JSON(http.StatusOK, map[string]interface{}{
				"request": u,
				"output":  out,
			})
		}

	})

	log.Infoln([]string{"AtpService-endpoint started"})
	router.Start(fmt.Sprintf(":%s", port))
}
