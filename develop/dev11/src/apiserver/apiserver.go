package apiserver

import (
	"dev11/src/store"
	"fmt"
	"net/http"
	"github.com/sirupsen/logrus"
)

// APIServer ...
type APIServer struct {
	config    *Config
	logger    *logrus.Logger
	router    *http.ServeMux
	routerLog http.HandlerFunc
	store     *store.Store
}

// New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: http.NewServeMux(),
		store: store.New(),
	}
}

func (s *APIServer) Start() error {
	s.configureRouter()
	s.configureRouterLog()

	s.logger.Info("starting server")

	return http.ListenAndServe(s.config.Host+s.config.Port, s.routerLog)
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/create_event", s.createEvent)
	s.router.HandleFunc("/update_event", s.updateEvent)
	s.router.HandleFunc("/delete_event", s.deleteEvent)
	s.router.HandleFunc("/events_for_day", s.eventsForDay)
	s.router.HandleFunc("/events_for_week", s.eventsForWeek)
	s.router.HandleFunc("/events_for_month", s.eventsForMonth) 
}

func (s *APIServer) configureRouterLog() {
	s.routerLog = http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			s.router.ServeHTTP(w, r)
			logMessage := fmt.Sprintf("Method: %s URI: %s", r.Method, r.RequestURI)
			s.logger.Infoln(logMessage)
		},
	)
}