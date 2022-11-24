function UserCreator(name, counter) {
  this.name = name;
  this.counter = counter;
}

UserCreator.prototype.incrementCounter = function() {
  this.counter++;
}

const user = new UserCreator("flleonx", 3);
user.incrementCounter();
console.log(user);
