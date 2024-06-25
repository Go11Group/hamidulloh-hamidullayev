package main

import (
    "context"
    "log"
    "net"

    pb "dodi/proto/genproto"

    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
)

type server struct {
    pb.UnimplementedLibraryServiceServer
    books map[string]*pb.Book
}

func (s *server) AddBook(ctx context.Context, req *pb.AddBookRequest) (*pb.AddBookResponse, error) {
    book := &pb.Book{
        BookId:        generateID(), // generateID() - unique ID generatsiyasi uchun funksiya
        Title:         req.Title,
        Author:        req.Author,
        YearPublished: req.YearPublished,
    }
    s.books[book.BookId] = book
    return &pb.AddBookResponse{BookId: book.BookId}, nil
}

func (s *server) SearchBook(ctx context.Context, req *pb.SearchBookRequest) (*pb.SearchBookResponse, error) {
    var foundBooks []*pb.Book
    for _, book := range s.books {
        if book.Title == req.Query || book.Author == req.Query {
            foundBooks = append(foundBooks, book)
        }
    }
    return &pb.SearchBookResponse{Books: foundBooks}, nil
}

func (s *server) BorrowBook(ctx context.Context, req *pb.BorrowBookRequest) (*pb.BorrowBookResponse, error) {
    if _, exists := s.books[req.BookId]; exists {
        return &pb.BorrowBookResponse{Success: true}, nil
    }
    return &pb.BorrowBookResponse{Success: false}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterLibraryServiceServer(s, &server{books: make(map[string]*pb.Book)})
    reflection.Register(s)

    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

func generateID() string {
    // ID generatsiyasi uchun funksiya
    return "unique-id"
}
