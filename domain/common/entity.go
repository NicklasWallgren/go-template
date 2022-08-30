package common

type PrimaryID uint32

type EntityConstraint interface {
	Id() PrimaryID
}

// type EntityConstraint interface { // TODO, not available until GO 1.19+
// 	~struct{ Entity }
// }

type Entity struct {
	ID PrimaryID // Uid
}

func (e Entity) Id() PrimaryID { // nolint:revive,stylecheck
	return e.ID
}
