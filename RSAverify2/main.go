// RSAverify2 project main.go
package main

import (
	"RSAverify2/prover"
	"RSAverify2/verifier"
	"fmt"
	"math/rand"
	"time"

	"github.com/gansidui/priority_queue"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	startTime := prover.Get_currentTime()

	mq := priority_queue.New()

	//生成泊松业务流mq
	for i := 0; i < 100; i++ {
		//分配两对密钥对
		var Alice prover.Prover
		Alice.Gen_key()
		Alice.ArrivalTime = int64(prover.Gen_poissonTime(1000)) + prover.Get_currentTime()
		mq.Push(&Alice)
	}

	for mq.Len() > 0 {
		Alice := mq.Pop().(*prover.Prover)
		for Alice.ArrivalTime > prover.Get_currentTime() {

		}
		var Bob verifier.Verifier
		Bob.Gen_key()
		temp := Bob.PrivateKey
		Bob.PrivateKey = Alice.PrivateKey
		Alice.PrivateKey = temp

		Alice.Nn = Bob.N
		Bob.Nn = Alice.N

		//开始密钥协商
		Alice.A = prover.Gen_rand()
		Alice.Encrypt(Alice.A)
		prover.Delay(10000)
		Bob.Dcode(Alice.Stext)
		Bob.B = verifier.Gen_rand()
		Bob.C = verifier.Cal(Bob.B, Bob.Dtext)
		Bob.Encrypt(Bob.C)
		prover.Delay(10000)
		Alice.Dcode(Bob.Stext)
		Alice.B = prover.Cal(Alice.A, Alice.Dtext)
		Alice.Encrypt(Alice.B)
		prover.Delay(10000)
		Bob.Dcode(Alice.Stext)
		if Bob.Judge() == true {
			fmt.Println("check success")
		}
		fmt.Println(Bob.B)
		fmt.Println(Alice.B)
	}
	endTime := prover.Get_currentTime()
	fmt.Println("running time of the whole program ", endTime-startTime)
}
