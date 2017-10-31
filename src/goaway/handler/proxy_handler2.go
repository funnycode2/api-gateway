package handler

import (
	"github.com/valyala/fasthttp"
	"sync"
	"net"
	"bufio"
	"github.com/labstack/gommon/log"
)

type ProxyHandler2 struct{
	readerPool sync.Pool
	writerPool sync.Pool
}


func (h *ProxyHandler2) Matches(url string) bool {
	return true
}

func (h *ProxyHandler2) Handle(
	req *fasthttp.Request,
	res *fasthttp.Response,
	ctx *fasthttp.RequestCtx,
	host string) {

	upReq := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(upReq)
	upReq.Reset()
	req.CopyTo(upReq)
	upReq.SetHost(host)
	log.Printf("forwarding to host: " + string(upReq.Host()))

	upRes := fasthttp.AcquireResponse()
	upRes.Reset()
	defer fasthttp.ReleaseResponse(upRes)

	fasthttp.Do(upReq, upRes)

	upRes.Header.CopyTo(&res.Header)
	res.AppendBody(upRes.Body())
}

func (c *ProxyHandler2) acquireWriter(conn *net.Conn) *bufio.Writer {
	writer := c.writerPool.Get()
	if writer == nil {
		newWriter := bufio.NewWriterSize(*conn, BUF_SIZE)
		newWriter.Reset(*conn)
		return newWriter
	}
	return writer.(*bufio.Writer)
}

func (c *ProxyHandler2) releaseWriter(writer *bufio.Writer) {
	c.writerPool.Put(writer)
}

func (c *ProxyHandler2) acquireReader(conn *net.Conn) *bufio.Reader {
	reader := c.readerPool.Get()
	if reader == nil {
		newReader := bufio.NewReaderSize(*conn, BUF_SIZE)
		newReader.Reset(*conn)
		return newReader
	}
	return reader.(*bufio.Reader)
}

func (c *ProxyHandler2) releaseReader(reader *bufio.Reader) {
	c.readerPool.Put(reader)
}