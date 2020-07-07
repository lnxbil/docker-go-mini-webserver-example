# Minimal Webserver with Request Header Output

This minimal working example (MWE) shows how you can create a simple docker
container with a single go binary that provides a webserver running on port
`8081` and reports back all request headers.

I personally use this image to test [Traefik](https://containo.us/traefik/)
configurations like this:

```yaml
version: '3.7'

services:
  webserver:
    image: lnxbil/docker-go-mini-webserver-example:latest
    restart: unless-stopped
    networks:
    - traefik-web
    labels:
    - traefik.enable=true
    - traefik.http.routers.http-plain.tls=true
    - traefik.http.routers.http-plain.entrypoints=https
    - traefik.http.routers.http-plain.rule=Host(`test0.mydomain`)
    - traefik.http.routers.http-header1.tls=true
    - traefik.http.routers.http-header1.entrypoints=https
    - traefik.http.routers.http-header1.rule=Host(`test1.mydomain`)
    - traefik.http.routers.http-header1.middlewares=http-header1-add-header
    - traefik.http.middlewares.http-header1-add-header.headers.customrequestheaders.X-FE-Landing-Page-Code=TestHeader1

networks:
  traefik-web:
    external: true
```

This is only the local configuration, not the Traefik configuration itself.
