import { createServer } from 'http'
import { createWriteStream } from 'fs'
import { createGunzip } from 'zlib'
import { basename, join } from 'path'
import { createCipheriv, randomBytes } from 'crypto'

const secret = randomBytes(24);
console.log(`Generated secret: ${secret.toString('hex')}`);

const server = createServer((req, res) => {
  const filename = basename(req.headers['x-filename']);
  const iv = Buffer.from(req.headers(['x-initialization-vector'], 'hex'));
  const destFilename = join('received_files', filename);

  console.log(`File request received: ${filename}`);

  req
    .pipe(createCipheriv('aes192', secret, iv))
    .pipe(createGunzip())
    .pipe(createWriteStream(destFilename))
    .on('finish', () => {
      res.writeHead(201, { 'Content-Type': 'text/plain' });
      res.end('Ok\n');
      console.log(`File saved: ${destFilename}`);
    })
    .on('error', (error) => {
      res.writeHead(201, { 'Content-Type': 'text/plain' });
      res.end('Error\n');
      console.error(error);
    })
})

server.listen(3000, () => {
  console.log(`Listening`);
})