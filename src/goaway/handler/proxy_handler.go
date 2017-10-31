package handler

import (
	"github.com/valyala/fasthttp"
	"sync"
	"net"
	"bufio"
)

type ProxyHandler struct{
	readerPool sync.Pool
	writerPool sync.Pool
}

const (
	BUF_SIZE = 4096
)


func (h *ProxyHandler) Matches(url string) bool {
	return true
}

func (h *ProxyHandler) Handle(
	req *fasthttp.Request,
	res *fasthttp.Response,
	ctx *fasthttp.RequestCtx,
	host string) {
	conn, _ := fasthttp.Dial(host)

	writer := h.acquireWriter(&conn)
	defer h.releaseWriter(writer)
	req.Write(writer)
	writer.Flush()

	reader := h.acquireReader(&conn)
	defer h.releaseReader(reader)

	upRes := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(upRes)

	upRes.Reset()
	upRes.Read(reader)

	upRes.Header.CopyTo(&res.Header)
	res.AppendBody(upRes.Body())
}

func (c *ProxyHandler) acquireWriter(conn *net.Conn) *bufio.Writer {
	writer := c.writerPool.Get()
	if writer == nil {
		newWriter := bufio.NewWriterSize(*conn, BUF_SIZE)
		newWriter.Reset(*conn)
		return newWriter
	}
	return writer.(*bufio.Writer)
}

func (c *ProxyHandler) releaseWriter(writer *bufio.Writer) {
	c.writerPool.Put(writer)
}

func (c *ProxyHandler) acquireReader(conn *net.Conn) *bufio.Reader {
	reader := c.readerPool.Get()
	if reader == nil {
		newReader := bufio.NewReaderSize(*conn, BUF_SIZE)
		newReader.Reset(*conn)
		return newReader
	}
	return reader.(*bufio.Reader)
}

func (c *ProxyHandler) releaseReader(reader *bufio.Reader) {
	c.readerPool.Put(reader)
}