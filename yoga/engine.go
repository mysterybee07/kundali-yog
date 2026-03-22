package yoga

func AnyCondition(conditions ...Condition) Condition {
	return func(chart Chart) bool {
		for _, condition := range conditions {
			if condition != nil && condition(chart) {
				return true
			}
		}
		return false
	}
}

func AllConditions(conditions ...Condition) Condition {
	return func(chart Chart) bool {
		for _, condition := range conditions {
			if condition == nil || !condition(chart) {
				return false
			}
		}
		return true
	}
}
