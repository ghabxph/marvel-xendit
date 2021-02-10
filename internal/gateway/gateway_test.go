package gateway

import (
	"os"
	"io/ioutil"
	"net/http/httptest"
	"testing"
	"github.com/ghabxph/marvel-xendit/internal/live"
	"github.com/ghabxph/marvel-xendit/internal/memorydb"
	"github.com/ghabxph/marvel-xendit/internal/testutils"
	"github.com/gofiber/fiber/v2/utils"
)

func TestGatewayEndpoints(t *testing.T) {

	// Set config path
	os.Setenv(live.CONFIG_PATH_KEY, "../../config.yaml")

	// MemoryDB Instance
	db := memorydb.GetInstance()

	// Populate dataset in memorydb
	testutils.PrepareDataset(db)

	// Create gateway instance
	gateway := GetInstance(db)

	// Create fiber instance
	fiber := gateway.Fiber()

	t.Run("Get all characters through /characters", func(t *testing.T) {
		resp, _ := fiber.Test(httptest.NewRequest("GET", "/characters", nil))
		body, _ := ioutil.ReadAll(resp.Body)
		utils.AssertEqual(t, 200, resp.StatusCode)
		utils.AssertEqual(t, testutils.GetTestCharacters(), string(body))
	})

	t.Run("Get a character through /characters/:id", func(t *testing.T) {
		resp, _ := fiber.Test(httptest.NewRequest("GET", "/characters/1009146", nil))
		body, _ := ioutil.ReadAll(resp.Body)
		utils.AssertEqual(t, 200, resp.StatusCode)
		utils.AssertEqual(t, testutils.GetTestCharacter(), string(body))
	})
}
