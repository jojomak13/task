package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"task/core"
	"task/models"
	"task/requests"
	"task/utils"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var data requests.ProductRequest

	json.NewDecoder(r.Body).Decode(&data)

	if err := core.NewValidator(data); err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]interface{}{
			"errors": err,
		})
		return
	}

	product, err := models.CreateProduct(core.DB, data)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var data requests.ProductRequest

	json.NewDecoder(r.Body).Decode(&data)

	if err := core.NewValidator(data); err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]interface{}{
			"errors": err,
		})
		return
	}

	product, err := models.UpdateProduct(core.DB, id, data)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	err = models.DeleteProduct(core.DB, id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Product deleted successfully"})
}

func ListProducts(w http.ResponseWriter, r *http.Request) {
	products, err := models.ListProducts(core.DB)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, products)
}

func GetProductSales(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	var fromDate, toDate time.Time
	var err error

	fromDateStr := query.Get("from_date")
	if fromDateStr != "" {
		fromDate, err = time.Parse("2006-01-02", fromDateStr)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid from_date format. Use YYYY-MM-DD")
			return
		}
	}

	toDateStr := query.Get("to_date")
	if toDateStr != "" {
		toDate, err = time.Parse("2006-01-02", toDateStr)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid to_date format. Use YYYY-MM-DD")
			return
		}
		// Set to end of day
		toDate = toDate.Add(24*time.Hour - time.Second)
	}

	username := query.Get("username")

	sales, err := models.GetProductSales(core.DB, fromDate, toDate, username)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"data": sales,
	})
}
