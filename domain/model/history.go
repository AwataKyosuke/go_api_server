package model

import (
	"time"
)

// History 履歴情報
type History struct {
	date       time.Time
	content    string
	amount     int
	bank       string
	majorType  string
	mediumType string
	memo       string
}

// NewHistory コンストラクタ
func NewHistory() *History {
	return &History{}
}

// Date 記入日を返す
func (h *History) Date() time.Time {
	return h.date
}

// SetDate 記入日を設定する。許容しない日付の場合エラーを返す。
func (h *History) SetDate(value time.Time) error {
	h.date = value
	return nil
}

// Content 内容を返す
func (h *History) Content() string {
	return h.content
}

// SetContent 内容を設定する。許容しない内容の場合エラーを返す
func (h *History) SetContent(value string) error {
	h.content = value
	return nil
}

// Amount 金額を返す
func (h *History) Amount() int {
	return h.amount
}

// SetAmount 金額を設定する。許容しない金額の場合エラーを返す
func (h *History) SetAmount(value int) error {
	h.amount = value
	return nil
}

// Bank 金融機関を返す
func (h *History) Bank() string {
	return h.bank
}

// SetBank 金融機関を設定する。許容しない金融機関の場合エラーを返す
func (h *History) SetBank(value string) error {
	h.bank = value
	return nil
}

// MajorType 大項目を返す
func (h *History) MajorType() string {
	return h.majorType
}

// SetMajorType 大項目を設定。許容しない大項目の場合エラーを返す
func (h *History) SetMajorType(value string) error {
	h.majorType = value
	return nil
}

// MediumType 中項目を返す
func (h *History) MediumType() string {
	return h.mediumType
}

// SetMediumType 中項目を設定。許容しない中項目の場合エラーを返す
func (h *History) SetMediumType(value string) error {
	h.mediumType = value
	return nil
}

// Memo メモを返す
func (h *History) Memo() string {
	return h.memo
}

// SetMemo メモを設定。許容しないメモの場合はエラーを返す
func (h *History) SetMemo(value string) error {
	h.memo = value
	return nil
}
