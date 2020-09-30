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

func NewHistory() *History {
	return &History{}
}

func (h *History) GetDate() time.Time {
	return h.date
}

func (h *History) SetDate(value time.Time) error {
	h.date = value
	return nil
}

func (h *History) GetContent() string {
	return h.content
}

func (h *History) SetContent(value string) error {
	h.content = value
	return nil
}

func (h *History) GetAmount() int {
	return h.amount
}

func (h *History) SetAmount(value int) error {
	h.amount = value
	return nil
}

func (h *History) GetBank() string {
	return h.bank
}

func (h *History) SetBank(value string) error {
	h.bank = value
	return nil
}

func (h *History) GetMajorType() string {
	return h.majorType
}

func (h *History) SetMajorType(value string) error {
	h.majorType = value
	return nil
}

func (h *History) GetMediumType() string {
	return h.mediumType
}

func (h *History) SetMediumType(value string) error {
	h.mediumType = value
	return nil
}

func (h *History) GetMemo() string {
	return h.memo
}

func (h *History) SetMemo(value string) error {
	h.memo = value
	return nil
}
