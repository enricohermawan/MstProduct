package do

import (
	"context"
	doEntity "product/internal/entity/do"
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

// GetAllHeaderJSONDO ...
func (d Data) GetAllHeaderJSONDO(ctx context.Context, noTransf string) (doEntity.TransfH, error) {
	var json doEntity.JSONDO
	var endpoint = d.baseURL + "/CariDO/" + noTransf
	_, err := d.client.GetJSON(ctx, endpoint, nil, &json)
	if err != nil {
		return json.Header, errors.Wrap(err, "[DATA][GetJsonDOByNoTransf]")
	}

	return json.Header, err
}

// GetAllDetailJSONDO ...
func (d Data) GetAllDetailJSONDO(ctx context.Context, noTransf string) ([]doEntity.TransfD, error) {
	var json doEntity.JSONDO
	var endpoint = d.baseURL + "/CariDO/" + noTransf
	_, err := d.client.GetJSON(ctx, endpoint, nil, &json)
	if err != nil {
		return json.Detail, errors.Wrap(err, "[DATA][GetJsonDOByNoTransf]")
	}

	return json.Detail, err
}
