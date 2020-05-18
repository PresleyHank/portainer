package edgetemplates

import (
	"encoding/json"
	"net/http"

	httperror "github.com/portainer/libhttp/error"
	"github.com/portainer/libhttp/response"
	"github.com/portainer/portainer/api"
	"github.com/portainer/portainer/api/http/client"
)

type templateFileFormat struct {
	Version   string               `json:"version"`
	Templates []portainer.Template `json:"templates"`
}

// GET request on /api/edgetemplates
func (handler *Handler) edgeTemplateList(w http.ResponseWriter, r *http.Request) *httperror.HandlerError {
	settings, err := handler.SettingsService.Settings()
	if err != nil {
		return &httperror.HandlerError{http.StatusInternalServerError, "Unable to retrieve settings from the database", err}
	}

	url := portainer.DefaultTemplatesURL
	if settings.TemplatesURL != "" {
		url = settings.TemplatesURL
	}

	var templateData []byte
	templateData, err = client.Get(url, 10)
	if err != nil {
		return &httperror.HandlerError{http.StatusInternalServerError, "Unable to retrieve external templates", err}
	}

	var templateFile templateFileFormat

	err = json.Unmarshal(templateData, &templateFile)
	if err != nil {
		return &httperror.HandlerError{http.StatusInternalServerError, "Unable to parse template file", err}
	}

	filteredTemplates := make([]portainer.Template, 0)

	for _, template := range templateFile.Templates {
		if template.Type == portainer.EdgeStackTemplate {
			filteredTemplates = append(filteredTemplates, template)
		}
	}

	return response.JSON(w, filteredTemplates)
}
