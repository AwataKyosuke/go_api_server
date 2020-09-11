package model

import "github.com/pkg/errors"

// Assets 資産情報
type Assets struct {
	name   string
	amount int
	bank   string
}

func NewAssets() *Assets {
	return &Assets{}
}

func (a *Assets) SetName(value string) error {
	if len(value) == 0 {
		return errors.WithStack(errors.New("名称は1文字以上である必要があります。"))
	}
	a.name = value
	return nil
}

func (a *Assets) SetAmount(value int) error {
	if value < 0 {
		return errors.WithStack(errors.New("残高は0以上の整数である必要があります。"))
	}
	a.amount = value
	return nil
}

func (a *Assets) SetBank(value string) error {
	if len(value) == 0 {
		return errors.WithStack(errors.New("保有金融機関は1文字以上である必要があります。"))
	}
	a.bank = value
	return nil
}

func (a *Assets) IsValid() bool {
	if len(a.name) == 0 {
		return false
	}
	if a.amount < 0 {
		return false
	}
	if len(a.bank) == 0 {
		return false
	}
	return true
}
