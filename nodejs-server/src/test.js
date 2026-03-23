const http = require('http');
const assert = require('assert');

function testServer() {
  const feeling = process.env.FEELING || '$FEELING';
  const awesomness = process.env.AWESOMNESS || '$AWESOMNESS';
  const expected = `Howdy! I feeeeel ${feeling} today. I will tell you my secret of awesomeness: "${awesomness}".`;
  
  const options = {
    hostname: 'localhost',
    port: 8080,
    path: '/',
    method: 'GET'
  };

  const req = http.request(options, (res) => {
    let data = '';
    
    res.on('data', (chunk) => {
      data += chunk;
    });

    res.on('end', () => {
      assert(res.statusCode === 200, `Expected status 200, got ${res.statusCode}`);
      assert.strictEqual(data, expected, `Expected exact response, got: ${data}`);
      console.log('✓ Test passed: Server responds with updated message');
      console.log(`  Response: ${data.trim()}`);
      process.exit(0);
    });
  });

  req.on('error', (e) => {
    console.error('✗ Test failed:', e.message);
    process.exit(1);
  });

  req.end();

  // Timeout after 5 seconds
  setTimeout(() => {
    console.error('✗ Test failed: Server did not respond');
    process.exit(1);
  }, 5000);
}

// Run test
testServer();
