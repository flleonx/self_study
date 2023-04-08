import { readFileSync } from "fs";

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

const fromNullable = (x) => (x != null ? Right(x) : Left());

const tryCatch = (f) => {
  try {
    return Right(f());
  } catch (error) {
    return Left(error);
  }
};

// const getPort = () => {
//   try {
//     const str = readFileSync("config.json");
//     const config = JSON.parse(str);
//     return config.port;
//   } catch (error) {
//     return 5000;
//   }
// };

// const getPort = () =>
//   tryCatch(() => readFileSync("config.json"))
//     .map((contents) => JSON.parse(contents))
//     .map((config) => config.port)
//     .fold(
//       () => 8080,
//       (x) => x
//     );

const myReadFileSync = (path) => tryCatch(() => readFileSync(path));

const getPort = () =>
  myReadFileSync("config.json")
    .map((contents) => JSON.parse(contents))
    .map((config) => config.port)
    .fold(
      () => 8080,
      (x) => x
    );

const result = getPort();

console.log("result", result);
