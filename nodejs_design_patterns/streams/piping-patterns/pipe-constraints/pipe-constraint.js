import { createReadStream, createWriteStream } from 'fs';
import { Transform, pipeline } from 'stream';
import { strict as assert } from 'assert';

const streamA = createReadStream('test.json');

const streamB = new Transform({
  transform(chunk, encoding, done) {
    this.push(chunk.toString().toUpperCase());
    done();
  }
});

const streamC = createWriteStream('test-uppercase.json');

const pipelineReturn = pipeline(
  streamA,
  streamB,
  streamC,
  () => {
    // Handle errors here
  }
)

assert.strictEqual(streamC, pipelineReturn); // valid
const pipeReturn = streamA.pipe(streamB).pipe(streamC);
assert.strictEqual(streamC, pipeReturn); // valid
