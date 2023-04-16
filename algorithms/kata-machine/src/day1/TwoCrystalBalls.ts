export default function two_crystal_balls(breaks: boolean[]): number {
  const offset = Math.floor(Math.sqrt(breaks.length));
  
  for(let i = offset; i <= breaks.length - 1; i += offset) {
    if (breaks[i]) {
      for (let j = i - offset; j <= i; j++) {
        if (breaks[j]) {
          return j;
        }
      }  
    }
  }

  return -1;
}
