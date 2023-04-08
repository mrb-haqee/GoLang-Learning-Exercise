package main

import "fmt"

func EmailInfo(email string) string {
	Domain, TLD := "", ""
	n := len(email)
	for i := 0; i < n; i++ {
		if string(email[i]) == "@" {
			i++
			for {
				Domain += string(email[i])
				i++
				if string(email[i]) == "." {
					break
				}
			}
			i++
			for {
				TLD += string(email[i])
				i++
				if i >= n {
					break
				}
			}
		}
	}
	return fmt.Sprintf("Domain: %s dan TLD: %s", Domain, TLD) // TODO: replace this

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(EmailInfo("adadmin@yahoo.co.id"))

}
