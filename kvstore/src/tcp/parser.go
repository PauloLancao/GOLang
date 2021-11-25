package tcp

import (
	"logging"
	"response"
	"strings"
)

// ParamKeys const
const (
	ParamCmd  string = "cmd"
	ParamKey  string = "key"
	ParamBody string = "body"
)

// Parser struct
type Parser struct {
	Content string
}

// ParsePipe func commandline
func (p Parser) ParsePipe() (interface{}, error) {
	errCtx := "Parser::ParsePipe"
	uuid := logging.UUID()
	clArgs := strings.Split(p.Content, "|")

	logging.Msgf(uuid, errCtx, logging.Input, "Parse Content %s:", p.Content)

	if clArgs == nil || len(clArgs) == 0 {
		logging.Msg(uuid, errCtx, logging.Error, response.ErrorCodeText(response.ErrArgumentException))
		return nil, response.New(response.ErrArgumentException, errCtx)
	}

	elementMap := make(map[string]interface{})

	for _, ss := range clArgs {
		innerSplit := strings.Split(ss, "=")

		if len(innerSplit) != 2 {
			logging.Msg(uuid, errCtx, logging.Error, response.ErrorCodeText(response.ErrArgumentException))
			return nil, response.New(response.ErrArgumentException, errCtx)
		}

		trimKey := strings.TrimSpace(innerSplit[0])
		trimValue := strings.TrimSpace(innerSplit[1])

		elementMap[trimKey] = trimValue
	}

	return elementMap, nil
}
