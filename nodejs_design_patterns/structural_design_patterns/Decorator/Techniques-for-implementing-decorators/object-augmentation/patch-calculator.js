/*
  Object decoration can also be achieved by simply attaching new methods
  directly to the decorated object (monkey patching).
*/

import { StackCalculator } from "../stack-calculator.js";

function patchCalculator(calculator) {
  // new method
  calculator.add = function() {
    const addend2 = calculator.getValue();
    const addend1 = calculator.getValue();
    const result = addend1 + addend2;
    calculator.putValue(result);
    return result;
  }

  // modified method
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

/*
  patchCalculator() is mutating the original calculator object and then
  returning it.
*/
const calculator = new StackCalculator();
const enhancedCalculator = patchCalculator(calculator);

enhancedCalculator.putValue(4);
enhancedCalculator.putValue(3);
console.log(enhancedCalculator.add());
enhancedCalculator.putValue(2);
console.log(enhancedCalculator.multiply());
