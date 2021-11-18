let process = 0;

function longestPalindrome(s) {
  let longestPalindrome = "";
  let longestLength = 0;
  for (let i = 0; i < s.length; i++) {
    for (let j = 1; j <= s.length; j++) {
      if (i < j) {
        let sliceOfString = s.slice(i, j);
        // check if the sliceOfString is a palindrome
        let reversed = "";
        let isPalindrome = true;
        for (let c = 0; c < sliceOfString.length; c++) {
          process++;
          if (
            sliceOfString[c] !== sliceOfString[sliceOfString.length - 1 - c]
          ) {
            isPalindrome = false;
            break;
          }
        }

        if (isPalindrome) {
          if (sliceOfString.length > longestLength) {
            longestLength = sliceOfString.length;
            longestPalindrome = sliceOfString;
          }
        }
      }
    }
  }
  return longestPalindrome;
}

const str = "babaqweqwedbabaweqwe123dbabadasdasdbabadbabadbabadbabaqweqwedbabaweqwe123dbabadasdasdbabadbabadbabad";
console.log(
  longestPalindrome(str),
  str.length,
  Math.pow(str.length, 2),
  process
);
