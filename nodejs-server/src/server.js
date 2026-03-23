const http = require('http');

const port = process.env.PORT || 8080;

function responseMessage() {
  const feeling = process.env.FEELING || '$FEELING';
  const awesomness = process.env.AWESOMNESS || '$AWESOMNESS';
  return `Howdy! I feeeeel ${feeling} today. I will tell you my secret of awesomeness: "${awesomness}".`;
}

const server = http.createServer((req, res) => {
  res.writeHead(200, { 'Content-Type': 'text/plain' });
  res.end(responseMessage());
});

server.listen(port, '0.0.0.0', () => {
  console.log(`Server running on http://0.0.0.0:${port}`);
});
