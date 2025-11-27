package db

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"uniqueIndex" json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Branches  []Branch       `json:"branches,omitempty"`
}

type Branch struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"index" json:"name"`
	ProjectID uint           `json:"project_id"`
	Type      string         `json:"type"` // "feature", "merge", etc.
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Tags      []Tag          `json:"tags,omitempty"`
}

type Tag struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"uniqueIndex" json:"name"`
	BranchID  uint           `json:"branch_id"`
	Branch    Branch         `json:"branch,omitempty"`
	CreatedBy string         `json:"created_by"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Deployment struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	TagID       uint      `json:"tag_id"`
	Tag         Tag       `json:"tag"`
	Environment string    `json:"environment"` // "staging", "production", etc.
	DeployedBy  string    `json:"deployed_by"`
	DeployedAt  time.Time `json:"deployed_at"`
}
