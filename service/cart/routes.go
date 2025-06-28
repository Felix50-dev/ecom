package cart

import (
	"net/http"

	"github.com/Felix50-dev/ecom/service/auth"
	"github.com/Felix50-dev/ecom/types"
	"github.com/Felix50-dev/ecom/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	orderStore   types.OrderStore
	userStore    types.UserStore
	productStore types.ProductStore
}

func NewHandler(
	orderStore types.OrderStore,
	userStore types.UserStore,
	productStore types.ProductStore,
) *Handler {
	return &Handler{
		orderStore:   orderStore,
		userStore:    userStore,
		productStore: productStore,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/cart/checkout", auth.WithJWTAuth(h.handleCheckout, h.userStore)).Methods(http.MethodPost)
}

func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())

	var cart types.CartCheckoutPayload
	if err := utils.ParseJSON(r, &cart); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	productIds, err := getCartItemsIDs(cart.Items)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// get products
	products, err := h.productStore.GetProductsByID(productIds)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	orderID, totalPrice, err := h.createOrder(products, cart.Items, userID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"total_price": totalPrice,
		"order_id":    orderID,
	})

}
