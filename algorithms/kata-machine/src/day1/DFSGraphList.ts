function walk(
  graph: WeightedAdjacencyList,
  curr: number,
  needle: number,
  seen: boolean[],
  path: number[],
): boolean {
  if (seen[curr]) {
    return false;
  }

  // prev
  path.push(curr);

  if (curr === needle) {
    return true;
  }

  seen[curr] = true;
  // recurse
  const list = graph[curr];
  for (let i = 0; i < list.length; i++) {
    const edge = list[i];
    const found = walk(graph, edge.to, needle, seen, path);

    if (found) {
      return true;
    }
  }
  // post
  path.pop();

  return false;
}

export default function dfs(
  graph: WeightedAdjacencyList,
  source: number,
  needle: number,
): number[] | null {
  const seen: boolean[] = new Array(graph.length).fill(false);
  const path: number[] = [];
  const found = walk(graph, source, needle, seen, path);

  if (found) {
    return path;
  }

  return null;

  /*
    if (path.length === 0) {
      return null;
    }

    return path;
  */
}
