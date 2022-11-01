import { StackCalculator } from '../stack-calculator.js';

/*
  The Proxy object inherits the prototype of the subject, therefore
  running `safeCalculator instanceof StackCalculator` will return true.
*/

const safeCalculatorHandler = {
  get: (target, property) => {
    console.log('PROPERTY', property);
    if (property === 'divide') {
      // proxied method
      return function() {
        // aditional validation logic
        const divisor = target.peekValue();

        if (divisor === 0) {
          throw Error('Division by 0');
        }

        // if valid delegates to the subject
        return target.divide();
      }
    }

    // delegated methods and properties
    return target[property];
  }
}

const calculator = new StackCalculator();

const safeCalculator = new Proxy(
  calculator,
  safeCalculatorHandler
);

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

