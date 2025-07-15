package models

type Todo struct {
	ID    int    `json:"id"    gorm:"column:id;primaryKey;autoIncrement;notNull"`
	Title string `json:"title" gorm:"column:title;notNull"`
	Done  bool   `json:"done"  gorm:"column:done;notNull"`
}
