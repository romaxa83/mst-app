package commands

import "github.com/romaxa83/mst-app/gateway/internal/dto/author"

type AuthorCmds struct {
	CreateAuthor CreateAuthorCmdHandler
}

func NewAuthorCmds(createAuthor CreateAuthorCmdHandler) *AuthorCmds {
	return &AuthorCmds{
		CreateAuthor: createAuthor,
	}
}

type CreateAuthorCmd struct {
	CreateDto *dto.CreateAuthorDto
}

func NewCreateAuthorCmd(dto *dto.CreateAuthorDto) *CreateAuthorCmd {
	return &CreateAuthorCmd{CreateDto: dto}
}
