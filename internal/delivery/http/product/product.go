package product

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	doEntity "product/internal/entity/do"
	pEntity "product/internal/entity/product"
	"product/pkg/response"
)

// IProductSvc is an interface to Skeleton Service
// Masukkan function dari service ke dalam interface ini
type IProductSvc interface {
	TampilDetailMP(ctx context.Context, kode string) (pEntity.MstProduct, error)
	// TampilDetailReceiveByNoReceive(ctx context.Context, NoTranrc string) (pEntity.MstProduct, error)
	TampilAllHeaderDataReceive(ctx context.Context) ([]pEntity.HeaderRC, error)
	TampilDataByNoReceive(ctx context.Context, NoTranrc string) (pEntity.JSONRCByNoReceive, error)
	TampilDataDO(ctx context.Context, noTransf string) (doEntity.JSONDO, error)
	InsertDataFromAPI(ctx context.Context, noTransf string) error
	EditDetailOrderByNoTransfandProcode(ctx context.Context, detail doEntity.TransfD, noTransf string, procod string) error
	PrintReceive(ctx context.Context, noTransf string, NoTranrc string) ([]pEntity.JSONPrintReceive, error)
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
		case 2:
			_, NoTranrcOK := paramMap["NoTranrc"]
			_, NoTransfOK := paramMap["NoTransf"]
			if NoTranrcOK && NoTransfOK {
				var (
					NoTranrc string
					noTransf string
				)
				NoTranrc = r.FormValue("NoTranrc")
				noTransf = r.FormValue("NoTransf")
				result, err = h.ProductSvc.PrintReceive(context.Background(), noTransf, NoTranrc)
			}
		case 1:
			_, getKodeOK := paramMap["kode"]
			_, NoTranrcOK := paramMap["NoTranrc"]
			_, NoTransfOK := paramMap["NoTransf"]
			if getKodeOK {
				var (
					kode string
				)

				kode = r.FormValue("kode")
				result, err = h.ProductSvc.TampilDetailMP(context.Background(), kode)
			} else if NoTranrcOK {
				var (
					NoTranrc string
				)
				NoTranrc = r.FormValue("NoTranrc")
				result, err = h.ProductSvc.TampilDataByNoReceive(context.Background(), NoTranrc)
			} else if NoTransfOK {
				var (
					noTransf string
				)

				noTransf = r.FormValue("NoTransf")
				result, err = h.ProductSvc.TampilDataDO(context.Background(), noTransf)
			}

		case 0:
			result, err = h.ProductSvc.TampilAllHeaderDataReceive(context.Background())

		}

	// Check if request method is POST
	case http.MethodPost:
		paramMap := r.URL.Query()
		len := len(paramMap)
		switch len {
		case 3:
			var editProduct doEntity.EditProduct
			_, NoTransfOK := r.URL.Query()["NoTransf"]
			_, ProcodOK := r.URL.Query()["Procod"]
			// _, totalPriceOK := r.URL.Query()["totalPrice"]
			body, _ := ioutil.ReadAll(r.Body)
			// totalPrice, _ := strconv.Atoi(r.FormValue("totalPrice"))

			if NoTransfOK && ProcodOK {
				json.Unmarshal(body, &editProduct)
				err = h.ProductSvc.EditDetailOrderByNoTransfandProcode(context.Background(), editProduct.Detail, r.FormValue("NoTransf"), r.FormValue("Procod"))
			}
		case 1:
			var noTransf string
			var insert doEntity.JSONDO
			_, noTransfOK := r.URL.Query()["NoTransf"]
			body, _ := ioutil.ReadAll(r.Body)
			if noTransfOK {
				json.Unmarshal(body, &insert)
				err = h.ProductSvc.InsertDataFromAPI(context.Background(), noTransf)
			}

		}

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
