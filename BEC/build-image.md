```bash
docker build -t thenorstroem/furo-bec:v1.38.0-amd64 --build-arg ARCH=amd64/ .
docker build -t thenorstroem/furo-bec:v1.38.0-arm64v8 --build-arg ARCH=arm64v8/ .
```

```bash
docker push thenorstroem/furo-bec:v1.38.0-amd64
docker push thenorstroem/furo-bec:v1.38.0-arm64v8
```

```bash
docker manifest create \
thenorstroem/furo-bec:v1.38.0 \
--amend thenorstroem/furo-bec:v1.38.0-amd64 \
--amend thenorstroem/furo-bec:v1.38.0-arm64v8
```

```bash
docker manifest push thenorstroem/furo-bec:v1.38.0
```