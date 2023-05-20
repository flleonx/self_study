function maxArea(height: number[]): number {
  let i = 0;
  let j = height.length - 1;
  let maxArea = -1;

  while (i < j) {
    const area = Math.min(height[i], height[j]) * (j - i);
    maxArea = Math.max(maxArea, area);

    if (height[i] > height[j]) {
      j--;
    } else {
      i++;
    }
  }

  return maxArea;
}

// console.log("maxArea", maxArea([1, 8, 6, 2, 5, 4, 8, 3, 7]));
// console.log("maxArea", maxArea([1, 2, 1]));
// console.log("maxArea", maxArea([1, 8, 6, 2, 5, 4, 8, 25, 7]));
console.log("maxArea", maxArea([4, 6, 4, 4, 6, 2, 6, 7, 11, 2]));
