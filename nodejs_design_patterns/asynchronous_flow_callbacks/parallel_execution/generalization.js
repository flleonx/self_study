const tasks = [
  /* ... */
];

let completed = 0;
tasks.forEach((task) => {
  task(() => {
    if (++completed === tasks.length) {
      finish();
    }
  });
});

function finish() {
  // all the tasks completed
}
