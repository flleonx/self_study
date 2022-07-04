const quit = "qu";
const noquit = "qt";

const quRegex = /q(?=u)/;
const nquRegex = /q(?!u)/;

console.log(quit.match(quRegex));
console.log(noquit.match(nquRegex));

const sampleWord = "astronaut22";
// 5 or more characters and 2 or more digits
const pwRegex = /(?=\w{5})(?=\D*\d{2})/;
const result = pwRegex.test(sampleWord);
console.log(result);
