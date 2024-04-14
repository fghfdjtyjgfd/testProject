package model

import "time"

type ProfileSetting struct {
	UserRefID            string    `gorm:"column:user_ref_id" json:"user_ref_id"`
	NotiTagActive        bool      `gorm:"column:noti_tag_active" json:"noti_tag_active"`
	NotiCommentActive    bool      `gorm:"column:noti_comment_active" json:"noti_comment_active"`
	NotiMyActivityActive bool      `gorm:"column:noti_my_activity_active" json:"noti_my_activity_active"`
	NotiUserFollowActive bool      `gorm:"column:noti_user_follow_active" json:"noti_user_follow_active"`
	NotiRememberActive   bool      `gorm:"column:noti_remember_active" json:"noti_remember_active"`
	NotiPageFollowActive bool      `gorm:"column:noti_page_follow_active" json:"noti_page_follow_active"`
	NotiPostAllActive    bool      `gorm:"column:noti_post_all_active" json:"noti_post_all_active"`
	NotiPostFollowActive bool      `gorm:"column:noti_post_follow_active" json:"noti_post_follow_active"`
	NotiPageAllActive    bool      `gorm:"column:noti_page_all_active" json:"noti_page_all_active"`
	UpdateDate           time.Time `gorm:"column:update_date" json:"update_date"`
}

func (ProfileSetting) TableName() string {
	return "tbl_profile_setting"
}
