import { StackCalculator } from '../stack-calculator.js';

/*
  Object augmentation ( or monkey patching)
*/
function patchToSafeCalculator(calculator) {
  const divideOrig = calculator.divide;

  calculator.divide = () => {
    // additional validation logic
    const divisor = calculator.peekValue();

    if (divisor === 0) {
      throw new Error("Division by 0");
    }

    // if valid delegates to the subject
    return divideOrig.apply(calculator);
  }

  return calculator;
}

const calculator = new StackCalculator();
const safeCalculator = new patchToSafeCalculator(calculator);

calculator.putValue(2);
calculator.putValue(3);
console.log(calculator.multiply());

safeCalculator.putValue(2);
console.log(safeCalculator.multiply());

calculator.putValue(0);
console.log(calculator.divide());

safeCalculator.clear();
calculator.putValue(4);
calculator.putValue(0);
console.log(safeCalculator.divide());

