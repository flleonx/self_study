const hello = "  Hello, World!  ";
// const wsRegex = /\s(?!W)/g;
const wsRegex = /^\s+|\s+$/g;
const result = hello.replace(wsRegex, "");
console.log(result);
