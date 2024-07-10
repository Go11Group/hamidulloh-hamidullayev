package postgres

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq"
	pb "EcoTrack/community/genproto/protos"
)

// TestCreateGroup funksiyasi CreateGroup funksiyasini test qiladi
func TestCreateGroup(t *testing.T) {
	// sqlmock yordamida yangi mock database yaratamiz
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error opening a stub database connection: %v", err)
	}
	defer db.Close() // Test tugagandan keyin database-ni yopamiz

	// CommunityRepo strukturasining yangi obyektini yaratamiz
	s := NewCommunityRepo(db)
	// Test uchun CreateGroupRequest yaratamiz
	req := &pb.CreateGroupRequest{
		Name:        "Test Group",
		Description: "Test Description",
		CreatedBy:   1,
	}

	// mock.ExpectQuery - SQL so'rovi uchun kutishlarni o'rnatamiz
	mock.ExpectQuery("INSERT INTO groups").
		WithArgs(req.Name, req.Description, req.CreatedBy). // Kutishlarni so'rov argumentlari bilan tekshiramiz
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1)) // So'rov natijasini qaytarish uchun

	// CreateGroup funksiyasini chaqiramiz va natijani tekshiramiz
	resp, err := s.CreateGroup(req)
	if err != nil {
		t.Errorf("unexpected error: %v", err) // Agar xatolik yuzaga kelsa, error message qaytaramiz
	}

	// Natijani tekshiramiz
	if resp.Group.Id != 1 {
		t.Errorf("expected id 1, got %v", resp.Group.Id) // Agar natija kutilganidek bo'lmasa, error message qaytaramiz
	}
}
