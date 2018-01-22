package entities

type Biodata struct {
	emailId string
	age string
	name string
}

func (b *Biodata) SetEmailID(emailId string){
	b.emailId = emailId
}

func (b *Biodata) SetAge(age string){
	b.age = age
}

func (b *Biodata) SetName(name string){
	b.name = name
}

func (b Biodata) GetEmailID() string{
	return b.emailId
}

func (b Biodata) GetAge() string {
	return b.age
}

func (b Biodata) GetName() string {
	return b.name
}