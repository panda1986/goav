package avcodec

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>

import "C"

func (ctxt *ParserContext) Width() int {
    return int(ctxt.width)
}
