package main

func uniq(s []string) []string {
	i := 0
	for _, v := range s[1:] {
		if s[i] != v {
			i++
			s[i] = v
		}
	}
	return s[:i+1]
}
