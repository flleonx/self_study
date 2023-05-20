// function twoSum(numbers: number[], target: number): number[] {
//   let i = 0;
//   while (i < numbers.length) {
//     let j = i + 1;

//     while (j < numbers.length) {
//       if (numbers[i] + numbers[j] === target) {
//         return [i + 1, j + 1];
//       }
//       j++;
//     }

//     i++;
//   }

//   return [];
// }

function twoSum(numbers: number[], target: number): number[] {
  let i = 0;
  let j = numbers.length - 1;
  while (i < j) {
    if (numbers[i] + numbers[j] < target) {
      i++;
      continue;
    }

    if (numbers[i] + numbers[j] > target) {
      j--;
      continue;
    }

    return [i + 1, j + 1];
  }

  return [];
}
