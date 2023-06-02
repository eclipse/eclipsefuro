```bash
docker build --pull -t thenorstroem/furo-bec:v1.38.0-amd64 .
```
```bash
docker build --pull -t thenorstroem/furo-bec:v1.38.0-arm64v8 --platform=linux/arm/v7 .
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