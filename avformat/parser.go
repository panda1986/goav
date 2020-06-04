package avformat

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>

import "C"

func (ctxt *CodecParserContext) Width() int {
    return int(ctxt.width)
}
