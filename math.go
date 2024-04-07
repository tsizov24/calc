package main

func count(n1, n2 int, operator byte) int {
	switch operator {
	case '+':
		return n1 + n2
	case '-':
		return n1 - n2
	case '*':
		return n1 * n2
	default:
		return n1 / n2
	}
}
