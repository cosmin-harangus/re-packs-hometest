package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/cosmin-harangus/re-packs-hometest/algo"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(startCmd)
}

var calculator *algo.Calculator
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the application",
	Run: func(cmd *cobra.Command, args []string) {
		// init the calculator with the pack sizes
		packSizes := readPackSizes(LOCATION)
		calculator = algo.NewCalculator(packSizes)
		// start the server
		StartServer()
	},
}

func StartServer() {
	r := gin.Default()

	r.GET("/", getIndexHandler)
	r.GET("/api/packs", getPackSizesHandler)
	r.POST("/api/packs", setPackSizesHandler)
	r.GET("/api/order/fulfill", orderHandler)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", PORT), r); err != nil {
		log.Fatal().Err(err).Msg("Unable to start the server")
	}
}

// getIndexHandler returns the index.html file
func getIndexHandler(c *gin.Context) {
	c.File("./public/index.html")
}

// readPackSizes reads the pack sizes from the given location
// and returns them as a slice of integers
// Normally I would use a database to keep the sizes and persist them
// between server restarts but for a test I think it's not needed.
func readPackSizes(location string) []int {
	packSizesStr, err := os.ReadFile(location)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to read pack sizes")
	}
	var packSizes []int
	err = json.Unmarshal(packSizesStr, &packSizes)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to unmarshal pack sizes")
	}
	return packSizes
}

func getPackSizesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"sizes": calculator.GetPackSizes()})
}

func setPackSizesHandler(c *gin.Context) {
	var packSizes []int
	if err := c.BindJSON(&packSizes); err != nil {
		log.Error().Err(err).Msg("Unable to bind pack sizes")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to bind pack sizes"})
		return
	}
	calculator.SetPackSizes(packSizes)
	c.JSON(http.StatusOK, gin.H{"sizes": calculator.GetPackSizes()})
}

// orderHandler handles the order request and returns the shipment information
func orderHandler(c *gin.Context) {
	order := c.Query("order")
	orderInt, err := strconv.Atoi(order)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to convert order to int")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to convert order to number"})
		return
	}
	// Calculate the shipment and return the result
	shipment := calculator.Calculate(orderInt)
	c.JSON(http.StatusOK, gin.H{"shipment": shipment})
}
