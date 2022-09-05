package common

import (
	"github.com/google/uuid"
	"go-starter/config"
	"time"
)

var swsreport = config.GetDb("swsreport")

type LogParams struct {
	Id       string    `json:"id" gorm:"column:id;primaryKey"`
	UserId   string    `json:"user_id" gorm:"column:user_id"`
	Action   string    `json:"action" gorm:"column:action"`
	CreateAt time.Time `json:"create_at" gorm:"column:create_at"`
}

func SaveLog(userId string, action string) {
	var logParams = LogParams{
		Id:       uuid.New().String(),
		UserId:   userId,
		Action:   action,
		CreateAt: time.Now(),
	}
	swsreport.Table("user_log").Create(logParams)
}
