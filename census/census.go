package census

type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident{Name: name, Age: age, Address: address}
}

func (r *Resident) HasRequiredInfo() bool {
	return len(r.Name) > 0 && len(r.Address["street"]) > 0
}

func (r *Resident) Delete() {
	*r = Resident{}
}

func Count(residents []*Resident) int {
	count := 0

	for _, resident := range residents {
		if resident.HasRequiredInfo() {
			count++
		}
	}
	
	return count
}
