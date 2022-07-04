import http from 'http';

const server = http.createServer();

server.on('request', (req, res) => {
  res.end('<h1>Hello</h1>');
});

server.listen(56000);
