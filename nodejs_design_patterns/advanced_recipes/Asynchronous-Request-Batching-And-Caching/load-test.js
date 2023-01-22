import superagent from 'superagent'

const start = Date.now()
let count = 20
let pending = count
const interval = 200
const query = process.argv[2] ? process.argv[2] : 'product=book'
const PORT = 57000;

function sendRequest () {
  superagent.get(`http://localhost:${PORT}?${query}`)
    .then(result => {
      console.log(result.status, result.body)
      if (!--pending) {
        console.log(`All completed in: ${Date.now() - start}ms`)
      }
    })

  if (--count) {
    setTimeout(sendRequest, interval)
  }
}

sendRequest()
