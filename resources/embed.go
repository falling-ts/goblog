package resources

import "embed"

//go:embed views/**/*
var TplFS embed.FS
