FROM alpine:3.4

COPY ./build/congress /

CMD ["/congress"]