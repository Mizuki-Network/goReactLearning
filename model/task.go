package model

import "time"

type Task struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"` //空の値を許可しない
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	//どのユーザーがタスクを追加したか分かるように以下にユーザー情報も追加。
	//"foreignKeyにUser型のIDを設定し一対価の関係を設定できる
	//constraint:Ondelete:CASCADEによりユーザーを削除時に同一ユーザーに紐づくTaskを全て削除する
	User   User `json:"user" gorm:"foreignKey:UserID; constraint:OnDelete:CASCADE"`
	UserID uint `json:"user_id" gorm:"not null"`
}

// クライアントからのGETリクエスト時にクライアントへ返すデータ構造を定義する
type TaskResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
