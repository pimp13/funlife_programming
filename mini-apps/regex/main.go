package main

import "fmt"

func main() {
	// mystr := "#f1f1f1"
	// matched, err := regexp.MatchString(`/^#(?:[a-f\d]{6})$/i`, mystr)
	// regex , err := regexp.Compile(`/^#(?:[a-f\d]{6})$/i`)
	// mactched2, err := regexp.Match(`/[a-z0-9]/`, []byte(mystr))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(matched)
	// fmt.Println(regex)
	// fmt.Println(mactched2)

	fmt.Println(xor_encrypt("hello"))

}

func xor_encrypt(in string) string {
	key := "hello"
	var output byte

	for i := 0; i < len(in); i++ {
		output += in[i] ^ key[i%len(key)]
	}

	fmt.Println(output)
	return string(output)
}
