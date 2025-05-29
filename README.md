# websockify

[![Docker Image Version (latest semver)](https://img.shields.io/docker/v/isayme/websockify?sort=semver&style=flat-square)](https://hub.docker.com/r/isayme/websockify)
![Docker Image Size (latest semver)](https://img.shields.io/docker/image-size/isayme/websockify?sort=semver&style=flat-square)
![Docker Pulls](https://img.shields.io/docker/pulls/isayme/websockify?style=flat-square)

# Usage

```
// server
server --listen :6080 --vnc [vnchost]:5900 --web /path/to/vnc.html

// browser
http://127.0.0.1:6080/vnc.html?host=127.0.0.1&port=6080
```
