package helpers

func hash_password(password string) (string) {
	sum := sha256.Sum256([]byte(password))
	out := fmt.Sprintf("%x", sum)
	return out
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}