const repeatStr = "javascript,go";

const repeatRegex = /(?=.*javascript)(?=.*go)/;
console.log(repeatStr.match(repeatRegex));
