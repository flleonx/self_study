// function isValid(s: string): boolean {
//   const stack: string[] = [];

//   for (const char of s) {
//     if (char === "(" || char === "{" || char === "[") {
//       stack.push(char);
//       continue;
//     }

//     if (char === ")") {
//       const openToken = stack.pop();
//       if (openToken === "(") {
//         continue;
//       }

//       return false;
//     }

//     if (char === "}") {
//       const openToken = stack.pop();
//       if (openToken === "{") {
//         continue;
//       }

//       return false;
//     }

//     if (char === "]") {
//       const openToken = stack.pop();
//       if (openToken === "[") {
//         continue;
//       }

//       return false;
//     }
//   }

//   if (stack.length === 0) {
//     return true;
//   }

//   return false;
// }

function isValid(s: string): boolean {
  const stack: string[] = [];
  const bracketsMap =  new Map<string, string>([[")", "("], ["}", "{"], ["]", "["]]);

  for (const char of s) {
    if (char === "(" || char === "{" || char === "[") {
      stack.push(char);
      continue;
    }

    const lastBracket = stack.pop();
    const matchBracket = bracketsMap.get(char);

    if (lastBracket !== matchBracket) {
      return false;
    }
  }

  if (stack.length === 0) {
    return true;
  }

  return false;
}
