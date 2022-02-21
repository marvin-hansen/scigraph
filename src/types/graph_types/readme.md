# Golang DAG[ T ] (Directed Acyclic Graph)

This is a generic DAG based largely on the work of Hashicorp. The intent of this library is to allow for anyone to create a DAG using any datatype so long as it implements `Hashcode() string`, that does not require casting the interface upon getting from the dag.
The files in this code were taken from [this commit](https://github.com/hashicorp/terraform/tree/8b6522169fcf5ea991d2ff45087526ebccf664d6)