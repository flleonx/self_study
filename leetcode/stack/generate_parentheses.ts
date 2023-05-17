function generateParenthesis(n: number): string[] {
  const stack: string[] = [];
  const res: string[] = [];

  function backtrack(openN: number, closedN: number) {
    // prev
    if (openN === n && closedN === n) {
      res.push(stack.join(""));
      return;
    }

    // recurse
    if (openN < n) {
      stack.push("(");
      backtrack(openN + 1, closedN);
      stack.pop();
    }

    // recurse (post)
    if (closedN < openN) {
      stack.push(")");
      backtrack(openN, closedN + 1);
      stack.pop();
    }
  }

  backtrack(0, 0);
  return res;
}

console.log("generateParenthesis", generateParenthesis(3));
