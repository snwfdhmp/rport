package rportserver

import (
	"fmt"
	"net/http"
	"time"

	"github.com/snwfdhmp/errlog"
	routing "github.com/snwfdhmp/ozzo-routing"
	"github.com/snwfdhmp/ozzo-routing/content"
	"github.com/snwfdhmp/ozzo-routing/cors"
	"github.com/spf13/afero"

	"github.com/sirupsen/logrus"
)

// Server represents interactions with a typical rport-server object
type Server interface {
	// Start loads storage from given path and
	// starts a blocking http listening process
	// on given port
	Start(serverPort int, storagePath string) error
}

type server struct {
	storage  *storage
	savePath string
}

// NewServer creates a new server object
func NewServer() Server {
	return &server{
		storage: newStorage(),
	}
}

func (s *server) Start(serverPort int, storagePath string) error {
	router := s.buildRouter()

	ok, err := afero.Exists(fs, storagePath)
	if err != nil {
		return fmt.Errorf("could not open file '%s': %s", storagePath, err)
	}
	if !ok {
		if err := newStorage().SaveTo(storagePath); err != nil {
			return fmt.Errorf("will not be able to save to '%s': %s", storagePath, err)
		}
	}

	s.storage, err = newStorageFrom(storagePath)
	if err != nil {
		return err
	}
	s.savePath = storagePath
	logrus.Infof("server: storage loaded from path '%s'.", storagePath)

	saveDelay := time.Second * 8
	go s.saveEvery(saveDelay)
	logrus.Infof("server: started auto-save every %.1f second.", saveDelay.Seconds())

	logrus.Infof("server: listening on port %d...", serverPort)
	panic(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", serverPort), router))
}

func (s *server) saveEvery(d time.Duration) {
	tick := time.NewTicker(d)

	for range tick.C {
		s.storage.SaveTo(s.savePath)
	}
}

func (s *server) buildRouter() *routing.Router {
	r := routing.New()

	r.Use(
		content.TypeNegotiator(content.JSON),
		cors.Handler(cors.Options{
			AllowOrigins: "*",
			AllowHeaders: "*",
			AllowMethods: "*",
		}),
	)

	r.Get("/<reportName>", s.handleGetReport)
	r.Post("/<reportName>", s.handleCreateReport)
	return r
}

func (s *server) handleCreateReport(c *routing.Context) error {
	reportName := c.Param("reportName")
	var payload map[string]interface{}
	if err := c.Read(&payload); errlog.Debug(err) {
		return err
	}

	createdReport := &report{Name: reportName, Payload: payload}

	if err := s.storage.Store(createdReport); err != nil {
		return err
	}

	return c.Write(createdReport)
}

func (s *server) handleGetReport(c *routing.Context) error {
	reportName := c.Param("reportName")

	items, err := s.storage.Get(reportName)
	if err != nil {
		return err
	}

	return c.Write(items)
}
