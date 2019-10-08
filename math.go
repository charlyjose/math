package math

// GCD returns the highest Common Factor of a number
func GCD(n1, n2 int) int {
	a := Factors(n1)
	b := Factors(n2)
	large := -1
	l1 := len(a)
	l2 := len(b)

	switch l1 > l2 {
	case false:
		{ // l1<l2
			for i := l2 - 1; i > -1; i-- {
				for j := l1 - 1; j > -1; j-- {
					if b[i] == a[j] {
						if large < a[i] {
							large = a[i]
						}
					}
				}
			}
		}
	case true:
		{ // l1>l2
			for i := l1 - 1; i > -1; i-- {
				for j := l2 - 1; j > -1; j-- {
					if a[i] == b[j] {
						if large < a[i] {
							large = a[i]
						}
					}
				}
			}
		}
	}
	return large
}

// Factors return array of factors of a number
func Factors(n int) []int {
	half := n / 2
	j := 0
	a := [...]int{}
	d := a[:]
	for i := 1; i <= half; i++ {
		if n%i == 0 {
			d = append(d, i)
			j++
		}
	}
	d = append(d, n)
	return d
}
