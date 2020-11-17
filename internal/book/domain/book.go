//go:generate protoc -I/usr/local/include -I. --go_out=Minternal/book/domain/book.proto=.:. --go_opt=paths=source_relative book.proto

package domain
