package structs

type TeachersNames struct {
	FakePk int      `reindex:"fake_PK,,pk" json:"-"`
	Names  []string `reindex:"names" json:"names"`
}
