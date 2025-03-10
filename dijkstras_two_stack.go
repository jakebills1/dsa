package dsa

import (
	"log"
	"slices"
	"strconv"
)

// evaluates expressions of the form ( 1 + ( ( 2 + 3 ) * ( 4 * 5 ) ) )
func evaluate(expr string) int {
	validOperators := []rune{'+', '-', '*', '/'}
	digits := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	ops := newArrayStack[rune]()
	vals := newArrayStack[int]()
	// push operators to ops
	// and values to vals
	// when operator is closing paren:
	// pop the last operator and the last two values,
	// apply the operation,
	// and push the result to vals
	for _, token := range expr {
		if token == ')' {
			// evaluate
			op, oErr := ops.pop()
			operand, vErr := vals.pop()
			if vErr != nil || oErr != nil {
				log.Fatalf("error in algorithm. expr = '%s', token = %#U, ops.pop() err = %s, vals.pop() err = %s", expr, token, oErr, vErr)
			}
			var result int
			switch op {
			case '+':
				operandTwo, o2Err := vals.pop()
				if o2Err != nil {
					log.Fatalf("error in algorithm. expr = '%s', token = %#U, vals.pop() err = %s", expr, token, o2Err)
				}
				result = operandTwo + operand
			case '-':
				operandTwo, o2Err := vals.pop()
				if o2Err != nil {
					log.Fatalf("error in algorithm. expr = '%s', token = %#U, vals.pop() err = %s", expr, token, o2Err)
				}
				result = operandTwo + operand
			case '*':
				operandTwo, o2Err := vals.pop()
				if o2Err != nil {
					log.Fatalf("error in algorithm. expr = '%s', token = %#U, vals.pop() err = %s", expr, token, o2Err)
				}
				result = operandTwo * operand
			case '/':
				operandTwo, o2Err := vals.pop()
				if o2Err != nil {
					log.Fatalf("error in algorithm. expr = '%s', token = %#U, vals.pop() err = %s", expr, token, o2Err)
				}
				result = operandTwo / operand
			}
			vals.push(result)
		} else if slices.Contains(validOperators, token) {
			ops.push(token)
		} else if slices.Contains(digits, token) {
			d, atoiErr := strconv.Atoi(string(token))
			if atoiErr != nil {
				log.Fatalf("error in algorithm. expr = '%s', token = %#U, atoi err = %s", expr, token, atoiErr)
			}
			vals.push(d)
		}
		// whitespace and open parens are no-ops
	}
	evaluation, _ := vals.pop()
	return evaluation
}
