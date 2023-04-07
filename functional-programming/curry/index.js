import { curry } from "ramda";

// const add = (x, y) => x + y;

// const curry = f => x => y => f(x,y);

const modulo = (x, y) => y % x;

const cModulo = curry(modulo);

const isOdd = cModulo(2);

const filter = curry((f, xs) => xs.filter(f));

const getOdds = filter(isOdd);

const replace = curry((regex, replacement, str) =>
  str.replace(regex, replacement)
);

const replaceVowels = replace(/[AEIOU]/gi, "!");

const result = replaceVowels('Hey I have words');

console.log("result", result);
