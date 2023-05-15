// type freqMap = Map<string, number>;

// function getFrequency(str: string): Map<string, number> {
//   const freqMap = new Map<string, number>();
//   for (const char of str) {
//     const savedFreq = freqMap.get(char);
//     if (savedFreq) {
//       freqMap.set(char, savedFreq + 1);
//       continue;
//     }

//     freqMap.set(char, 1);
//   }

//   if (str === "") {
//     freqMap.set(str, 1);
//   }

//   return freqMap;
// }

// function compareFreq(freq1: freqMap, freq2: freqMap, str1: string, str2: string): boolean {
//   if (str1.length !== str2.length) {
//     return false;
//   }

//   for (const [char1, charFreq1] of freq1) {
//     const charFreq2 = freq2.get(char1);

//     if (charFreq1 !== charFreq2) {
//       return false;
//     }
//   }

//   return true;
// }

// function groupAnagrams(strs: string[]): string[][] {
//   const strToFreqMap = new Map<
//     string,
//     { counter: number; map: Map<string, number> }
//   >();
//   const groupedAnagrams: string[][] = [];

//   for (let i = 0; i < strs.length; i++) {
//     const str = strs[i];

//     const currFreqState = strToFreqMap.get(str);
//     if (currFreqState) {
//       currFreqState.counter++;
//       continue;
//     }

//     strToFreqMap.set(str, {
//       map: getFrequency(str),
//       counter: 1,
//     });
//   }

//   for (let i = 0; i < strs.length; i++) {
//     const str = strs[i];

//     const currFreq = strToFreqMap.get(str);

//     if (!currFreq) {
//       continue;
//     }

//     const currStrAnagrams = [str];

//     for (let j = 0; j < strs.length; j++) {
//       if (i === j) {
//         continue;
//       }

//       const otherFreq = strToFreqMap.get(strs[j]);

//       if (!otherFreq) {
//         continue;
//       }

//       const areEqual = compareFreq(currFreq.map, otherFreq.map, str, strs[j]);

//       if (areEqual) {
//         otherFreq.counter--;
//         currStrAnagrams.push(strs[j]);

//         if (otherFreq.counter === 0) {
//           strToFreqMap.delete(strs[j]);
//         }
//       }
//     }

//     groupedAnagrams.push(currStrAnagrams);
//     strToFreqMap.delete(str);
//   }

//   return groupedAnagrams;
// }

function groupAnagrams(strs: string[]): string[][] {
  const groups = new Map<string, string[]>();

  for (const str of strs) {
    const sortedString = [...str].sort().join('');
    const group = groups.get(sortedString) || [];
    group.push(str);
    groups.set(sortedString, group);
  }

  return [...groups.values()];
}

console.log('groupAnagrams', groupAnagrams(['', '', '']));
