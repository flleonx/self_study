function isPalindrome(s: string): boolean {
  const str = s.replace(/[^a-zA-Z0-9]/g, "").toLowerCase().trim();

  let j = str.length - 1;
  for (let i = 0; i < str.length; i++) {
    if (str[i] !== str[j]) {
      return false;
    }

    if (i === j) {
      break;
    }

    j--;
  }

  return true;
}

console.log('isPalindrome', isPalindrome("A man, a plan, a canal: Panama"));

