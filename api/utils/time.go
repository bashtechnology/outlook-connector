package utils

import (
	"fmt"
	"strings"
	"time"
)

func StrToTime(currentTime string) time.Time {
	saoPauloLocation, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		return time.Time{}
	}
	format := "2006-01-02 15:04:05"
	parsedTime, err := time.ParseInLocation(format, currentTime, saoPauloLocation)
	if err != nil {
		// fmt.Println("Erro ao analisar a string de data e hora:", err)
		return time.Time{}
	}
	return parsedTime
}

func AdjustTimeString(timeString string) string {
	// Verifica se a string contém "T", indicando que está no formato ISO 8601
	if strings.Contains(timeString, "T") {
		// Substitui "T" por um espaço para corresponder ao formato esperado
		return strings.Replace(timeString, "T", " ", 1)
	}

	// Caso contrário, assume que a string já está no formato esperado
	return timeString
}

func StrToTimeErr(currentTime string) (time.Time, error) {
	saoPauloLocation, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		fmt.Println("Erro ao carregar o local de São Paulo:", err)
		return time.Time{}, err
	}
	parsedTime, err := time.Parse(time.RFC3339, currentTime)
	if err != nil {
		fmt.Println("Erro ao analisar a string de data e hora:", err)
		return time.Time{}, err
	}
	saoPauloTime := parsedTime.In(saoPauloLocation)
	return saoPauloTime, nil
}

func TimeToStr(currentTime time.Time) string {
	saoPauloLocation, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		fmt.Println("Erro ao carregar o local de São Paulo:", err)
		return ""
	}
	saoPauloTime := currentTime.In(saoPauloLocation)
	formattedTime := saoPauloTime.Format("2006-01-02 15:04:05")
	return formattedTime
}
func TimeToStrMs(currentTime time.Time) string {
	saoPauloLocation, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		fmt.Println("Erro ao carregar o local de São Paulo:", err)
		return ""
	}
	saoPauloTime := currentTime.In(saoPauloLocation)
	formattedTime := saoPauloTime.Format("2006-01-02 15:04:05.000")
	return formattedTime
}

func TimeToTimeLocal(currentTime time.Time) time.Time {
	saoPauloLocation, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		fmt.Println("Erro ao carregar o local de São Paulo:", err)
		return time.Time{}
	}
	saoPauloTime := currentTime.In(saoPauloLocation)
	return saoPauloTime
}

func GetYesterdayDateString() string {
	yesterday := time.Now().AddDate(0, 0, -1)
	yesterdayString := yesterday.Format("2006-01-02")
	return yesterdayString
}
func GetDateString() string {
	yesterday := time.Now()
	yesterdayString := yesterday.Format("2006-01-02")
	return yesterdayString
}
func GetTimeHoursToString(t int) string {
	now := time.Now()
	modifiedTime := now.Add(time.Duration(t) * time.Hour)
	formattedTime := modifiedTime.Format("2006-01-02")
	return formattedTime
}
func TimestampToSeg(timestamp int64) int64 {
	if timestamp > 1e10 {
		return timestamp / 1000
	}
	return timestamp
}
func TimestampToTime(timestamp int64) time.Time {
	var currentTime time.Time
	if timestamp > 1e10 {
		currentTime = time.Unix(timestamp/1000, 0)
	} else {
		currentTime = time.Unix(timestamp, 0)
	}

	saoPauloLocation, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		fmt.Println("Erro ao carregar o local de São Paulo:", err)
		return time.Time{}
	}
	saoPauloTime := currentTime.In(saoPauloLocation)
	return saoPauloTime
}
func DurationToString(seconds int64) string {
	duration := time.Duration(seconds) * time.Second
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds = seconds % 60

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}
