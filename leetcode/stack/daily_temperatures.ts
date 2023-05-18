// Time: O(n^2)
// function dailyTemperatures(temperatures: number[]): number[] {
//   const output: number[] = [];
//   for (let i = 0; i < temperatures.length; i++) {
//     let j = i;
//     while (j < temperatures.length) {
//       if (temperatures[i] < temperatures[j]) {
//         const distance = j - i;
//         output[i]= distance;
//         break;
//       }
//       j++;
//     }

//     if (!output[i]) {
//       output[i] = 0;
//     }
//   }

//   return output;
// };

// Time: O(n)
// Space: O(n)
function dailyTemperatures(temperatures: number[]): number[] {
  // Monotonic decreasing order
  const stack: { val: number; idx: number }[] = [];
  const output: number[] = new Array(temperatures.length).fill(0);
  console.log('output', output);

  for (let i = 0; i < temperatures.length; i++) {
    while (stack.length && stack[stack.length - 1].val < temperatures[i]) {
      if (stack[stack.length - 1].val < temperatures[i]) {
        const { idx } = stack.pop() as { val: number; idx: number };
        output[idx] = i - idx;
      }
    }

    stack.push({ val: temperatures[i], idx: i });
  }

  return output;
}

console.log(
  "dailyTemperatures",
  dailyTemperatures([73, 74, 75, 71, 69, 72, 76, 73])
);
