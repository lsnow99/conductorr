const corsproxy = require("../../functions/corsproxy");

async function testHandler() {
  let response = await corsproxy.handler({
    queryStringParameters: {
      url: "https://example.com",
    },
  })
  console.log(response)
}

testHandler();
