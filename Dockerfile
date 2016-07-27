FROM alpine:3.4

COPY congress /

ENV CONGRESS_EXCLUDES "kube-system,apigee,apigee-jenkins,calico-system,guardians"

CMD ["/congress"]