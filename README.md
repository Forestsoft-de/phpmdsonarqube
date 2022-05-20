### Description
This project converts phpmd json report into sonarqube format.


### Local Build

```bash
./build.sh
```

### Push image
./build.sh push

### Example Usage
```bash
docker run -v$(pwd):/workdir --rm  phpmdsonarqube -input /workdir/tmp/phpmd.json -output /workdir/tmp/sonarqube.json
```