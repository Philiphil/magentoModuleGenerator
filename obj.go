package main

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