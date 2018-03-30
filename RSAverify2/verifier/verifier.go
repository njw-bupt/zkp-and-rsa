package verifier

import (
	"RSAverify2/myprime"
	"math/rand"
)

type Verifier struct {
	P          int
	Q          int
	N          int
	Nn         int
	Orla       int
	PublicKey  int
	PrivateKey int
	A          int
	B          int
	C          int
	Stext      int
	Dtext      int
}

func Gen_rand() int {
	return rand.Intn(1000)
}

func Cal(a int, b int) int {
	return a ^ b
}

//10-1000的素数
func (r *Verifier) Gen_key() {
	mp := new(myprime.MyPrime)
	mp.Init()
	for {
		num1 := rand.Intn(mp.GetPrimeNum())
		r.P = int(mp.GetPrime(num1))
		for {
			num1 := rand.Intn(mp.GetPrimeNum())
			r.Q = int(mp.GetPrime(num1))
			if r.Q != r.P {
				break
			}
		}
		if r.Q > 10 && r.P > 10 {
			break
		}
	}
	r.N = r.P * r.Q

	r.Orla = (r.P - 1) * (r.Q - 1)
	for {
		num1 := rand.Intn(mp.GetPrimeNum())
		r.PublicKey = int(mp.GetPrime(num1))
		//r.PublicKey = rand.Intn(r.Orla)
		if r.Orla%r.PublicKey != 0 {
			break
		}
	}

	//扩展欧几里得求私钥
	var y int
	myprime.GcdEx(r.PublicKey, r.Orla, &r.PrivateKey, &y)
	if r.PrivateKey < 0 {

		r.PrivateKey = (r.PrivateKey%r.Orla + r.Orla) % r.Orla
	}
}

func (r *Verifier) Encrypt(m int) {
	r.Stext = myprime.ModPow(m, r.PublicKey, r.N)
}

func (r *Verifier) Dcode(c int) {
	r.Dtext = myprime.ModPow(c, r.PrivateKey, r.Nn)
}

func (r *Verifier) Judge() bool {
	if r.B == r.Dtext {
		return true
	} else {
		return false
	}
}
