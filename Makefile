#-I./account tells protoc which directories to look for imports
#make sure you use full path in option go_package to generate right imports in pb files
#use source_relative to use option go_package but generate files where you want them
#by default protoc will create dir structure based on go_package path if source relative not used
#use git tag v1.0.1 to version commits
#version is appended to go mod when downloaded
build:
	protoc \
		--proto_path=./customer --go_out=plugins=micro,paths=source_relative:./customer -I./customer -I./account \
		customer/customer.proto

	protoc \
		--proto_path=./account --go_out=plugins=micro,paths=source_relative:./account \
		account/account.proto