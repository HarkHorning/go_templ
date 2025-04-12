package handlers

import (
	"main/components"
	"main/models"
	"net/http"
)

func ListAllData(w http.ResponseWriter, r *http.Request) {
	component := components.Base(models.Topics)
	component.Render(r.Context(), w)
}
