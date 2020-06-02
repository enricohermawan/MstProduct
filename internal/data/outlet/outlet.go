package outlet

import (
	"context"

	outletEntity "product/internal/entity/outlet"

	"product/pkg/errors"
	"product/pkg/httpclient"
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

// GetOutletName ...
func (d Data) GetOutletName(ctx context.Context, outcode string) (outletEntity.Outlet, error) {
	var JSON outletEntity.JSONTerima
	var endpoint = d.baseURL + "TampilDataOutlet/" + outcode
	_, err := d.client.GetJSON(ctx, endpoint, nil, &JSON)
	if err != nil {
		return JSON.Data, errors.Wrap(err, "[DATA][GetOutletName]")
	}

	return JSON.Data, err
}
