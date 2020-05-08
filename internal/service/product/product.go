package retur

import (
	"context"

	"product/pkg/errors"
)

// Data ...
// Masukkan function dari package data ke dalam interface ini
type Data interface {
}

type doData interface {
	GetAllJSONDO(ctx context.Context, noTransf string) (reEntity.JSONDO, error)
}

// Service ...
// Tambahkan variable sesuai banyak data layer yang dibutuhkan
type Service struct {
	returData Data
	doData    doData
}

// New ...
// Tambahkan parameter sesuai banyak data layer yang dibutuhkan
func New(returData Data, doData doData) Service {
	// Assign variable dari parameter ke object
	return Service{
		returData: returData,
		doData:    doData,
	}
}

// TampilHeaderDO ...
func (s Service) TampilHeaderDO(ctx context.Context, noTransf string) (reEntity.TransFH, error) {
	var (
		tempResult reEntity.JSONDO
		err        error
	)

	tempResult, err = s.doData.GetAllJSONDO(ctx, noTransf)
	header := tempResult.TransFH
	if err != nil {
		return header, errors.Wrap(err, "[SERVICE][TampilHeaderDO]")
	}

	return header, err
}

// TampilDetailDO ...
func (s Service) TampilDetailDO(ctx context.Context, noTransf string) ([]reEntity.TransFD, error) {
	var (
		tempResult reEntity.JSONDO
		err        error
	)

	tempResult, err = s.doData.GetAllJSONDO(ctx, noTransf)
	details := tempResult.TransFD

	if err != nil {
		return details, errors.Wrap(err, "[SERVICE][TampilDetailDO]")
	}

	return details, err
}
