ARG BASE_IMAGE=golang:1.18.7
FROM ${BASE_IMAGE}
# optional environment: development、test、production;  From 为变量作用域
ARG ENVIRONMENT
ENV environment ${ENVIRONMENT}
ADD ./ /code
WORKDIR /code
RUN make build
RUN make clear
EXPOSE 8000
CMD ["./server"]