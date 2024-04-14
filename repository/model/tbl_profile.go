package model

import "time"

type PreProfile struct {
	UserRefId      string `gorm:"column:user_ref_id"`
	FirstName      string `gorm:"column:first_name"`
	LastName       string `gorm:"column:last_name"`
	ProfilePicture string `gorm:"column:profile_image"`

	// profile //
	CropProfileImage string  `gorm:"column:crop_profile_image"`
	XProfile         float32 `gorm:"column:x_profile_image"`
	YProfile         float32 `gorm:"column:y_profile_image"`
	ZoomProfile      float32 `gorm:"column:zoom_profile_image"`

	// avatar //
	AvatarImage     string  `gorm:"column:avatar_image"`
	CropAvatarImage string  `gorm:"column:crop_avatar_image"`
	XAvatar         float32 `gorm:"column:x_avatar_image"`
	YAvatar         float32 `gorm:"column:y_avatar_image"`
	ZoomAvatar      float32 `gorm:"column:zoom_avatar_image"`
}

type ProfileSchemae struct {
	UserRefID      string    `gorm:"column:user_ref_id" json:"user_ref_id"`
	UserName       string    `gorm:"column:user_name" json:"user_name"`
	FirstName      string    `gorm:"column:first_name" json:"first_name"`
	LastName       string    `gorm:"column:last_name" json:"last_name"`
	Birthday       string    `gorm:"column:birthday" json:"birthday"`
	Gender         string    `gorm:"column:gender" json:"gender"`
	Email          string    `gorm:"column:email" json:"email"`
	Password       string    `gorm:"column:password" json:"password"`
	Pin            string    `gorm:"column:pin" json:"pin"`
	IsActive       bool      `gorm:"column:is_active" json:"is_active"`
	Phone          string    `gorm:"column:phone" json:"phone"`
	SecretKey      string    `gorm:"column:secret_key" json:"secret_key"`
	ProfileImage   string    `gorm:"column:profile_image" json:"profile_image"`
	Biography      string    `gorm:"column:biography" json:"biography"`
	CreateDatetime time.Time `gorm:"column:create_datetime" json:"create_datetime"`
}

func (ProfileSchemae) TableName() string {
	return "tbl_profile"
}
