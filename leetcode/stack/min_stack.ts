class MinStack {
  private length: number;
  private minStack: number[];
  private stack: number[];

  constructor() {
    this.stack = [];
    this.length = 0;
    this.minStack = [];
  }

  push(val: number): void {
    this.stack[this.length] = val;
    this.length++;

    if (this.minStack.length === 0) {
      this.minStack = [val];
      return;
    }

    if (this.minStack[this.minStack.length - 1] >= val) {
      this.minStack.push(val);
    }
  }

  pop(): void {
    this.length--;
    const deletedElement = this.stack[this.length];

    if (
      this.minStack &&
      deletedElement === this.minStack[this.minStack.length - 1]
    ) {
      this.minStack.pop();
    }
  }

  top(): number {
    return this.stack[this.length - 1];
  }

  getMin(): number {
    return this.minStack[this.minStack.length - 1];
  }
}
