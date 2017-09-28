// Copyright 2016 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package client

import (
	"io"

	"golang.org/x/net/context"
)

type ChatAPIDecoder struct {
	handler ChatAPIHandler
}

func NewChatAPIDecoder(h ChatAPIHandler) *ChatAPIDecoder {
	return &ChatAPIDecoder{
		handler: h,
	}
}

func (d *ChatAPIDecoder) handle(ctx context.Context, c Call, w io.Writer) error {
	switch c.Params.Version {
	case 0, 1:
		return d.handleV1(ctx, c, w)
	default:
		return ErrInvalidVersion{version: c.Params.Version}
	}
}

func (d *ChatAPIDecoder) handleV1(ctx context.Context, c Call, w io.Writer) error {
	switch c.Method {
	case methodList:
		return d.handler.ListV1(ctx, c, w)
	case methodRead:
		return d.handler.ReadV1(ctx, c, w)
	case methodSend:
		return d.handler.SendV1(ctx, c, w)
	case methodEdit:
		return d.handler.EditV1(ctx, c, w)
	case methodDelete:
		return d.handler.DeleteV1(ctx, c, w)
	case methodAttach:
		return d.handler.AttachV1(ctx, c, w)
	case methodDownload:
		return d.handler.DownloadV1(ctx, c, w)
	case methodSetStatus:
		return d.handler.SetStatusV1(ctx, c, w)
	case methodMark:
		return d.handler.MarkV1(ctx, c, w)
	default:
		return ErrInvalidMethod{name: c.Method, version: 1}
	}
}
