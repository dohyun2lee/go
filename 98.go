package main

import ( 
	"fmt"
)

func main() {
	var N, a, sum, mid, round, d, ss int
	cnt := 1
	var A []int
	var count []int
	var B []int

	fmt.Println("N?")
	fmt.Scanln(&N)

	for i:=0;i<N;i++ {
		fmt.Scanln(&a)

		A = append(A, a)
	}

	for i:=0;i<N;i++ {
		for j:=i;j<N;j++ {
			if A[i]>A[j] {
				tmp := A[i]
				A[i] = A[j]
				A[j] = tmp
			} else {
				continue
			}
		}
	}

	for i:=0;i<N;i++ {
		sum += A[i]
	}

	mid = N/2 

	if ((sum*10)/N%10) > 4 {
		round = ((sum*10)/N+10)/10
	} else {
		round = sum / N
	}

	for i:=0;i<N-1;i++ {
		if A[i] == A[i+1] {
			cnt ++
		} else {
			count = append(count, cnt)
			cnt = 1
			B = append(B, A[i])
		}
	}

	for i:=0;i<len(count);i++ {
		c := 0
		if c < count[i] {
			c = count[i]
			d = i
		}
	}

	for i:=0;i<=d;i++ {
		ss += count[i]
	}


	fmt.Println(round)
	fmt.Println(A[mid])
	fmt.Println(A[ss])
	fmt.Println(A[N-1]-A[0])


}