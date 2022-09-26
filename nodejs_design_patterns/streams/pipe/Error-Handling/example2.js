// This allows us to ensure that all the allocated resources
// are properly released, and the error is handled gracefully

function handleError(err) {
  console.error(err);
  stream1.destroy();
  stream2.destroy();
}

stream1
  .on('error', handleError)
  .pipe(stream2)
  .on('error', handleError)
