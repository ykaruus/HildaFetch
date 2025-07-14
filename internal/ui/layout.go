package ui

import (
	"fmt"
	"hildafetch/internal/utils"
	"math/rand"

	"github.com/charmbracelet/lipgloss"
)

type LayoutService struct {
	config      *utils.Config
	borderTypes map[string]lipgloss.Border
}

func NewLayoutService(config *utils.Config) *LayoutService {

	return &LayoutService{
		config: config,
		borderTypes: map[string]lipgloss.Border{
			"normal":  lipgloss.NormalBorder(),
			"rounded": lipgloss.RoundedBorder(),
			"thick":   lipgloss.ThickBorder(),
			"hidden":  lipgloss.HiddenBorder(),
			"ascii":   lipgloss.ASCIIBorder(),
			"double":  lipgloss.DoubleBorder(),
			"block":   lipgloss.BlockBorder(),
		},
	}
}

func (layout *LayoutService) CreateColumText(str string) string {
	return lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(layout.config.ColorPrimary)).Render(str)
}

func (layout *LayoutService) FormatString(str string, separator string, value string) string {
	foo := layout.CreateColumText(str)
	return fmt.Sprintf("%v %s %v", foo, separator, value)
}

func (layout *LayoutService) CreateTable(str string, os_name string) string {

	layoutBorder := layout.borderTypes[layout.config.BorderType] // pega borda customizada do config.json

	if _, ok := layout.borderTypes[layout.config.BorderType]; !ok {
		fmt.Printf("the border style '%s' is not valid, seting rounded style.\n", layout.config.BorderType)
		layoutBorder = lipgloss.RoundedBorder()
	}

	frases := ""

	if len(layout.config.Frases) > 0 {
		frasesRand := rand.Intn(len(layout.config.Frases))
		frases = layout.config.Frases[frasesRand]

	}

	system_title := lipgloss.NewStyle().Bold(true).Align(lipgloss.Center).Foreground(lipgloss.Color(layout.config.UsernameColor)).Render(os_name)
	system_info := lipgloss.NewStyle().Width(50).Border(layoutBorder).BorderForeground(lipgloss.Color(layout.config.BorderColor)).Foreground(lipgloss.Color(layout.config.ColorSecundary)).Padding(1).Render(system_title + "\n" + str)
	system_leng := lipgloss.NewStyle().Align(lipgloss.Left).Italic(true).Render(frases)
	table := lipgloss.JoinVertical(lipgloss.Right, system_info, system_leng)

	row := lipgloss.JoinHorizontal(lipgloss.Center, table)

	return row

}

func (layout *LayoutService) CreateAscii(str string) string {
	ascii := lipgloss.NewStyle().Height(30).Width(45).Align(lipgloss.Center).MarginTop(1).Foreground(lipgloss.Color(layout.config.ColorPrimary)).Render(str)

	return ascii
}

func (layout *LayoutService) CreateFetch(ascii string, table string) string {
	row := lipgloss.JoinHorizontal(lipgloss.Center, ascii, table)

	return row
}
