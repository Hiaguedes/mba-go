package graph

import "github.com/hiaguedes/graphql-fullcycle/internal/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

type Resolver struct{
	CategoryDb *database.Category
	CourseDb *database.Course
}
