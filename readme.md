# Large Action Model Tools

Tools for my [Large Action Model](https://chat.openai.com/g/g-NoWMWIBMr-large-action-model) to utilize.

## Hello World

The hello world action can be executed by asking [Large Action Model](https://chat.openai.com/g/g-NoWMWIBMr-large-action-model):

    "Give me a greeting"

## ChatGPT OAuth Flow

ChatGPT does this request to the authorization url:

    https://careful-broadly-kitten.ngrok-free.app/finances/token?
        response_type=code&
		client_id=12910ccd&
		redirect_uri=https%3A%2F%2Fchat.openai.com%2Faip%2Fg-1bc9e788468a358cb5138906b5c6d706d0ecb2cc%2Foauth%2Fcallback&
		scope=all&
		state=8288346b-f30b-4aca-bcc8-0c99c5a6cff7

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

## NGINX Local Setup

### Start nginx with local config file

The command is:

    $ nginx -c nginx.conf

For this to work, all nginx config files need to be present and the `logs/` and `temp/` folders must exist (on windows at least).

### List nginx processes

    $ export MSYS_NO_PATHCONV=1
    $ tasklist /fi "imagename eq nginx.exe"

### Stop nginx

Fast shutdown:

    $ nginx -s stop 

Gracefull shutdown:

    $ nginx -s quit 

Kill tasks:

    $ export MSYS_NO_PATHCONV=1
    $ taskkill /f /im nginx.exe

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

## Database

We use [turso](https://turso.tech/) as database and [atlas](https://atlasgo.io/) (see 
[turso blog](https://blog.turso.tech/database-migrations-made-easy-with-atlas-df2b259862db))for migrations.

### Migrations (Atlas)

1. Create a schema file (similar to terraform):

atlas-example.hcl:
```hcl
table "users" {
    schema = schema.main

    column "id" {
        null = true
        type = text
    }
    column "username" {
        null = false
        type = text
    }

    primary_key {
        columns = [column.id]
    }
}
schema "main" {
}

```

2. Create the database and store url and token in an environment variable:

    $ turso db create atlas-example
    $ turso db show atlas-example --url
    libsql://atlas-example-glommer.turso.io
    $ export TURSO_DB_URL=libsql+wss://atlas-example-glommer.turso.io
    $ export TURSO_DB_TOKEN=$(turso db tokens create atlas-example)

3. Apply the migration

    $ atlas schema apply -u "${TURSO_DB_URL}?authToken=${TURSO_DB_TOKEN}" \
        --to file://atlas-example.hcl --exclude '_litestream_seq,_litestream_lock'
