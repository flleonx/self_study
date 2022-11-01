import { StackCalculator } from '../stack-calculator.js';

class SafeCalculator {
  constructor(calculator) {
    this.calculator = calculator;
  }

  // proxied method
  divide() {
    // additional validation logic
    const divisor = this.calculator.peekValue();

    if (divisor === 0) {
      throw new Error("Division by 0");
    }

    // if valid delegates to the subject
    return this.calculator.divide();
  }

  // delegated methods
  putValue(value) {
    return this.calculator.putValue(value);
  }

  getValue() {
    return this.calculator.peekValue();
  }

  clear() {
    return this.calculator.clear();
  }

  multiply() {
    return this.calculator.multiply();
  }
}

const calculator = new StackCalculator();
const safeCalculator = new SafeCalculator(calculator);

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

