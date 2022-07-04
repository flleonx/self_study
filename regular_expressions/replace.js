const wrongText = "The sky is silver";
const silverRegex = /silver/;
console.log(wrongText.replace(silverRegex, "blue"));

const codeCampStr = "Code Camp";
// console.log(codeCampStr.match(/(\w+)\s(\w+)/))
// $ means capture groups
console.log(codeCampStr.replace(/(\w+)\s(\w+)/, '$2 $1'));

const huhText = "This sandwich is good.";
const fixRegex = /good/;
const replaceText = "okey-dokey";
const result = huhText.replace(fixRegex, replaceText);
console.log(result);
