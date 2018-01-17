#!/bin/bash

dos2unix  service-names-port-numbers.csv
cat  service-names-port-numbers.csv | \
	awk -F, '{ 
			if($1 =="") {  next }
		
				if($3=="") {
					printf("_%s._%s:%d\n", $1, "tcp", $2 ) 
				} else { 
					printf("_%s._%s:%d\n", $1, $3, $2 ) 
				}  
	}'  > mdns_service.all



