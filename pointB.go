package main
import "fmt"

type elektronik struct {
	ID      int
	nama    string
	ruangan string
	watt    float64
	durasi  float64
}

type dataElektronik[1000]elektronik
func input(T *dataElektronik, n int){
	for i := 0; i < n ; i++ {
		fmt.Scan(&T[i].ID)
	}
}