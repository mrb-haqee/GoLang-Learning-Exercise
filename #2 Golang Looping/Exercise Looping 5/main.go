package main

import "fmt"

func ReverseWord(str string) string {

	cast := []rune(str)
	n := len(cast)
	hasil := []rune{}

	for i := 0; i < n; i++ { // main function
		if cast[i] == rune(32) { //mengecek spasi
			// fmt.Println("ini i ke-", i)
			j := i - 1
			for { // eksekusi
				if cast[j] == rune(32) { //mengecek spasi di awal kata misal: <spasi>.k.a.t.a
					// fmt.Printf("Print huruf-%s pada indeks j ke-%d\n", string(cast[j]), j)
					hasil = append(hasil, rune(32))
					break
				} else if j == 0 { // mengecek kata di awal
					// fmt.Printf("Print huruf-%s pada indeks j ke-%d\n", string(cast[j]), j)
					hasil = append(hasil, rune(cast[j]))
					hasil = append(hasil, rune(32))
					break
				} else {
					// fmt.Printf("Print huruf-%s pada indeks j ke-%d\n", string(cast[j]), j)
					hasil = append(hasil, cast[j])
					j--
				}
				// fmt.Println("index j ke-", j)
			}

		} else if i == n-1 {
			// fmt.Println("ini i ke-", i)
			j := i
			for {
				if cast[j] == rune(32) { // mengecek kata di akhir
					break
				} else {
					hasil = append(hasil, cast[j])
					j--
				}
			}
		}

	}
	// fmt.Println(hasil)
	// fmt.Println(string(hasil))
	// fmt.Println(int(hasil[0])-32)
	// hasil[0]=rune(int(hasil[0])-32)
	// hasil[3]=rune(int(hasil[3])+32)
	// fmt.Println(string(hasil))
	for i := 0; i < 2; i++ {
		n:= check(int(hasil[i]))
		m:= check(int(hasil[i+1]))
		if n == true && m == true {
			break
		} else {
			for i := 0; i < len(hasil); i++ { // mengganti huruf besar dan kecil
				uppercase := check(int(hasil[i]))

				// space ASCII = 32
				// space ASCII = A-65 Z-90, a-97 z-122

				if uppercase == true {
					j := i
					if j == 0 {
						continue
					}

					for j >= 0 {
						// fmt.Println("indeks j ke-", j)
						if j == i {
							hasil[j] = rune(int(hasil[j]) + 32)
						} else if j == 0 {
							hasil[j] = rune(int(hasil[j]) - 32)
							break
						} else if hasil[j] == rune(32) {
							j++
							hasil[j] = rune(int(hasil[j]) - 32)
							break
						}
						j--
					}
					// fmt.Println("indeks j ke-", j)
					// fmt.Println(uppercase)
				}
			}
		}
	}

	return string(hasil) // TODO: replace this

}

func check(a int) bool { // fungsi untuk mengecek uppercase
	upper := []int{}
	for i := 65; i > 0; i++ {
		upper = append(upper, i)
		if i == 90 {
			break
		}
	}
	result := false
	for i := 0; i < len(upper); i++ {
		if upper[i] == a {
			result = true
		}
	}
	return result
}

// gunakan untuk melakukan debug
func main() {
	// fmt.Println(ReverseWord("A bird fly to Germany and got a worm"))
	fmt.Println(ReverseWord("A Nama Kamu Siapa"))
	fmt.Println(ReverseWord("Nama Kamu Siapa"))
	fmt.Println(ReverseWord("SIAPA NAMA KAMU"))
}
