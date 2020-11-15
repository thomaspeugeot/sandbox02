import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { MarshallerDB } from '../marshaller-db'
import { MarshallerService } from '../marshaller.service'


@Component({
  selector: 'app-marshaller-adder',
  templateUrl: './marshaller-adder.component.html',
  styleUrls: ['./marshaller-adder.component.css']
})
export class MarshallerAdderComponent implements OnInit {

	marshaller = {} as MarshallerDB;






  constructor(
    private marshallerService: MarshallerService, 
	  ) {
  }

  ngOnInit(): void {

  }


  add(): void {



    this.marshallerService.postMarshaller( this.marshaller )
    .subscribe(marshaller => {
		this.marshallerService.MarshallerServiceChanged.next("post")
		
		this.marshaller = {} // reset fields
	    console.log("marshaller added")
    });
  }
}
