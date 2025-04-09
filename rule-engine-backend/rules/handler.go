package rules

import (
	"encoding/json"
	"log"
	"net/http"
	"rule-engine-backend/engine"
)

func CreateRule(w http.ResponseWriter, r *http.Request) {
	var rule Rule
	if err := json.NewDecoder(r.Body).Decode(&rule); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	AddRule(rule)
	w.WriteHeader(http.StatusCreated)
	log.Printf("Rule created: %+v\n", rule)
}

func GetRules(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(GetAllRules())
}

func HandleEvent(w http.ResponseWriter, r *http.Request) {
	var event map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, rule := range GetAllRules() {
		if engine.EvaluateCondition(rule.Condition, event) {
			engine.ExecuteAction(rule.Action)
		}
	}

	w.WriteHeader(http.StatusOK)
}
