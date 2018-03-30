package myprime

import (
	"math"
)

const MAXN int = 1000
const PRIME_NUM int = 168

type MyPrime struct {
	pflag [MAXN + 1]bool
	prime [PRIME_NUM + 1]int
}

func (mp *MyPrime) ESieve(n int) {
	mp.pflag[2] = true
	for i := 3; i <= n; i += 2 {
		mp.pflag[i] = true
	}

	end := int(math.Sqrt(float64(n)))

	for i := 3; i <= end; i += 2 {
		step := i + i
		if mp.pflag[i] {
			j := i + step
			for j <= n {
				mp.pflag[j] = false
				j += step
			}
		}
	}

	j := 2
	mp.prime[0] = 1
	mp.prime[1] = 2
	for i := 3; i <= n; i += 2 {
		if mp.pflag[i] {
			mp.prime[j] = i
			j++
		}
	}
}

func (mp *MyPrime) Init() {
	mp.ESieve(MAXN)
}

func (mp *MyPrime) GetPrime(n int) int {
	return mp.prime[n]
}

func (mp *MyPrime) GetPrimeNum() int {
	return PRIME_NUM
}

func ModPow(a int, b int, mod int) int {
	var res int
	res = 1
	for i := b; i > 0; i-- {
		res = res * a
		if res > mod {
			res = res % mod
		}
	}
	return res
}
func GcdEx(a int, b int, x *int, y *int) int {
	var r int

	if b == 0 {
		*x = 1
		*y = 0
		return a
	}
	r = GcdEx(b, a%b, x, y) /* r = GCD(a, b) = GCD(b, a%b) */
	t := *x
	*x = *y
	*y = t - ((a / b) * (*y))
	return r

}
