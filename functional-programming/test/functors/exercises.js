import QUnit from "qunit";

const Box = (x) => ({
  map: (f) => Box(f(x)),
  fold: (f) => f(x),
  toString: () => `Box(${x})`,
  chain: (f) => f(x),
});

// Exercise: Box
// Goal: Refactor each example using Box
// Keep these tests passing!
// Bonus points: no curly braces

// Ex1: Using Box, refactor moneyToFloat to be unnested.
// =========================
// const moneyToFloat = str =>
//   parseFloat(str.replace(/\$/, ''))

const moneyToFloat = (str) =>
  Box(str)
    .map((x) => x.replace(/\$/, ""))
    .fold((x) => parseFloat(x));

QUnit.test("Ex1: moneyToFloat", (assert) => {
  assert.equal(String(moneyToFloat("$5.00")), 5);
});

// Ex2: Using Box, refactor percentToFloat to remove assignment
// =========================
// const percentToFloat = str => {
//   const float = parseFloat(str.replace(/\%/, ''))
//   return float * 0.0100
// }

const percentToFloat = (str) =>
  Box(str)
    .map((s) => s.replace(/\%/, ""))
    .map((s) => parseFloat(s))
    .fold((f) => f * 0.01);

QUnit.test("Ex2: percentToFloat", (assert) => {
  assert.equal(String(percentToFloat("20%")), 0.2);
});

// Ex3: Using Box, refactor applyDiscount (hint: each variable introduces a new Box)
// =========================
// const applyDiscount = (price, discount) => {
//   const cents = moneyToFloat(price);
//   const savings = percentToFloat(discount);
//   return cents - cents * savings;
// };

// const applyDiscount = (price, discount) =>
//   Box(moneyToFloat(price)).fold((cents) =>
//     Box(percentToFloat(discount)).fold((savings) => cents - cents * savings)
//   );

const applyDiscount = (price, discount) =>
  Box(moneyToFloat(price)).chain((cents) =>
    Box(percentToFloat(discount)).map((savings) => cents - cents * savings)
  ).fold(x => x)

QUnit.test("Ex3: Apply discount", (assert) => {
  assert.equal(String(applyDiscount("$5.00", "20%")), 4);
});
