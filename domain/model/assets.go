package model

import (
	"github.com/pkg/errors"
)

// Assets 資産情報
type Assets struct {
	name   string
	amount int
	bank   string
}

// NewAssets コンストラクタ
func NewAssets() *Assets {
	return &Assets{}
}

// GetName 名称を返す
func (a *Assets) GetName() string {
	return a.name
}

// SetName 名称を設定する。許容しない文字列の場合エラーを返す
func (a *Assets) SetName(value string) error {
	if len(value) == 0 {
		return errors.WithStack(errors.New("名称は1文字以上である必要があります。"))
	}
	a.name = value
	return nil
}

// GetAmount 残高を返す
func (a *Assets) GetAmount() int {
	return a.amount
}

// SetAmount 残高を設定する。許容しない数値の場合はエラーを返す
func (a *Assets) SetAmount(value int) error {
	if value < 0 {
		return errors.WithStack(errors.New("残高は0以上の整数である必要があります。"))
	}
	a.amount = value
	return nil
}

// GetBank 金融機関を返す
func (a *Assets) GetBank() string {
	return a.bank
}

// SetBank 金融機関を設定する。許容しない文字列の場合はエラーを返す
func (a *Assets) SetBank(value string) error {
	if len(value) == 0 {
		return errors.WithStack(errors.New("保有金融機関は1文字以上である必要があります。"))
	}
	a.bank = value
	return nil
}

// IsValid それぞれのデータが正しいか確認する。正しい場合はtrue、正しくない場合はfalseを返す
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
