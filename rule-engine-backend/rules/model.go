package rules

type Rule struct {
	ID        string `json:"id"`
	Condition string `json:"condition"` // Example: "temp > 40"
	Action    string `json:"action"`    // Example: "alert('High Temp')"
}
