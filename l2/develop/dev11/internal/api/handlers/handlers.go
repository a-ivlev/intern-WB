package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"l2/develop/dev11/internal/app/event"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/google/uuid"
)

var (
	ErrBadMetod    = errors.New("bad method")
	ErrBadRequest  = errors.New("bad request")
	ErrServer      = errors.New("server error")
	ErrCreateEvent = errors.New("error when creating")
	ErrUpdateEvent = errors.New("error when updating")
	ErrDeleteEvent = errors.New("error when deleting")
	ErrParseForm   = errors.New("parse form error")
	ErrNotFound    = errors.New("not found")
)

type Router struct {
	*http.ServeMux
	*event.EventStore
}

func NewRouter(es *event.EventStore) *Router {
	r := &Router{
		ServeMux:   http.NewServeMux(),
		EventStore: es,
	}

	r.Handle("/create_event", r.LogMiddleware(http.HandlerFunc(r.CreateEvent)))
	r.Handle("/update_event", r.LogMiddleware(http.HandlerFunc(r.UpdateEvent)))
	r.Handle("/delete_event", r.LogMiddleware(http.HandlerFunc(r.DeleteEvent)))
	r.Handle("/events_for_day", r.LogMiddleware(http.HandlerFunc(r.EventsForDay)))
	r.Handle("/events_for_week", r.LogMiddleware(http.HandlerFunc(r.EventsForWeek)))
	r.Handle("/events_for_mount", r.LogMiddleware(http.HandlerFunc(r.EventsForMount)))
	return r
}

type Event struct {
	ID          uuid.UUID
	UserID      int64
	EventDate   time.Time
	Description string
}

type params struct {
	raw url.Values
}

func (p *params) GetString(field string) string {
	return p.raw.Get(field)
}

func (p *params) GetInt64(field string) int64 {
	rawInt := p.raw.Get(field)
	resInt, err := strconv.ParseInt(rawInt, 0, 10)
	if err != nil {
		log.Printf("[ ERROR ] %s: %s\n", ErrParseForm.Error(), err)
	}

	return resInt
}

func (p *params) GetData(field string) time.Time {
	rawData := p.raw.Get(field)
	resData, err := time.Parse("2006-01-02", rawData)
	if err != nil {
		log.Printf("[ ERROR ] %s: %s\n", ErrParseForm.Error(), err)
	}

	return resData
}

func (p *params) GetUUID(field string) uuid.UUID {
	rawData := p.raw.Get(field)
	resUUID, err := uuid.Parse(rawData)
	if err != nil {
		log.Printf("[ ERROR ] %s: %s\n", ErrParseForm.Error(), err)
	}

	return resUUID
}

type EventsList []Event

func (el EventsList) String() string {
	v, _ := json.Marshal(el)
	return string(v)
}

func (rt *Router) LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			if err := r.ParseForm(); err != nil {
				log.Printf("[ ERROR ] %s: %s\n", ErrParseForm.Error(), err)
				http.Error(w, fmt.Sprintf("{\"error\": \"%s\"}\n", ErrParseForm.Error()), http.StatusBadRequest)
				return
			}

			param := &params{r.Form}

			userID := param.GetInt64("user_id")
			eventDate := param.GetData("date")
			description := param.GetString("description")

			log.Printf("[ INFO ] metod: %s, remote addr: %s, request: %s, user_id: %d, date: %s, description: %s.\n", r.Method, r.RemoteAddr, r.RequestURI, userID, eventDate, description)

			next.ServeHTTP(w, r)
		},
	)
}

func (rt *Router) CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("[ WARNING ] %s\n", ErrBadMetod.Error())
		http.Error(w, fmt.Sprintf("{\"error\": \"%s\"}\n", ErrBadMetod.Error()), http.StatusMethodNotAllowed)
		return
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Printf("[ ERROR ] %s: %s\n", ErrCreateEvent.Error(), err)
		}
	}()

	if err := r.ParseForm(); err != nil {
		log.Printf("[ ERROR ] %s: %s\n", ErrParseForm.Error(), err)
		http.Error(w, fmt.Sprintf("{\"error\": \"%s\"}\n", ErrParseForm.Error()), http.StatusBadRequest)
		return
	}

	param := &params{r.Form}

	eventCore := event.Event{
		UserID:      param.GetInt64("user_id"),
		EventDate:   param.GetData("date"),
		Description: param.GetString("description"),
	}

	creteEvent, err := rt.EventStore.CreateEvent(r.Context(), eventCore)
	if err != nil {
		log.Printf("[ ERROR ] %s: %s\n", ErrCreateEvent.Error(), err)
		http.Error(w, fmt.Sprintf("{\"error\": \"%s\"}\n", ErrCreateEvent.Error()), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err = fmt.Fprintf(w, "{\"result\":%s}\n", creteEvent); err != nil {
		log.Printf("[ ERROR ] %s: %s\n", ErrServer.Error(), err)
		http.Error(w, fmt.Sprintf("{\"error\": \"%s\"}\n", ErrServer.Error()), http.StatusInternalServerError)
	}
}

func (rt *Router) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("[ WARNING ] %s\n", ErrBadMetod.Error())
		http.Error(w, fmt.Sprintf("{\"error\": \"%s\"}\n", ErrBadMetod.Error()), http.StatusMethodNotAllowed)
		return
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Printf("[ ERROR ] %s: %s\n", ErrCreateEvent.Error(), err)
		}
	}()

	if err := r.ParseForm(); err != nil {
		log.Printf("[ ERROR ] %s: %s\n", ErrParseForm.Error(), err)
		http.Error(w, fmt.Sprintf("{\"error\": \"%s\"}\n", ErrParseForm.Error()), http.StatusBadRequest)
		return
	}

	param := &params{
		r.Form,
	}

	eventCore := event.Event{
		ID:          param.GetUUID("id"),
		UserID:      param.GetInt64("user_id"),
		EventDate:   param.GetData("date"),
		Description: param.GetString("description"),
	}

	updEvent, err := rt.EventStore.UpdateEvent(r.Context(), eventCore)
	if err != nil {
		log.Printf("[ ERROR ] %s: %s\n", ErrUpdateEvent.Error(), err)
		http.Error(w, fmt.Sprintf("{\"error\": \"%s\"}\n", ErrUpdateEvent.Error()), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err = fmt.Fprintf(w, "{\"result\":%s}\n", updEvent); err != nil {
		log.Printf("[ ERROR ] %s: %s\n", ErrServer.Error(), err)
		http.Error(w, fmt.Sprintf("{\"error\": \"%s\"}\n", ErrServer.Error()), http.StatusInternalServerError)
	}
}

func (rt *Router) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("[ WARNING ] %s\n", ErrBadMetod.Error())
		http.Error(w, fmt.Sprintf("{\"error\": \"%s\"}\n", ErrBadMetod.Error()), http.StatusMethodNotAllowed)
		return
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Printf("[ ERROR ] %s: %s\n", ErrCreateEvent.Error(), err)
		}
	}()

	if err := r.ParseForm(); err != nil {
		log.Printf("[ ERROR ] %s: %s\n", ErrParseForm.Error(), err)
		http.Error(w, fmt.Sprintf("{\"error\": \"%s\"}\n", ErrParseForm.Error()), http.StatusBadRequest)
		return
	}

	param := &params{
		r.Form,
	}

	eventID := param.GetUUID("id")

	err := rt.EventStore.DeleteEvent(r.Context(), eventID)
	if err != nil {
		log.Printf("[ ERROR ] %s: %s\n", ErrDeleteEvent.Error(), err)
		http.Error(w, fmt.Sprintf("{\"error\": \"%s\"}\n", ErrDeleteEvent.Error()), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := fmt.Fprintf(w, "{\"result\":\"event id %s deleted.\"}\n", eventID); err != nil {
		log.Printf("[ ERROR ] %s: %s\n", ErrServer.Error(), err)
		http.Error(w, fmt.Sprintf("{\"error\": \"%s\"}\n", ErrServer.Error()), http.StatusInternalServerError)
	}
}

func (rt *Router) EventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Printf("[ WARNING ] %s\n", ErrBadMetod.Error())
		http.Error(w, fmt.Sprintf("{\"error\": \"%s\"}\n", ErrBadMetod.Error()), http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		log.Printf("[ ERROR ] %s: %s\n", ErrParseForm.Error(), err)
		http.Error(w, fmt.Sprintf("{\"error\": \"%s\"}\n", ErrParseForm.Error()), http.StatusBadRequest)
		return
	}

	param := &params{
		r.Form,
	}
	userID := param.GetInt64("user_id")
	eventDate := param.GetData("date")

	eventsList, err := rt.EventStore.SearchEvent(r.Context(), userID, eventDate, event.EventsDay)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, fmt.Sprintf("{\"error\": \"%s\"}\n", ErrNotFound), http.StatusNotFound)
		} else {
			http.Error(w, "error when reading", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err = fmt.Fprintf(w, "{\"result\":%s}\n", eventsList); err != nil {
		log.Printf("[ ERROR ] %s: %s", ErrServer.Error(), err)
		http.Error(w, fmt.Sprintf("{\"error\": \"%s\"}\n", ErrServer.Error()), http.StatusInternalServerError)
	}
}

func (rt *Router) EventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Printf("[ WARNING ] %s", ErrBadMetod.Error())
		http.Error(w, fmt.Sprintf("{\"error\": \"%s\"}\n", ErrBadMetod.Error()), http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		log.Printf("[ ERROR ] %s: %s\n", ErrParseForm.Error(), err)
		http.Error(w, fmt.Sprintf("{\"error\": \"%s\"}\n", ErrParseForm.Error()), http.StatusBadRequest)
		return
	}

	param := &params{
		r.Form,
	}
	userID := param.GetInt64("user_id")
	eventDate := param.GetData("date")

	eventsList, err := rt.EventStore.SearchEvent(r.Context(), userID, eventDate, event.EventsWeek)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, fmt.Sprintf("{\"error\": \"%s\"}\n", ErrNotFound), http.StatusNotFound)
		} else {
			http.Error(w, "error when reading", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err = fmt.Fprintf(w, "{\"result\":%s}\n", eventsList); err != nil {
		log.Printf("[ ERROR ] %s: %s", ErrServer.Error(), err)
		http.Error(w, fmt.Sprintf("{\"error\": \"%s\"}\n", ErrServer.Error()), http.StatusInternalServerError)
	}
}

func (rt *Router) EventsForMount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Printf("[ WARNING ] %s", ErrBadMetod.Error())
		http.Error(w, fmt.Sprintf("{\"error\": \"%s\"}\n", ErrBadMetod.Error()), http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		log.Printf("[ ERROR ] %s: %s", ErrParseForm.Error(), err)
		http.Error(w, fmt.Sprintf("{\"error\": \"%s\"}\n", ErrParseForm.Error()), http.StatusBadRequest)
		return
	}

	param := &params{
		r.Form,
	}
	userID := param.GetInt64("user_id")
	eventDate := param.GetData("date")

	eventsList, err := rt.EventStore.SearchEvent(r.Context(), userID, eventDate, event.EventsMount)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, fmt.Sprintf("{\"error\": \"%s\"}\n", ErrNotFound), http.StatusNotFound)
		} else {
			http.Error(w, "error when reading", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err = fmt.Fprintf(w, "{\"result\":%s}\n", eventsList); err != nil {
		log.Printf("[ ERROR ] %s: %s", ErrServer.Error(), err)
		http.Error(w, fmt.Sprintf("{\"error\": \"%s\"}\n", ErrServer.Error()), http.StatusInternalServerError)
	}
}
