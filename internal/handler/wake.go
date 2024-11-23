package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mstcl/wolly/internal/packet"
)

func Wake(e echo.Context) error {
	var err error

	if err := e.Request().ParseForm(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("parse form: %v", err))
	}

	form := e.Request().Form

	var isUi bool

	isUi, err = strconv.ParseBool(form.Get("ui"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("parse bool: %v", err))
	}

	if !isUi {
		if err := handlePacket(form.Get("mac")); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
	}

	keys := []string{"a", "b", "c", "d", "e", "f"}

	var addr strings.Builder
	for idx, v := range keys {
		addr.WriteString(form.Get(v))

		if idx != len(keys)-1 {
			addr.WriteString(":")
		}
	}

	if err := handlePacket(addr.String()); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return nil
}

func handlePacket(m string) error {
	p, err := packet.New(m)
	if err != nil {
		return fmt.Errorf("creating packet: %v", err)
	}

	rep, err := p.Marshal()
	if err != nil {
		return fmt.Errorf("marshalling packet: %v", err)
	}

	return packet.Send(rep)
}
