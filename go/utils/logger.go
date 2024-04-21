/*
 * Smidgen
 *
 * API for interacting with Smidgen.
 *
 *   Smidgen aims to simplify and automate common tasks that logisticians
 *   conduct on a daily basis so they can focus on the effective distribution
 *   of materiel, as well as maintain an accurate record keeping book of
 *   receiving, issuance, audits, surpluses, amongst other logistical tasks.
 *   Copyright (C) 2024  Jose Hernandez
 *
 *   This program is free software: you can redistribute it and/or modify
 *   it under the terms of the GNU General Public License as published by
 *   the Free Software Foundation, either version 3 of the License, or
 *   (at your option) any later version.
 *
 *   This program is distributed in the hope that it will be useful,
 *   but WITHOUT ANY WARRANTY; without even the implied warranty of
 *   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *   GNU General Public License for more details.
 *
 *   You should have received a copy of the GNU General Public License
 *   along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package smidgen

import (
	"net/http"
	"os"
	"time"

	"github.com/charmbracelet/lipgloss"
	logger "github.com/charmbracelet/log"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		Log().Info(
			"[%s %s] %s TTE: %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

func Log() *logger.Logger {
	styles := logger.DefaultStyles()
	styles.Levels[logger.ErrorLevel] = lipgloss.NewStyle().
		SetString("ERROR").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("#ff0000")).
		Foreground(lipgloss.Color("0"))
	styles.Levels[logger.FatalLevel] = lipgloss.NewStyle().
		SetString("FATAL").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("#ff0000")).
		Foreground(lipgloss.Color("0"))
	styles.Levels[logger.WarnLevel] = lipgloss.NewStyle().
		SetString("WARN").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("#ffff00")).
		Foreground(lipgloss.Color("0"))
	styles.Levels[logger.InfoLevel] = lipgloss.NewStyle().
		SetString("INFO").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("#00ff00")).
		Foreground(lipgloss.Color("0"))

	logger := logger.NewWithOptions(os.Stderr, logger.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
	})

	logger.SetStyles(styles)
	return logger
}
