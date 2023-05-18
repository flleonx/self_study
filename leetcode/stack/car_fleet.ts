function getTime(pos1: number, pos2: number, v1: number, v2: number): number {
  return (pos2 - pos1) / (v1 - v2);
}

interface CarInfo {
  position: number;
  speed: number;
}

function carFleet(target: number, position: number[], speed: number[]): number {
  const sorted: CarInfo[] = [];

  for (const [idx, pos] of position.entries()) {
    const sp = speed[idx];
    sorted.push({ position: pos, speed: sp });
  }

  sorted.sort((a, b) => b.position - a.position);
  const stack: CarInfo[] = [];

  for (let i = 0; i < sorted.length; i++) {
    if (stack.length === 0) {
      stack.push(sorted[i]);
      continue;
    }

    const currCar = stack[stack.length - 1];
    const next = sorted[i];
    stack.push(next);
    const time = getTime(
      currCar.position,
      next.position,
      currCar.speed,
      next.speed
    );
    const meetingPoint = currCar.position + currCar.speed * time;

    if (time > 0 && meetingPoint <= target) {
      stack.pop();
    }
  }

  return stack.length;
}

console.log("carFleet", carFleet(12, [10, 8, 0, 5, 3], [2, 4, 1, 1, 3]));
// console.log("carFleet", carFleet(100, [0, 2, 4], [4, 2, 1]));
// console.log("carFleet", carFleet(10, [0, 4, 2], [2, 1, 3]));
