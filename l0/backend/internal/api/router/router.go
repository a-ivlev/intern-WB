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
	// suid := r.URL.Query().Get("uid")
	// if suid == "" {
	// 	http.Error(w, "bad request", http.StatusBadRequest)
	// 	return
	// }
	tmpl, _ := template.ParseFiles("backend/internal/web/template/home.template.html")
	// if err != nil {
	// 	//tmpl.ExecuteTemplate(w, "Home", "template parsing error")
	// 	log.Println("template parsing error")
	// }

	suid := r.FormValue("orderUID")

	order, err := rt.store.GetOrder(r.Context(), suid)
	if err != nil {

		//fmt.Printf("error order uid %s: %s\n", suid, err)
		//http.Error(w, "error when creating", http.StatusInternalServerError)
		//return
	}

	//_ = json.NewEncoder(w).Encode(order)
	tmpl.Execute(w, order)
}
