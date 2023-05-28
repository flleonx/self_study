function binary_search(
  nums: number[],
  l: number,
  r: number,
  target: number
): number {
  const middle = l + Math.floor((r - l) / 2);

  if (nums[middle] === target) {
    return middle;
  }

  if (l >= r) {
    return -1;
  }

  if (nums[middle] < target) {
    return binary_search(nums, middle + 1, r, target);
  } else {
    return binary_search(nums, l, middle - 1, target);
  }
}

function search(nums: number[], target: number): number {
  return binary_search(nums, 0, nums.length - 1, target);
}

console.log("search", search([-1, 0, 3, 5, 9, 12], 9));
console.log("search", search([-1, 0, 3, 5, 9, 12], 2));
console.log("search", search([5], 5));
console.log("search", search([2, 5], 0));
