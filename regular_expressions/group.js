const repeatStr = "regex regex";
// capture group
const repeatRegex = /(\w+)\s\1/;
console.log(repeatRegex.test(repeatStr));
console.log(repeatStr.match(repeatRegex));

// const repeatNum = "42 42 42";
// const reRegex = /^(\d+)\s\1\s\1$/;
// console.log(repeatNum.match(reRegex));
// console.log(reRegex.test(repeatNum));
