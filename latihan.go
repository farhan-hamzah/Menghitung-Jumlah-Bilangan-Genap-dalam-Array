package main
import "fmt"
const NMAX int = 100
func main(){
	var A[NMAX]int
	var n, i int
	fmt.Scan(&n)
	for i = 0; i <n; i++{
		fmt.Scan(&A[i])
	}
	hasil := findNumber(n, A[:n])
	fmt.Print(hasil)
}
func findNumber(n int, A[] int)int{
	var i, genap int
	for i = 0; i<n; i++{
		if A[i]%2 == 0{
			genap++
		}
	}
	return genap
}