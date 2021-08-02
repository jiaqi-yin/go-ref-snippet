package main

import "fmt"

type Person struct {
	Id   int
	Name string
	Next *Person
}

func (p *Person) ShowMe() {
	fmt.Printf("person %d found in link %d\n", p.Id, p.Id%7)
}

type PersonLink struct {
	Head *Person
}

func (pl *PersonLink) Insert(p *Person) {
	current := pl.Head
	var prev *Person

	if current == nil {
		pl.Head = p
		return
	}

	for {
		if current != nil {
			if current.Id > p.Id {
				break
			}
			prev = current
			current = current.Next
		} else {
			break
		}
	}

	if prev == nil {
		pl.Head = p
		p.Next = current
	} else {
		prev.Next = p
		p.Next = current
	}

}

func (pl *PersonLink) ShowLink(num int) {
	if pl.Head == nil {
		fmt.Printf("link %d is empty\n", num)
		return
	}

	current := pl.Head
	fmt.Printf("link %d ", num)
	for {
		if current != nil {
			fmt.Printf("[id=%d name=%s] -> ", current.Id, current.Name)
			current = current.Next
		} else {
			break
		}
	}
	fmt.Println("")
}

func (pl *PersonLink) FindById(id int) *Person {
	current := pl.Head
	for {
		if current == nil {
			break
		}
		if current != nil && current.Id == id {
			return current
		}
		current = current.Next
	}
	return nil
}

func (pl *PersonLink) DeleteById(id int) bool {
	current := pl.Head
	var prev *Person
	for {
		if current == nil {
			return false
		}
		if current != nil && current.Id == id {
			// Delete
			if prev == nil {
				pl.Head = current.Next
			} else {
				prev.Next = current.Next
			}
			return true
		}
		prev = current
		current = current.Next
	}
}

type HashTable struct {
	LinkArr [7]PersonLink
}

func (ht *HashTable) Insert(p *Person) {
	linkNo := ht.HashFun(p.Id)
	ht.LinkArr[linkNo].Insert(p)
}

func (ht *HashTable) ShowAll() {
	fmt.Println("********** Show all **********")
	for i := 0; i < len(ht.LinkArr); i++ {
		ht.LinkArr[i].ShowLink(i)
	}
}

func (ht *HashTable) FindById(id int) *Person {
	linkNo := ht.HashFun(id)
	return ht.LinkArr[linkNo].FindById(id)
}

func (ht *HashTable) DeleteById(id int) bool {
	linkNo := ht.HashFun(id)
	return ht.LinkArr[linkNo].DeleteById(id)
}

func (ht *HashTable) HashFun(id int) int {
	return id % 7
}

func main() {
	var hashTable HashTable

	person1 := &Person{
		Id:   100,
		Name: "Tom",
	}

	person2 := &Person{
		Id:   80,
		Name: "Jerry",
	}

	person3 := &Person{
		Id:   9,
		Name: "John",
	}

	person4 := &Person{
		Id:   87,
		Name: "Jane",
	}

	hashTable.Insert(person1)
	hashTable.Insert(person2)
	hashTable.Insert(person3)
	hashTable.Insert(person4)

	hashTable.ShowAll()

	person := hashTable.FindById(1)
	if person == nil {
		fmt.Printf("person %d not found\n", 1)
	} else {
		person.ShowMe()
	}

	person = hashTable.FindById(80)
	if person == nil {
		fmt.Printf("person %d not found\n", 80)
	} else {
		person.ShowMe()
	}

	deleted := hashTable.DeleteById(1)
	fmt.Printf("Delete person %d: %t\n", 1, deleted)
	deleted = hashTable.DeleteById(100)
	fmt.Printf("Delete person %d: %t\n", 100, deleted)
	deleted = hashTable.DeleteById(80)
	fmt.Printf("Delete person %d: %t\n", 80, deleted)

	hashTable.ShowAll()
}
