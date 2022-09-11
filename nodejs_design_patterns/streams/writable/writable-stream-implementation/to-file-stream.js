import { Writable } from 'stream';
import fs from 'fs';
import { promises as fsp } from 'fs';
import { dirname } from 'path';

export class ToFileStream extends Writable {
  constructor(options) {
    super({ ...options, objectMode: true });
  }

  _write(chunk, encoding, cb) {
    if (!fs.existsSync(dirname(chunk.path))) {
      fsp.mkdir(dirname(chunk.path))
        .then(() => console.log('Directory created'))
        .catch(cb)
    } 

    fsp.writeFile(chunk.path, chunk.content)
      .then(() => cb())
      .catch(cb)
  }

}
