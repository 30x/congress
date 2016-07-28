FROM alpine:3.4

COPY congress /

ENV CONGRESS_EXCLUDES "default,kube-system,apigee,apigee-jenkins,calico-system,guardians"
ENV CONGRESS_SELECTOR "runtime=shipyard"
ENV CONGRESS_IGNORE_SELECTOR "runtime=other"


CMD ["/congress"]