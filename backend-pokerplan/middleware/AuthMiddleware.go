func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("userID")

		// Check if the session is valid
		if userID == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Session expired. Please log in again."})
			c.Abort() // Stop further processing
			return
		}

		c.Next() // Proceed to the next handler
	}
}