import { Task } from "../types.js";

const t1 = Task((rej, res) => res(2))
  .map((two) => two + 1)
  .map((three) => three * 2);

t1.fork(console.error, console.log);
