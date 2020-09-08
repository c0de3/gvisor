// automatically generated by stateify.

package fuse

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (x *Request) StateTypeName() string {
	return "pkg/sentry/fsimpl/fuse.Request"
}

func (x *Request) StateFields() []string {
	return []string{
		"requestEntry",
		"id",
		"hdr",
		"data",
	}
}

func (x *Request) beforeSave() {}

func (x *Request) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.requestEntry)
	m.Save(1, &x.id)
	m.Save(2, &x.hdr)
	m.Save(3, &x.data)
}

func (x *Request) afterLoad() {}

func (x *Request) StateLoad(m state.Source) {
	m.Load(0, &x.requestEntry)
	m.Load(1, &x.id)
	m.Load(2, &x.hdr)
	m.Load(3, &x.data)
}

func (x *Response) StateTypeName() string {
	return "pkg/sentry/fsimpl/fuse.Response"
}

func (x *Response) StateFields() []string {
	return []string{
		"opcode",
		"hdr",
		"data",
	}
}

func (x *Response) beforeSave() {}

func (x *Response) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.opcode)
	m.Save(1, &x.hdr)
	m.Save(2, &x.data)
}

func (x *Response) afterLoad() {}

func (x *Response) StateLoad(m state.Source) {
	m.Load(0, &x.opcode)
	m.Load(1, &x.hdr)
	m.Load(2, &x.data)
}

func (x *futureResponse) StateTypeName() string {
	return "pkg/sentry/fsimpl/fuse.futureResponse"
}

func (x *futureResponse) StateFields() []string {
	return []string{
		"opcode",
		"ch",
		"hdr",
		"data",
	}
}

func (x *futureResponse) beforeSave() {}

func (x *futureResponse) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.opcode)
	m.Save(1, &x.ch)
	m.Save(2, &x.hdr)
	m.Save(3, &x.data)
}

func (x *futureResponse) afterLoad() {}

func (x *futureResponse) StateLoad(m state.Source) {
	m.Load(0, &x.opcode)
	m.Load(1, &x.ch)
	m.Load(2, &x.hdr)
	m.Load(3, &x.data)
}

func (x *inodeRefs) StateTypeName() string {
	return "pkg/sentry/fsimpl/fuse.inodeRefs"
}

func (x *inodeRefs) StateFields() []string {
	return []string{
		"refCount",
	}
}

func (x *inodeRefs) beforeSave() {}

func (x *inodeRefs) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.refCount)
}

func (x *inodeRefs) afterLoad() {}

func (x *inodeRefs) StateLoad(m state.Source) {
	m.Load(0, &x.refCount)
}

func (x *requestList) StateTypeName() string {
	return "pkg/sentry/fsimpl/fuse.requestList"
}

func (x *requestList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (x *requestList) beforeSave() {}

func (x *requestList) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.head)
	m.Save(1, &x.tail)
}

func (x *requestList) afterLoad() {}

func (x *requestList) StateLoad(m state.Source) {
	m.Load(0, &x.head)
	m.Load(1, &x.tail)
}

func (x *requestEntry) StateTypeName() string {
	return "pkg/sentry/fsimpl/fuse.requestEntry"
}

func (x *requestEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (x *requestEntry) beforeSave() {}

func (x *requestEntry) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.next)
	m.Save(1, &x.prev)
}

func (x *requestEntry) afterLoad() {}

func (x *requestEntry) StateLoad(m state.Source) {
	m.Load(0, &x.next)
	m.Load(1, &x.prev)
}

func init() {
	state.Register((*Request)(nil))
	state.Register((*Response)(nil))
	state.Register((*futureResponse)(nil))
	state.Register((*inodeRefs)(nil))
	state.Register((*requestList)(nil))
	state.Register((*requestEntry)(nil))
}