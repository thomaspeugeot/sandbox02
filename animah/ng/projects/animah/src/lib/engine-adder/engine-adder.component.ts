import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { EngineDB } from '../engine-db'
import { EngineService } from '../engine.service'


@Component({
  selector: 'app-engine-adder',
  templateUrl: './engine-adder.component.html',
  styleUrls: ['./engine-adder.component.css']
})
export class EngineAdderComponent implements OnInit {

	engine = {} as EngineDB;






  constructor(
    private engineService: EngineService, 
	  ) {
  }

  ngOnInit(): void {

  }


  add(): void {



    this.engineService.postEngine( this.engine )
    .subscribe(engine => {
		this.engineService.EngineServiceChanged.next("post")
		
		this.engine = {} // reset fields
	    console.log("engine added")
    });
  }
}
