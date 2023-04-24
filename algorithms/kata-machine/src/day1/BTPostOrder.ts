function walk(ptr: BinaryNode<number> | null, path: number[]): void {
  if (!ptr) {
    return;
  }

  walk(ptr.left, path);
  walk(ptr.right, path);
  path.push(ptr.value);
}

export default function post_order_search(head: BinaryNode<number>): number[] {
  const path: number[] = [];
  walk(head, path);

  return path;
}
