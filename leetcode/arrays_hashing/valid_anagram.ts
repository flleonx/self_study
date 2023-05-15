function getFrequency(str: string): Map<string, number> {
  const freqMap = new Map<string, number>();

  for (let char of str) {
    const currFreq = freqMap.get(char);
    if (currFreq) {
      freqMap.set(char, currFreq + 1)
    }
    freqMap.set(char, 1);
  }

  return freqMap;
}

function isAnagram(s: string, t: string): boolean {
  if (s.length !== t.length) {
    return false;
  }

  const freqS: Map<string, number> = getFrequency(s);
  const freqT: Map<string, number> = getFrequency(t);

  for (const [key, freq] of Object.entries(freqS)) {
    const freqCharT = freqT.get(key);

    if (freqCharT !== freq) {
      return false;
    }
  }

  return true;
}
