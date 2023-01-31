package handlers

import (
	"encoding/json"
	propertiesdto "housy/dto/properties"
	dto "housy/dto/result"
	"housy/models"
	"housy/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerProperty struct {
	PropertyRepository repositories.PropertyRepository
}

func HandlerProperty(PropertyRepository repositories.PropertyRepository) *handlerProperty {
	return &handlerProperty{PropertyRepository}
}

func (h *handlerProperty) FindProperties(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	properties, err := h.PropertyRepository.FindProperties()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: properties}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerProperty) GetProperty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	property, err := h.PropertyRepository.GetProperty(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: property}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerProperty) AddProperty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(propertiesdto.AddPropertyRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	property := models.Property{
		Name:      request.Name,
		CityId:    request.CityId,
		Address:   request.Address,
		Price:     request.Price,
		TypeRent:  request.TypeRent,
		Amenities: request.Amenities,
		Bedroom:   request.Bedroom,
		Bathroom:  request.Bathroom,
	}

	data, err := h.PropertyRepository.AddProperty(property)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseProperty(data)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerProperty) UpdateProperty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(propertiesdto.UpdatePropertyRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	property, err := h.PropertyRepository.GetProperty(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Name != "" {
		property.Name = request.Name
	}
	if request.Address != "" {
		property.Address = request.Address
	}
	if request.TypeRent != "" {
		property.TypeRent = request.TypeRent
	}
	if request.Amenities != nil {
		property.Amenities = request.Amenities
	}
	if request.Price != 0 {
		property.Price = request.Price
	}
	if request.Bedroom != 0 {
		property.Bedroom = request.Bedroom
	}
	if request.Bathroom != 0 {
		property.Bathroom = request.Bathroom
	}
	// if request.Image != "" {
	// 	house.Image = request.Image
	// }

	data, err := h.PropertyRepository.UpdateProperty(property)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseProperty(data)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerProperty) DeleteProperty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	property, err := h.PropertyRepository.GetProperty(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.PropertyRepository.DeleteProperty(property)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseProperty(data)}
	json.NewEncoder(w).Encode(response)
}

func convertResponseProperty(u models.Property) models.PropertyResponse {
	return models.PropertyResponse{
		ID:        u.ID,
		Name:      u.Name,
		CityId:    u.CityId,
		City:      u.City,
		Address:   u.Address,
		Price:     u.Price,
		TypeRent:  u.TypeRent,
		Amenities: u.Amenities,
		Bedroom:   u.Bedroom,
		Bathroom:  u.Bathroom,
	}
}