package demands

type IssuesByPeriodDTO struct {
	JQL          *string `json:"jql,omitempty"`
	Total        int     `json:"total"`
	PeriodValues []int   `json:"values"`
}

type IssuesAnalytics struct {
	OverallProgress   float64   `json:"overallProgress"`
	ProgressPerPeriod []float64 `json:"progressPerPeriod"`
	CreatedTotal      int       `json:"createdTotal"`
	ResolvedTotal     int       `json:"resolvedTotal"`
	// PendingTotal possui informação atemporal.
	PendingTotal int `json:"pendingTotal"`
}

type IssuesDetailsByProject struct {
	Total                int              `json:"total"`
	IssuesTypes          []string         `json:"issuesTypes"`
	TotalByType          []int            `json:"totalByType"`
	TotalByPeriod        []int            `json:"totalByPeriod"`
	TotalByTypeAndPeriod []map[string]int `json:"totalByTypeAndPeriod"`
}

type GetIssuesByPeriodProject struct {
	Projects               []string                 `json:"projects"`
	ProjectsAvatars        []string                 `json:"projectsAvatars"`
	IssuesDetailsByProject []IssuesDetailsByProject `json:"issuesDetailsByProject"`
}

type GetIssuesByPeriodResponse struct {
	Created  IssuesByPeriodDTO `json:"created"`
	Resolved IssuesByPeriodDTO `json:"resolved"`
	// Pending possui informação temporal.
	Pending   IssuesByPeriodDTO        `json:"pending"`
	Periods   []string                 `json:"periods"`
	Analytics IssuesAnalytics          `json:"analytics"`
	Project   GetIssuesByPeriodProject `json:"project"`
}

func (s *GetIssuesByPeriodResponse) DoAnalysis() {
	s.Analytics.ProgressPerPeriod = make([]float64, len(s.Periods))

	for i := range s.Periods {
		s.Created.Total += s.Created.PeriodValues[i]
		s.Resolved.Total += s.Resolved.PeriodValues[i]
		s.Pending.Total += s.Pending.PeriodValues[i]

		if s.Resolved.PeriodValues[i] > 0 && s.Created.PeriodValues[i] > 0 {
			s.Analytics.ProgressPerPeriod[i] += float64(s.Resolved.PeriodValues[i]) / float64(s.Created.PeriodValues[i]) * 100
		}
	}

	s.Analytics.CreatedTotal = s.Created.Total
	s.Analytics.ResolvedTotal = s.Resolved.Total
	s.Analytics.PendingTotal = s.Analytics.CreatedTotal - s.Analytics.ResolvedTotal

	if s.Resolved.Total > 0 && s.Created.Total > 0 {
		s.Analytics.OverallProgress = float64(s.Resolved.Total) / float64(s.Created.Total) * 100
	}
}

func (s *GetIssuesByPeriodResponse) FixEmpty() {
	for k, v := range s.Project.IssuesDetailsByProject {
		if v.Total == 0 {
			defaultLength := len(s.Periods)

			s.Project.IssuesDetailsByProject[k].IssuesTypes = make([]string, defaultLength)
			s.Project.IssuesDetailsByProject[k].TotalByType = make([]int, defaultLength)
			s.Project.IssuesDetailsByProject[k].TotalByPeriod = make([]int, defaultLength)
			s.Project.IssuesDetailsByProject[k].TotalByTypeAndPeriod = make([]map[string]int, defaultLength)
		}
	}
}
