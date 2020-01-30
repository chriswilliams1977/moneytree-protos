#to import files add extra includes when generating e.g -I./account
build:
	protoc \
		--proto_path=./customer --go_out=plugins=micro:./customer -I./customer -I./account \
		customer/customer.proto

	protoc \
		--proto_path=./account --go_out=plugins=micro:./account \
		account/account.proto