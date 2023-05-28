function minEatingSpeed(piles: number[], h: number): number {
  const max = Math.max(...piles);

  let k = max;
  let l = 1;
  let r = max;
  while (l <= r) {
    let middle = l + Math.floor((r - l) / 2);
    let hours = 0;
    for (const pile of piles) {
      hours += Math.ceil(pile / middle);
    }

    if (hours <= h) {
      k = Math.min(middle, k);
      r = middle - 1;
    } else {
      l = middle + 1;
    }
  }

  return k;
}

console.log("minEatingSpeed", minEatingSpeed([3, 6, 7, 11], 8));
console.log("minEatingSpeed", minEatingSpeed([30, 11, 23, 4, 20], 5));
console.log("minEatingSpeed", minEatingSpeed([30, 11, 23, 4, 20], 6));
console.log("minEatingSpeed", minEatingSpeed([312884470], 312884469));
