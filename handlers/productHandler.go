package handlers

import (
	"encoding/json"
	"fmt"
	"go-mongo/models"
	mongocon "go-mongo/mongo-con"
	"go-mongo/repository/impl"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateProduct(rw http.ResponseWriter, rq *http.Request) {
	var product models.Product
	err := json.NewDecoder(rq.Body).Decode(&product)
	if err != nil {
		responseJsonData(rw, http.StatusInternalServerError, map[string]string{"Eror": "Parse json when get data from API " + err.Error()})
		return
	}
	pRepo := impl.NewProductRepo(mongocon.ConnectMongoDB().DB.Database("product_api"))
	err = pRepo.CreateProduct(&product)
	if err != nil {
		responseJsonData(rw, http.StatusInternalServerError, map[string]string{"Eror": "Can not save data to DB " + err.Error()})
		return
	}
	responseJsonData(rw, http.StatusOK, product)

}

func GetAllProduct(rw http.ResponseWriter, rq *http.Request) {
	pRepo := impl.NewProductRepo(mongocon.ConnectMongoDB().DB.Database("product_api"))
	listProduct, err := pRepo.FindAll()

	if err != nil {
		responseJsonData(rw, http.StatusInternalServerError, map[string]string{"Eror": "Can not get data from db " + err.Error()})
		return
	}
	responseJsonData(rw, http.StatusOK, listProduct)
}

func GetProductByName(rw http.ResponseWriter, rq *http.Request) {
	pRepo := impl.NewProductRepo(mongocon.ConnectMongoDB().DB.Database("product_api"))
	var product map[string]string
	json.NewDecoder(rq.Body).Decode(&product)
	fmt.Printf("Data %v\n", product["name"])
	listProduct, err := pRepo.GetProductByName(product["name"])

	if err != nil {
		responseJsonData(rw, http.StatusInternalServerError, map[string]string{"Eror": "Can not get data from db " + err.Error()})
		return
	}
	responseJsonData(rw, http.StatusOK, listProduct)
}

func UpdateProduct(rw http.ResponseWriter, rq *http.Request) {
	var product models.Product
	err := json.NewDecoder(rq.Body).Decode(&product)
	if err != nil {
		responseJsonData(rw, http.StatusInternalServerError, map[string]string{"Eror": "Parse json when get data from API " + err.Error()})
		return
	}
	params := mux.Vars(rq)
	id := params["id"]

	pRepo := impl.NewProductRepo(mongocon.ConnectMongoDB().DB.Database("product_api"))
	err = pRepo.EditProductById(&product, id)
	if err != nil {
		responseJsonData(rw, http.StatusInternalServerError, map[string]string{"Eror": "Can not update product " + err.Error()})
		return
	}
	responseJsonData(rw, http.StatusOK, product)

}

func responseJsonData(rw http.ResponseWriter, status int, object interface{}) error {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	err := json.NewEncoder(rw).Encode(object)
	return err
}
