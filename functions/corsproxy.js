// Simple Netlify lambda function to proxy a GET request to bypass CORS

const https = require('https')

// Function to perform get request using Node standard library
async function doGet(url, maxRedirects=10) {
    return new Promise((resolve, reject) => {
        https.get(url, async (res) => {
            let data = [];

            if(res.statusCode === 301 || res.statusCode === 302) {
                if (maxRedirects == 0) {
                    reject("max redirects reached")
                }
                let response = await doGet(res.headers.location, maxRedirects - 1)
                resolve(response)
            }
        
            res.on('data', chunk => {
                data.push(chunk);
            });
        
            res.on('end', () => {
                resolve(Buffer.concat(data).toString())
            });
        }).on('error', err => {
            reject(err)
        });
    })
}

// Lambda handler
exports.handler = async (event, context) => {
    const url = event.queryStringParameters['url']

    try {
        let response = await doGet(url)
        return {
            statusCode: 200,
            body: response
        }
    } catch (error) {
        return {
            statusCode: 502,
            body: error
        }
    }   
}
