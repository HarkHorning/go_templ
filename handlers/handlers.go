package handlers

import (
	"main/components"
	"net/http"
)

func ListAllData(w http.ResponseWriter, r *http.Request) {
	component := components.Base()
	component.Render(r.Context(), w)
}
