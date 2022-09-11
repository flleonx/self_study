import { Writable } from 'stream';
import fs, { promises as fsp } from 'fs';
import { dirname, join } from 'path';

const tfs = new Writable({
  objectMode: true,
  write(chunk, _encoding, cb) {
    if (!fs.existsSync(dirname(chunk.path))) {
      fsp.mkdir(dirname(chunk.path))
        .then(() => console.log('Folder created'))
        .catch(cb)
    }

    fsp.writeFile(chunk.path, chunk.content)
      .then(() => cb())
      .catch(cb)
  }
})

tfs.write({
  path: join('./files', 'file1.txt'), content: 'Hello',
}, () => console.log('First file'));

tfs.write({
  path: join('./files', 'file2.txt'), content: 'Node.js'
}, () => console.log('Second file'));

tfs.write({
  path: join('./files', 'file3.txt'), content: 'streams'
}, () => console.log('Third file'));

tfs.end(() => console.log('All files created'));
