package rules

import (
	"encoding/json"
	"log"
	"net/http"
	"rule-engine-backend/db"
	"rule-engine-backend/engine"
)

func CreateRule(w http.ResponseWriter, r *http.Request) {
	var rule Rule
	if err := json.NewDecoder(r.Body).Decode(&rule); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dbConn, err := db.ConnectDB()
	if err != nil {
		http.Error(w, "Failed to connect to DB", http.StatusInternalServerError)
		return
	}
	defer dbConn.Close()

	if err := AddRule(dbConn, rule); err != nil {
		http.Error(w, "Failed to add rule", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	log.Printf("Rule created: %+v\n", rule)
}

func GetRules(w http.ResponseWriter, r *http.Request) {
	dbConn, err := db.ConnectDB()
	if err != nil {
		http.Error(w, "Failed to connect to DB", http.StatusInternalServerError)
		return
	}
	defer dbConn.Close()

	rules, err := GetAllRules(dbConn)
	if err != nil {
		http.Error(w, "Failed to get rules", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(rules)
}

func HandleEvent(w http.ResponseWriter, r *http.Request) {
	var event map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dbConn, err := db.ConnectDB()
	if err != nil {
		http.Error(w, "Failed to connect to DB", http.StatusInternalServerError)
		return
	}
	defer dbConn.Close()

	rules, err := GetAllRules(dbConn)
	if err != nil {
		http.Error(w, "Failed to get rules", http.StatusInternalServerError)
		return
	}

	for _, rule := range rules {
		if engine.EvaluateCondition(rule.Condition, event) {
			engine.ExecuteAction(rule.Action)
		}
	}

	w.WriteHeader(http.StatusOK)
}
