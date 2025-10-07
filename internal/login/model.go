package login

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/zalando/go-keyring"
)

type TokenMsg string

type Model struct {
	form *huh.Form
}

func NewModel() Model {
	return Model{
		form: huh.NewForm(huh.NewGroup(huh.NewInput().
			Key("token").
			Placeholder("Token").
			EchoMode(huh.EchoModePassword).
			Validate(huh.ValidateNotEmpty()),
		)),
	}
}

func (m Model) Init() tea.Cmd {
	return m.form.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.form.State {
	case huh.StateCompleted:
		token := m.form.GetString("token")
		go keyring.Set(keyringName, "token", token)
		return m, func() tea.Msg {
			return TokenMsg(token)
		}
	}

	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
	}

	return m, cmd
}

func (m Model) View() string {
	return m.form.View()
}
