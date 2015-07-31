
# OAuth profile extract

Library used to make HTTP request and extract the common profile from the response.

NB! To make the request make sure you have the `access_token`.

Providers supported:

  - Facebook - `facebook`
  - Github - `github`

## Usage

Take a look at the test file to see example.

```
profile, err := Extract("facebook", os.Getenv("FB_TEST_TOKEN"))
```

## TODO

  - Add more providers currently Google, Windows Live other providers are not supported
