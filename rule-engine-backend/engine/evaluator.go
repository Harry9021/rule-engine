package engine

import (
	"github.com/Knetic/govaluate"
)

func EvaluateCondition(condition string, event map[string]interface{}) bool {
	expr, err := govaluate.NewEvaluableExpression(condition)
	if err != nil {
		return false
	}
	result, err := expr.Evaluate(event)
	if err != nil {
		return false
	}
	return result == true
}
