import { pipeline } from 'stream';
import { createReadStream, createWriteStream } from 'fs';
import split from 'split';
import superagent from 'superagent';
import ParallelTransform from 'parallel-transform';

async function parallelCb(url, done) {
  if (!url)  {
    return done();
  }
  
  console.log('URL:', url);

  try {
    await superagent.head(url, { timeout: 5 * 1000 });
    this.push(`${url} is up\n`);
  } catch (error) {
    this.push(`${url} is down\n`);
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
  ParallelTransform(4, parallelCb),
  createWriteStream('results.txt'),
  doneCb
)
