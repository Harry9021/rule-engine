package rules

import (
	"database/sql"
)

func AddRule(db *sql.DB, rule Rule) error {
	_, err := db.Exec("INSERT INTO rules (id, condition, action) VALUES (?, ?, ?)", rule.ID, rule.Condition, rule.Action)
	return err
}

func GetAllRules(db *sql.DB) ([]Rule, error) {
	rows, err := db.Query("SELECT id, condition, action FROM rules")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rules := []Rule{}
	for rows.Next() {
		var rule Rule
		if err := rows.Scan(&rule.ID, &rule.Condition, &rule.Action); err != nil {
			return nil, err
		}
		rules = append(rules, rule)
	}
	return rules, nil
}

func GetRuleByID(db *sql.DB, id string) (Rule, bool, error) {
	var rule Rule
	err := db.QueryRow("SELECT id, condition, action FROM rules WHERE id = ?", id).Scan(&rule.ID, &rule.Condition, &rule.Action)
	if err != nil {
		if err == sql.ErrNoRows {
			return Rule{}, false, nil
		}
		return Rule{}, false, err
	}
	return rule, true, nil
}

func UpdateRuleInDB(db *sql.DB, rule Rule) error {
	_, err := db.Exec("UPDATE rules SET condition = ?, action = ? WHERE id = ?", rule.Condition, rule.Action, rule.ID)
	if err != nil {
		return err
	}
	return nil
}


func DeleteRule(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE FROM rules WHERE id = ?", id)
	return err
}
