package notes

func notesAvg(values ...int) float64 {
	var sum int

	for _, value := range values {

		if value < 0 {
			return -1

		} else {
			sum += value
		}

	}

	avgResult := float64(sum) / float64(len(values))

	return float64(avgResult)

}
