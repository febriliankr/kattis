package main

import "fmt"

func longestPalindrome(s string) string {
	var longestPalindrome string
	var longestLength int = 0
	for i := 0; i < len(s); i++ {
		for j := 1; j <= len(s); j++ {
			if i < j {
				sliceOfString := s[i:j]
				// check if the sliceOfString is a palindrome
				var reversed string
				for _, char := range sliceOfString {
					reversed = string(char) + reversed
				}

				isPalindrome := sliceOfString == reversed

				if isPalindrome {
					if len(sliceOfString) > longestLength {
						longestLength = len(sliceOfString)
						longestPalindrome = sliceOfString
					}
				}
			}

		}
	}
	return longestPalindrome
}

func main() {
	longestPalindrome := longestPalindrome("civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth")
	fmt.Println(longestPalindrome)
}
