import fs from 'fs';
import lazystream from 'lazystream';

const lazyURandom = new lazystream.Readable(function (options) {
  return fs.createReadStream('/dev/urandom');
});

lazyURandom.on('readable',() => {
    let chunk
    console.log('New data avaliable');
    while ((chunk = lazyURandom.read()) !== null) {
      console.log(
        `Chunk read (${chunk.length} bytes): "${chunk.toString()}"`
      );
    }
  })
  .on('end', () => console.log('End of stream'))
