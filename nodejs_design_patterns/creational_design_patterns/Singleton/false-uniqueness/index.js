import { getDbInstance as getDbFromA } from 'package-a';
import { getDbInstance as getDbFromB } from 'package-b';
console.log(global)
const a = this;
console.log(a);

const isSame = getDbFromA() === getDbFromB()
console.log('Is the db instance in package-a the same ' +
  `as package-b? ${isSame ? 'YES' : 'NO'}`)
