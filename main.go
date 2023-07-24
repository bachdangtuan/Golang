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
			fmt.Println("Số i là", i)
		} else {
			fmt.Println("Đây là ", i)
		}

	}
}

func forWhile() {
	var i int = 2
	for i < 10 {
		sum := i + 1
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

// Giai Phuong Trinh bac 2

func tinhDelta(a float64, b float64, c float64) float64 {

	return b*b - (4 * a * c)

}

func PhuongTrinhBac2() {
	type ThamSo struct {
		a float64
		b float64
		c float64
	}

	type nghiem struct {
		x1 float64
		x2 float64
	}

	cacThamSo := ThamSo{
		a: 12,
		b: 500,
		c: 40,
	}

	delta := tinhDelta(cacThamSo.a, cacThamSo.b, cacThamSo.c)

	if delta < 0 {
		fmt.Println("Phương trình vô nghiệm")
	} else if delta > 0 {
		ketQua := nghiem{
			x1: (-(cacThamSo.b) + math.Sqrt(delta)) / 2 * cacThamSo.a,
			x2: (-(cacThamSo.b) - math.Sqrt(delta)) / 2 * cacThamSo.a,
		}
		fmt.Println("Phương trình có 2 nghiệm là ", ketQua)
	} else {
		ketQua := nghiem{
			x1: -(cacThamSo.b) / 2 * cacThamSo.a,
			x2: -(cacThamSo.b) / 2 * cacThamSo.a,
		}
		fmt.Println("Phương trình có nghiệm kép", ketQua)
	}

	fmt.Println("Gia Tri delta la", tinhDelta(cacThamSo.a, cacThamSo.b, cacThamSo.c))
}

// Pointer to structs
func Pointer() {
	// pointer thường (giá trị con trỏ được lấy từ biến)
	x := 10
	p := &x
	fmt.Println("con trỏ p có giá trị là:", *p)
	// Gán giá trị của cho biến từ con trỏ
	*p = 20
	fmt.Println("Giá trị mới của x:", x)

	// Con trỏ của struct
	type structPointer struct {
		x int
		y int
	}
	// 19,371,673 hits

	demoPointer := structPointer{
		x: 50,
		y: 100,
	}
	pointer := &demoPointer
	fmt.Println("Đọc con trỏ:", pointer)

}

// Struc Literals
func StructLiterals() {
	type Vertex struct {
		X, Y int
	}

	var (
		v1 = Vertex{1, 2} // has type Vertex
		v2 = Vertex{X: 1} // Y:0 is implicit
		v3 = Vertex{}     // X:0 and Y:0
		p  = Vertex{1, 2} // has type *Vertex
	)
	fmt.Println(v1, p, v2, v3)
}

func array() {
	// Khai báo 1 mảng trên arrayDemo với số lượng index là 2
	var arrayDemo [2]string
	// Gán arrayDemo
	arrayDemo[0] = "ukraine"
	arrayDemo[1] = "russia"
	// Đọc cả mảng và đọc các phần tử trong mảng
	fmt.Println(arrayDemo, arrayDemo[0])
	primes := [6]any{2, 3, 5, "việt nam", 11, 13}
	fmt.Println(primes)
}

func sliceArrays() {
	primes := [6]any{2, 3, 5, "việt nam", 11, 13}

	// tạo ra 1 biến mới để hứng mảng được cắt
	var sliceArray []any = primes[1:4]

	fmt.Println(sliceArray)
}

func SlicesReferences() {
	// Cho mảng
	name := [4]any{"Nguyễn", "Chó", "Văn", "Mèo"}
	a := name[0:1]
	b := name[1:2]
	// Đọc mảng
	fmt.Println(name, a, b)
	b[0] = "XXX"
	a[0] = "Vãi lz"
	fmt.Println(name, b, a)

}

func sliceLiterals() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

func sliceDefault() {
	s := []int{2, 3, 5, 7, 11, 13}
	s = s[:4]
	fmt.Println(s)
}

func infoArray() {
	name := [4]string{"Loz", "Mẹ", "Mày", "To"}
	// lấy độ dài
	lengthName := len(name)
	capName := cap(name)
	fmt.Println(lengthName, capName)
}

// Make
func makeSlice() {
	// Tạo 1 mảng rỗng có độ dài bằng 5
	a := make([]int, 5)
	b := make([]int, 0, 5)
	fmt.Println(a, b)

}

// Gắn vào
func AppendArray() {
	array := []any{}
	array = append(array, 1, 2, 3, 4, 5, 6, "sting")
	fmt.Println(array)
}

func main() {
	fmt.Println("Gô Lăng 2")
	AppendArray()
	// makeSlice()

	// infoArray()
	// sliceDefault()
	// sliceLiterals()
	// SlicesReferences()
	// sliceArrays()
	// array()
	// StructLiterals()
	// Pointer()
	// PhuongTrinhBac2()
	// forWhile()
	// forTiepDien2()
	// forTiepDien()
	// fmt.Println(tinhTong(4, 5))
	// tinhDienTichTamGiac(5,6,5)
	// // docso()
	// CheckSoThuTu()
}
