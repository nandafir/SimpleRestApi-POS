package utils

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
)

type UniversalDTO struct {
	Data interface{} `json:"data"`
}

func ErrorLog(flag string, err error) {
	fmt.Println("============================================")
	log.Error(fmt.Sprintf("ERRORS:::%s => %s", flag, err))
}

func ErrorLogWithParams(flag string, err error, universalDTO UniversalDTO) {
	byteData, _ := json.Marshal(universalDTO)
	fmt.Println("============================================")
	log.Error(fmt.Sprintf("ERRORS:::%s => %s ::: %s", flag, err, string(byteData)))
}
