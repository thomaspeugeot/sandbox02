import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { UmlscDB } from '../umlsc-db'
import { UmlscService } from '../umlsc.service'


@Component({
  selector: 'app-umlsc-adder',
  templateUrl: './umlsc-adder.component.html',
  styleUrls: ['./umlsc-adder.component.css']
})
export class UmlscAdderComponent implements OnInit {

	umlsc = {} as UmlscDB;






  constructor(
    private umlscService: UmlscService, 
	  ) {
  }

  ngOnInit(): void {

  }


  add(): void {



    this.umlscService.postUmlsc( this.umlsc )
    .subscribe(umlsc => {
		this.umlscService.UmlscServiceChanged.next("post")
		
		this.umlsc = {} // reset fields
	    console.log("umlsc added")
    });
  }
}
