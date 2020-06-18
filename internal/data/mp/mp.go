package mp

import (
	"context"

	pEntity "product/internal/entity/product"
	"product/pkg/httpclient"

	"github.com/pkg/errors"
)

// Data ...
type Data struct {
	client  *httpclient.Client
	baseURL string
}

// New return an ads resource
func New(client *httpclient.Client, baseURL string) Data {
	d := Data{
		client:  client,
		baseURL: baseURL,
	}

	return d
}

// GetAllJSONMP ...
func (d Data) GetAllJSONMP(ctx context.Context, kode string) (pEntity.MstProduct, error) {
	var json pEntity.JSONTerima
	var endpoint = d.baseURL + "/Product?kode=" + kode

	_, err := d.client.GetJSON(ctx, endpoint, nil, &json)

	if err != nil {
		return json.Data, errors.Wrap(err, "[DATA][GetJsonMPByProcod]")
	}

	return json.Data, err
}
