// Unique Email Adresses

func numUniqueEmails(emails []string) int {
	uniqueEmails := map[string]bool{}
	for _, x := range emails {
		emailSplitted := strings.Split(x, "@")
		local, domain := emailSplitted[0], emailSplitted[1]
		formattedLocal := strings.ReplaceAll(strings.Split(local, "+")[0], ".", "")
		email := formattedLocal + "@" + domain
		uniqueEmails[email] = true
	}
	return len(uniqueEmails)
}
