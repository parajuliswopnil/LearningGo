package main

import (
	"fmt"
)

func returnAnyValue(V any) {
	fmt.Println("The variable that is passed is: ", V)
}

func returnKeysInMap[K comparable, V any](m map[K]V) []K {
	keyList := make([]K, len(m))
	for keys, _ := range m {
		keyList = append(keyList, keys)
	}
	return keyList
}

type List[T any] struct {
	head, tail *elements[T]
}

type elements[T any] struct {
	next     *elements[T]
	previous *elements[T]
	value    T
}

func (list *List[T]) push(val T) {
	if list.tail == nil {
		list.head = &elements[T]{value: val}
		list.tail = list.head
	} else {
		list.tail.next = &elements[T]{value: val, previous: list.tail}
		list.tail = list.tail.next
	}
}

func (list *List[T]) pop() {
	if list.tail.previous == nil {
		list.head = nil
		list.tail = nil
	} else {
		list.tail = list.tail.previous
		list.tail.next = nil
	}

}

func (list *List[T]) getListElements() []T {
	var listElements []T
	if list.head == nil {
		return listElements
	} else {
		for e := list.head; e != nil; e = e.next {
			listElements = append(listElements, e.value)
		}
		return listElements
	}
}

func generics() {
	returnAnyValue(10)
	returnAnyValue("Swopnil")

	var personInfo = map[string]int{"swopnil": 1, "RandomGuy": 2}
	returnAnyValue(personInfo)

	fmt.Println(returnKeysInMap(personInfo))

	customList := List[int]{}

	customList.push(1)

	customList.push(2)

	customList.push(3)

	// fmt.Println(customList.head)
	// fmt.Println(customList.head.next)
	// fmt.Println(customList.tail)

	// customList.push(4)
	// customList.push(5)

	fmt.Println(customList.getListElements())

	customList.pop()
	fmt.Println(customList.tail)

	fmt.Println(customList.getListElements())

}
