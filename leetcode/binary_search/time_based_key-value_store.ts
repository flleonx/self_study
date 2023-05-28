interface Payload {
  [key: string]: {
    values: string[];
    timestamps: number[];
  };
}

class TimeMap {
  private payload: Payload;

  constructor() {
    this.payload = {};
  }

  set(key: string, value: string, timestamp: number): void {
    if (this.payload[key]) {
      const { values, timestamps } = this.payload[key];
      values.push(value);
      timestamps.push(timestamp);
      return;
    }

    this.payload[key] = {
      values: [value],
      timestamps: [timestamp],
    };
  }

  get(key: string, timestamp: number): string {
    if (!this.payload[key]) {
      return "";
    }

    const { values, timestamps } = this.payload[key];

    if (timestamp < timestamps[0]) {
      return "";
    }

    const index = this.getIndex(timestamps, timestamp);

    return values[index];
  }

  private getIndex(timestamps: number[], target: number): number {
    let l = 0;
    let r = timestamps.length - 1;

    while (l <= r) {
      const middle = l + Math.floor((r - l) / 2);

      if (timestamps[middle] === target) {
        return middle;
      } else if (target > timestamps[middle]) {
        l = middle + 1;
      } else {
        r = middle - 1;
      }
    }

    return l < timestamps.length - 1 ? l : r;
  }
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * var obj = new TimeMap()
 * obj.set(key,value,timestamp)
 * var param_2 = obj.get(key,timestamp)
 */

// const timeMap = new TimeMap();
// timeMap.set("foo", "bar", 1)
// console.log(timeMap.get("foo", 1));
// console.log(timeMap.get("foo", 3));
// timeMap.set("foo", "bar2", 4);
// console.log(timeMap.get("foo", 4));
// console.log(timeMap.get("foo", 5));

const timeMap = new TimeMap();
timeMap.set("love", "high", 10)
timeMap.set("love", "low", 20)
console.log(timeMap.get("love", 5));
console.log(timeMap.get("love", 10));
console.log(timeMap.get("love", 15));
console.log(timeMap.get("love", 20));
console.log(timeMap.get("love", 25));
