import { Transform } from 'stream';

export class ReplaceStream extends Transform {
  constructor(searchStr, replaceStr, options) {
    super({ ...options });
    this.searchStr = searchStr;
    this.replaceStr = replaceStr;
    this.tail = '';
  }

  _transform(chunk, encoding, callback) {
    const pieces = (this.tail + chunk).split(this.searchStr);
    const lastPiece = pieces[pieces.length - 1];

    // If there is not match, the n-1 letters of the searchStr
    // could be contained in the actual chunk
    const tailLen = this.searchStr.length - 1;
    this.tail = lastPiece.slice(-tailLen);

    pieces[pieces.length - 1] = lastPiece.slice(0, -tailLen);

    this.push(pieces.join(this.replaceStr));

    callback();
  }

  _flush(callback) {
    this.push(this.tail);
    callback();
  }

}
