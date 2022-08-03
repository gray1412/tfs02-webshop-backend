package controller

import (
	"net/http"
	"project-backend/brand/response"
	"project-backend/database"
	resultResponse "project-backend/util/response"
)

func getAllBrand() []response.Brand {
	brands := []response.Brand{}
	db := database.ConnectDB()
	sql := "SELECT b.id, b.name,bi.url,bi.alt FROM brands b, brand_images bi WHERE b.image_id = bi.id"
	db.Raw(sql).Scan(&brands)
	return brands
}
func GetAllBrand(w http.ResponseWriter, r *http.Request) {
	brand := getAllBrand()
	resultResponse.RespondWithJSON(w, 200, 1, "", brand)
}
