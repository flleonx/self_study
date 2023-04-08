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

const tryCatch = (f) => {
  try {
    return Right(f());
  } catch (error) {
    return Left(error);
  }
};

const myReadFileSync = (path) => tryCatch(() => readFileSync(path));
const parseJSON = contents => tryCatch(() => JSON.parse(contents));

const getPort = () =>
  myReadFileSync("config.json")
    .chain((contents) => parseJSON(contents))
    .map((config) => config.port)
    .fold(
      () => 8080,
      (x) => x
    );

const result = getPort();

console.log("result", result);
