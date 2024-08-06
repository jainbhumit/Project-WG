package component

import (
	"bufio"
	"encoding/json"
	"file/models"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

const dailyStatusFile = "daily_status.json"

// ReadDailyStatus reads daily status from a file
func ReadDailyStatus() (map[string][]models.DailyStatus, error) {
	statusMap := make(map[string][]models.DailyStatus)
	file, err := os.Open(dailyStatusFile)
	if err != nil {
		if os.IsNotExist(err) {
			return statusMap, nil // File doesn't exist, return empty map
		}
		return nil, err
	}
	defer file.Close()

	var statusList []struct {
		Username      string               `json:"username"`
		DailyStatuses []models.DailyStatus `json:"daily_statuses"`
	}

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&statusList); err != nil {
		return nil, err
	}

	for _, s := range statusList {
		statusMap[s.Username] = s.DailyStatuses
	}

	return statusMap, nil
}

// writes daily status to a file
func WriteDailyStatus(statusMap map[string][]models.DailyStatus) error {
	var statusList []struct {
		Username      string               `json:"username"`
		DailyStatuses []models.DailyStatus `json:"daily_statuses"`
	}

	for username, statuses := range statusMap {
		statusList = append(statusList, struct {
			Username      string               `json:"username"`
			DailyStatuses []models.DailyStatus `json:"daily_statuses"`
		}{
			Username:      username,
			DailyStatuses: statuses,
		})
	}

	data, err := json.MarshalIndent(statusList, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(dailyStatusFile, data, 0644)
}

// UpdateDailyStatus updates the daily status for a user
func UpdateDailyStatus(username string) error {
	statusMap, err := ReadDailyStatus()
	if err != nil {
		return err
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the new status:")
	status, _ := reader.ReadString('\n')
	status = strings.TrimSpace(status)         // Trim whitespace
	now := time.Now().Truncate(24 * time.Hour) // Truncate to the start of the day
	newStatus := models.DailyStatus{
		Date:   now,
		Status: status,
	}

	statuses := statusMap[username]
	for i, s := range statuses {
		if s.Date.Equal(now) {
			// Update the existing status for today
			statuses[i] = newStatus
			statusMap[username] = statuses
			return WriteDailyStatus(statusMap)
		}
	}

	// Add new status if it doesn't exist for today
	statusMap[username] = append(statuses, newStatus)
	return WriteDailyStatus(statusMap)
}

// ShowDailyStatus displays the daily status for a user
func ShowDailyStatus(username string) {
	statusMap, err := ReadDailyStatus()
	if err != nil {
		fmt.Println("Error reading daily status:", err)
		return
	}

	statuses, ok := statusMap[username]
	if !ok {
		fmt.Println("No daily status found for user:", username)
		return
	}

	fmt.Printf("Daily status for %s:\n", username)
	for _, s := range statuses {
		fmt.Printf("Date: %s\n", s.Date.Format("2006-01-02"))
		fmt.Printf("Status: %s\n", s.Status)
	}
}
