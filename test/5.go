package main

func main() {
}

func reverseLeftWords(s string, n int) string {
	s1 := s[0:n]
	s2 := s[n:]
	return s2 + s1
}
