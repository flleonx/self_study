function topKFrequent(nums: number[], k: number): number[] {
  const frequencies = new Map<number, number>();

  for (let i = 0; i < nums.length; i++) {
    const currNum = nums[i];
    const foundFreq = frequencies.get(currNum);

    if (foundFreq) {
      frequencies.set(currNum, foundFreq + 1);
      continue;
    }

    frequencies.set(currNum, 1);
  }

  const sortedArray: { value: number; freq: number }[] = [];
  for (const [currNum, currFreq] of frequencies) {
    sortedArray.push({ value: currNum, freq: currFreq });
  }

  sortedArray.sort((a, b) => b.freq - a.freq);

  return sortedArray.slice(0, k).map((el) => el.value);
}
