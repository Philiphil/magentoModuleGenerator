package main
import (
	"fmt"
	"os"
	"strings"
)
/*
	cron,router,controler,model,rewrite,bloc,layout
*/

type Module struct{
	Archi string
	Namespace string
	Name string
	HasObserver bool
	IsEventCustom bool
	Event string
	HasHelper bool
}
func main(){
	createModule(makeModule());
	
}

func makeModule()(m Module){
	m=getArchi(m)
	fmt.Println("module Namespace ?")
	fmt.Scan(&m.Namespace)
	fmt.Println("module Name ?")
	fmt.Scan(&m.Name)
	fmt.Println("has module observer ?")
	m.HasObserver = askBool();
	if m.HasObserver{
		fmt.Println("Is observer custom ?")
		m.IsEventCustom = askBool();
		fmt.Println("event name ?")
		fmt.Scan(&m.Name)
	}
	return
}

func askBool()(b bool){
	var v string
	fmt.Scan(&v)
	switch(v){
		case "yes":return true
		case "y" : return true
		case "true" : return true
		case "oui" : return true
		case "o" :return true
		default :return false;
	}
}

func getArchi(m Module)(Module){
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var mapi = explore(pwd)
	var localfolder = "app"+string(os.PathSeparator)+"code"+string(os.PathSeparator)+"local"
	for key,_ := range mapi {
		if strings.Contains(key,localfolder){
			m.Archi = key[0:strings.Index(key,localfolder)+len(localfolder)+1]
			return m;
		}
	}
	return m
}

func createModule(m Module){
/*

*/
}