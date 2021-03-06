package repository

import "context"

type NoSQL interface {
	Store(context.Context, string, interface{}) error
	Update(context.Context, string, interface{}, interface{}) error
	FindAll(context.Context, string, interface{}, interface{}) error
	FindOne(context.Context, string, interface{}, interface{}, interface{}) error
}
