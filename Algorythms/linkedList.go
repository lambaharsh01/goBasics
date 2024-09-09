package main;
import "fmt";

func main(){

	list1 := newLinkedList();
	list1.print()
	list1.insert(1)
	list1.insert(2)
	list1.insert(3)
	list1.insert(4)
	list1.print()
	


}

type node struct{
	data int
	next *node
}

type linkedList struct{
	head *node
}

// when you create a new struct instance using the & operator, it initializes the struct with zero values for all its fields. In this case, the head field of the linkedList struct is a pointer to a node, and when you create a new linkedList instance, the head field is initialized to nil.
func newLinkedList() *linkedList{ //passes a pointer to a list it is going to return
 return &linkedList{nil} //the list is null
}



func (exisitingList *linkedList) insert(data int){  //The receiver type is specified as exisitingList *linkedList, which means that the method operates on a pointer to a linkedList instance.
													//The * symbol in exisitingList *linkedList is called the receiver pointer. It indicates that the method operates on a pointer to a linkedList instance, rather than a copy of the instance.
	var newNode node = node{
		data: data,
		next:nil,
	}

	if exisitingList.head==nil{
		exisitingList.head= &newNode;  // take the address of newNode
		// By using the & operator, you're taking the address of newNode, which gives you a pointer to a node value. This pointer can then be assigned to existingList.head, which is a *node field.
	}else{
		currentNode := exisitingList.head;
		 for currentNode.next!=nil{
			currentNode=currentNode.next;
		 }
		 currentNode.next= &newNode
	}
	
}


func (existingList *linkedList) print(){
	// {
	// var currentNode node= existingList.head;
	// The error message is telling you that you can't use existingList.head (which is a *node value) as a node value in the variable declaration.
	// The reason for this is that existingList.head is a pointer to a node, not a node value itself. In Go, when you declare a variable without the * operator, it's a value, not a pointer.
	// }

	var currentNode *node=  existingList.head;
	if currentNode==nil{
		fmt.Println("List is Empty");
		return;
	}

	for currentNode!= nil{
		fmt.Println(currentNode, "lOL");
		currentNode=currentNode.next;
	}
	return;

}