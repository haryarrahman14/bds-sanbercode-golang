package helper

func TambahData(name string, phones *[]string) {
	*phones = append(*phones, name)
}