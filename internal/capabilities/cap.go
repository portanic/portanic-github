package capabilties

import (
	"net/http"

	"github.com/labstack/echo/v4"
	cap_v1 "github.com/portanic/api/capabilities/v1"
)

// Define the service interface, implementing business logic
type PluginService struct{}

// Business logic implementation for GetCapabilities
func (s *PluginService) GetCapabilities(c echo.Context) error {
	var req cap_v1.GetCapabilitiesRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, &cap_v1.GetCapabilitiesResponse{
			Success: false,
			Message: "Invalid request",
		})
	}

	// Logic to determine the capabilities based on plugin_id
	var capabilities []cap_v1.Capability
	if req.PluginId == "test_plugin" {
		capabilities = []cap_v1.Capability{cap_v1.Capability_DATA, cap_v1.Capability_BLOCK, cap_v1.Capability_CREATE_ISSUE}
	} else {
		capabilities = []cap_v1.Capability{cap_v1.Capability_PAGE}
	}

	// Prepare the response
	resp := &cap_v1.GetCapabilitiesResponse{
		Success: true,
		Message: "Capabilities retrieved successfully",
		Capabilities: &cap_v1.PluginCapabilities{
			Capabilities: capabilities,
		},
	}

	// Send response
	return c.JSON(http.StatusOK, resp)
}
