package models

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

//Log Struct can be customized however we like
type Log struct {
	AppVersion string
	IP         string
	UserID     int
	Gender     string
	Method     string
	Status     int
	URL        string
	Country    string
	TimeStamp  string
}

//Generates Data to Log struct using gofakeit library
func (log *Log) GenerateLog(flag bool) Log {

	if flag {
		log.AppVersion = gofakeit.AppVersion()
		log.IP = gofakeit.IPv4Address()
		log.UserID = gofakeit.Number(1, 100301)
		log.Gender = gofakeit.Gender()
		log.Method = gofakeit.HTTPMethod()
		log.Status = gofakeit.HTTPStatusCode()
		log.URL = gofakeit.URL()
		log.Country = gofakeit.Country()
		log.TimeStamp = time.Now().Local().Format("2006/01/02 03:04:05") //time format YY/MM/DD HH:MM:SS
	}
	return *log
}

// func (log *Log) StopGeneration() bool {

// 	return true
// }

//  # paths:
//   #   - /var/log/*.log
