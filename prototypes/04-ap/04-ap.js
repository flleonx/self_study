class UserCreator {
  constructor(name, counter) {
    this.name = name;
    this.counter = counter;
  }

  incrementCounter() {
    this.counter++;
  }
}

const user = new UserCreator("flleonx", 7);
user.incrementCounter();
console.log(user);
