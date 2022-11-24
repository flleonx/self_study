/*
  Using closures each user creates incrementCounter function
  independently.

  A really bad approach regarding memory efficiency
*/

function userCreator(name, counter) {
  const newUser = {};
  newUser.name = name;
  newUser.counter = counter;
  newUser.incrementCounter = function() {
    newUser.counter++;
  }

  return newUser;
}

const user = userCreator("flleonx", 2);
user.incrementCounter();
console.log(user);
