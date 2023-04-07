import { curry } from 'ramda';

const add = (x, y) => x + y;

const toUpper = str => str.toUpperCase();

const exclaim = str => str + '!';

const first = xs => xs[0];

// f . g = x => f(g(x))
// Composition is associative
const compose = (f, g) => x => f(g(x));

const shout = compose(exclaim, toUpper);

const firstL = compose(first, shout);

console.log(shout('tears'));
console.log(firstL('tears'));
