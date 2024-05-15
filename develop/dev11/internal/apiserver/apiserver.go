package apiserver

import (
	"github.com/zetsbu0/wbschool_exam_L2/develop/dev11/internal/calendar"
	"log"
	"net/http"
	"time"
)

type Storage interface {
	CreateEvent(event *calendar.Event)
	UpdateEvent(ID int, Time time.Time, Name string) error
	DeleteEvent(ID int) (*calendar.Event, error)
	DailyEvents() []calendar.Event
	WeeklyEvents() []calendar.Event
	MonthlyEvents() []calendar.Event
}

type ApiServer struct {
	config     *Config
	router     *http.ServeMux
	middleware *Middleware
	calendar   Storage
}

func New(config *Config) *ApiServer {
	router := http.NewServeMux()
	middleware := NewMiddleware(router)
	cal := calendar.NewCalendar()

	return &ApiServer{
		config:     config,
		router:     router,
		middleware: middleware,
		calendar:   cal,
	}
}

func (s *ApiServer) Start() error {
	log.Println("Starting API Server on port", s.config.BindAddr)
	s.configureRouter()
	return http.ListenAndServe(s.config.BindAddr, s.middleware)
}

func (s *ApiServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello)

	s.router.HandleFunc("POST /create_event", s.createEvent)
	s.router.HandleFunc("POST /update_event", s.updateEvent)
	s.router.HandleFunc("POST /delete_event", s.deleteEvent)

	s.router.HandleFunc("GET /events_for_day", s.showDailyEvents)
	s.router.HandleFunc("GET /events_for_week", s.showWeeklyEvents)
	s.router.HandleFunc("GET /events_for_month", s.showMonthlyEvents)

}

//POST /create_event
//POST /update_event
//POST /delete_eventj
//GET /events_for_day
//GET /events_for_week
//GET /events_for_month
