package questoes

func isFibonacci(number int) bool {
	fibonacci := []int{0, 1}
	for fibonacci[len(fibonacci)-1] < number {

		fibonacci = append(fibonacci, fibonacci[len(fibonacci)-1]+fibonacci[len(fibonacci)-2])
	}

	if number == fibonacci[len(fibonacci)-1] || number == 0 {
		return true
	} else {
		return false
	}
}
