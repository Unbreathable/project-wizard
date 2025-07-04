package service

import (
	"github.com/Liphium/project-wizard/neogate"
	"github.com/go-playground/validator/v10"
)

var Instance *neogate.Instance
var Validate *validator.Validate = validator.New()
