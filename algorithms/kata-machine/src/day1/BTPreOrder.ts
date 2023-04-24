function walk(ptr: BinaryNode<number> | null, path: number[]): void {
  if (!ptr) {
    return;
  }

  path.push(ptr.value);

  walk(ptr.left, path);
  walk(ptr.right, path);
}

export default function pre_order_search(head: BinaryNode<number>): number[] {
  const path: number[] = [];
  walk(head, path);

  return path;
}
