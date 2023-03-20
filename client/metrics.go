package tower

type MetricsService struct {
	client *Client
}

const metricsAPIEndpoint = "/api/v2/metrics/"

// ListMetrics shows list of awx metrics.
func (m *MetricsService) ListMetrics() (*MetricsResponse, error) {
	result := new(MetricsResponse)
	resp, err := m.client.Requester.GetJSON(metricsAPIEndpoint, result, map[string]string{"format": "json"})
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
