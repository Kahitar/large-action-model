# Large Action Model Tools

Tools for my [Large Action Model](https://chat.openai.com/g/g-NoWMWIBMr-large-action-model) to utilize.

## Hello World

The hello world action can be executed by asking [Large Action Model](https://chat.openai.com/g/g-NoWMWIBMr-large-action-model):

    "Give me a greeting"


## Localhost Development

ChatGPT Actions require the API to be reachable from the web. Just running the app in localhost does not work.
To still be able to test it locally, you can use [ngrok](https://dashboard.ngrok.com/), which gives you a static
url and when you run it's executable locally, will serve your localhost to that static url. You can then use that
url in your ChatGPT action.

Create an account with ngrok (free tier is enough) and follow the installation guide. Than run ngrok locally:

    $ ngrok http --domain=<YOUR_NGROK_URL> 8080

Now run the API locally. For example for the hello world go app:

    $ cd hello-world
    $ go run main.go


## OpenAPI 3.0

ChatGPT Actions require the OpenAPI spec to be in version 3.0. So far I found no tool to generate a 3.0 spec
from my go code. But a good start is to use the [swag](https://github.com/swaggo/swag) go library and modify
it from there. ChatGPT can do the conversion fro you.

The spec also needs to contain the "operationId" field (unique identifier for an endpoint), which is not generated
automatically. Add it next to the 'summary' and 'description' entries under the endpoint.

Here is an example spec that works:

```json
{
  "openapi": "3.0.0",
  "info": {
    "description": "This is a sample server for a hello world API.",
    "version": "1.0",
    "title": "Hello World API",
    "contact": {}
  },
  "servers": [
    {
      "url": "<YOUR_API_URL>"
    }
  ],
  "paths": {
    "/hello": {
      "post": {
        "tags": [
          "hello"
        ],
        "summary": "Return a random greeting",
        "description": "Get a random greeting message",
        "operationId": "getRandomGreeting",
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "type": "object",
                "properties": {}
              }
            }
          }
        },
        "responses": {
          "200": {
            "description": "OK",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "additionalProperties": {
                    "type": "string"
                  }
                }
              }
            }
          }
        }
      }
    }
  }
}
```

This still gives an error `In context=('paths', '/hello', '200', 'response', 'content', 'application/json', 'schema'), object schema missing properties`,
but it works.
