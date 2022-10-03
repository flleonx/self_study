import { Transform } from 'stream';

export class LimitedParallelStream extends Transform {
  constructor(concurrency, userTransform, opts) {
    super({ ...opts, objectMode: true });
    this.userTransform = userTransform;
    this.running = 0;
    this.terminateCb = null;
    // new
    this.concurrency = concurrency;
    this.continueCb = null;
  }

  _transform(chunk, encoding, done) {
    this.running++;

    this.userTransform(
      chunk,
      encoding,
      this.push.bind(this),
      this._onComplete.bind(this)
    )

    if (this.running < this.concurrency) {
      done();
    } else {
      console.log('Max concurrent task reached');
      this.continueCb = done;
    }
  }

  _flush(done) {
    if (this.running > 0) {
      this.terminateCb = done;
    } else {
      done();
    }
  }

  _onComplete(err) {
    this.running--;

    if (err) {
      return this.emit('error', err);
    }

    const tmpCb = this.continueCb;
    this.continueCb = null;
    tmpCb && tmpCb();

    if (this.running === 0) {
      this.terminateCb && this.terminateCb();
    }
  }
}
