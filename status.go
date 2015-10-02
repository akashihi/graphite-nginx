/*
   conntrack-logger
   Copyright (C) 2015 Denis V Chapligin <akashihi@gmail.com>
   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.
   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"regexp"
	"strings"
)

type Status struct {
	Active  string
	Accept  string
	Handle  string
	Request string
	Read    string
	Write   string
	Wait    string
}

var ActiveRX, _ = regexp.Compile(`Active connections:\s+(\d+)\s+`)
var ConectionsRX, _ = regexp.Compile(`(\d+)\s+(\d+)\s+(\d+)`)
var SocketRX, _ = regexp.Compile(`Reading\:\s+(\d+)\s+Writing\:\s+(\d+)\s+Waiting\:\s+(\d+)\s+`)

func parse(page string) Status {
	var result = Status{}

	var statusData = strings.Split(page, "\n")
	for _, element := range statusData {
		result = parseActive(result, element)
		result = parseConnections(result, element)
		result = parseSockets(result, element)
	}

	return result
}

func parseActive(status Status, line string) Status {
	if ActiveRX.MatchString(line) {
		status.Active = ActiveRX.FindStringSubmatch(line)[1]
	}
	return status
}

func parseConnections(status Status, line string) Status {
	if ConectionsRX.MatchString(line) {
		status.Accept = ConectionsRX.FindStringSubmatch(line)[1]
		status.Handle = ConectionsRX.FindStringSubmatch(line)[2]
		status.Request = ConectionsRX.FindStringSubmatch(line)[3]
	}
	return status
}

func parseSockets(status Status, line string) Status {
	if SocketRX.MatchString(line) {
		status.Read = SocketRX.FindStringSubmatch(line)[1]
		status.Write = SocketRX.FindStringSubmatch(line)[2]
		status.Wait = SocketRX.FindStringSubmatch(line)[3]
	}
	return status
}
