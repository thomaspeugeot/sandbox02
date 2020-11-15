import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { UpdateStateDB } from '../updatestate-db'
import { UpdateStateService } from '../updatestate.service'


@Component({
  selector: 'app-updatestate-adder',
  templateUrl: './updatestate-adder.component.html',
  styleUrls: ['./updatestate-adder.component.css']
})
export class UpdateStateAdderComponent implements OnInit {

	updatestate = {} as UpdateStateDB;






  constructor(
    private updatestateService: UpdateStateService, 
	  ) {
  }

  ngOnInit(): void {

  }


  add(): void {



    this.updatestateService.postUpdateState( this.updatestate )
    .subscribe(updatestate => {
		this.updatestateService.UpdateStateServiceChanged.next("post")
		
		this.updatestate = {} // reset fields
	    console.log("updatestate added")
    });
  }
}
