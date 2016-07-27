IMAGE_VERSION=0.1.0

build-and-package: compile-linux build-image

compile-linux:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o congress

build-image:
	docker build -t thirtyx/congress .

push-to-hub:
	docker tag -f thirtyx/congress thirtyx/congress:$(IMAGE_VERSION)
	docker push thirtyx/congress:$(IMAGE_VERSION)