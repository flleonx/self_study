import EventEmitter from 'events';

const ticker = (number, callback) => {
  const emitter = new EventEmitter();
  const totalTicks = number/0.05;
  let actualTicks = 0;
  const tickCounter = () => {

    try {
      if (Date.now() % 5 === 0) {
        throw new Error('Divisible by 5');
      }
    } catch (err) {
      process.nextTick(() => emitter.emit('error', err));
      return callback(err, null);
    }

    if (totalTicks === actualTicks) {
      emitter.emit('tick', actualTicks);
      callback(actualTicks);
    } else {
      actualTicks++;
      setImmediate(() => emitter.emit('tick', actualTicks));
      setTimeout(tickCounter, 500);
    }
  }
  tickCounter();

  return emitter;
};

const sub = ticker(1, (err, finalTick) => {
  if (err) {
    return console.log(`Callback error ${err}`);
  }

  console.log(`Final: ${finalTick}`);
});

sub.on('tick', tick => console.log(`Tick number: ${tick}`));
sub.on('error', err => console.log(`Event error ${err}`));
