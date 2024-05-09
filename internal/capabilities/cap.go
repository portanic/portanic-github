package capabilties

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Define the service interface, implementing business logic
type PluginService struct{}

// Business logic implementation for GetCapabilities
func (s *PluginService) GetCapabilities(c echo.Context) error {
	var req pb.GetCapabilitiesRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, &pb.GetCapabilitiesResponse{
			Success: false,
			Message: "Invalid request",
		})
	}

	// Logic to determine the capabilities based on plugin_id
	var capabilities []pb.Capability
	if req.PluginId == "test_plugin" {
		capabilities = []pb.Capability{pb.Capability_DATA, pb.Capability_BLOCK, pb.Capability_CREATE_ISSUE}
	} else {
		capabilities = []pb.Capability{pb.Capability_PAGE}
	}

	// Prepare the response
	resp := &pb.GetCapabilitiesResponse{
		Success: true,
		Message: "Capabilities retrieved successfully",
		Capabilities: &pb.PluginCapabilities{
			Capabilities: capabilities,
		},
	}

	// Send response
	return c.JSON(http.StatusOK, resp)
}
