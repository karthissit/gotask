package main

import (
	"fmt"
)

func printMap(nodeFamilies map[string][]string){
	for key, value := range nodeFamilies {
		fmt.Println("Family \""+key+"\" has nodes ", value)
	}
	fmt.Println()
}

func checkFamilyExists(nodeFamilies map[string][]string, family string) bool{
	_,ok :=nodeFamilies[family]
	if(ok){return true}
	return false
}

func addFamily(nodeFamilies map[string][]string, family string) map[string][]string{
	nodeFamilies[family]=[]string{}
	return nodeFamilies
}

func checkNodeExists(nodes []string, node string) int{
	var index = -1
	for i:=0;i<len(nodes);i++{
		if(nodes[i]== node){
			index = i;
		}
	}

	return index
}

func addFamilyNodes(nodeFamilies map[string][]string, family string, node string) map[string][]string{
	var arr = []string{node}
	if(checkNodeExists(nodeFamilies[family], node) < 0){
		nodeFamilies[family]=append(nodeFamilies[family], arr...)
		fmt.Println("Node added to family")
	}else{
		fmt.Println("Cannot add duplicate Node to family")
	}
	return nodeFamilies
}

func deleteNode(nodes []string, i int) []string{
	copy(nodes[i:], nodes[i+1:])
	nodes[len(nodes)-1] = ""
	nodes = nodes[:len(nodes)-1]
	return nodes
}

func featureAddFamily(nodeFamilies map[string][]string) map[string][]string {
	var family string
	fmt.Println("Enter the Family name to be added")	
    fmt.Scanln(&family)
	if(!checkFamilyExists(nodeFamilies,family)){
		nodeFamilies=addFamily(nodeFamilies,family)
		fmt.Println("Family added")
		fmt.Println()
		fmt.Println("List of Families and its node are as below")
		printMap(nodeFamilies)
	}else{
		fmt.Println("Family exists, cannot add duplicate family")
		fmt.Println()		
	}

	return nodeFamilies
}

func featureAddNodes(nodeFamilies map[string][]string) map[string][]string {
	var family,node string
	fmt.Println("Enter the Family name and Node name to be added as space separated values")
	fmt.Scanln(&family,&node)

	if(!checkFamilyExists(nodeFamilies,family)){
		nodeFamilies=addFamily(nodeFamilies,family)		
	}

	fmt.Println()
	fmt.Println("List of Families and its node are as below")
	printMap(nodeFamilies)

	return addFamilyNodes(nodeFamilies,family,node) 
}

func featureDeleteDeprecatedNode(nodeFamilies map[string][]string) map[string][]string {
	var family,node string
	fmt.Println("Enter the Family name and depricated Node name to be deleted as space seprated values")
	fmt.Scanln(&family,&node)

	if(!checkFamilyExists(nodeFamilies,family)){
		fmt.Println("Family \""+family+"\" does not exist")	
	}else{
		var index = checkNodeExists(nodeFamilies[family],node)
		if(index < 0){
			fmt.Println("Node \""+node+"\" does not exist in given family ",family)
		}else{
			nodeFamilies[family] = deleteNode(nodeFamilies[family],index)
			fmt.Println()
			fmt.Println("Node \""+node+"\" deleted from family ",family)
			fmt.Println()
			fmt.Println("List of Families and its node are as below")
			printMap(nodeFamilies)
		}
	}
}


func main() {

	nodeFamilies := make(map[string][]string)
	nodeFamilies["SR"]=[]string{"7750","7450","7950","7650"}
	nodeFamilies["IP"]=[]string{"6750","6450","6950","6650"}
	nodeFamilies["NK"]=[]string{"5750","5450","5950","5650"}

	fmt.Println("List of Families and its node are as below")
	printMap(nodeFamilies)

	nodeFamilies = featureAddFamily(nodeFamilies)

	nodeFamilies = featureAddNodes(nodeFamilies)

	nodeFamilies = featureDeleteDeprecatedNode(nodeFamilies)	

}
