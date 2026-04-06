package tui

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"toofan/internal/theme"
)

type profileData struct {
	Tests     int
	Time      time.Duration
	Best      map[string]map[int]float64 // mode -> dur -> wpm
	Recent    []testEntry
	Activity  map[string]int
	RecentAvg float64
}

type testEntry struct {
	Date time.Time
	WPM  float64
	Dur  int
	Acc  float64
	Mode string
}

func loadProfile() profileData {
	pd := profileData{
		Best:     make(map[string]map[int]float64),
		Activity: make(map[string]int),
	}
	pd.Best["words"] = make(map[int]float64)
	pd.Best["code"] = make(map[int]float64)

	home, err := os.UserHomeDir()
	if err != nil {
		return pd
	}
	dataDir := filepath.Join(home, ".toofan")

	f, err := os.Open(filepath.Join(dataDir, "results.txt"))
	if err != nil {
		return pd
	}
	defer f.Close()

	var all []testEntry
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		e, ok := parseResultLine(sc.Text())
		if !ok {
			continue
		}
		all = append(all, e)
		pd.Tests++
		pd.Time += time.Duration(e.Dur) * time.Second

		if pd.Best[e.Mode] == nil {
			pd.Best[e.Mode] = make(map[int]float64)
		}
		if e.WPM > pd.Best[e.Mode][e.Dur] {
			pd.Best[e.Mode][e.Dur] = e.WPM
		}
		pd.Activity[e.Date.Format("2006-01-02")]++
	}

	if len(all) > 80 {
		pd.Recent = all[len(all)-80:]
	} else {
		pd.Recent = all
	}

	if len(all) >= 10 {
		sum := 0.0
		for i := len(all) - 10; i < len(all); i++ {
			sum += all[i].WPM
		}
		pd.RecentAvg = sum / 10.0
	} else if len(all) > 0 {
		sum := 0.0
		for _, e := range all {
			sum += e.WPM
		}
		pd.RecentAvg = sum / float64(len(all))
	}

	return pd
}

func parseResultLine(line string) (testEntry, bool) {
	parts := strings.Split(line, "|")
	if len(parts) < 5 {
		return testEntry{}, false
	}

	date, err := time.Parse("2006-01-02 15:04", strings.TrimSpace(parts[0]))
	if err != nil {
		return testEntry{}, false
	}

	wpmStr := strings.TrimSpace(parts[1])
	wpmStr = strings.TrimSuffix(wpmStr, "wpm")
	wpmStr = strings.TrimSpace(wpmStr)
	wpm, _ := strconv.ParseFloat(wpmStr, 64)

	accStr := strings.TrimSpace(parts[2])
	accStr = strings.TrimSuffix(accStr, "%")
	accStr = strings.TrimSpace(accStr)
	acc, _ := strconv.ParseFloat(accStr, 64)

	durStr := strings.TrimSpace(parts[3])
	durStr = strings.TrimSuffix(durStr, "s")
	durStr = strings.TrimSpace(durStr)
	dur, _ := strconv.Atoi(durStr)

	modeStr := strings.TrimSpace(parts[4])

	return testEntry{Date: date, WPM: wpm, Dur: dur, Acc: acc, Mode: modeStr}, true
}

func (m model) handleProfile(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	m.active = screenTyping
	return m, nil
}

func cleanLegacyLang(lang string) string {
	for _, suffix := range []string{"15s", "30s", "60s", "120s"} {
		lang = strings.TrimSuffix(lang, suffix)
	}
	switch lang {
	case "javascript":
		return "js"
	case "typescript":
		return "ts"
	}
	if len(lang) > 9 {
		return lang[:9]
	}
	return lang
}

func (m model) viewProfile(p theme.Palette) string {
	dim := lipgloss.NewStyle().Foreground(p.Foreground)
	val := lipgloss.NewStyle().Foreground(p.Typed).Bold(true)
	hi := lipgloss.NewStyle().Foreground(p.Accent)

	title := val.Render("_hello friend")

	gridWidth := 74
	if m.width > 0 && m.width < 80 {
		gridWidth = m.width - 6
	}
	paneWidth := (gridWidth - 2) / 2

	paneStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(p.Foreground).
		Padding(1, 2)

	hours := int(m.prof.Time.Hours())
	mins := int(m.prof.Time.Minutes()) % 60
	timeStr := fmt.Sprintf("%dm", mins)
	if hours > 0 {
		timeStr = fmt.Sprintf("%dh %dm", hours, mins)
	}

	avgStr := dim.Render("-")
	if m.prof.Tests > 0 {
		avgStr = val.Render(fmt.Sprintf("%.0f wpm", m.prof.RecentAvg))
	}

	overviewInner := lipgloss.JoinVertical(lipgloss.Left,
		hi.Render("overview"),
		"",
		dim.Render("tests       ")+val.Render(fmt.Sprintf("%d", m.prof.Tests)),
		dim.Render("time        ")+val.Render(timeStr),
		dim.Render("avg speed   ")+avgStr,
	)

	cw := 5
	lw := 6

	durLabels := lipgloss.JoinHorizontal(lipgloss.Left,
		col(lw, ""),
		col(cw, dim.Render("15s")),
		col(cw, dim.Render("30s")),
		col(cw, dim.Render("60s")),
		col(cw, dim.Render("120s")),
	)

	buildBestsLine := func(label string, data map[int]float64) string {
		cells := []string{col(lw, hi.Render(label))}
		for _, d := range []int{15, 30, 60, 120} {
			if wpm, ok := data[d]; ok {
				cells = append(cells, col(cw, val.Render(fmt.Sprintf("%.0f", wpm))))
			} else {
				cells = append(cells, col(cw, dim.Render("-")))
			}
		}
		return lipgloss.JoinHorizontal(lipgloss.Left, cells...)
	}

	wordsLine := buildBestsLine("words", m.prof.Best["words"])
	codeLine := buildBestsLine("code", m.prof.Best["code"])

	bestInner := lipgloss.JoinVertical(lipgloss.Left,
		hi.Render("personal bests"),
		"",
		durLabels,
		wordsLine,
		codeLine,
	)

	h1 := lipgloss.Height(overviewInner)
	h2 := lipgloss.Height(bestInner)
	if h1 < h2 {
		overviewInner += strings.Repeat("\n", h2-h1)
	} else if h2 < h1 {
		bestInner += strings.Repeat("\n", h1-h2)
	}

	overviewBox := paneStyle.Width(paneWidth).Render(overviewInner)
	bestBox := paneStyle.Width(paneWidth).Render(bestInner)
	topRow := lipgloss.JoinHorizontal(lipgloss.Top, overviewBox, "  ", bestBox)

	var histRows []string
	header := lipgloss.JoinHorizontal(lipgloss.Left,
		col(7, hi.Render("wpm")),
		col(7, hi.Render("acc")),
		col(7, hi.Render("type")),
		col(10, hi.Render("lang")),
		col(6, hi.Render("time")),
		col(13, hi.Render("date")),
	)
	histRows = append(histRows, header, "")

	limit := 10
	if len(m.prof.Recent) < limit {
		limit = len(m.prof.Recent)
	}

	for i := len(m.prof.Recent) - 1; i >= len(m.prof.Recent)-limit; i-- {
		e := m.prof.Recent[i]
		dstr := e.Date.Format("02 Jan 15:04")

		modeType := "words"
		modeLang := "english"
		if strings.HasPrefix(e.Mode, "code:") {
			modeType = "code"
			modeLang = cleanLegacyLang(strings.TrimPrefix(e.Mode, "code:"))
		}

		durStr := "∞"
		if e.Dur > 0 {
			durStr = fmt.Sprintf("%ds", e.Dur)
		}

		row := lipgloss.JoinHorizontal(lipgloss.Left,
			col(7, val.Render(fmt.Sprintf("%.0f", e.WPM))),
			col(7, dim.Render(fmt.Sprintf("%.0f%%", e.Acc))),
			col(7, dim.Render(modeType)),
			col(10, dim.Render(modeLang)),
			col(6, dim.Render(durStr)),
			col(13, dim.Render(dstr)),
		)
		histRows = append(histRows, row)
	}

	fullWidth := paneWidth*2 + 2
	histBox := paneStyle.Width(fullWidth).Render(
		lipgloss.JoinVertical(lipgloss.Left,
			hi.Render("recent tests"),
			"",
			lipgloss.JoinVertical(lipgloss.Left, histRows...),
		),
	)

	heatmapStr := heatGrid(m.prof.Activity, p, fullWidth)
	heatBox := paneStyle.Width(fullWidth).Render(
		lipgloss.JoinVertical(lipgloss.Left,
			hi.Render("activity map"),
			"",
			heatmapStr,
		),
	)

	body := lipgloss.JoinVertical(lipgloss.Left,
		title,
		"",
		topRow,
		"",
		histBox,
		"",
		heatBox,
	)

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, body)
}

func heatGrid(activity map[string]int, p theme.Palette, width int) string {
	now := time.Now()

	peak := 1
	for _, n := range activity {
		if n > peak {
			peak = n
		}
	}

	labels := []string{"mon", "tue", "wed", "thu", "fri", "sat", "sun"}
	dows := []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday, time.Sunday}

	c0 := lipgloss.NewStyle().Foreground(lipgloss.Color("#333333"))
	c1 := lipgloss.NewStyle().Foreground(p.Foreground)
	c2 := lipgloss.NewStyle().Foreground(p.Typed)
	c3 := lipgloss.NewStyle().Foreground(p.Accent)
	c4 := lipgloss.NewStyle().Foreground(p.Success).Bold(true)
	colors := []lipgloss.Style{c0, c1, c2, c3, c4}

	dim := lipgloss.NewStyle().Foreground(p.Foreground)

	innerWidth := width - 4
	weeks := (innerWidth - 5) / 2
	if weeks < 1 {
		weeks = 1
	}

	var rows []string
	for i, dow := range dows {
		var row strings.Builder
		row.WriteString(dim.Render(fmt.Sprintf("%3s  ", labels[i])))

		for w := weeks - 1; w >= 0; w-- {
			d := now.AddDate(0, 0, -w*7)
			for d.Weekday() != dow {
				d = d.AddDate(0, 0, -1)
			}

			n := activity[d.Format("2006-01-02")]
			idx := 0
			if n > 0 {
				idx = int(math.Ceil(float64(n) / float64(peak) * 4))
				if idx > 4 {
					idx = 4
				}
			}
			row.WriteString(colors[idx].Render("■") + " ")
		}
		rows = append(rows, row.String())
	}

	var legend strings.Builder
	legend.WriteString(dim.Render("Less "))
	legend.WriteString(colors[0].Render("■") + " ")
	legend.WriteString(colors[1].Render("■") + " ")
	legend.WriteString(colors[2].Render("■") + " ")
	legend.WriteString(colors[3].Render("■") + " ")
	legend.WriteString(colors[4].Render("■") + " ")
	legend.WriteString(dim.Render("More"))

	return strings.Join(rows, "\n") + "\n\n" + legend.String()
}
