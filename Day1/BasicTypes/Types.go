package main

import "fmt"

func main() {
	//	var (
	//		i int
	//		j int = 3
	//		k     = 4
	//	)
	//	l := 5
	//	fmt.Println(i, j, k, l)

	var (
		b         bool = true //false
		s         string
		漢字        = "Hepp"
		övreGräns = 13

		// Signed int
		i   int // 32 bitar på ett 32-bitarssystem, 64 bitar på ett 64 bitars
		i8  int8
		i16 int16
		i32 int32
		i64 int64

		// Unsigned int
		ui   uint // 32 bitar på ett 32-bitarssystem, 64 bitar på ett 64 bitars
		ui8  uint8
		ui16 uint16
		ui32 uint32
		ui64 uint64

		// Unsigned int pointer
		uiptr uintptr // 32 bitar på ett 32-bitarssystem, 64 bitar på ett 64 bitars

		//float
		f32 float32
		f64 float64

		// Complex
		c64  complex64
		c128 complex128

		// Alias
		by byte // Alias för uint8
		r  rune // Alias för int32
	)
	fmt.Println(b, s, 漢字, övreGräns, i, i8, i16, i32, i64, ui, ui8, ui16, ui32, ui64, uiptr, f32, f64, c64, c128, by, r)
}
