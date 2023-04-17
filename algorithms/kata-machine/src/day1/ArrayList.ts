export default class ArrayList<T> {
  public length: number;
  private list: T[];

  constructor(capacity: number) {
    this.length = 0;
    this.list = Array(capacity);
  }

  prepend(item: T): void {
    this.length++;

    for (let i = this.length; i >= 0; i--) {
      this.list[i] = this.list[i - 1];
    }

    this.list[0] = item;
  }

  insertAt(item: T, idx: number): void {
    for (let i = idx; i < this.length; i++) {
      this.list[i + 1] = this.list[i];
    }

    this.length++;

    this.list[idx] = item;
  }

  append(item: T): void {
    const idx = this.length;
    this.list[idx] = item;
    this.length++;
  }

  remove(item: T): T | undefined {
    let idx = undefined;
    for (let i = 0; i < this.length; i++) {
      if (this.list[i] === item) {
        idx = i;
      }
    }

    return idx !== undefined ? this.removeAt(idx) : undefined;
  }

  get(idx: number): T | undefined {
    return this.list[idx];
  }

  removeAt(idx: number): T | undefined {
    const item = this.list[idx];

    for (let i = idx; i < this.length; i++) {
      this.list[i] = this.list[i + 1];
    }

    this.length--;

    return item;
  }
}
