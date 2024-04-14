package model

import "time"

type LoginStatus struct {
	ForceLogout bool   `json:"forceLogout"`
	ClientIP    string `json:"ip_address"`
}

type LoginRequest struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Device   string
	Platform string
	ClientIP string
}

type LoginHistory struct {
	UserRefID string    `gorm:"column:user_ref_id" json:"userRefId"`
	Device    string    `gorm:"column:x_device_model" json:"devicemodel"`
	Platform  string    `gorm:"column:x_platform" json:"platform"`
	Time      time.Time `gorm:"column:datetime_login" json:"datelogin"`
	ClientIP  string    `gorm:"column:ip_address" json:"ip_address"`
}

type LoginHistoryResponse struct {
	UserRefID string `gorm:"column:user_ref_id" json:"userRefId"`
	Device    string `gorm:"column:x_device_model" json:"devicemodel"`
	Platform  string `gorm:"column:x_platform" json:"platform"`
	Time      string `gorm:"column:datetime_login" json:"datelogin"`
	ClientIP  string `gorm:"column:ip_address" json:"ip_address"`
}

type LoginResponse struct {
	UserRefID            string         `json:"userRefId"`
	UserName             string         `json:"username"`
	FirstName            string         `json:"firstname"`
	LastName             string         `json:"lastname"`
	ProfileName          string         `json:"profileName"`
	Birthday             string         `json:"birthday"`
	Gender               string         `json:"gender"`
	Email                string         `json:"email"`
	SecretKey            string         `json:"secretKey"`
	ProfileImage         ProfilePicture `json:"profileImage"`
	Biography            string         `json:"biography"`
	NotiTagActive        bool           `json:"notiTagActive"`
	NotiCommentActive    bool           `json:"notiCommentActive"`
	NotiRememberActive   bool           `json:"notiRememberActive"`
	NotiMyActivityActive bool           `json:"notiMyActivityActive"`
	NotiUserFollowActive bool           `json:"notiUserFollowActive"`
	NotiPageFollowActive bool           `json:"notiPageFollowActive"`
	NotiPostAllActive    bool           `json:"notiPostAllActive"`
	SessionKey           string         `json:"sessionKey"`
}

type ProfileLogin struct {
	UserRefID            string
	UserName             string
	FirstName            string
	LastName             string
	Birthday             string
	Gender               string
	Email                string
	Password             string
	Pin                  string
	IsActive             string
	SecretKey            string
	ProfileImage         string
	Biography            string
	NotiTagActive        bool
	NotiCommentActive    bool
	NotiMyActivityActive bool
	NotiUserFollowActive bool
	NotiRememberActive   bool
	NotiPageFollowActive bool
	NotiPostAllActive    bool
	NotiPostFollowActive bool
	NotiPageAllActive    bool
}

type RegRequest struct {
	UserName  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Birthday  string `json:"birthday"`
	Gender    string `json:"gender"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
}

type ProfilePicture struct {
	CropPosition   CropPosition `json:"cropPosition"`
	ImageUrl       string       `json:"imageUrl"`
	ImageUrlCroped string       `json:"imageUrlCroped"`
}

type CropPosition struct {
	X    float32 `json:"x"`
	Y    float32 `json:"y"`
	Zoom float32 `json:"zoom"`
}
