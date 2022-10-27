package mapper

import (
	"github.com/EricMesquita/tcc-golang/internal/app/domain/book"
	"github.com/EricMesquita/tcc-golang/internal/app/transport/inbound"
)

func BookFromCreateBookRequest(request *inbound.CreateBookRequest) *book.Book {
	return &book.Book{
		Name: request.Name,
	}
}

func BookFromUpdateBookRequest(id int64, request *inbound.UpdateBookRequest) *book.Book {
	return &book.Book{
		Id:   id,
		Name: request.Name,
	}
}
