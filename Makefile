#-I./account tells protoc which directories to look for imports
#make sure you use full path in option go_package to generate right imports in pb files
#use source_relative to use option go_package but generate files where you want them
#by default protoc will create dir structure based on go_package path if source relative not used
#UPDATING IMPORTS
#Create new tag: git tag v1.0.1
#Push tags git push --tags
#Download latest packages: go get github.com/chriswilliams1977/moneytree-protos@v1.0.22
#version is appended to go mod when downloaded
build:
	protoc \
		--proto_path=./customer --go_out=plugins=micro,paths=source_relative:./customer -I./customer -I./account -I./address -I./transaction \
		customer/customer.proto

	protoc \
		--proto_path=./account --go_out=plugins=micro,paths=source_relative:./account  -I./account -I./transaction \
		account/account.proto

	protoc \
    		--proto_path=./transaction --go_out=plugins=micro,paths=source_relative:./transaction \
    		transaction/transaction.proto

	protoc \
		--proto_path=./address --go_out=plugins=micro,paths=source_relative:./address \
		address/address.proto

	protoc \
		--proto_path=./bank --go_out=plugins=micro,paths=source_relative:./bank  -I./bank -I./address \
		bank/bank.proto

