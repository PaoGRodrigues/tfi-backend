package alert

import (
	"encoding/json"
	"fmt"
)

const cybersecurity = "Cybersecurity"

func ParseAlerts(alerts []Alert) []string {
	messages := []string{}

	for _, alert := range alerts {
		if alert.Category == cybersecurity {
			b, err := json.Marshal(alert)
			if err != nil {
				fmt.Println(err)
			}
			messages = append(messages, string(b))
		}
	}
	return messages

}
