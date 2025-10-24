package metadata

import "errors"

var ErrFileNotFound = errors.New("file not found")
var ErrInvalidOperation = errors.New("invalid operation")
var ErrDuplicateName = errors.New("name already used")
var ErrFolderNotEmpty = errors.New("cannot perform operation, folder not empty")
var ErrNotFolder = errors.New("not a folder")
var ErrInvalidPath = errors.New("inavlid path format")