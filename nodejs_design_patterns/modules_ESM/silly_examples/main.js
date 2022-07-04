// main.js
import * as a from './a.js'
import * as b from './b.js'
console.log('a ->', a)
console.log('b ->', b)
const obj = {
  a: 1,
  b: 2
};
obj.myself = obj;
console.log(obj);
