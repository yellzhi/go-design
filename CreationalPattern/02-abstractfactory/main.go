package main

func getMainAndDetail(fac DAOFactory){
	fac.CreateOrderDetailDAO().SaveOrderDetail()
	fac.CreaterOrderMainDAO().SaveOrderMain()
}

func main(){
	//test RDB
	var fac DAOFactory
	fac = &RDBDAOFactory{}
	getMainAndDetail(fac)

	// test xml
	fac = &XMLDAOFactory{}
	getMainAndDetail(fac)
}

