function binary_search(
  nums: number[],
  l: number,
  r: number,
  target: number
): number {
  const middle = l + Math.floor((r - l) / 2);

  if (target === nums[middle]) {
    return middle;
  }

  if (l > r) {
    return - 1;
  }

  if (nums[l] > nums[r]) {
    const path1 = binary_search(nums, l, middle - 1, target);
    const path2 = binary_search(nums, middle + 1, r, target);
    if (path1 !== -1) return path1;
    if (path2 !== -1) return path2;
    return -1;
  } else if (target < nums[middle]) {
    return binary_search(nums, l, middle - 1, target);
  } else if (target > nums[middle]) {
    return binary_search(nums, middle + 1, r, target);
  }

  return -1;
}

function search(nums: number[], target: number): number {
  let l = 0;
  let r = nums.length - 1;
  return binary_search(nums, l, r, target);
}

// console.log("search", search([4,5,6,7,0,1,2], 0));
// console.log("search", search([4,5,6,7,0,1,2], 3));
// console.log("search", search([1], 0));
console.log("search", search([1, 3], 2));
// console.log("search", search([3, 5, 1], 3));
