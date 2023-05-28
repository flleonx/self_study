function findMin(nums: number[]): number {
  const min = nums[0];
  const max = nums[nums.length - 1];

  let l = 0;
  let r = nums.length - 1;
  let currMin = Infinity;
  while (l <= r) {
    const middle = l + Math.floor((r - l) / 2);
    if (nums[middle] <= max) {
      currMin = Math.min(currMin, nums[middle]);
      r = middle - 1;
    } else {
      l = middle + 1;
    }
  }

  return currMin;
}

console.log("findMin", findMin([3, 4, 5, 1, 2]));
console.log("findMin", findMin([4, 5, 6, 7, 0, 1, 2]));
console.log("findMin", findMin([11, 13, 15, 17]));
