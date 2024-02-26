package cmd

import (
	"finances/handlers"
	"finances/oauth"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// Define the port flag default value
var port string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "finances",
	Short: "Finances for Large Action Model (LAM) tool",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your application. For example: Cobra is a CLI library for Go that empowers applications.`,
	// The code that will be executed when the root command runs
	Run: runCommand,
}

func runCommand(cmd *cobra.Command, args []string) {
    fmt.Printf("Starting finances server on port %s\n", port)
    mux := http.NewServeMux()

    log.Println("Starting oauth server")
    oauth.InitializeOauthServer(mux)

    log.Println("Starting finance server")
    mux.HandleFunc("GET /", financeHandler)

    if err := http.ListenAndServe("127.0.0.1:" + port, mux); err != nil {
        fmt.Println("Error starting server:", err)
    }
}

func financeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("%s %s\n", r.Method, r.Host)
    handlers.AddStockHandler(w, r)
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// Here you will define your flags and configuration settings.
	rootCmd.PersistentFlags().StringVarP(&port, "port", "p", "9002", "Port to listen on")
}

