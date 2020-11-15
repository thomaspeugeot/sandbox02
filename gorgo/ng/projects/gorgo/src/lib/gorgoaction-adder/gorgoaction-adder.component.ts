import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { GorgoactionDB } from '../gorgoaction-db'
import { GorgoactionService } from '../gorgoaction.service'


@Component({
  selector: 'app-gorgoaction-adder',
  templateUrl: './gorgoaction-adder.component.html',
  styleUrls: ['./gorgoaction-adder.component.css']
})
export class GorgoactionAdderComponent implements OnInit {

	gorgoaction = {} as GorgoactionDB;






  constructor(
    private gorgoactionService: GorgoactionService, 
	  ) {
  }

  ngOnInit(): void {

  }


  add(): void {



    this.gorgoactionService.postGorgoaction( this.gorgoaction )
    .subscribe(gorgoaction => {
		this.gorgoactionService.GorgoactionServiceChanged.next("post")
		
		this.gorgoaction = {} // reset fields
	    console.log("gorgoaction added")
    });
  }
}
