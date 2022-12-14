package clock

import (
	"time"
	"github.com/beevik/ntp"
)

// GetDate - функция, возвращающее время из NTP хоста
func GetDate(host string) (time.Time, error) {
	t, err := ntp.Time(host)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}