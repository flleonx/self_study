/*
  A popular application of the Revealing Constructor pattern is in the
  JavaScript Promise class.

  When we create a new Promise from scratch, its constructor accepts
  as input an executor function that will receive the resolve() and
  reject() functions used to mutate the internal state of the Promise.
*/
return new Promise((resolve, reject) => {
  // ...
})
/*
  Once created, the Promise state cannot be altered by any other
  means. All we can do is receive its fulfilment value or rejection reason.
*/
