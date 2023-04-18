type Node<T> = {
  value: T;
  prev?: Node<T>;
  next?: Node<T>;
};

export default class DoublyLinkedList<T> {
  public length: number;
  private head?: Node<T>;
  private tail?: Node<T>;

  constructor() {
    this.length = 0;
    this.head = undefined;
    this.tail = undefined;
  }

  prepend(item: T): void {
    const node: Node<T> = {
      value: item,
    };

    this.length++;

    if (!this.head) {
      this.head = this.tail = node;
      return;
    }

    node.next = this.head;
    this.head.prev = node;
    this.head = node;
  }

  insertAt(item: T, idx: number): void {
    if (idx > this.length) {
      throw new Error("Out of bounds");
    } else if (idx === this.length) {
      this.append(item);
    } else if (idx === 0) {
      this.prepend(item);
    }

    const node: Node<T> = {
      value: item,
    };

    this.length++;

    let currentNode = this.head;

    for (let i = 0; i < idx && currentNode; i++) {
      currentNode = currentNode.next;
    }

    currentNode = currentNode as Node<T>;

    node.next = currentNode;
    node.prev = currentNode.prev;
    currentNode.prev = node;

    if (node.prev) {
      node.prev.next = node;
    }
  }

  append(item: T): void {
    const node: Node<T> = {
      value: item,
    };

    this.length++;

    if (!this.tail) {
      this.head = this.tail = node;
      return;
    }

    node.prev = this.tail;
    this.tail.next = node;
    this.tail = node;
  }

  remove(item: T): T | undefined {
    let currentNode = this.head;

    for (let i = 0; i < this.length && currentNode; i++) {
      if (currentNode.value === item) {
        break;
      }
      currentNode = currentNode.next;
    }

    if (!currentNode) {
      return undefined;
    }

    return this.removeNode(currentNode);
  }

  get(idx: number): T | undefined {
    return this.getAt(idx)?.value;
  }

  removeAt(idx: number): T | undefined {
    if (idx > this.length) {
      throw new Error("Out of bounds");
    }

    const currentNode = this.getAt(idx);

    if (!currentNode) {
      return undefined;
    }

    return this.removeNode(currentNode);
  }

  private getAt(idx: number): Node<T> | undefined {
    let currentNode = this.head;

    for (let i = 0; i < idx && currentNode; i++) {
      currentNode = currentNode.next;
    }

    return currentNode;
  }

  private removeNode(node: Node<T>): T | undefined {
    this.length--;

    if (this.length === 0) {
      this.head = this.tail = undefined;
      return node.value;
    }

    if (node.prev) {
      node.prev.next = node.next;
    }

    if (node.next) {
      node.next.prev = node.prev;
    }

    if (node === this.head) {
      this.head = node.next;
    }

    if (node === this.tail) {
      this.tail = node.prev;
    }

    node.next = node.prev = undefined;

    return node.value;
  }
}
