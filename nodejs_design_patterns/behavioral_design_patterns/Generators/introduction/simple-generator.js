function* twoWayGenerator() {
  const what = yield null;
  yield "Hello " + what;
}

const twoWay = twoWayGenerator();
console.log(twoWay.next());
console.log(twoWay.next("world"));
