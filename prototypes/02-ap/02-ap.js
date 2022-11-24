function userCreator(name, counter) {
  const newUser = Object.create(functionStore);
  newUser.name = name;
  newUser.counter = counter;

  return newUser;
}

const functionStore = {
  incrementCounter: function() { this.counter++ }
}

const user = userCreator("flleonx", 5);
user.incrementCounter();
console.log(user);
