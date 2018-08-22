package http

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/drive/v3"
)

type filteredFileListRequest struct {
	Filter string `json:"filter"`
}

type errorResponse struct {
	Error string `json:"error"`
}

type filesResponse struct {
	Files interface{} `json:"files"`
}

// Server is a GIN rest server
type Server struct {
	drive  *drive.Service
	server *gin.Engine
	port   int
}

// New creates new instance of server
func New(port int, drive *drive.Service) *Server {
	return &Server{
		port:  port,
		drive: drive,
	}
}

// Start starts http rest-server
func (s *Server) Start() error {
	s.server = gin.New()
	corsConf := cors.DefaultConfig()
	corsConf.AllowAllOrigins = true
	s.server.Use(cors.New(corsConf))
	// сюда еще можно добавить проверку флага из конфига на интерактивный режим

	s.server.POST("/files", s.postGetFilteredFileList)

	return s.server.Run(fmt.Sprintf(":%d", s.port))
}

func (s *Server) postGetFilteredFileList(c *gin.Context) {
	req := new(filteredFileListRequest)

	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{"wrong request"})
		return
	}

	r, err := s.drive.Files.
		List().
		Q(fmt.Sprintf("name contains '%s'", req.Filter)).
		Fields("nextPageToken, files(createdTime, name, webViewLink)").
		Do()

	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{err.Error()})
		return
	}

	c.JSON(200, filesResponse{r.Files})
}
