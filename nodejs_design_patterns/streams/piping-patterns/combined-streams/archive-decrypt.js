import { createReadStream, createWriteStream } from 'fs';
import { pipeline } from 'stream';
import { createDecryptAndDecompress } from './combined-streams.js';

const [,, password, source] = process.argv;
const iv = Buffer.from('12edbc22116c8bfabcffa9d463ca9356', 'hex');
const destination = `${source.split('.')[0]}.txt`;

pipeline(
  createReadStream(source),
  createDecryptAndDecompress(password, iv),
  createWriteStream(destination),
  (err) => {
    if (err) {
      console.error(err);
      process.exit(1);
    }
    console.log(`File decrypted and decompressed`);
  }
)
