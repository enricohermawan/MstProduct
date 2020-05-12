package product

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"product/pkg/response"

	pEntity "product/internal/entity/product"
)

// IProductSvc is an interface to Skeleton Service
// Masukkan function dari service ke dalam interface ini
type IProductSvc interface {
	TampilDetailMP(ctx context.Context, kode string) (pEntity.MstProduct, error)
}

type (
	// Handler ...
	Handler struct {
		ProductSvc IProductSvc
	}
)

// New for bridging product handler initialization
func New(is IProductSvc) *Handler {
	return &Handler{
		ProductSvc: is,
	}
}

// ProductHandler will receive request and return response
func (h *Handler) ProductHandler(w http.ResponseWriter, r *http.Request) {
	var (
		resp     *response.Response
		result   interface{}
		metadata interface{}
		err      error
		errRes   response.Error
	)
	resp = &response.Response{}
	defer resp.RenderJSON(w, r)

	switch r.Method {
	// Check if request method is GET
	case http.MethodGet:

		paramMap := r.URL.Query()
		len := len(paramMap)
		switch len {
		case 1:
			_, getKodeOK := paramMap["kode"]
			if getKodeOK {
				var (
					kode string
				)

				kode = r.FormValue("kode")
				fmt.Println("Masuk ke Handler")
				result, err = h.ProductSvc.TampilDetailMP(context.Background(), kode)
			}
		}

	// Check if request method is POST
	case http.MethodPost:

	// Check if request method is PUT
	case http.MethodPut:

	// Check if request method is DELETE
	case http.MethodDelete:

	default:
		err = errors.New("404")
	}

	// If anything from service or data return an error
	if err != nil {
		// Error response handling
		errRes = response.Error{
			Code:   101,
			Msg:    "Data Not Found",
			Status: true,
		}
		// If service returns an error
		if strings.Contains(err.Error(), "service") {
			// Replace error with server error
			errRes = response.Error{
				Code:   201,
				Msg:    "Failed to process request due to server error",
				Status: true,
			}
		}

		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		resp.Error = errRes
		return
	}

	resp.Data = result
	resp.Metadata = metadata
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
}
