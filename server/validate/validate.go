package validate

import "github.com/gin-gonic/gin/binding"

func Init () {
	binding.Validator = new(defaultValidator)
}
