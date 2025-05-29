package main

import "errors"

// NOTE: Guard Clauses:
// - leverage the ability to return early from a function, to make nested conditions 1-D.
// instead of using if/else chains, we just return early from the function at the end of each conditional block.
func divisor(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("cant divide by zero")
	}
	return dividend / divisor, nil
}

// using guard clauses:
// 1. nested conditional logic
func getInsuranceAmount(status insuranceStatus) int {
	amount := 0
	if !status.hasInsurance() {
		amount = 1
	} else {
		if status.isTotaled() {
			amount = 10000
		} else {
			if status.isDented() {
				amount = 160
				if status.isBigDent() {
					amount = 270
				}
			} else {
				amount = 0
			}
		}
	}
	return amount
}

// 2. guard clauses
func getInsuranceAmount(status insuranceStatus) int {
	if !status.hasInsurance() {
		return 1
	}
	if status.isTotaled() {
		return 10000
	}
	if !status.isDented() {
		return 0
	}
	if status.isBigDent() {
		return 270
	}
	return 160
}
