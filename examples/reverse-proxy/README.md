# reverse-proxy

In this example, we build a reverse proxy. Run github api with your worker!

## Deploy

```shell
make docker
```

```shell
make publish
```

## Try

Fetch stargazers

```shell
curl https://${YOUR_WORKER_URL}/repos/aki-0421/transport.wasm/stargazers
```