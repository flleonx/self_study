function iterate(index) {
  if (index === tasks.length) {
    return finish();
  }
  const task = tasks[index];
  task(() => iterate(index + 1));
}
function finish() {
  // iteration completed
}
iterate(0);

// Even more generalized
iterateSeries(collection, iteratorCallback, finalCallback)

function interateSeries(collection, iteratorCallback, finalCallback) {
  function iterate(index) {
    console.log('Iterate function');
    if (index === collection.length) {
      console.log('Exit condition');
      return finalCallback();
    }

    iteratorCallback(collection[index], (err) => {
      if (err) {
        return cb(err);
      }
      
      iterate(index + 1);
    })
  }
  iterate(0);
};
