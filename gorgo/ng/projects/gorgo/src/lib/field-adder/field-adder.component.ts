import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { FieldDB } from '../field-db'
import { FieldService } from '../field.service'


@Component({
  selector: 'app-field-adder',
  templateUrl: './field-adder.component.html',
  styleUrls: ['./field-adder.component.css']
})
export class FieldAdderComponent implements OnInit {

	field = {} as FieldDB;






  constructor(
    private fieldService: FieldService, 
	  ) {
  }

  ngOnInit(): void {

  }


  add(): void {



    this.fieldService.postField( this.field )
    .subscribe(field => {
		this.fieldService.FieldServiceChanged.next("post")
		
		this.field = {} // reset fields
	    console.log("field added")
    });
  }
}
