package engine

import (
	"fmt"
	"strings"
)

func ExecuteAction(action string) {
	if strings.HasPrefix(action, "alert") {
		fmt.Println("[ALERT]:", action)
	} else {
		fmt.Println("[ACTION EXECUTED]:", action)
	}
}
