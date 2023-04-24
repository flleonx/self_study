function walk(ptr: BinaryNode<number> | null, path: number[]): void {
  if (!ptr) {
    return;
  }

  walk(ptr.left, path);
  path.push(ptr.value);
  walk(ptr.right, path);
}

export default function in_order_search(head: BinaryNode<number>): number[] {
  const path: number[] = [];
  walk(head, path);

  return path;
}
