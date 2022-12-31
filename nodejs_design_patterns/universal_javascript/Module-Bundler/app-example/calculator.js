import { parser } from "./parser.js";
import { resolver } from "./resolver.js";

export function calculator(expr) {
  return resolver(parser(expr));
}
