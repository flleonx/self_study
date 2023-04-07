import {
  curry,
  compose,
  filter,
  map,
  toLower,
  reverse,
  trim,
  split,
  join
} from "ramda";

const concat = curry((y, x) => x + y);

const toUpper = (str) => str.toUpperCase();

const exclaim = (str) => str + "!";

const first = (xs) => xs[0];

const shout = compose(concat("!"), toUpper);

console.log(shout("tears"));

const doStuff = compose(
  join(""),
  filter((x) => x.length > 3),
  reverse,
  map(trim),
  split(" "),
  toLower
);

console.log("doStuff", doStuff("Chain Dot Com "));
