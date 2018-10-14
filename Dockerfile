FROM fedora:28
ADD . /code
WORKDIR /code
ENV GOPATH=/gopath
RUN mkdir /gopath
RUN dnf -y install libexif-devel git golang
RUN go get github.com/xiam/exif github.com/nfnt/resize github.com/gorilla/sessions github.com/gorilla/mux github.com/disintegration/imaging
