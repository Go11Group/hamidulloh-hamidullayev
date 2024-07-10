package file

import (
	"context"

	pb "dodi/proto/genproto"
)

type Sserver struct {
	pb.UnimplementedLibraryServiceServer
	books map[string]*pb.Book
}

func generateID() string {
	return "unique-id"
}

func (s *Sserver) AddBook(ctx context.Context, req *pb.AddBookRequest) (*pb.AddBookResponse, error) {
	book := &pb.Book{
		BookId:        generateID(),
		Title:         req.Title,
		Author:        req.Author,
		YearPublished: req.YearPublished,
	}
	s.books[book.BookId] = book
	return &pb.AddBookResponse{BookId: book.BookId}, nil
}

func (s *Sserver) SearchBook(ctx context.Context, req *pb.SearchBookRequest) (*pb.SearchBookResponse, error) {
	var foundBooks []*pb.Book
	for _, book := range s.books {
		if book.Title == req.Query || book.Author == req.Query {
			foundBooks = append(foundBooks, book)
		}
	}
	return &pb.SearchBookResponse{Books: foundBooks}, nil
}

func (s *Sserver) BorrowBook(ctx context.Context, req *pb.BorrowBookRequest) (*pb.BorrowBookResponse, error) {
	if _, exists := s.books[req.BookId]; exists {
		return &pb.BorrowBookResponse{Success: true}, nil
	}
	return &pb.BorrowBookResponse{Success: false}, nil
}
