package api

import (
	"net/http"

	"github.com/google/uuid"
	jsonutils "github.com/reinheimermat/gobid/internal/json_utils"
	"github.com/reinheimermat/gobid/internal/usecase/product"
)

func (api *Api) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	data, problems, err := jsonutils.DecodeValidJSON[product.CreateProductRequest](r)

	if err != nil {
		jsonutils.EncodeJSON(w, r, http.StatusUnprocessableEntity, problems)
		return
	}

	userID, ok := api.Sessions.Get(r.Context(), "AuthenticatedUserId").(uuid.UUID)

	if !ok {
		jsonutils.EncodeJSON(w, r, http.StatusInternalServerError, map[string]any{
			"error": "could not get user from session",
		})
		return
	}

	id, err := api.ProductService.CreateProduct(
		r.Context(),
		userID,
		data.ProductName,
		data.Description,
		data.Baseprice,
		data.AuctionEnd,
	)

	if err != nil {
		jsonutils.EncodeJSON(w, r, http.StatusInternalServerError, map[string]any{
			"error": "could not create product",
		})
		return
	}

	jsonutils.EncodeJSON(w, r, http.StatusCreated, map[string]any{
		"message":    "product created successfully",
		"product_id": id,
	})
}
