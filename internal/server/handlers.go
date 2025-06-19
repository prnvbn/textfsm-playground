package server

import (
	"context"
)

type NameResponse struct {
	Body struct {
		Message string `json:"message"`
	} `json:"body"`
}

func (s *Server) NameHandler(ctx context.Context, params *struct {
	Name string `path:"name"`
}) (*NameResponse, error) {
	resp := NameResponse{}
	resp.Body.Message = "Hello, " + params.Name + "!"
	return &resp, nil
}
