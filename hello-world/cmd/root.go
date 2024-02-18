package cmd

import (
	"fmt"
	"hello/handlers"
	"net/http"

	"github.com/spf13/cobra"
)

// Define the port flag default value
var port string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hello-world",
	Short: "Hello world for Large Action Model (LAM) tool",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your application. For example: Cobra is a CLI library for Go that empowers applications.`,
	// The code that will be executed when the root command runs
	Run: runCommand,
}

func runCommand(cmd *cobra.Command, args []string) {
    fmt.Printf("Starting server on port %s\n", port)
    mux := http.NewServeMux()

    mux.HandleFunc("/hello/", helloHandler)

    if err := http.ListenAndServe("127.0.0.1:" + port, mux); err != nil {
        fmt.Println("Error starting server:", err)
    }
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("%s %s\n", r.Method, r.Host)
    if r.Method == "GET" {
        handlers.HelloGetHandler(w, r)
    } else {
        handlers.HelloPostHandler(w, r)
    }
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// Here you will define your flags and configuration settings.
	rootCmd.PersistentFlags().StringVarP(&port, "port", "p", "8080", "Port to listen on")
}

