/**
 * Copyright (c) 2011 ~ 2015 Deepin, Inc.
 *               2013 ~ 2015 jouyouyun
 *
 * Author:      jouyouyun <jouyouwen717@gmail.com>
 * Maintainer:  jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, see <http://www.gnu.org/licenses/>.
 **/

package timedate

import (
	"pkg.linuxdeepin.com/dde-daemon/timedate/zoneinfo"
	"time"
)

/**
 * SetDate Set the system clock to the specified.
 * The time may be specified in the format '2015' '1' '1' '18' '18' '18' '8'.
 **/
func (m *Manager) SetDate(year, month, day, hour, min, sec, nsec int32) error {
	loc, err := time.LoadLocation(m.Timezone)
	if err != nil {
		logger.Debugf("Load location '%s' failed: %v", m.Timezone, err)
		return err
	}
	ns := time.Date(int(year), time.Month(month), int(day),
		int(hour), int(min), int(sec), int(nsec), loc).UnixNano()
	return m.SetTime(ns/1000, false)
}

/**
 * SetTime Set the system clock to the specified.
 *
 * usec: pass a value of microseconds since 1 Jan 1970 UTC.
 * relative: if true, the passed usec value will be added to the current system time; if false, the current system time will be set to the passed usec value.
 **/
func (m *Manager) SetTime(usec int64, relative bool) error {
	err := m.td1.SetTime(usec, relative, true)
	if err != nil {
		logger.Debug("SetTime failed:", err)
	}

	return err
}

/**
 * SetNTP To control whether the system clock is synchronized with the network.
 **/
func (m *Manager) SetNTP(useNTP bool) error {
	err := m.td1.SetNTP(useNTP, true)
	if err != nil {
		logger.Debug("SetNTP failed:", err)
	}

	return err
}

/**
 * SetLocalRTC To control whether the RTC is the local time or UTC.
 * Time standards are divided into: localtime and UTC.
 * UTC standard will automatically adjust the daylight saving time.
 *
 * localRTC: whether to use local time.
 * fixSystem: if true, will use the RTC time to adjust the system clock; if false, the system time is written to the RTC taking the new setting into account.
 **/
func (m *Manager) SetLocalRTC(localRTC, fixSystem bool) error {
	err := m.td1.SetLocalRTC(localRTC, fixSystem, true)
	if err != nil {
		logger.Debug("SetLocalRTC failed:", err)
	}

	return err
}

/**
 * SetTimezone Set the system time zone to the specified value.
 * Valid timezones you may parse from /usr/share/zoneinfo/zone.tab.
 *
 * zone: pass a value like "Asia/Shanghai" to set the timezone.
 **/
func (m *Manager) SetTimezone(zone string) error {
	err := m.td1.SetTimezone(zone, true)
	if err != nil {
		logger.Debug("SetTimezone failed:", err)
		return err
	}

	return m.AddUserTimezone(zone)
}

/**
 * AddUserTimezone Add the specified time zone to user time zone list.
 **/
func (m *Manager) AddUserTimezone(zone string) error {
	if !zoneinfo.IsZoneValid(zone) {
		logger.Debug("Invalid zone:", zone)
		return zoneinfo.ErrZoneInvalid
	}

	oldList := m.UserTimezones.Get()
	newList, added := addItemToList(zone, oldList)
	if added {
		m.settings.SetStrv(settingsKeyTimezoneList, newList)
	}
	return nil
}

/**
 * DeleteUserTimezone Delete the specified time zone from user time zone list.
 **/
func (m *Manager) DeleteUserTimezone(zone string) error {
	if !zoneinfo.IsZoneValid(zone) {
		logger.Debug("Invalid zone:", zone)
		return zoneinfo.ErrZoneInvalid
	}

	oldList := m.UserTimezones.Get()
	newList, deleted := deleteItemFromList(zone, oldList)
	if deleted {
		m.settings.SetStrv(settingsKeyTimezoneList, newList)
	}
	return nil
}

/**
 * GetZoneInfo Get ZoneInfo of the specified time zone.
 **/
func (m *Manager) GetZoneInfo(zone string) (zoneinfo.ZoneInfo, error) {
	info, err := zoneinfo.GetZoneInfo(zone)
	if err != nil {
		logger.Debugf("Get zone info for '%s' failed: %v", zone, err)
		return zoneinfo.ZoneInfo{}, err
	}
	info.Desc = getZoneDesc(zone)

	return *info, nil
}

/**
 * GetZoneInfo Get all ZoneInfo in the specified list.
 **/
func (m *Manager) GetZoneList() []string {
	var list []string
	for _, zdesc := range zoneWhiteList {
		if !zoneinfo.IsZoneValid(zdesc.zone) {
			continue
		}
		list = append(list, zdesc.zone)
	}

	return list
}