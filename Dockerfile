# PURPOSE: This Dockerfile will prepare an Amazon Linux docker image with everything needed to compile binary wkhtmltopdf

FROM amazonlinux

WORKDIR /
RUN yum update -y

# Download && install Wkhtmltopdf
RUN yum -y install openssl-devel bzip2-devel libffi-devel wget tar gzip make gcc-c++
RUN wget https://github.com/wkhtmltopdf/packaging/releases/download/0.12.6-1/wkhtmltox-0.12.6-1.amazonlinux2.x86_64.rpm
RUN yum -y install wkhtmltox-0.12.6-1.amazonlinux2.x86_64.rpm

# Test wkhtmltopdf
RUN wkhtmltopdf http://google.com google.pdf

# Copy Wkhtmltopdf && its dependecies
RUN mkdir layer
WORKDIR /layer
RUN cp /usr/local/bin/wkhtmltopdf wkhtmltopdf
RUN ldd -u /usr/local/bin/wkhtmltopdf | grep lib64 | xargs cp -t ./

# Make zip file for AWS layer
RUN yum -y install zip
RUN zip -r wkhtmltopdf.zip ./* -x "wkhtmltopdf.zip"

# check archive(zip) info
RUN ls -l wkhtmltopdf.zip
