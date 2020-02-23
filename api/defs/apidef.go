package defs

import "time"

//data models
type User struct {
	ID        uint   `gorm:"AUTO_INCREMENT" json:"id"`
	LoginName string `gorm:"type:varchar(64);UNIQUE" form:"user_name" json:"user_name" `
	Pwd       string `gorm:"type:text;not_null" form:"pwd" json:"pwd" `
}
type VideoInfo struct {
	ID           string     `gorm:"type:varchar(64);not_null" json:"id" `
	AuthorID     uint       `json:"author_id" `
	Name         string     `gorm:"type:text" json:"name" `
	DisplayCtime string     `gorm:"type:text" json:"display_ctime" `
	CreateTime   *time.Time `gorm:"default:current_timestamp"`
}
type Comment struct {
	ID       string     `gorm:"type:varchar(64);not_null" json:"id" `
	VideoID  string     `gorm:"type:varchar(64)" json:"video_id" `
	AuthorID uint       `json:"author_id" `
	Content  string     `gorm:"type:text" json:"content" `
	Time     *time.Time `gorm:"default:current_timestamp"`
}

type Session struct {
	ID        string `gorm:"type:varchar(255);not_null;primary_key"`
	TTL       string `gorm:"type:tinytext"`
	LoginName string `gorm:"type:varchar(64)"`
}

//response
type SignedUp struct {
	Success   bool   `json:"success"`
	SessionID string `json:"session_id"`
}
type SignedIn struct {
	Success   bool   `json:"success"`
	SessionID string `json:"session_id"`
}
type VideosInfo struct {
	Videos []*VideoInfo `json:"videos"`
}

type CommentForList struct {
	Comment
	AuthorName string `json:"author" `
}

type Comments struct {
	Comments []*CommentForList `json:"comments"`
}
