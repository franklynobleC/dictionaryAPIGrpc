package service

import "context"

type Repository interface {
	SearchWords(ctx context.Context, getWords *Wordrequest) (*WordResponse, error)
}
