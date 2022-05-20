### Local Build

```bash
docker build -t forestsoft/phpmdsonarqube .
```

### Run
```bash
docker run -v$(pwd):/workdir --rm  phpmdsonarqube -input /workdir/tmp/phpmd.json -output /workdir/tmp/sonarqube.json
```