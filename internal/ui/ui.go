package ui

// The UI runner.
type UI struct {
	ui uiInterface
}

// Show the configuration.
func (s *UI) ShowConfig() {
	s.ui.ShowConfig()
}

// Return the UI.
func NewUI(appsGetter appsGetter) *UI {
	return &UI{
		ui: NewCmd(appsGetter.GetApps()),
	}
}
