package common

type PrimaryId uint32

type EntityConstraint interface {
	Id() PrimaryId
}

//type EntityConstraint interface { // TODO, not available until GO 1.19
//	~struct{ Entity }
//}

type Entity struct {
	ID PrimaryId // Uid
}

func (e Entity) Id() PrimaryId {
	return e.ID
}
