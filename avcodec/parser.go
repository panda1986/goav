package avcodec

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>

import "C"

func (ctxt *ParserContext) Width() int {
    return int(ctxt.width)
}

func (ctxt *ParserContext) Height() int {
    return int(ctxt.height)
}

func (ctxt *ParserContext) PixFmt() PixelFormat {
    return (PixelFormat)(ctxt.format)
}

func (ctxt *ParserContext) FieldOrder() AvFieldOrder {
    return (AvFieldOrder)(ctxt.field_order)
}
