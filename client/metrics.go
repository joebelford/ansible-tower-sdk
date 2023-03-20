package tower

type MetricsService struct {
	client *Client
}

type MetricsResponse struct {
	AwxSystemInfo                Metric `json:"awx_system_info"`
	AwxOrganizationsTotal        Metric `json:"awx_organizations_total"`
	AwxUsersTotal                Metric `json:"awx_users_total"`
	AwxTeamsTotal                Metric `json:"awx_teams_total"`
	AwxInventoriesTotal          Metric `json:"awx_inventories_total"`
	AwxProjectsTotal             Metric `json:"awx_projects_total"`
	AwxJobTemplatesTotal         Metric `json:"awx_job_templates_total"`
	AwxWorkflowJobTemplatesTotal Metric `json:"awx_workflow_job_templates_total"`
	AwxHostsTotal                Metric `json:"awx_hosts_total"`
	AwxSchedulesTotal            Metric `json:"awx_schedules_total"`
	AwxInventoryScriptsTotal     Metric `json:"awx_inventory_scripts_total"`
	AwxSessionsTotal             Metric `json:"awx_sessions_total"`
	AwxCustomVirtualenvsTotal    Metric `json:"awx_custom_virtualenvs_total"`
	AwxRunningJobsTotal          Metric `json:"awx_running_jobs_total"`
	AwxPendingJobsTotal          Metric `json:"awx_pending_jobs_total"`
	AwxStatusTotal               Metric `json:"awx_status_total"`
	AwxInstanceCapacity          Metric `json:"awx_instance_capacity"`
	AwxInstanceCPU               Metric `json:"awx_instance_cpu"`
	AwxInstanceMemory            Metric `json:"awx_instance_memory"`
	AwxInstanceInfo              Metric `json:"awx_instance_info"`
	AwxInstanceLaunchTypeTotal   Metric `json:"awx_instance_launch_type_total"`
	AwxInstanceStatusTotal       Metric `json:"awx_instance_status_total"`
	AwxInstanceConsumedCapacity  Metric `json:"awx_instance_consumed_capacity"`
	AwxInstanceRemainingCapacity Metric `json:"awx_instance_remaining_capacity"`
	AwxLicenseInstanceTotal      Metric `json:"awx_license_instance_total"`
	AwxLicenseInstanceFree       Metric `json:"awx_license_instance_free"`
}

const metricsAPIEndpoint = "/api/v2/metrics/?format=json"

// ListMetrics shows list of awx metrics.
func (p *MetricsService) ListMetrics() (*MetricsResponse, error) {
	result := new(MetricsResponse)
	resp, err := p.client.Requester.GetJSON(metricsAPIEndpoint, result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
