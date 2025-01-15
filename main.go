package main

/*
DYNAMIC PROGRAMMING

Grashopper

1 2 3 4 5 6 7 ... n
A grashopper can jump either 1 or 2 cells ahead,
How many ways (K) are there to reach the N-th cell ?

Kn = Kn-1 + Kn-2
*/
import (
	"fmt"
)

func numberOfTrajectories(n int) int{
	var K = make([]int, n+1)
	K[0] = 0
	K[1] = 1
	for i := 2; i <= n; i++ {
		K[i] = K[i-1] + K[i-2]
	}
	return K[n]
}

func main() {

	var finish int;
	fmt.Scan(&finish)
	fmt.Printf("Grashopper has %d trajectories from 1 to %d\n", numberOfTrajectories(finish),  finish)
	
	//5
    //Grashopper has 5 trajectories from 1 to 5 

}
