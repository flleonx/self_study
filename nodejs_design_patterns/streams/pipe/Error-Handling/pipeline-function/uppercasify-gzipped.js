import { createGzip, createGunzip } from 'zlib';
import { Transform, pipeline } from 'stream';

const uppercasify = new Transform({
  transform(chunk, _encoding, callback) {
    this.push(chunk.toString().toUpperCase());
    callback();
  }
})

pipeline(
  process.stdin,
  createGunzip(),
  uppercasify,
  createGzip(),
  process.stdout,
  (err) => {
    if (err) {
      console.error(err);
      process.exit(1)
    }
  }
)
