```bash
docker build -t thenorstroem/furo-bec:manifest-amd64 --build-arg ARCH=amd64/ .
docker build -t thenorstroem/furo-bec:manifest-arm64v8 --build-arg ARCH=arm64v8/ .
```

```bash
docker push thenorstroem/furo-bec:manifest-amd64
docker push thenorstroem/furo-bec:manifest-arm64v8
```

```bash
docker manifest create \
thenorstroem/furo-bec:latest \
--amend thenorstroem/furo-bec:manifest-amd64 \
--amend thenorstroem/furo-bec:manifest-arm64v8
```

```bash
docker manifest push thenorstroem/furo-bec:latest
```