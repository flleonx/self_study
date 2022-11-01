const evenNumbers = new Proxy([], {
  get: (_target, index) => index * 2,
  has: (_target, number) => number % 2 === 0
});

console.log(2 in evenNumbers);
console.log(5 in evenNumbers);
console.log(evenNumbers[7]);
