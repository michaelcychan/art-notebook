package models

type SavedData struct {
	Source   string   `json:"source" gorm:"column:source"`
	ID       int      `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	SourceId string   `json:"source-id" gorm:"column:source_id"`
	User     string   `json:"user" gorm:"column:username"`
	Tag      []string `json:"tag" gorm:"column:tag;type:text[]"`
	Note     string   `json:"note" gorm:"column:note"`
}
