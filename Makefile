pb-gen:
	protoc -I protofiles/ protofiles/${service}.proto --go_out=plugins=grpc:services/${service}/proto && \
	protoc -I protofiles/ protofiles/${service}.proto --go_out=plugins=grpc:services/graphql/proto