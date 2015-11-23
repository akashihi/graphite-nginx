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
	"fmt"
	"github.com/marpaia/graphite-golang"
)

func sendMetrics(status Status, index int, useIndex bool, config Configuration) {
	var Graphite, err = graphite.NewGraphite(config.MetricsHost, config.MetricsPort)
	if err != nil {
		log.Error("Can't connect to graphite collector: %v", err)
		return
	}

	var nginxSuffix = "."
	if useIndex {
		nginxSuffix = fmt.Sprintf(".%d.", index)
	}

	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".nginx", nginxSuffix, "active"), status.Active)
	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".nginx", nginxSuffix, "accept"), status.Accept)
	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".nginx", nginxSuffix, "handle"), status.Handle)
	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".nginx", nginxSuffix, "request"), status.Request)
	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".nginx", nginxSuffix, "read"), status.Read)
	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".nginx", nginxSuffix, "write"), status.Write)
	Graphite.SimpleSend(fmt.Sprint(config.MetricsPrefix, ".nginx", nginxSuffix, "wait"), status.Wait)
}
