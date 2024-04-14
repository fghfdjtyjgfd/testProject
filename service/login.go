package service

import (
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	model "github.com/jarbza/backend-api/repository/model"
	"github.com/jarbza/errs"
	"github.com/jinzhu/gorm"
)

type Databaser interface {
	Login(req model.LoginRequest) (*model.ProfileLogin, error)
	InsertLoginHistory(req model.LoginRequest, userRefId string) error
	GetLoginRecord(req model.LoginHistory) (*model.LoginHistoryResponse, error)
	GetLoginRecords(userRefId string) ([]model.LoginHistoryResponse, error)
	Person(userRefId string) (*model.PreProfile, error)
}

type Servicer interface {
	ValidateLogin(c *gin.Context, req model.LoginRequest) (*model.LoginResponse, error)
	GetRecords(c *gin.Context, userRefId string) (*[]model.LoginHistoryResponse, error)
	CheckStatus(c *gin.Context, req model.LoginHistory) *model.LoginStatus
}

type Service struct {
	database Databaser
}

func NewService(database Databaser) *Service {
	return &Service{
		database: database,
	}
}

const GGStore = "https://storage.googleapis.com/jarb-bucket-001/"

func (s *Service) CheckStatus(c *gin.Context, req model.LoginHistory) *model.LoginStatus {
	res := model.LoginStatus{
		ForceLogout: true,
	}
	record, err := s.database.GetLoginRecord(req)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &res
	}
	res.ForceLogout = false
	res.ClientIP = record.ClientIP
	return &res
}

func (s *Service) GetRecords(c *gin.Context, userRefId string) (*[]model.LoginHistoryResponse, error) {
	var record []model.LoginHistoryResponse
	record, err := s.database.GetLoginRecords(userRefId)
	if err != nil {
		return &[]model.LoginHistoryResponse{}, errs.NewBadRequest("Problem on login-history", "")
	}

	return &record, nil
}

func (s *Service) ValidateLogin(c *gin.Context, req model.LoginRequest) (*model.LoginResponse, error) {

	res, err := s.database.Login(req)
	if err != nil {
		return &model.LoginResponse{}, errs.NewBadRequest("Invalid username or password invalid", "")
	}

	err = s.database.InsertLoginHistory(req, res.UserRefID)
	if err != nil {
		return &model.LoginResponse{}, errs.NewBadRequest("Problem on login-history", "")
	}

	token := `Bearer ` + createToken(res)

	user_info, err := s.database.Person(res.UserRefID)
	if err != nil {
		return nil, errs.New(http.StatusConflict, "20009", "Can not identify user", "")
	}
	userProfileImage := profileImageBuilder(*user_info)

	return &model.LoginResponse{
		UserRefID:            res.UserRefID,
		UserName:             res.UserName,
		FirstName:            res.FirstName,
		LastName:             res.LastName,
		ProfileName:          res.FirstName + " " + res.LastName,
		Birthday:             res.Birthday,
		Gender:               res.Gender,
		Email:                res.Email,
		SecretKey:            res.SecretKey,
		ProfileImage:         userProfileImage,
		Biography:            res.Biography,
		NotiTagActive:        res.NotiTagActive,
		NotiCommentActive:    res.NotiCommentActive,
		NotiMyActivityActive: res.NotiMyActivityActive,
		NotiUserFollowActive: res.NotiUserFollowActive,
		NotiRememberActive:   res.NotiRememberActive,
		NotiPageFollowActive: res.NotiPageFollowActive,
		NotiPostAllActive:    res.NotiPostAllActive,
		// NotiPostFollowActive: res.NotiPostFollowActive,
		// NotiPageAllActive:    res.NotiPageAllActive,
		SessionKey: token,
	}, nil
}

func createToken(res *model.ProfileLogin) string {

	token := jwt.New(jwt.SigningMethodHS256)
	atClaims := token.Claims.(jwt.MapClaims)

	atClaims["userRefId"] = res.UserRefID
	atClaims["username"] = res.UserName
	atClaims["firstname"] = res.FirstName
	atClaims["lastname"] = res.LastName
	atClaims["profileName"] = res.FirstName + " " + res.LastName
	atClaims["birthday"] = res.Birthday
	atClaims["gender"] = res.Gender
	atClaims["email"] = res.Email
	atClaims["secretKey"] = res.SecretKey
	atClaims["biography"] = res.Biography
	atClaims["exp"] = time.Now().Add(time.Hour * 2160).Unix()

	tokenjwt, _ := token.SignedString([]byte("secret"))

	return tokenjwt
}

func URLpadding(s string) string {
	if s != "" {
		return GGStore + s
	} else {
		return ""
	}
}

func profileImageBuilder(user_info model.PreProfile) model.ProfilePicture {
	cropPosition := model.CropPosition{
		X:    user_info.XProfile,
		Y:    user_info.YProfile,
		Zoom: user_info.ZoomProfile,
	}
	return model.ProfilePicture{
		ImageUrl:       URLpadding(user_info.ProfilePicture),
		ImageUrlCroped: URLpadding(user_info.CropProfileImage),
		CropPosition:   cropPosition,
	}
}
