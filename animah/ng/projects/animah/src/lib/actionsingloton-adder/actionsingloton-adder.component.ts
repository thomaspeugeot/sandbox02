import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { ActionSinglotonDB } from '../actionsingloton-db'
import { ActionSinglotonService } from '../actionsingloton.service'


@Component({
  selector: 'app-actionsingloton-adder',
  templateUrl: './actionsingloton-adder.component.html',
  styleUrls: ['./actionsingloton-adder.component.css']
})
export class ActionSinglotonAdderComponent implements OnInit {

	actionsingloton = {} as ActionSinglotonDB;






  constructor(
    private actionsinglotonService: ActionSinglotonService, 
	  ) {
  }

  ngOnInit(): void {

  }


  add(): void {



    this.actionsinglotonService.postActionSingloton( this.actionsingloton )
    .subscribe(actionsingloton => {
		this.actionsinglotonService.ActionSinglotonServiceChanged.next("post")
		
		this.actionsingloton = {} // reset fields
	    console.log("actionsingloton added")
    });
  }
}
