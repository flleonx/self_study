function threeSum(nums: number[]): number[][] {
  nums.sort((a, b) => a - b);
  const output: number[][] = [];

  for (let i = 0; i < nums.length - 2; i++) {
    if (i > 0 && nums[i] === nums[i - 1]) {
      continue;
    }

    let l = i + 1;
    let r = nums.length - 1;

    while (l < r) {
      const threeSum = nums[i] + nums[l] + nums[r];
      if (threeSum > 0) {
        r--;
      } else if (threeSum < 0) {
        l++;
      } else {
        output.push([nums[i], nums[l], nums[r]]);
        l++;
        while (nums[l] === nums[l - 1] && l < r) {
          l++;
        }
      }
    }
  }

  return output;
}

// console.log("threeSum", threeSum([0, 0, 0]));
console.log("threeSum", threeSum([-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4]));
