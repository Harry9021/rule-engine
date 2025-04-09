package rules

var rules = make(map[string]Rule)

func AddRule(rule Rule) {
	rules[rule.ID] = rule
}

func GetAllRules() []Rule {
	all := []Rule{}
	for _, rule := range rules {
		all = append(all, rule)
	}
	return all
}

func GetRuleByID(id string) (Rule, bool) {
	rule, exists := rules[id]
	return rule, exists
}

func DeleteRule(id string) {
	delete(rules, id)
}
