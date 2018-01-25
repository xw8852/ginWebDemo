package v1

import(  "github.com/gin-gonic/gin" )

func setupRounter() *gin.Engine {
	r := gin.Default()

	return r
}
func main() {
	r := setupRounter()
	r.Run("miapp/api/v1/")
}