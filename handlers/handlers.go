package handlers

import (
	"main/components"
	"net/http"
)

func ListAllData(w http.ResponseWriter, r *http.Request) {
	component := components.Base("GoProject")
	component.Render(r.Context(), w)
}
