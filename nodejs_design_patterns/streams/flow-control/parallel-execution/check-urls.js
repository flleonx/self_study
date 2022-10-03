import { pipeline } from 'stream';
import { createReadStream, createWriteStream } from 'fs';
import split from 'split';
import superagent from 'superagent';
import { ParallelStream } from './parallel-stream.js';
import { LimitedParallelStream } from './limited-parallel-stream.js';

const parallelCb = async (url, _enc, push, done) => {
  if (!url)  {
    return done();
  }

  try {
    await superagent.head(url, { timeout: 5 * 1000 });
    push(`${url} is up\n`);
  } catch (error) {
    push(`${url} is down\n`);
  }

  done();
};

const doneCb = (err) => {
  if (err) {
    console.error(err);
    process.exit(1);
  }
  console.log('All urls have been checked');
}

pipeline(
  createReadStream(process.argv[2]),
  split(),
  // new ParallelStream(parallelCb),
  new LimitedParallelStream(2, parallelCb),
  createWriteStream('results.txt'),
  doneCb
)
