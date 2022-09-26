import { createReadStream } from 'fs';
import { createBrotliCompress } from 'zlib';
import { PassThrough } from 'stream';
import { basename } from 'path';
import { upload } from './upload.js';

const filepath = process.argv[2];
const filename = basename(filepath);
const contentStream = new PassThrough();

upload(`${filename}.br`, contentStream)
  then((response) => {
    console.log(`Server response: ${response.data}`);
  })
  .catch((err) => {
    console.error(err);
    process.exit(1)
  })

createReadStream(filepath)
  .pipe(createBrotliCompress())
  .pipe(contentStream)


// We can return a writable stream if we want
// to push the data
// Using a PassThrough stream

function createUploadStream(filename) {
  const connector = new PassThrough();
  upload(filename, connector);
  return connector;
}

const upload = createUploadStream('a-file.txt');
upload.write('Hello World');
upload.end();
