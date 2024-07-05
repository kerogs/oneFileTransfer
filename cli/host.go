package cli

import (
    "fmt"
    "log"
    "net/http"
    "path/filepath"
)

var (
    server *http.Server
    mux    = http.NewServeMux() // ServeMux global
)

func init() {
    // Register routes during initialization
    mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Default handler for root path
        http.Error(w, "Not found", http.StatusNotFound)
    }))
}

func Host(hostMode *bool, commands []string, cliPrefix *string, folderPath *string) {
    if len(commands) < 3 {
        fmt.Println("Invalid arguments")
        return
    }

    if commands[2] == "start" && commands[3] == "-d" && len(commands) >= 5 && !*hostMode {
        fmt.Println("Starting hosting... on " + commands[4])
        *cliPrefix = "host"
        *hostMode = true
        *folderPath = commands[4]

        startServer(commands[4])

    } else if commands[2] == "stop" && *hostMode {
        fmt.Println("Stopping hosting...")
        *cliPrefix = "user"
        *hostMode = false
        *folderPath = ""

        stopServer()

    } else {
        fmt.Println("Args detection error")
        return
    }
}

func startServer(folderPath string) {
    // Resolve absolute path
    absPath, err := filepath.Abs(folderPath)
    if err != nil {
        log.Fatal("Error resolving folder path:", err)
    }

    // Create file server
    fileServer := http.FileServer(http.Dir(absPath))

    // Register handler with global ServeMux
    mux.Handle("/files/", http.StripPrefix("/files/", fileServer))

    // Start HTTP server
    server = &http.Server{Addr: ":8080", Handler: mux} // Use global mux
    go func() {
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatal("Error starting HTTP server:", err)
        }
    }()
    fmt.Printf("Server started. Hosting %s on port 8080.\n", absPath)
}

func stopServer() {
    if server != nil {
        if err := server.Shutdown(nil); err != nil {
            log.Fatal("Error stopping HTTP server:", err)
        }
        fmt.Println("Server stopped.")
        
        // Reset ServeMux
        resetServeMux()

    } else {
        fmt.Println("Server is not running.")
    }
}

func resetServeMux() {
    // Create a new ServeMux to replace the global one
    mux = http.NewServeMux()
    mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Default handler for root path
        http.Error(w, "Not found", http.StatusNotFound)
    }))
    fmt.Println("ServeMux reset.")
}
