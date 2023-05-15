// function longestConsecutive(nums: number[]): number {
//   const set = new Set<number>(nums);
//   const sequences = new Map<number, number[]>();

//   for (const value of set) {
//     const prev = value - 1;
//     let next = value + 1;

//     if (set.has(prev)) {
//       continue;
//     }

//     const sequence = [value];
//     while (set.has(next)) {
//       sequence.push(next);
//       next++;
//     }

//     sequences.set(value, sequence);
//   }

//   let maxSequence = 0;

//   for (const sequence of sequences.values()) {
//     if (sequence.length > maxSequence) {
//       maxSequence = sequence.length;
//     }
//   }

//   return maxSequence;
// }

function longestConsecutive(nums: number[]): number {
  const set = new Set<number>(nums);
  let longestSequence = 0;

  for (const value of set) {
    const prev = value - 1;
    let next = value + 1;

    if (set.has(prev)) {
      continue;
    }

    let currSequenceLength = 1;

    while (set.has(next)) {
      currSequenceLength++;
      next++;
    }

    longestSequence = Math.max(longestSequence, currSequenceLength);
  }

  return longestSequence;
}

console.log("longestConsecutive", longestConsecutive([100, 4, 200, 1, 3, 2]));
console.log(
  "longestConsecutive",
  longestConsecutive([0, 3, 7, 2, 5, 8, 4, 6, 0, 1])
);
