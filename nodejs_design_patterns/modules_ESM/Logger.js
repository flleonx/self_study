export default class Logger {
  constructor (name) {
    this.name = name;
  }
  static sum() {
    return 1 + 9;
  }
  log (message) {
    console.log(`[${this.name}] ${message}`);
  }
}

console.log(Logger.sum());
