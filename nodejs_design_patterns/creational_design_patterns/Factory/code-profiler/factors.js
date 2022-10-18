import { createProfiler } from './profiler.js';

function getAllFactors(intNumber) {
  const profiler = createProfiler( `Finding all factors of ${intNumber}`);

  profiler.start();

  const factors = [];

  for (let factor = 1; factor <= intNumber; factor++) {
    if ((intNumber % factor === 0)) {
      factors.push(factor);
    }
  }

  profiler.end();

  return factors;
}

const myNumber = process.argv[2];
const myFactors = getAllFactors(myNumber);

console.log(`Factors of ${myNumber} are: `, myFactors);
