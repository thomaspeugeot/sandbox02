import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { ActionDB } from '../action-db'
import { ActionService } from '../action.service'


@Component({
  selector: 'app-action-adder',
  templateUrl: './action-adder.component.html',
  styleUrls: ['./action-adder.component.css']
})
export class ActionAdderComponent implements OnInit {

	action = {} as ActionDB;






  constructor(
    private actionService: ActionService, 
	  ) {
  }

  ngOnInit(): void {

  }


  add(): void {



    this.actionService.postAction( this.action )
    .subscribe(action => {
		this.actionService.ActionServiceChanged.next("post")
		
		this.action = {} // reset fields
	    console.log("action added")
    });
  }
}
