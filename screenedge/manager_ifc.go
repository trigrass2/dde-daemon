/*
 * Copyright (C) 2014 ~ 2018 Deepin Technology Co., Ltd.
 *
 * Author:     jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package screenedge

import (
	"dbus/com/deepin/wm"
	ddbus "pkg.deepin.io/dde/daemon/dbus"
)

// Enable desktop edge zone detected
//
// 是否启用桌面边缘热区功能
func (m *Manager) EnableZoneDetected(enable bool) {
	if !ddbus.IsSessionBusActivated("com.deepin.wm") {
		logger.Warning("Deepin window manager not running")
		return
	}

	obj, err := wm.NewWm("com.deepin.wm", "/com/deepin/wm")
	if err != nil {
		logger.Warning("[EnableZoneDetected] Failed to connect wm dbus:", err)
		return
	}

	err = obj.EnableZoneDetected(enable)
	wm.DestroyWm(obj)
	if err != nil {
		logger.Warning("[EnableZoneDetected] failed to toggle zone:", err)
	}
}

// Set left-top edge action
func (m *Manager) SetTopLeft(value string) {
	m.settings.SetEdgeAction(TopLeft, value)
}

// Get left-top edge action
func (m *Manager) TopLeftAction() string {
	return m.settings.GetEdgeAction(TopLeft)
}

// Set left-bottom edge action
func (m *Manager) SetBottomLeft(value string) {
	m.settings.SetEdgeAction(BottomLeft, value)
}

// Get left-bottom edge action
func (m *Manager) BottomLeftAction() string {
	return m.settings.GetEdgeAction(BottomLeft)
}

// Set right-top edge action
func (m *Manager) SetTopRight(value string) {
	m.settings.SetEdgeAction(TopRight, value)
}

// Get right-top edge action
func (m *Manager) TopRightAction() string {
	return m.settings.GetEdgeAction(TopRight)
}

// Set right-bottom edge action
func (m *Manager) SetBottomRight(value string) {
	m.settings.SetEdgeAction(BottomRight, value)
}

// Get right-bottom edge action
func (m *Manager) BottomRightAction() string {
	return m.settings.GetEdgeAction(BottomRight)
}
