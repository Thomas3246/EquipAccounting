package datevalidate

import "time"

func ValidateDate(dateStr string) bool {
	// Пытаемся распарсить строку по заданному формату
	_, err := time.Parse("2006.01.02", dateStr)
	return err == nil
}
