package redirect

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Redirect(c *gin.Context, targetURL string) {
	// Log the redirection attempt
	log.Printf("Redirecting to: %s", targetURL)

	// Check if the request is coming from HTMX
	if c.GetHeader("HX-Request") != "" {
		// Respond with an HTMX-specific header to trigger a client-side redirect
		c.Header("HX-Redirect", targetURL)
		c.Status(http.StatusOK)
	} else {
		// Regular redirection for non-HTMX requests
		c.Redirect(http.StatusFound, targetURL)
	}
}
