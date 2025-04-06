package handlers

import (
	"go_templ/components"
	"net/http"
)

func ListAllData(w http.ResponseWriter, r *http.Request) {
	component := components.Base()
	component.Render(r.Context(), w)
}
