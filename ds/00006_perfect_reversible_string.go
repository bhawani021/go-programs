// Check if reverse of all substring exists in given string
// An efficient solution for this problem is based on the fact that
// reverse of all substrings of ‘str’
// will exist in ‘str’ if and only if the entire string ‘str’ is palindrome.
package ds

func IsStringPerfectReversible(str string) bool {

	min, max := 0, len(str)-1

	for min < max {
		if str[min] != str[max] {
			return false
		}
		min++
		max--
	}

	return true
}
