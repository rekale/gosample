package users

import (
	"big_projects/counts"
	"big_projects/pagination"
	"html/template"
	"net/http"
	"strconv"

	"github.com/tokopedia/sqlt"
)

type handler struct {
	db   *sqlt.DB
	view *template.Template
}

func NewHandler(db *sqlt.DB, view *template.Template) handler {
	return handler{db, view}
}

// Home handler
func (h *handler) Home(w http.ResponseWriter, r *http.Request) {
	var users []User
	var isSearch bool
	var paginate pagination.SimplePagination

	name := r.FormValue("name")
	lastID, _ := strconv.Atoi(r.FormValue("last_id"))

	if name == "" {
		users, paginate = GetUser(h.db, 10, lastID)
	} else {
		users, paginate = GetUserByName(h.db, name, 10, lastID)
		isSearch = true
	}

	data := struct {
		IsSearch  bool
		Users     []User
		CountView int
		Paginate  *pagination.SimplePagination
		Name      string
	}{
		isSearch,
		users,
		counts.Get(),
		&paginate,
		name,
	}

	h.view.ExecuteTemplate(w, "index.html", data)
}
