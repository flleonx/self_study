function productExceptSelf(nums: number[]): number[] {
  const output: number[] = [];

  for (let i = 0; i < nums.length; i++) {
    if (i === 0) {
      output[i] = 1;
      continue;
    }

    output[i] = output[i - 1] * nums[i - 1];
  }

  let post = 1;
  for (let i = nums.length - 1; i >= 0; i--) {
    output[i] = output[i] * post;
    post = post * nums[i];
  }

  return output;
}

// console.log("productExceptSelf", productExceptSelf([-1, 1, 0, -3, 3]));
console.log("productExceptSelf", productExceptSelf([1, -1]));
