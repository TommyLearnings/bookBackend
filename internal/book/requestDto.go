package book

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type SaveRequestBody struct {
	Id          int64     `json:"id"`
	ReferenceId uuid.UUID `json:"referenceId"`

	// 系統/版本欄位
	Version       int64     `json:"version"`
	CreatedBy     int64     `json:"createdBy"`
	LastUpdatedBy int64     `json:"lastUpdatedBy"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`

	// 業務詳細欄位
	Isbn        string    `json:"isbn"`
	Title       string    `json:"title"`
	Episode     string    `json:"episode"`
	Author      string    `json:"author"`
	PublishDate time.Time `json:"publishDate"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"imageUrl"`
}

func (n *SaveRequestBody) Validate() (record *Record, errs error) {
	// 1. 基本必填欄位驗證
	if n.Author == "" {
		errs = errors.Join(errs, fmt.Errorf("author is empty"))
	}
	if n.Title == "" {
		errs = errors.Join(errs, fmt.Errorf("title is empty"))
	}

	if n.Description == "" { // 原 Summary 改為 Description
		errs = errors.Join(errs, fmt.Errorf("description is empty"))
	}
	if n.Isbn == "" { // 新增 Isbn 驗證
		errs = errors.Join(errs, fmt.Errorf("isbn is empty"))
	}

	// 2. 時間驗證
	// 因為 n.CreatedAt 已經是 time.Time，直接檢查是否為零值即可
	if n.CreatedAt.IsZero() {
		errs = errors.Join(errs, fmt.Errorf("created_at is required"))
	}

	// 如果有錯誤則返回
	if errs != nil {
		return nil, errs
	}

	// 5. 轉換為 Record Model
	return &Record{
		Id:            n.Id,
		ReferenceId:   n.ReferenceId,
		Isbn:          n.Isbn,
		Title:         n.Title,
		Episode:       n.Episode,
		Author:        n.Author,
		PublishDate:   n.PublishDate,
		Description:   n.Description,
		ImageUrl:      n.ImageUrl,
		CreatedAt:     n.CreatedAt,
		UpdatedAt:     n.UpdatedAt,
		Version:       n.Version,
		CreatedBy:     n.CreatedBy,
		LastUpdatedBy: n.LastUpdatedBy,
	}, nil
}
