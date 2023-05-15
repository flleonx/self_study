// function twoSum(nums: number[], target: number): number[] {
//   for (let i = 0; i < nums.length; i++) {
//     for (let j = 0; j < nums.length; j++) {
//       if (nums[i] + nums[j] === target && i !== j) {
//         return [i, j];
//       }
//     }
//   }

//   return [];
// }

function twoSum(nums: number[], target: number): number[] {
  const visitedNumbers = new Map<number, number>();
  for (let i = 0; i < nums.length; i++) {
    const remaining = target - nums[i];
    const foundIdx = visitedNumbers.get(remaining);
    if (!foundIdx && foundIdx !== 0) {
      visitedNumbers.set(nums[i], i);
      continue;
    }

    if (foundIdx >= 0) {
      return [foundIdx, i];
    }
  }

  return [];
}
