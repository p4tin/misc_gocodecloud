package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "log"
    "time"
)

// simulate some private data
var secrets = gin.H{
    "foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
    "austin": gin.H{"email": "austin@example.com", "phone": "666"},
    "lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func main() {
    r := gin.Default()

    // Group using gin.BasicAuth() middleware
    // gin.Accounts is a shortcut for map[string]string
    authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
        "foo":    "bar",
        "austin": "1234",
        "lena":   "hello2",
        "manu":   "4321",
    }))

    // /admin/secrets endpoint
    // hit "localhost:8080/admin/secrets
    authorized.GET("/secrets", func(c *gin.Context) {
        // get user, it was setted by the BasicAuth middleware
        user := c.MustGet(gin.AuthUserKey).(string)
        if secret, ok := secrets[user]; ok {
            c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
        } else {
            c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
        }
    })

    r.GET("/long_async", func(c *gin.Context) {
        // create copy to be used inside the goroutine
        c_cp := c.Copy()
        go func() {
            // simulate a long task with time.Sleep(). 5 seconds
            time.Sleep(5 * time.Second)
            c_cp.JSON(http.StatusOK, "{\"long_asych\": \"completed\"}")
            // note than you are using the copied context "c_cp", IMPORTANT
            log.Println("Done! in path " + c_cp.Request.URL.Path)
        }()
    })

    // Listen and server on 0.0.0.0:8080
    //http.ListenAndServe(":8080", r)
    //Start the webserver
    //log.Println("Starting Web Server on " + *addr)
    err := http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", r)
    if err != nil {
        log.Printf("Web Server cannot start: %v\n", err)
    }
}
