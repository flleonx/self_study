/*
  This is a simplified implementation of the Change Observer pattern.
  More advanced implementations support multiple observers and use more
  traps to catch other types of mutation, such as field deletions or
  changes of prototype. Moreover, this implementation does not recursively
  create proxies for nested objects or arrays - a more advance implementation
  takes care of these cases as well.
*/

export function createObservable(target, observer) {
  const observable = new Proxy(target, {
    set(obj, prop, value) {
      if (value !== obj[prop]) {
        const prev = obj[prop];
        obj[prop] = value;
        observer({ prop, prev, curr: value });
      }

      return true;
    }
  })

  return observable;
}
