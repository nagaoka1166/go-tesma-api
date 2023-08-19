// ファイル名: app/domain/entity/faculty.go
package entity

type Faculty struct {
	ID    int `json:"id"`
	Name  string `json:"name"`
	Order int    `json:"order"`
}