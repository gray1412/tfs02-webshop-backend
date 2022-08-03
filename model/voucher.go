package model

import (
	"errors"
	"project-backend/util/constant"
	"time"
)

type Voucher struct {
	Id            int       `json:"id"`
	Code          string    `json:"code"`
	Discount      float64   `json:"discount"` // giam
	Unit          string    `json:"unit"`     // persent || usd
	MaxSaleAmount float64   `json:"max_sale_amount"`
	Description   string    `json:"description"`
	TimeEnd       time.Time `json:"time_end"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (v *Voucher) GetByCode(code string) error {
	err := db.Where("code = ?", code).Take(v).Error
	if err != nil{
		return errors.New(constant.ERROR_VOUCHER_NOT_EXISTS)
	}
	return nil
}
func (v *Voucher) CheckVoucher(code string) error {
	err := v.GetByCode(code)
	if err != nil {
		return err
	}
	if !time.Now().Before(v.TimeEnd) {
		return errors.New(constant.ERROR_VOUCHER_EXPIRED)
	}
	return nil
}
