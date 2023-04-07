import { curry, compose } from 'ramda';

const add = (x, y) => x + y;

const toUpper = str => str.toUpperCase();

const exclaim = str => str + '!';

const first = xs => xs[0];

// pipe = (f, g) => x => g(f(x))
const shout = compose(first, exclaim, toUpper);

console.log(shout('tears'));
