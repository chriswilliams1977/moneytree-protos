build:
	protoc \
	--go_out=plugins=micro:./customer -I./customer \
	customer/customer.proto

	protoc \
		--go_out=plugins=micro:./account -I./account -I./customer \
		account/account.proto