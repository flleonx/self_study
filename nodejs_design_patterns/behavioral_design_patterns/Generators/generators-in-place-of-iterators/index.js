import { Matrix } from "./matrix.js";

const matrix2x2 = new Matrix([
  ["11", "12"],
  ["21", "22"],
]);

const iterator = matrix2x2[Symbol.iterator]();
let iteratorResult = iterator.next();

while (!iteratorResult.done) {
  console.log(iteratorResult.value);
  iteratorResult = iterator.next();
}

for (const element of matrix2x2) {
  console.log(element);
}

console.log(matrix2x2);
const flattenedMatrix = [...matrix2x2];
console.log(flattenedMatrix);

const [oneOne, oneTwo, twoOne, twoTwo] = matrix2x2;
console.log(oneOne, oneTwo, twoOne, twoTwo);
