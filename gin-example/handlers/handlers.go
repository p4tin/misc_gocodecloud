package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	example "misc/gin-example"
)

type HealthResponse struct {
	TimeStamp   string `json:"timestamp"`
	AppName     string `json:"appName"`
	Branch      string `json:"githubBranch"`
	BuildNumber string `json:"jenkinsBuildNumber"`
	GitHash     string `json:"githubHash"`
	Environment string `json:"environment"`
	Status      string `json:"status"`
}

type Service struct {
	Config example.Config
}

func CreateService(config example.Config) Service {
	return Service{
		Config: config,
	}
}

func (s Service) HealthHandler(c *gin.Context) {

	serv := HealthResponse{
		TimeStamp:   time.Now().UTC().Format(time.RFC822Z),
		Status:      "ok",
		AppName:     s.Config.AppName,
		Branch:      s.Config.Branch,
		BuildNumber: s.Config.BuildNumber,
		GitHash:     s.Config.GitHash,
		Environment: s.Config.Environment,
	}

	c.JSON(http.StatusOK, gin.H{"status": serv})
}
