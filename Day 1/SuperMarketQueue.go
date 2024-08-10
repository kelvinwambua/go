/*
There is a queue for the self-checkout tills at the supermarket. Your task is write a function to calculate the total time required for all the customers to check out!

input
customers: an array of positive integers representing the queue. Each integer represents a customer, and its value is the amount of time they require to check out.
n: a positive integer, the number of checkout tills.
output
The function should return an integer, the total time required.
*/
package kata

func QueueTime(customers []int, n int) int {
	if n <= 0 {
		return 0
	}

	if len(customers) == 0 {
		return 0
	}
	tills := make([]int, n)

	for _, customer := range customers {
		minIndex := 0
		for i := 1; i < n; i++ {
			if tills[i] < tills[minIndex] {
				minIndex = i
			}
		}
		tills[minIndex] += customer
	}

	maxTime := 0
	for _, time := range tills {
		if time > maxTime {
			maxTime = time
		}
	}

	return maxTime
}
