import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { ClassdiagramDB } from '../classdiagram-db'
import { ClassdiagramService } from '../classdiagram.service'


@Component({
  selector: 'app-classdiagram-adder',
  templateUrl: './classdiagram-adder.component.html',
  styleUrls: ['./classdiagram-adder.component.css']
})
export class ClassdiagramAdderComponent implements OnInit {

	classdiagram = {} as ClassdiagramDB;






  constructor(
    private classdiagramService: ClassdiagramService, 
	  ) {
  }

  ngOnInit(): void {

  }


  add(): void {



    this.classdiagramService.postClassdiagram( this.classdiagram )
    .subscribe(classdiagram => {
		this.classdiagramService.ClassdiagramServiceChanged.next("post")
		
		this.classdiagram = {} // reset fields
	    console.log("classdiagram added")
    });
  }
}
