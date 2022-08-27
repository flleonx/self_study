process.stdin
  .on('data', () => {
    console.log('New data avaliable');
    console.log(
      `Chunk read (${chunk.length} bytes): "${chunk.toString()}"`
    );
  })
  .on('end', () => console.log('End of stream'))
