# About
This app forwards contact forms via webhook. It receives form submissions via http and forwards them to a defined
webhook target.

# Usage
```
./go-contact -url "<URL>"
```
```
  -debug
    	Enable debug mode
  -messageKey string
    	The webhook target will receive the message value of a form submission with this key in the json data (default "message")
  -nameKey string
    	The webhook target will receive the name value of a form submission with this key in the json data (default "name")
  -port int
    	Port to listen to for form submissions (default 8080)
  -url string
    	URL to receive form submissions (webhook target)

```

# Example
Execute the application with the following config:
```
./go-contact -url "http://localhost/webhook-receiver" -port 8085 -nameKey "value1" -messageKey "value2"
```

It will start the app and listen on port 8085. After a contact form submission is received it will be forwarded
by the following http post request to your defined url:
```
POST /webhook-receiver HTTP/1.1
Host: localhost
Content-Length: 45
Content-Type: application/json

{
  "value2": "Hi there",
  "value1": "Max"
}
```


# Build
To build executables for multiple platforms you can use the build script at `scripts/build.sh`.