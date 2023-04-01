package main

import "fmt"

func main() {
	/* ⁡⁣⁣⁢HELLO WORLD TRYOUT⁡⁡ */
	// fmt.Println("Hello World")

	/* ⁡⁢⁣⁣NUMBER TRYOUT⁡ */
	// fmt.Println("Satu : ", 1)
	// fmt.Println("Dua  :", 2)
	// fmt.Println("Tiga Koma Lima :", 3.5)

	/* ⁡⁣⁣⁢BOOLEAN TRYOUT⁡ */
	// fmt.Println("Benar = ", true)
	// fmt.Println("Salah = ", false)

	/* STRING TRYOUT */
	// fmt.Println("Calvin")
	// fmt.Println("Calvin Jefen")
	// fmt.Println("Calvin Jefen Kumambong")

	// fmt.Println(len("Calvin"))
	// fmt.Println("Calvin Jefen"[0])
	// fmt.Println("Calvin Jefen Kumambong"[1])

	/* ⁡⁣⁣⁢⁡⁢⁣⁣VARIABLE TRYOUT⁡⁡ */
	// var name string
	// name = "Calvin Jefen"
	// fmt.Println(name)

	// name = "Calvin Kumambong"
	// fmt.Println(name)

	// var myname = "Calvin Jefen1"
	// fmt.Println(myname)

	// myname = "Calvin Kumambong1"
	// fmt.Println(myname)

	// Name := "Calvin Jefen2"
	// fmt.Println(Name)

	// Name = "Calvin Kumambong2"
	// fmt.Println(Name)

	// var (
	// 	firstname = "Calvin Jefen"
	// 	lastname  = "Kumambong"
	// )

	// fmt.Println(firstname)
	// fmt.Println(lastname)

	/* ⁡⁣⁣⁢CONSTANT TRYOUT⁡⁡ */
	// const firstname string = "Calvin Jefen"
	// const lastname = "Kumambong"
	// const value = 1000

	// const (
	// 	friend string = "Lucky"
	// 	Mother        = "Yes"
	// )

	// fmt.Println(firstname)
	// fmt.Println(lastname)
	// fmt.Println(value)

	/* ⁡⁣⁣⁢⁡⁢⁣⁣DATA TYPE CONVERSION TRYOUT⁡⁡ */

	// var nilai32 int32 = 100000
	// var nilai64 int64 = int64(nilai32)

	// var nilai8 int8 = int8(nilai32)

	// fmt.Println(nilai32)
	// fmt.Println(nilai64)
	// fmt.Println(nilai8)

	// var name = "Calvin"
	// var C byte = name[0]
	// var Cstring string = string(C)

	// fmt.Println(name)
	// fmt.Println(C)
	// fmt.Println(Cstring)

	/* ⁡⁣⁣⁢DATA TYPE DECLARATION TRYOUT ⁡*/
	// type NoKTP string
	// type Married bool

	// var noktpCal NoKTP = "065738232834"
	// var marriedstatus Married = false

	// fmt.Println(noktpCal)
	// fmt.Println(marriedstatus)

	/* MATHEMATICAL OPERATION TRYOUT */
	// var result = 10 + 10
	// fmt.Println(result)

	// var a = 10
	// var b = 10
	// var c = a * b
	// fmt.Println(c)

	// var i = 10
	// i += 10
	// fmt.Println(i)

	// i++ //i = i+1
	// fmt.Println(i)

	// var negative = -100
	// var positive = 100
	// fmt.Println(negative)
	// fmt.Println(positive)

	/* ⁡⁢⁣⁣COMPARATION OPERATION TRYOUT⁡ */
	// var name1 = "Calvin"
	// var name2 = "Calvin"

	// var result = name1 == name2
	// fmt.Println(result)

	// var name3 = "A"
	// var name4 = "B"

	// var result2 = name3 > name4
	// fmt.Println(result2)

	// var value1 = 100
	// var value2 = 200

	// fmt.Println(value1 > value2)
	// fmt.Println(value1 < value2)
	// fmt.Println(value1 == value2)
	// fmt.Println(value1 != value2)

	/* ⁡⁣⁣⁢BOOLEAN OPERATION TRYOUT⁡ */
	// var ujian = 80
	// var absensi = 80

	// var lulusUjian = ujian > 80
	// var lulusAbsensi = absensi > 80

	// var lulus = lulusAbsensi && lulusUjian
	// fmt.Println(lulus)

	/* ⁡⁢⁣⁣ARRAY DATA TYPE TRYOUT⁡ */
	// var names [3]string

	// names[0] = "Calvin"
	// names[1] = "Jefen"
	// names[2] = "Kumambong"

	// // var calvin = names[0]
	// fmt.Println(names[0])
	// fmt.Println(names[1])
	// fmt.Println(names[2])

	// var values = [3]int{
	// 	90,
	// 	95,
	// 	80,
	// }
	// fmt.Println(values)

	// fmt.Println(len(names))
	// fmt.Println(len(values))

	/* ⁡⁣⁣⁢SLICE DATA TYPE TRYOUT⁡ */
	// var months = [...]string{
	// 	"Januari",
	// 	"Februari",
	// 	"Maret",
	// 	"April",
	// 	"Mei",
	// 	"Juni",
	// 	"Juli",
	// 	"Agustus",
	// 	"September",
	// 	"Oktober",
	// 	"November",
	// 	"December",
	// }

	// var slice1 = months[4:7]
	// fmt.Println(slice1)
	// fmt.Println(len(slice1))
	// fmt.Println(cap(slice1))

	// var slice2 = months[10:]
	// fmt.Println(slice2)

	// var slice3 = append(slice2, "Calvin")
	// fmt.Println(slice3)
	// slice3[1] = "Bukan Desember"
	// fmt.Println(slice3)

	// fmt.Println(slice2)
	// fmt.Println(months)

	// newslice := make([]string, 2, 5)

	// newslice[0] = "Calvin"
	// newslice[1] = "Jefen"

	// fmt.Println(newslice)
	// fmt.Println(len(newslice))
	// fmt.Println(cap(newslice))

	// copySlice := make([]string, len(newslice), cap(newslice))
	// copy(copySlice, newslice)
	// fmt.Println(copySlice)

	/* ⁡⁢⁣⁣MAP DATA TYPE TRYOUT⁡ */
	fmt.Println()
}
