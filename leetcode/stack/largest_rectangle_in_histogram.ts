interface ElementInfo {
  height: number;
  index: number;
}

function largestRectangleArea(heights: number[]): number {
  const stack: ElementInfo[] = [];
  let maxArea = -1;

  for (const [idx, height] of heights.entries()) {
    let newIndex = idx;
    while (stack.length > 0 && stack[stack.length - 1].height > height) {
      const top = stack.pop() as ElementInfo;
      const area = top.height * (idx - top.index);
      maxArea = Math.max(maxArea, area);
      newIndex = top.index;
    }

    stack.push({ height, index: newIndex });
  }

  if (stack.length > 0) {
    const limit = heights.length;
    for (const el of stack) {
      const area = el.height * (limit - el.index);
      maxArea = Math.max(maxArea, area);
    }
  }

  return maxArea;
}

console.log("area", largestRectangleArea([2, 1, 5, 6, 2, 3]));
console.log("area", largestRectangleArea([2, 4]));
console.log("area", largestRectangleArea([3, 6, 5, 7, 4, 8, 1, 0]));
