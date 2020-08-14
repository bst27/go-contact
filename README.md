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


# Build
To build executables for multiple platforms you can use the build script at `scripts/build.sh`.