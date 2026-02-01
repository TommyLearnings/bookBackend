package book

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Record struct {
	bun.BaseModel `bun:"table:bi100"`

	// 基礎欄位
	Id          int       `bun:"id,pk,autoincrement"`
	ReferenceId uuid.UUID `bun:"reference_id,notnull,unique,default:uuid_generate_v4()"`

	Version       int64     `bun:"version,default:0"`
	CreatedBy     int64     `bun:"created_by,notnull"`      // 假設儲存 User ID
	LastUpdatedBy int64     `bun:"last_updated_by,notnull"` // 假設儲存 User ID
	CreatedAt     time.Time `bun:"date_created,nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:"last_updated,nullzero,notnull,default:current_timestamp"`

	Isbn        string    `bun:"isbn,notnull,unique"`
	Title       string    `bun:"title,notnull"`
	Episode     string    `bun:"episode"`
	Author      string    `bun:"author"`
	PublishDate time.Time `bun:"publish_date"`
	Description string    `bun:"description"`
	ImageUrl    string    `bun:"image_url"`
}
