/*
 * This file is part of arduino-create-agent.
 *
 * arduino-create-agent is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA  02110-1301  USA
 *
 * As a special exception, you may use this file as part of a free software
 * library without restriction.  Specifically, if other files instantiate
 * templates or use macros or inline functions from this file, or you compile
 * this file and link it with other files to produce an executable, this
 * file does not by itself cause the resulting executable to be covered by
 * the GNU General Public License.  This exception does not however
 * invalidate any other reasons why the executable file might be covered by
 * the GNU General Public License.
 *
 * Copyright 2017 ARDUINO AG (http://www.arduino.cc/)
 */
package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("tools_v1", func() {
	BasePath("/v1/tools")

	Action("list", func() {
		Description("Show the installed tools")
		Routing(GET("/"))
		Response(OK, CollectionOf(ToolV1))
	})
	Action("download", func() {
		Description("Downloads a tool in the system")
		Routing(POST("/"))
		Payload(ToolDownloadV1)
		Response(OK, ToolV1)
	})
})

var ToolV1 = MediaType("application/vnd.arduino.agent.tools.tool+json", func() {
	Attributes(func() {
		Attribute("path", String, "Path of the installed tool", func() {
			Example("~/.arduino/create/avrdude/latest")
		})
		Attribute("name", String, "Name of the installed tool", func() {
			Example("avrdude")
		})
		Attribute("packager", String, "Packager of the installed tool", func() {
			Example("packager")
		})
		Attribute("version", String, "Version of the installed tool", func() {
			Example("latest")
		})
	})

	Required("path", "name", "version", "packager")

	View("default", func() {
		Attribute("packager")
		Attribute("path")
		Attribute("name")
		Attribute("version")
	})
})

var ToolDownloadV1 = Type("tool.download", func() {
	Attribute("name", String, "Name of the tool", func() {
		Example("avrdude")
	})
	Attribute("packager", String, "Packager of the tool", func() {
		Example("arduino")
	})
	Attribute("version", String, "Version of the tool", func() {
		Example("6.0.1-arduino2")
	})
	Attribute("url", String, "Url of the tool to download", func() {
		Example("http://downloads.arduino.cc/tools/avrdude-6.0.1-arduino2-x86_64-pc-linux-gnu.tar.bz2")
	})
	Attribute("checksum", String, "The checksum of the tool to download", func() {
		Example("SHA-256:71117cce0096dad6c091e2c34eb0b9a3386d3aec7d863d2da733d9e5eac3a6b1")
	})
	Attribute("signature", String, "The signature of the url of the tool", func() {
		Example("2d701b4efbc8cec62dc299cde01730c5eebcf23d7e4393db8cf7744a9bf1d3de")
	})

	Required("packager", "name", "version", "url", "checksum", "signature")
})
