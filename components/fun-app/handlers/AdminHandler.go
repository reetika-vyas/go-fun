package handlers

import (
	util2 "github.com/amanhigh/go-fun/common/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdminHandler struct {
	Shutdown *util2.GracefullShutdown `inject:""`
}

/**
go test fun-app_test.go fun-app.go -coverprofile=coverage.out
curl http://localhost:8080/admin/stop
go tool cover -func=coverage.out
*/
func (self *AdminHandler) Stop(c *gin.Context) {
	self.Shutdown.Stop()
	//TODO:Add check to prevent accidental stop
	c.JSON(http.StatusOK, "Stop Started")
}
