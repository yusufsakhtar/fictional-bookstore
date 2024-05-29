package handlers

import (
	"errors"

	"github.com/gorilla/mux"
	"github.com/yusufsakhtar/playstation-assignment/internal/repository"
	"github.com/yusufsakhtar/playstation-assignment/internal/service"

	"encoding/json"
	"net/http"
)

func GetCart(cartSvc *service.CartService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		cart, err := cartSvc.GetCart(repository.GetCartInput{ID: id})
		if err != nil {
			if errors.Is(err, repository.ErrCartNotFound) {
				http.Error(w, "cart not found", http.StatusNotFound)
				return
			}
			http.Error(w, "internal server error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(cart)
	}
}

func CheckoutCart(checkoutSvc *service.CheckoutService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		response, err := checkoutSvc.CheckoutCart(service.CheckoutInput{ID: id})
		if err != nil {
			if errors.Is(err, repository.ErrCartNotFound) {
				http.Error(w, "cart not found", http.StatusNotFound)
				return
			}
			http.Error(w, "internal server error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func ListCarts(cartSvc *service.CartService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		carts, err := cartSvc.ListCarts()
		if err != nil {
			http.Error(w, "internal server error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(carts)
	}
}
