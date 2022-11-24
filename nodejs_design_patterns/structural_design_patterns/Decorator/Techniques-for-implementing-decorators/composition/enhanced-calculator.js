import { StackCalculator } from "../stack-calculator.js";

class EnhancedCalculator {
  constructor(calculator) {
    this.calculator = calculator;
  }

  // new method
  add() {
    const addend2 = this.getValue();
    const addend1 = this.getValue();
    const result = addend1 + addend2;
    this.putValue(result);
    return result;
  }

  // modified method
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
    return this.calculator.getValue();
  }

  peekValue() {
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
const enhancedCalculator = new EnhancedCalculator(calculator);

enhancedCalculator.putValue(4);
enhancedCalculator.putValue(3);
console.log(enhancedCalculator.add());
enhancedCalculator.putValue(2);
console.log(enhancedCalculator.multiply());
