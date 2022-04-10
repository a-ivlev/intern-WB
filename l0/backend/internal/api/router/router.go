package router

import (
	"html/template"

	"intern-WB/l0/backend/internal/app/repository"

	"net/http"
)

type Router struct {
	*http.ServeMux
	store *repository.OrdersRepo
}

func NewRouter(st *repository.OrdersRepo) *Router {
	r := &Router{
		ServeMux: http.NewServeMux(),
		store:       st,
	}

	r.Handle("/", http.HandlerFunc(r.Home))

	return r
}

func (rt *Router) Home(w http.ResponseWriter, r *http.Request) {

	tmpl, _ := template.ParseFiles("backend/internal/web/template/home.template.html")

	suid := r.FormValue("orderUID")

	order, err := rt.store.GetOrder(r.Context(), suid)
	if err != nil {

	}

	tmpl.Execute(w, order)
}
