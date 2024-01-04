package tax

func taxCalculator(salary float64) float64 {

	if salary > 50000.0 {
		return salary * (0.83)

	} else if salary > 150000.0 {
		return salary * (0.73)

	} else {
		return salary
	}

}
