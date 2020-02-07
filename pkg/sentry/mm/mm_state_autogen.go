// automatically generated by stateify.

package mm

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (x *aioManager) beforeSave() {}
func (x *aioManager) save(m state.Map) {
	x.beforeSave()
	m.Save("contexts", &x.contexts)
}

func (x *aioManager) afterLoad() {}
func (x *aioManager) load(m state.Map) {
	m.Load("contexts", &x.contexts)
}

func (x *ioResult) beforeSave() {}
func (x *ioResult) save(m state.Map) {
	x.beforeSave()
	m.Save("data", &x.data)
	m.Save("ioEntry", &x.ioEntry)
}

func (x *ioResult) afterLoad() {}
func (x *ioResult) load(m state.Map) {
	m.Load("data", &x.data)
	m.Load("ioEntry", &x.ioEntry)
}

func (x *AIOContext) beforeSave() {}
func (x *AIOContext) save(m state.Map) {
	x.beforeSave()
	if !state.IsZeroValue(x.dead) {
		m.Failf("dead is %v, expected zero", x.dead)
	}
	m.Save("results", &x.results)
	m.Save("maxOutstanding", &x.maxOutstanding)
	m.Save("outstanding", &x.outstanding)
}

func (x *AIOContext) load(m state.Map) {
	m.Load("results", &x.results)
	m.Load("maxOutstanding", &x.maxOutstanding)
	m.Load("outstanding", &x.outstanding)
	m.AfterLoad(x.afterLoad)
}

func (x *aioMappable) beforeSave() {}
func (x *aioMappable) save(m state.Map) {
	x.beforeSave()
	m.Save("AtomicRefCount", &x.AtomicRefCount)
	m.Save("mfp", &x.mfp)
	m.Save("fr", &x.fr)
}

func (x *aioMappable) afterLoad() {}
func (x *aioMappable) load(m state.Map) {
	m.Load("AtomicRefCount", &x.AtomicRefCount)
	m.Load("mfp", &x.mfp)
	m.Load("fr", &x.fr)
}

func (x *fileRefcountSet) beforeSave() {}
func (x *fileRefcountSet) save(m state.Map) {
	x.beforeSave()
	var root *fileRefcountSegmentDataSlices = x.saveRoot()
	m.SaveValue("root", root)
}

func (x *fileRefcountSet) afterLoad() {}
func (x *fileRefcountSet) load(m state.Map) {
	m.LoadValue("root", new(*fileRefcountSegmentDataSlices), func(y interface{}) { x.loadRoot(y.(*fileRefcountSegmentDataSlices)) })
}

func (x *fileRefcountnode) beforeSave() {}
func (x *fileRefcountnode) save(m state.Map) {
	x.beforeSave()
	m.Save("nrSegments", &x.nrSegments)
	m.Save("parent", &x.parent)
	m.Save("parentIndex", &x.parentIndex)
	m.Save("hasChildren", &x.hasChildren)
	m.Save("keys", &x.keys)
	m.Save("values", &x.values)
	m.Save("children", &x.children)
}

func (x *fileRefcountnode) afterLoad() {}
func (x *fileRefcountnode) load(m state.Map) {
	m.Load("nrSegments", &x.nrSegments)
	m.Load("parent", &x.parent)
	m.Load("parentIndex", &x.parentIndex)
	m.Load("hasChildren", &x.hasChildren)
	m.Load("keys", &x.keys)
	m.Load("values", &x.values)
	m.Load("children", &x.children)
}

func (x *fileRefcountSegmentDataSlices) beforeSave() {}
func (x *fileRefcountSegmentDataSlices) save(m state.Map) {
	x.beforeSave()
	m.Save("Start", &x.Start)
	m.Save("End", &x.End)
	m.Save("Values", &x.Values)
}

func (x *fileRefcountSegmentDataSlices) afterLoad() {}
func (x *fileRefcountSegmentDataSlices) load(m state.Map) {
	m.Load("Start", &x.Start)
	m.Load("End", &x.End)
	m.Load("Values", &x.Values)
}

func (x *ioList) beforeSave() {}
func (x *ioList) save(m state.Map) {
	x.beforeSave()
	m.Save("head", &x.head)
	m.Save("tail", &x.tail)
}

func (x *ioList) afterLoad() {}
func (x *ioList) load(m state.Map) {
	m.Load("head", &x.head)
	m.Load("tail", &x.tail)
}

func (x *ioEntry) beforeSave() {}
func (x *ioEntry) save(m state.Map) {
	x.beforeSave()
	m.Save("next", &x.next)
	m.Save("prev", &x.prev)
}

func (x *ioEntry) afterLoad() {}
func (x *ioEntry) load(m state.Map) {
	m.Load("next", &x.next)
	m.Load("prev", &x.prev)
}

func (x *MemoryManager) save(m state.Map) {
	x.beforeSave()
	if !state.IsZeroValue(x.active) {
		m.Failf("active is %v, expected zero", x.active)
	}
	if !state.IsZeroValue(x.captureInvalidations) {
		m.Failf("captureInvalidations is %v, expected zero", x.captureInvalidations)
	}
	m.Save("p", &x.p)
	m.Save("mfp", &x.mfp)
	m.Save("layout", &x.layout)
	m.Save("privateRefs", &x.privateRefs)
	m.Save("users", &x.users)
	m.Save("vmas", &x.vmas)
	m.Save("brk", &x.brk)
	m.Save("usageAS", &x.usageAS)
	m.Save("lockedAS", &x.lockedAS)
	m.Save("dataAS", &x.dataAS)
	m.Save("defMLockMode", &x.defMLockMode)
	m.Save("pmas", &x.pmas)
	m.Save("curRSS", &x.curRSS)
	m.Save("maxRSS", &x.maxRSS)
	m.Save("argv", &x.argv)
	m.Save("envv", &x.envv)
	m.Save("auxv", &x.auxv)
	m.Save("executable", &x.executable)
	m.Save("dumpability", &x.dumpability)
	m.Save("aioManager", &x.aioManager)
}

func (x *MemoryManager) load(m state.Map) {
	m.Load("p", &x.p)
	m.Load("mfp", &x.mfp)
	m.Load("layout", &x.layout)
	m.Load("privateRefs", &x.privateRefs)
	m.Load("users", &x.users)
	m.Load("vmas", &x.vmas)
	m.Load("brk", &x.brk)
	m.Load("usageAS", &x.usageAS)
	m.Load("lockedAS", &x.lockedAS)
	m.Load("dataAS", &x.dataAS)
	m.Load("defMLockMode", &x.defMLockMode)
	m.Load("pmas", &x.pmas)
	m.Load("curRSS", &x.curRSS)
	m.Load("maxRSS", &x.maxRSS)
	m.Load("argv", &x.argv)
	m.Load("envv", &x.envv)
	m.Load("auxv", &x.auxv)
	m.Load("executable", &x.executable)
	m.Load("dumpability", &x.dumpability)
	m.Load("aioManager", &x.aioManager)
	m.AfterLoad(x.afterLoad)
}

func (x *vma) beforeSave() {}
func (x *vma) save(m state.Map) {
	x.beforeSave()
	var realPerms int = x.saveRealPerms()
	m.SaveValue("realPerms", realPerms)
	m.Save("mappable", &x.mappable)
	m.Save("off", &x.off)
	m.Save("dontfork", &x.dontfork)
	m.Save("mlockMode", &x.mlockMode)
	m.Save("numaPolicy", &x.numaPolicy)
	m.Save("numaNodemask", &x.numaNodemask)
	m.Save("id", &x.id)
	m.Save("hint", &x.hint)
}

func (x *vma) afterLoad() {}
func (x *vma) load(m state.Map) {
	m.Load("mappable", &x.mappable)
	m.Load("off", &x.off)
	m.Load("dontfork", &x.dontfork)
	m.Load("mlockMode", &x.mlockMode)
	m.Load("numaPolicy", &x.numaPolicy)
	m.Load("numaNodemask", &x.numaNodemask)
	m.Load("id", &x.id)
	m.Load("hint", &x.hint)
	m.LoadValue("realPerms", new(int), func(y interface{}) { x.loadRealPerms(y.(int)) })
}

func (x *pma) beforeSave() {}
func (x *pma) save(m state.Map) {
	x.beforeSave()
	m.Save("off", &x.off)
	m.Save("translatePerms", &x.translatePerms)
	m.Save("effectivePerms", &x.effectivePerms)
	m.Save("maxPerms", &x.maxPerms)
	m.Save("needCOW", &x.needCOW)
	m.Save("private", &x.private)
}

func (x *pma) afterLoad() {}
func (x *pma) load(m state.Map) {
	m.Load("off", &x.off)
	m.Load("translatePerms", &x.translatePerms)
	m.Load("effectivePerms", &x.effectivePerms)
	m.Load("maxPerms", &x.maxPerms)
	m.Load("needCOW", &x.needCOW)
	m.Load("private", &x.private)
}

func (x *privateRefs) beforeSave() {}
func (x *privateRefs) save(m state.Map) {
	x.beforeSave()
	m.Save("refs", &x.refs)
}

func (x *privateRefs) afterLoad() {}
func (x *privateRefs) load(m state.Map) {
	m.Load("refs", &x.refs)
}

func (x *pmaSet) beforeSave() {}
func (x *pmaSet) save(m state.Map) {
	x.beforeSave()
	var root *pmaSegmentDataSlices = x.saveRoot()
	m.SaveValue("root", root)
}

func (x *pmaSet) afterLoad() {}
func (x *pmaSet) load(m state.Map) {
	m.LoadValue("root", new(*pmaSegmentDataSlices), func(y interface{}) { x.loadRoot(y.(*pmaSegmentDataSlices)) })
}

func (x *pmanode) beforeSave() {}
func (x *pmanode) save(m state.Map) {
	x.beforeSave()
	m.Save("nrSegments", &x.nrSegments)
	m.Save("parent", &x.parent)
	m.Save("parentIndex", &x.parentIndex)
	m.Save("hasChildren", &x.hasChildren)
	m.Save("keys", &x.keys)
	m.Save("values", &x.values)
	m.Save("children", &x.children)
}

func (x *pmanode) afterLoad() {}
func (x *pmanode) load(m state.Map) {
	m.Load("nrSegments", &x.nrSegments)
	m.Load("parent", &x.parent)
	m.Load("parentIndex", &x.parentIndex)
	m.Load("hasChildren", &x.hasChildren)
	m.Load("keys", &x.keys)
	m.Load("values", &x.values)
	m.Load("children", &x.children)
}

func (x *pmaSegmentDataSlices) beforeSave() {}
func (x *pmaSegmentDataSlices) save(m state.Map) {
	x.beforeSave()
	m.Save("Start", &x.Start)
	m.Save("End", &x.End)
	m.Save("Values", &x.Values)
}

func (x *pmaSegmentDataSlices) afterLoad() {}
func (x *pmaSegmentDataSlices) load(m state.Map) {
	m.Load("Start", &x.Start)
	m.Load("End", &x.End)
	m.Load("Values", &x.Values)
}

func (x *SpecialMappable) beforeSave() {}
func (x *SpecialMappable) save(m state.Map) {
	x.beforeSave()
	m.Save("AtomicRefCount", &x.AtomicRefCount)
	m.Save("mfp", &x.mfp)
	m.Save("fr", &x.fr)
	m.Save("name", &x.name)
}

func (x *SpecialMappable) afterLoad() {}
func (x *SpecialMappable) load(m state.Map) {
	m.Load("AtomicRefCount", &x.AtomicRefCount)
	m.Load("mfp", &x.mfp)
	m.Load("fr", &x.fr)
	m.Load("name", &x.name)
}

func (x *vmaSet) beforeSave() {}
func (x *vmaSet) save(m state.Map) {
	x.beforeSave()
	var root *vmaSegmentDataSlices = x.saveRoot()
	m.SaveValue("root", root)
}

func (x *vmaSet) afterLoad() {}
func (x *vmaSet) load(m state.Map) {
	m.LoadValue("root", new(*vmaSegmentDataSlices), func(y interface{}) { x.loadRoot(y.(*vmaSegmentDataSlices)) })
}

func (x *vmanode) beforeSave() {}
func (x *vmanode) save(m state.Map) {
	x.beforeSave()
	m.Save("nrSegments", &x.nrSegments)
	m.Save("parent", &x.parent)
	m.Save("parentIndex", &x.parentIndex)
	m.Save("hasChildren", &x.hasChildren)
	m.Save("keys", &x.keys)
	m.Save("values", &x.values)
	m.Save("children", &x.children)
}

func (x *vmanode) afterLoad() {}
func (x *vmanode) load(m state.Map) {
	m.Load("nrSegments", &x.nrSegments)
	m.Load("parent", &x.parent)
	m.Load("parentIndex", &x.parentIndex)
	m.Load("hasChildren", &x.hasChildren)
	m.Load("keys", &x.keys)
	m.Load("values", &x.values)
	m.Load("children", &x.children)
}

func (x *vmaSegmentDataSlices) beforeSave() {}
func (x *vmaSegmentDataSlices) save(m state.Map) {
	x.beforeSave()
	m.Save("Start", &x.Start)
	m.Save("End", &x.End)
	m.Save("Values", &x.Values)
}

func (x *vmaSegmentDataSlices) afterLoad() {}
func (x *vmaSegmentDataSlices) load(m state.Map) {
	m.Load("Start", &x.Start)
	m.Load("End", &x.End)
	m.Load("Values", &x.Values)
}

func init() {
	state.Register("mm.aioManager", (*aioManager)(nil), state.Fns{Save: (*aioManager).save, Load: (*aioManager).load})
	state.Register("mm.ioResult", (*ioResult)(nil), state.Fns{Save: (*ioResult).save, Load: (*ioResult).load})
	state.Register("mm.AIOContext", (*AIOContext)(nil), state.Fns{Save: (*AIOContext).save, Load: (*AIOContext).load})
	state.Register("mm.aioMappable", (*aioMappable)(nil), state.Fns{Save: (*aioMappable).save, Load: (*aioMappable).load})
	state.Register("mm.fileRefcountSet", (*fileRefcountSet)(nil), state.Fns{Save: (*fileRefcountSet).save, Load: (*fileRefcountSet).load})
	state.Register("mm.fileRefcountnode", (*fileRefcountnode)(nil), state.Fns{Save: (*fileRefcountnode).save, Load: (*fileRefcountnode).load})
	state.Register("mm.fileRefcountSegmentDataSlices", (*fileRefcountSegmentDataSlices)(nil), state.Fns{Save: (*fileRefcountSegmentDataSlices).save, Load: (*fileRefcountSegmentDataSlices).load})
	state.Register("mm.ioList", (*ioList)(nil), state.Fns{Save: (*ioList).save, Load: (*ioList).load})
	state.Register("mm.ioEntry", (*ioEntry)(nil), state.Fns{Save: (*ioEntry).save, Load: (*ioEntry).load})
	state.Register("mm.MemoryManager", (*MemoryManager)(nil), state.Fns{Save: (*MemoryManager).save, Load: (*MemoryManager).load})
	state.Register("mm.vma", (*vma)(nil), state.Fns{Save: (*vma).save, Load: (*vma).load})
	state.Register("mm.pma", (*pma)(nil), state.Fns{Save: (*pma).save, Load: (*pma).load})
	state.Register("mm.privateRefs", (*privateRefs)(nil), state.Fns{Save: (*privateRefs).save, Load: (*privateRefs).load})
	state.Register("mm.pmaSet", (*pmaSet)(nil), state.Fns{Save: (*pmaSet).save, Load: (*pmaSet).load})
	state.Register("mm.pmanode", (*pmanode)(nil), state.Fns{Save: (*pmanode).save, Load: (*pmanode).load})
	state.Register("mm.pmaSegmentDataSlices", (*pmaSegmentDataSlices)(nil), state.Fns{Save: (*pmaSegmentDataSlices).save, Load: (*pmaSegmentDataSlices).load})
	state.Register("mm.SpecialMappable", (*SpecialMappable)(nil), state.Fns{Save: (*SpecialMappable).save, Load: (*SpecialMappable).load})
	state.Register("mm.vmaSet", (*vmaSet)(nil), state.Fns{Save: (*vmaSet).save, Load: (*vmaSet).load})
	state.Register("mm.vmanode", (*vmanode)(nil), state.Fns{Save: (*vmanode).save, Load: (*vmanode).load})
	state.Register("mm.vmaSegmentDataSlices", (*vmaSegmentDataSlices)(nil), state.Fns{Save: (*vmaSegmentDataSlices).save, Load: (*vmaSegmentDataSlices).load})
}