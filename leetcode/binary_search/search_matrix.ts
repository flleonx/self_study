function getRow(matrix: number[][], target: number): number[] {
  let l = 0;
  let r = matrix.length - 1;
  while (l <= r) {
    const middle = l + Math.floor((r - l) / 2);
    const min = matrix[middle][0];
    const max = matrix[middle][matrix[0].length - 1];
    if (target >= min && target <= max) {
      return matrix[middle];
    } else if (target > max) {
      l = middle + 1;
    } else {
      r = middle - 1;
    }
  }

  return [];
}

function searchMatrix(matrix: number[][], target: number): boolean {
  const row = getRow(matrix, target);

  if (row.length === 0) {
    return false;
  }

  let l = 0;
  let r = row.length - 1;
  while (l <= r) {
    const middle = l + Math.floor((r - l) / 2);

    if (row[middle] === target) {
      return true;
    } else if (target > row[middle]) {
      l = middle + 1;
    } else {
      r = middle - 1;
    }
  }

  return false;
}

console.log("searchMatrix", searchMatrix([[1,3,5,7],[10,11,16,20],[23,30,34,60]], 3));
console.log("searchMatrix", searchMatrix([[1,3,5,7],[10,11,16,20],[23,30,34,60]], 13));

