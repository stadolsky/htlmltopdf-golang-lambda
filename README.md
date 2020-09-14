#### Build package of lambda function and zip it

`$ GOOS=linux go build -o build/main cmd/main.go & go version`

`$ zip output/function.zip build/main`

#### Build Docker image Wkhtmltopdf factory and prepare files and build AWS Layer

`$docker build --tag=wkhtmltopdf-layer-factory:latest .` 

#### Copy AWS Layer from the docker container

`docker run --rm -it -v $(pwd):/data wkhtmltopdf-layer-factory cp /layer/wkhtmltopdf.zip /data/output`

 
**In the end all files will be located in `output` dir**

```Output```:
 - `function.zip` - this file is used to create Lambda function
 - `wkhtmltopdf.zip` - this file is used to create Lambda Layer Wkhtmltopdf  