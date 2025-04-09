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

func UpdateRule(w http.ResponseWriter, r *http.Request) {
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

	_, exists, err := GetRuleByID(dbConn, rule.ID)
	if err != nil {
		http.Error(w, "Failed to check rule existence", http.StatusInternalServerError)
		return
	}
	
	if !exists {
		http.Error(w, "Rule not found", http.StatusNotFound)
		return
	}

	if err := UpdateRuleInDB(dbConn, rule); err != nil {
		http.Error(w, "Failed to update rule", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	log.Printf("Rule updated: %+v\n", rule)
}

func DeleteRuleHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Rule ID is required", http.StatusBadRequest)
		return
	}

	dbConn, err := db.ConnectDB()
	if err != nil {
		http.Error(w, "Failed to connect to DB", http.StatusInternalServerError)
		return
	}
	defer dbConn.Close()

	_, exists, err := GetRuleByID(dbConn, id)
	if err != nil {
		http.Error(w, "Failed to check rule existence", http.StatusInternalServerError)
		return
	}
	
	if !exists {
		http.Error(w, "Rule not found", http.StatusNotFound)
		return
	}

	if err := DeleteRule(dbConn, id); err != nil {
		http.Error(w, "Failed to delete rule", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	log.Printf("Rule deleted: %s\n", id)
}

