package apiserver

import (
	"encoding/json"
	"fmt"
	"github.com/zetsbu0/wbschool_exam_L2/develop/dev11/internal/calendar"
	"github.com/zetsbu0/wbschool_exam_L2/develop/dev11/internal/utils"
	"io"
	"log"
	"net/http"
	"strconv"
)

func (s *ApiServer) handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func (s *ApiServer) createEvent(w http.ResponseWriter, r *http.Request) {
	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/x-www-form-urlencoded" {
		http.Error(w, "invalid data", http.StatusBadRequest)
		return
	}

	event := &calendar.Event{}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("error happened :(", err)
		http.Error(w, "Error parsing body", 400)
		return
	}

	err = json.Unmarshal(body, &event)
	if err != nil {
		log.Println("error happened :(", err)
		http.Error(w, "Error unmarshalling json", http.StatusInternalServerError)
		return
	}

	s.calendar.CreateEvent(event)

	utils.SendResult(w, []byte("new event created"))
}

func (s *ApiServer) updateEvent(w http.ResponseWriter, r *http.Request) {
	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/x-www-form-urlencoded" {
		http.Error(w, "invalid data", http.StatusBadRequest)
		return
	}

	ID, Time, Name, err := utils.ParseUpdateRequest(r)
	if err != nil {
		log.Println("error happened :(", err)
		http.Error(w, "Error parsing request", 400)
		return
	}

	err = s.calendar.UpdateEvent(ID, Time, Name)
	if err != nil {
		log.Println("error happened :(", err)
		http.Error(w, "Error updating event", http.StatusInternalServerError)
		return
	}

	utils.SendResult(w, []byte(fmt.Sprintf("event #%d updated", ID)))
}

func (s *ApiServer) deleteEvent(w http.ResponseWriter, r *http.Request) {
	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/x-www-form-urlencoded" {
		http.Error(w, "invalid data", http.StatusBadRequest)
		return
	}

	err := r.ParseForm()
	if err != nil {
		log.Println("error happened :(", err)
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	ID, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		log.Println("error happened :(", err)
		http.Error(w, "Error converting id", http.StatusInternalServerError)
		return
	}

	deleted, err := s.calendar.DeleteEvent(ID)
	if err != nil {
		log.Println("error happened :(", err)
		http.Error(w, "Error deleting data", http.StatusInternalServerError)
		return
	}

	utils.SendResult(w, []byte(fmt.Sprintf("event #%d (%s, %v) removed", deleted.ID, deleted.Name, deleted.Time)))

}

func (s *ApiServer) showDailyEvents(w http.ResponseWriter, r *http.Request) {

	data, err := calendar.SerializeEventSlice(s.calendar.WeeklyEvents())
	if err != nil {
		log.Println("error happened :(", err)
		http.Error(w, "Error serializing slice", http.StatusInternalServerError)
		return
	}

	utils.SendResult(w, data)
}

func (s *ApiServer) showWeeklyEvents(w http.ResponseWriter, r *http.Request) {
	data, err := calendar.SerializeEventSlice(s.calendar.WeeklyEvents())
	if err != nil {
		utils.SendError(w, err, http.StatusServiceUnavailable)
		return
	}

	utils.SendResult(w, data)
}

func (s *ApiServer) showMonthlyEvents(w http.ResponseWriter, r *http.Request) {
	data, err := calendar.SerializeEventSlice(s.calendar.MonthlyEvents())
	if err != nil {
		utils.SendError(w, err, http.StatusServiceUnavailable)
		return
	}

	utils.SendResult(w, data)
}
