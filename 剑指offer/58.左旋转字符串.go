package main

func reverseLeftWords(s string, n int) string {
	str1 := []byte(s)
	len := len(str1)
	var tmp []byte
	for i := 0; i < n; i++ {
		tmp = append(tmp, str1[i])
	}
	str1 = append(str1[n:len], tmp...)
	return string(str1)
}

// func main() {

// }
