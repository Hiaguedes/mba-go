package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/hiaguedes/sqlc/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

type CourseDb struct {
	conn *sql.DB
	* db.Queries
}

type CourseParams struct {
	ID          string
	Name        string
	Description sql.NullString
	CategoryID  string
	Price       float64
}

type CategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

func NewCourseDb(conn *sql.DB) *CourseDb {
	return &CourseDb{
		conn:    conn,
		Queries: db.New(conn),
	}
}

func (c *CourseDb) RunInTransaction(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := db.New(tx)
	err = fn(q)

	if err != nil {
		if errRb := tx.Rollback(); errRb != nil {
			return fmt.Errorf("tx original error: %v, rollback err: %v", err, errRb)
		}
		return err
	}

	return tx.Commit()
}

func (c *CourseDb) CreateCourseWithCategory(ctx context.Context, courseParams CourseParams, categoryParams CategoryParams) error {
	return c.RunInTransaction(ctx, func(q *db.Queries) error {
		err := q.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          categoryParams.ID,
			Name:        categoryParams.Name,
			Description: categoryParams.Description,
		})

		if err != nil {
			return err
		}

		return q.CreateCourse(ctx, db.CreateCourseParams{
			ID:          courseParams.ID,
			Name:        courseParams.Name,
			Description: courseParams.Description,
			CategoryID:  categoryParams.ID,
			Price: 		 courseParams.Price,
		})
	})
}

func main() {
	ctx := context.Background()
	conn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	// courseArgs := CourseParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "BackEnd With Golang",
	// 	Description: sql.NullString{String: "Description for BackEnd With Golang", Valid: true},
	// 	Price: 100.10,
	// }

	// categoryArgs := CategoryParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Programming",
	// 	Description: sql.NullString{String: "Description for Programming", Valid: true},
	// }

	// courseDb := NewCourseDb(conn)

	// err = courseDb.CreateCourseWithCategory(ctx, courseArgs, categoryArgs)
	// if err != nil {
	// 	panic(err)
	// }

	queries := db.New(conn)
	courses, err := queries.ListCourses(ctx)

	for _, course := range courses {
		fmt.Printf("Course: %v, Category: %v\n", course, course.CategoryName)
	}

}