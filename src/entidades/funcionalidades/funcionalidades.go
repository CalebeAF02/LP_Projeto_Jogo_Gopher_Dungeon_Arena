package funcionalidades

func Simetria(s1 string, s2 string, c1 string, c2 string) bool {
	return (s1 == c1 && s2 == c2) || (s1 == c2 && s2 == c1)
}
