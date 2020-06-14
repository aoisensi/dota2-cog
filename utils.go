package main

func containsString(s []string, e string) bool {
	for _, s := range s {
		if s == e {
			return true
		}
	}
	return false
}
