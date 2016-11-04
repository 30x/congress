IMAGE_VERSION=0.1.4

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
	kubectl replace -f congress-dev.yaml  --namespace=apigee

clean:
	rm congress

run-local-binary:
	env CONGRESS_ISOLATE_NAMESPACE=true \
	CONGRESS_EXCLUDES=kube-system \
	CONGRESS_ROUTING_NAMESPACE=apigee \
	CONGRESS_ROUTING_LABEL=Name \
	CONGRESS_ROUTING_POLICY_NAME=allow-routing \
	CONGRESS_IGNORE_SELECTOR="congress=exclude" \
	./congress