import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { PositionDB } from '../position-db'
import { PositionService } from '../position.service'


@Component({
  selector: 'app-position-adder',
  templateUrl: './position-adder.component.html',
  styleUrls: ['./position-adder.component.css']
})
export class PositionAdderComponent implements OnInit {

	position = {} as PositionDB;






  constructor(
    private positionService: PositionService, 
	  ) {
  }

  ngOnInit(): void {

  }


  add(): void {



    this.positionService.postPosition( this.position )
    .subscribe(position => {
		this.positionService.PositionServiceChanged.next("post")
		
		this.position = {} // reset fields
	    console.log("position added")
    });
  }
}
