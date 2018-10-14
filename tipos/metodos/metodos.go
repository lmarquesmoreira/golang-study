package metodos

import (
	"fmt"
	"strings"
)

type People struct {
	name string
	surname string
}

func (p People) getFullName() string {
	return p.name + " " + p.surname;
}

func (p *People) setFullName(fullName string) {
	nameSplited := strings.Split(fullName, " ")
	p.name = nameSplited[0]
	p.surname = nameSplited[1]
}

func ExecuteDemoMetodos(){
	p1 := People{"Pedro", "Silva"}
	fmt.Println(p1.getFullName())
	p1.setFullName("Maria Pereira")
	fmt.Println(p1.getFullName())
}


