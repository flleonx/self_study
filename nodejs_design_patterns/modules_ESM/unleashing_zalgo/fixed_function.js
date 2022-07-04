import { readFile } from 'fs';

const cache = new Map();

export default function inconsistentRead(filename, cb) {
  if (cache.has(filename)) {
    // invoked synchronously
    // It pushes the callback at the beginning of the event queue
    process.nextTick(() => cb(cache.get(filename)));
  } else {
    // asynchronous function
    readFile(filename, 'utf8', (err, data) => {
      cache.set(filename, data);
      cb(data);
    })
  }
}
