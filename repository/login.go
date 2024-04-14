package repository

import (
	"errors"
	"time"

	model "github.com/jarbza/backend-api/repository/model"
	"github.com/jinzhu/gorm"
)

type GromDB struct {
	db *gorm.DB
}

func New(db *gorm.DB) *GromDB {
	return &GromDB{db: db}
}

func (db *GromDB) Person(userRefId string) (*model.PreProfile, error) {
	var response model.PreProfile
	err := db.db.Table("tbl_profile").
		Where(`tbl_profile.user_ref_id = ?`, userRefId).First(&response).Error
	if err != nil {
		return &response, err
	}
	return &response, nil
}

func (db *GromDB) Login(req model.LoginRequest) (*model.ProfileLogin, error) {
	var response model.ProfileLogin
	if err := db.db.Table("tbl_profile").Select(`tbl_profile.user_ref_id, tbl_profile.user_name, tbl_profile.first_name, tbl_profile.last_name,
	 tbl_profile.birthday, tbl_profile.gender, tbl_profile.email, tbl_profile.password, tbl_profile.pin, tbl_profile.is_active, tbl_profile.secret_key,
	 tbl_profile.profile_image, tbl_profile.biography, tbl_profile.create_datetime,tbl_profile_setting.noti_tag_active, tbl_profile_setting.noti_comment_active, 
	 tbl_profile_setting.noti_my_activity_active, tbl_profile_setting.noti_user_follow_active,tbl_profile_setting.noti_remember_active, tbl_profile_setting.noti_page_follow_active,
	 tbl_profile_setting.noti_post_all_active, tbl_profile_setting.noti_post_follow_active, tbl_profile_setting.noti_page_all_active`).
		Joins(`left join tbl_profile_setting on tbl_profile_setting.user_ref_id = tbl_profile.user_ref_id`).
		Where(`(tbl_profile.email = ? or tbl_profile.user_name = ?) && tbl_profile.password= ?`,
			req.UserName, req.UserName, req.Password).Scan(&response).Error; err != nil {
		return &response, err
	}
	return &response, nil
}

func (db *GromDB) InsertLoginHistory(req model.LoginRequest, userRefId string) error {

	data := maploginhistory(req, userRefId)
	_, err := db.GetLoginRecord(data)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			var record model.LoginHistory
			record.UserRefID = userRefId
			record.Device = req.Device
			record.Platform = req.Platform
			record.Time = time.Now()
			record.ClientIP = req.ClientIP

			db.db.Table("tbl_login_history").Create(&record)
			return nil
		} else {
			return err
		}
	}
	return nil
}

func (db *GromDB) GetLoginRecord(req model.LoginHistory) (*model.LoginHistoryResponse, error) {
	var record model.LoginHistoryResponse
	err := db.db.Table("tbl_login_history").
		Where(`tbl_login_history.user_ref_id = ? and tbl_login_history.x_device_model = ? and tbl_login_history.x_platform = ?`,
			req.UserRefID, req.Device, req.Platform).Find(&record).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &record, err
	} else {
		return &record, nil
	}
}

func (db *GromDB) GetLoginRecords(userRefId string) ([]model.LoginHistoryResponse, error) {
	var record []model.LoginHistoryResponse
	err := db.db.Table("tbl_login_history").Select(`tbl_login_history.user_ref_id, tbl_login_history.x_device_model, tbl_login_history.x_platform, tbl_login_history.datetime_login`).
		Where(`tbl_login_history.user_ref_id = ?`, userRefId).Find(&record).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return record, err
	} else {
		return record, nil
	}
}

func maploginhistory(req model.LoginRequest, userRefId string) model.LoginHistory {
	data := model.LoginHistory{
		UserRefID: userRefId,
		Device:    req.Device,
		Platform:  req.Platform,
	}
	return data
}
