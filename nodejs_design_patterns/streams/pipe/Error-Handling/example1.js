stream1
  .on('error', () => {})
  .pipe(stream2)
  .on('error', () => {})
