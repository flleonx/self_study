export default class MinHeap {
  public length: number;
  private heap: number[];

  constructor() {
    this.length = 0;
    this.heap = [];
  }

  insert(value: number): void {
    this.heap[this.length] = value;
    this.heapifyUp(this.length);
    this.length++;
  }

  delete(): number {
    if (this.length === 0) {
      return -1;
    }

    const out = this.heap[0];
    this.length--;

    if (this.length === 0) {
      this.heap = [];
      return out;
    }

    this.heap[0] = this.heap[this.length];
    this.heapifyDownIt(0);
    return out;
  }

  private heapifyUp(idx: number): void {
    if (idx === 0) {
      return;
    }

    const parentIdx = this.parent(idx);
    const parentValue = this.heap[parentIdx];
    const childValue = this.heap[idx];

    if (parentValue > childValue) {
      this.heap[parentIdx] = childValue;
      this.heap[idx] = parentValue;
      this.heapifyUp(parentIdx);
    }
  }

  private heapifyDown(idx: number): void {
    const lIdx = this.leftChild(idx);
    const rIdx = this.rightChild(idx);

    if (idx >= this.length || lIdx >= this.length) {
      return;
    }

    const leftValue = this.heap[lIdx];
    const rightValue = this.heap[rIdx];
    const currValue = this.heap[idx];

    if (leftValue > rightValue && currValue > rightValue) {
      this.heap[rIdx] = currValue;
      this.heap[idx] = rightValue;
      this.heapifyDown(rIdx);
    } else if (rightValue > leftValue && currValue > leftValue) {
      this.heap[lIdx] = currValue;
      this.heap[idx] = leftValue;
      this.heapifyDown(lIdx);
    }
  }

  private getSwap(currValue: number, leftValue: number, rightValue: number): { right: boolean; left: boolean} {
    return {
      right: currValue > rightValue && leftValue > rightValue,
      left: currValue > leftValue && rightValue > leftValue,
    }
  }
  
  private heapifyDownIt(idx: number): void {
    let lIdx = this.leftChild(idx);
    let rIdx = this.rightChild(idx);
    let currIndex = idx;

    let leftValue = this.heap[lIdx];
    let rightValue = this.heap[rIdx];
    const currValue = this.heap[idx];

    let swap = this.getSwap(currValue, leftValue, rightValue);

    while (swap.left || swap.right) {
      if (lIdx >= this.length || currIndex >= this.length) {
        return;
      }
      
      if (swap.right) {
        this.heap[rIdx] = currValue;
        this.heap[currIndex] = rightValue;
        currIndex = rIdx;
      } 

      if (swap.left) {
        this.heap[lIdx] = currValue;
        this.heap[currIndex] = leftValue;
        currIndex = lIdx;
      }

      lIdx = this.leftChild(currIndex);
      rIdx = this.rightChild(currIndex);
      leftValue = this.heap[lIdx];
      rightValue = this.heap[rIdx];
      swap = this.getSwap(currValue, leftValue, rightValue);
    }
  }

  private parent(idx: number): number {
    return Math.floor((idx - 1) / 2);
  }

  private leftChild(idx: number): number {
    return idx * 2 + 1;
  }

  private rightChild(idx: number): number {
    return idx * 2 + 2;
  }
}
