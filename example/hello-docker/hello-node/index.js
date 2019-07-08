const http = require('http');

const server = http.createServer(function (req, res) {
  console.log(`${req.method} ${req.url}`)
  res.write(`Hello Docker\nI'm nodejs server running on docker container`);
  res.end();
})

console.log("serve nodejs server on port 3000")
server.listen(3000)