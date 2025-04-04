const Right = (x) => ({
  chain: (f) => f(x),
  map: (f) => Right(f(x)),
  fold: (f, g) => g(x),
  toString: `Right(${x})`,
});

const Left = (x) => ({
  chain: (f) => Left(x),
  map: (f) => Left(x),
  fold: (f, g) => f(x),
  toString: `Left(${x})`,
});

// const findColor = (name) =>
//   ({ red: "#ff4444", blue: "#3b5998", yellow: "#fff68f" }[name]);

const findColor = (name) => {
  const found = { red: "#ff4444", blue: "#3b5998", yellow: "#fff68f" }[name];
  return found ? Right(found) : Left("missing");
};

const res = findColor("red")
  .map((x) => x.toUpperCase())
  .map((x) => x.slice(1))
  .fold(
    () => "no color!",
    (color) => color
  );

const _res = findColor("redd")
  .map((x) => x.toUpperCase())
  .map((x) => x.slice(1))
  .fold(
    () => "no color!",
    (color) => color
  );

console.log("res", res);
console.log("_res", _res);
