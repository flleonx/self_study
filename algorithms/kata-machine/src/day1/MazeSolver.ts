const dir = [
  [1, 0],
  [0, 1],
  [-1, 0],
  [0, -1],
];

function walk(
  maze: string[],
  wall: string,
  curr: Point,
  end: Point,
  seen: Map<string, boolean>,
  path: Point[],
) {
  const pairs = `${curr.y}-${curr.x}`;

  // beyond maze
  if (
    curr.x < 0 ||
    curr.x >= maze[0].length ||
    curr.y < 0 ||
    curr.y >= maze.length
  ) {
    return false;
  }

  // on a wall
  if (maze[curr.y][curr.x] === wall) {
    return false;
  }

  // already seen
  if (seen.has(pairs)) {
    return false;
  }

  // the end
  if (curr.x === end.x && curr.y === end.y) {
    path.push(curr);
    return true;
  }

  seen.set(pairs, true);
  path.push(curr);
  // recursive
  for (let i = 0; i < dir.length; i++) {
    const [y, x] = dir[i];
    const res = walk(
      maze,
      wall,
      { x: curr.x + x, y: curr.y + y },
      end,
      seen,
      path,
    );

    if (res) {
      return true;
    }
  }

  path.pop();
  return false;
}

export default function solve(
  maze: string[],
  wall: string,
  start: Point,
  end: Point,
): Point[] {
  const seen = new Map<string, boolean>();
  const path: Point[] = [];

  walk(maze, wall, start, end, seen, path);

  return path;
}
