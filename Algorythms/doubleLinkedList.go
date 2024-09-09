package main;
import "fmt"

func main(){

	list1 :=createDoubleLinkedList()
	list1.insert(1);
	list1.insert(2);
	list1.insert(3);
	list1.printAsc();
	list1.printDesc();
}


type node struct{
	data int
	prev *node
	next *node
}

type doubleLinkedList struct{
	head *node
	tail *node
}

func createDoubleLinkedList() doubleLinkedList {
	return doubleLinkedList{nil, nil}
}

func (list *doubleLinkedList) insert(data int){

	var newNode node= node{ data: data } 

	if list.head==nil{
		list.head=&newNode;
		list.tail=&newNode;
	}else{
		var currentNode *node= list.head; // *node because we dont need the node here but the refrence to the memeory location of the node
		
		for currentNode.next!=nil{
			currentNode= currentNode.next;
		}

		newNode.prev=currentNode; // not &currentNode because ^ above we did currentNode *node and the currentNode is already holding the address to the memory location
		currentNode.next= &newNode;
		list.tail=&newNode;
	}

	return

}

func (list *doubleLinkedList) printAsc(){

	if list.head==nil{
		fmt.Println("The List is Empty");
		return;
	}
	var currentNode *node= list.head;
	for currentNode!=nil{
		fmt.Println(currentNode);
		currentNode=currentNode.next;
	}
}


func (list *doubleLinkedList) printDesc(){

	if  list.tail==nil {
		fmt.Println("The List is Empty");
		return;
	}
	var currentNode *node= list.tail;
	for currentNode!=nil{
		fmt.Println(currentNode);
		currentNode=currentNode.prev;
	}
}


