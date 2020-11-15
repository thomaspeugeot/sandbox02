import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { MachineDB } from '../machine-db'
import { MachineService } from '../machine.service'


@Component({
  selector: 'app-machine-adder',
  templateUrl: './machine-adder.component.html',
  styleUrls: ['./machine-adder.component.css']
})
export class MachineAdderComponent implements OnInit {

	machine = {} as MachineDB;




	CleanedlaundryFormControl = new FormControl(false);


  constructor(
    private machineService: MachineService, 
	  ) {
  }

  ngOnInit(): void {

  }


  add(): void {

	this.machine.Cleanedlaundry = this.CleanedlaundryFormControl.value


    this.machineService.postMachine( this.machine )
    .subscribe(machine => {
		this.machineService.MachineServiceChanged.next("post")
		
		this.machine = {} // reset fields
	    console.log("machine added")
    });
  }
}
