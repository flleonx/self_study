import { StackCalculator } from '../stack-calculator.js';

/*
  Using an object literal and a factory function
*/
function createSafeCalculator(calculator) {
  return {
    // proxied method
    divide() {
      // additional validation logic
      const divisor = calculator.peekValue();

      if (divisor === 0) {
        throw new Error("Division by 0");
      }

      // if valid delegates to the subject
      return calculator.divide();
    },

    // delegated methods
    putValue(value) {
      return calculator.putValue(value);
    },
    getValue() {
      return calculator.getValue();
    },
    clear() {
      return calculator.clear();
    },
    multiply() {
      return calculator.multiply();
    }
  };
}

const calculator = new StackCalculator();
const safeCalculator = createSafeCalculator(calculator);

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

