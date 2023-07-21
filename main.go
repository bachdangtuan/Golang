package main

import (
	"fmt"
	"math"
)

func tinhTong(a int, b int) any {
	return a + b
}

var dientichtamgiac any

// Tính Diện Tích Tam Giác bất kỳ theo công thức Heroin
func tinhDienTichTamGiac(a float64, b float64, c float64) {
	var p float64
	p = (a + b + c) / 2
	dientichtamgiac = math.Sqrt(p * (p - a) * (p - b) * (p - c))
	fmt.Println("Diện tích tam giác là", dientichtamgiac)
}

// Vòng lặp for cơ bản
func docso() {
	for i := 0; i < 10; i++ {
		fmt.Println("In ra các số từ 1 đến 10", i)
	}
}

func CheckSoThuTu() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("Những con hàng chia hết cho 2", i)
		} else {
			fmt.Println(i, "Đéo được rồi")
		}
	}
}

// Vòng for tiếp diễn
func forTiepDien() {
	i := 1
	for i < 5 {
		fmt.Println("nhỏ hơn 5 thì luôn đúng nên chạy vòng lặp vô hạn")
	}
}
func forTiepDien2() {
	var i int = 6
	for ; i < 10; i++ {
		if i+2 <= 10 {
			fmt.Println("Số i là",i)
		} else {
			fmt.Println("Đây là ", i)
		}

	}
}

func forWhile()  {
	var i int = 2
	for i < 10 {
		sum:= i + 1
		switch sum {
			case 3:
				fmt.Println("Số là", i)
			case 7:
				fmt.Println("Số là", i)
			default:
				fmt.Println("Sai")
		}
		// i++
	}
}

func main() {
	fmt.Println("Gô Lăng")
	// forWhile()
	// forTiepDien2()
	// forTiepDien()
	// fmt.Println(tinhTong(4, 5))
	// tinhDienTichTamGiac(5,6,5)
	// // docso()
	// CheckSoThuTu()
}
