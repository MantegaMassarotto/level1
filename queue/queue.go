package queue

var numbers = []int{}

func Enqueue(number int) []int {
	numbers = append(numbers, number)
	return numbers
}

func Dequeue() []int {
	if len(numbers) > 0 {
		numbers = append(numbers[:0], numbers[0+1:]...)
	}
	return numbers
}
