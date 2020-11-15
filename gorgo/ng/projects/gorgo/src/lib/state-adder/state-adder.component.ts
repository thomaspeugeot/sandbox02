import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { StateDB } from '../state-db'
import { StateService } from '../state.service'


@Component({
  selector: 'app-state-adder',
  templateUrl: './state-adder.component.html',
  styleUrls: ['./state-adder.component.css']
})
export class StateAdderComponent implements OnInit {

	state = {} as StateDB;






  constructor(
    private stateService: StateService, 
	  ) {
  }

  ngOnInit(): void {

  }


  add(): void {



    this.stateService.postState( this.state )
    .subscribe(state => {
		this.stateService.StateServiceChanged.next("post")
		
		this.state = {} // reset fields
	    console.log("state added")
    });
  }
}
