package helper

import (
	"fmt"
	"time"
)

func GetKey(id string, endpoint string) string {
	return fmt.Sprintf("rate:%s:%s:%s", id, endpoint, time.Now().Format("2006-01-02"))
}
