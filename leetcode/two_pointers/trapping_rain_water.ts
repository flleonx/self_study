// interface BlockInfo {
//   index: number;
//   height: number;
// }

// function trap(height: number[]): number {
//   const stack = [];
//   let trapCounter = 0;

//   for (const [idx, currH] of height.entries()) {
//     let trap = 0;
//     if (stack.length > 0 && stack[0].height < currH) {
//       const partialMax = stack[0];
//       let counter = 0;
//       while (stack.length > 0) {
//         const item = stack.pop() as BlockInfo;
//         if (item.height < partialMax.height) {
//           counter += partialMax.height - item.height;
//         }
//       }
//       trap = counter;
//     }

//     trapCounter += trap;
//     stack.push({ index: idx, height: currH });
//   }

//   stack.reverse();

//   let i = 0;
//   let j = i + 1;
//   while (i < stack.length - 1 && j < stack.length - 1) {
//     if ((i === 0 && stack[i] < stack[j + 1]) || stack.length < 2) {
//       i++;
//     }

//     if (j === stack.length - 1) break;

//     if (stack[i].height >= stack[j].height) {
//       trapCounter += stack[i].height - stack[j].height;
//       j++;
//       continue;
//     }

//     i = j;
//     j++;
//   }

//   return trapCounter;
// }

// Two pointers time: O(n) space: O(1)
function trap(height: number[]): number {
  let l = 0;
  let r = height.length - 1;
  let lMax = height[l];
  let rMax = height[r];
  let trapCounter = 0;

  while (l < r) {
    if (lMax < rMax) {
      l++;
      lMax = Math.max(height[l], lMax);
      trapCounter += lMax - height[l];
    } else {
      r--;
      rMax = Math.max(height[r], rMax);
      trapCounter += rMax - height[r];
    }
  }

  return trapCounter;
}

console.log("trap", trap([1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]));
console.log("trap", trap([4, 2, 0, 3, 2, 5]));
console.log("trap", trap([5, 5, 1, 7, 1, 1, 5, 2, 7, 6]));
