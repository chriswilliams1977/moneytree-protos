#to import files add extra includes when generating e.g -I./account
build:
	protoc \
	--go_out=plugins=micro:./customer -I./customer -I./account \
	customer/customer.proto

	protoc \
		--go_out=plugins=micro:./account -I./account \
		account/account.proto