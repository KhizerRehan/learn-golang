package refresh

func Calculate(a int, b int) []float64 {
  // your code goes here
  // sum/diff/product/quotient

  sum := float64(a+b) 
  diff := float64(a-b)
  product := float64(a*b)
  quotient := float64(a/b)

  results := []float64{sum, diff, product, quotient}

	return results
}
