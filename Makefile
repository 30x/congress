IMAGE_VERSION=0.1.1

build-and-package: compile-linux build-image
build-deploy-dev: compile-linux build-image push-to-dev deploy-dev-image

compile-linux:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o congress

build-image:
	docker build -t thirtyx/congress .

push-to-dev:
	docker tag -f thirtyx/congress thirtyx/congress:dev
	docker push thirtyx/congress:dev

push-new-version:
	docker tag -f thirtyx/congress thirtyx/congress:$(IMAGE_VERSION)
	docker push thirtyx/congress:$(IMAGE_VERSION)

deploy-dev-image:
	kubectl create -f congress-dev.yaml  --namespace=apigee