package main

import (
    "context"
    "log"
    "math/rand"
    "os"
    "strconv"
    "time"

    pb "dodi/proto/genproto"

    "google.golang.org/grpc"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

    c := pb.NewLibraryServiceClient(conn)

    ctx := context.Background()

    // Ko'p kitoblarni qo'shish
    books := []pb.AddBookRequest{
        {Title: "Alpomish", Author: "Abdulla Qodiriy", YearPublished: 1923},
        {Title: "O'tgan kunlar", Author: "Abdulla Qodiriy", YearPublished: 1926},
        {Title: "Kecha va Kunduz", Author: "Abdulla Qahhor", YearPublished: 1944},
        {Title: "Shum bola", Author: "G'afur G'ulom", YearPublished: 1936},
        {Title: "Ikki eshik orasi", Author: "Abdulla Qahhor", YearPublished: 1966},
        {Title: "O'limdan kuchli", Author: "Cho'lpon", YearPublished: 1936},
        {Title: "Mehrobdan chayon", Author: "Abdulla Qodiriy", YearPublished: 1929},
        {Title: "Sariq devni minib", Author: "Xudoyberdi To'xtaboyev", YearPublished: 1972},
        {Title: "Qaldirg'och", Author: "Abdulla Qahhor", YearPublished: 1958},
        {Title: "Jimjitlik", Author: "Sa'dulla Siyoev", YearPublished: 1981},
    }

    var bookIDs []string
    for _, book := range books {
        addBookResponse, err := c.AddBook(ctx, &book)
        if err != nil {
            log.Fatalf("could not add book: %v", err)
        }
        if addBookResponse != nil {
            bookIDs = append(bookIDs, addBookResponse.BookId)
            log.Printf("Book ID: %s added with title: %s", addBookResponse.BookId, book.Title)
        }
    }

    // Limitni olish
    limit := len(bookIDs)
    if len(os.Args) > 1 {
        userLimit, err := strconv.Atoi(os.Args[1])
        if err == nil && userLimit > 0 && userLimit < limit {
            limit = userLimit
        }
    }

    // Tasodifiy kitoblarni tanlash
    rand.Seed(time.Now().Unix())
    selectedBookIDs := make([]string, 0, limit)
    selectedIndices := rand.Perm(len(bookIDs))[:limit]
    for _, idx := range selectedIndices {
        selectedBookIDs = append(selectedBookIDs, bookIDs[idx])
    }

    // Tanlangan kitoblarni ko'rsatish
    log.Println("Selected Books:")
    for i, bookID := range selectedBookIDs {
		for j, v := range books {
			if i == j{
				log.Println(bookID, v.Title)
			}
		}
        
    }

    // Tanlangan bir kitobni ijaraga olish
    selectedBookID := selectedBookIDs[rand.Intn(len(selectedBookIDs))]
    borrowBookResponse, err := c.BorrowBook(ctx, &pb.BorrowBookRequest{
        BookId: selectedBookID,
        UserId: "user-123",
    })
    if err != nil {
        log.Fatalf("could not borrow book: %v", err)
    }
    log.Printf("Borrow Success: %v for Book ID: %s", borrowBookResponse.Success, selectedBookID)
}
