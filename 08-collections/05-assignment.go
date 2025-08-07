/*
Accept a range from the user and print all the prime numbers between the given range
*/

/*
Refactor the logic for generating prime nos to a "generatePrimes()" function
Write "IsPrime()" that checks if the given number is a prime number or not
Print the generated prime numbers in the main() function

main -> generatedPrimes(range) -> IsPrime(no)
*/

package main

import "fmt"

func main() {
	var start, end int
	fmt.Println("Enter the range :")
	fmt.Scanln(&start, &end)
	primes := generatePrimes(start, end)
	for _, primeNo := range primes {
		fmt.Printf("Prime No : %d\n", primeNo)
	}
}

func generatePrimes(start, end int) []int {
	primes := make([]int, 0)
	for no := start; no <= end; no++ {
		if isPrime(no) {
			primes = append(primes, no)
		}
	}
	return primes
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
