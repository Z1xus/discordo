package root

import (
	"github.com/ayn2op/discordo/internal/home"
	"github.com/ayn2op/discordo/internal/login"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/zalando/go-keyring"
)

const (
	keyringService = "discordo"
	keyringUser    = "token"
)

type loginMsg struct{}

type Model struct {
	root tea.Model
}

func NewModel() Model {
	return Model{}
}

func (m Model) checkToken() tea.Msg {
	token, err := keyring.Get(keyringService, keyringUser)
	if err != nil {
		return loginMsg{}
	}

	return login.TokenMsg(token)
}

func (m Model) Init() tea.Cmd {
	return m.checkToken
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEscape:
			return m, tea.Quit
		}

	case loginMsg:
		m.root = login.NewModel()
		return m, m.root.Init()
	case login.TokenMsg:
		m.root = home.NewModel(string(msg))
		return m, m.root.Init()
	}

	var cmd tea.Cmd
	if m.root != nil {
		m.root, cmd = m.root.Update(msg)
	}

	return m, cmd
}

func (m Model) View() string {
	if m.root == nil {
		return "Loading"
	} else {
		return m.root.View()
	}
}
