import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { VerticeDB } from '../vertice-db'
import { VerticeService } from '../vertice.service'


@Component({
  selector: 'app-vertice-adder',
  templateUrl: './vertice-adder.component.html',
  styleUrls: ['./vertice-adder.component.css']
})
export class VerticeAdderComponent implements OnInit {

	vertice = {} as VerticeDB;






  constructor(
    private verticeService: VerticeService, 
	  ) {
  }

  ngOnInit(): void {

  }


  add(): void {



    this.verticeService.postVertice( this.vertice )
    .subscribe(vertice => {
		this.verticeService.VerticeServiceChanged.next("post")
		
		this.vertice = {} // reset fields
	    console.log("vertice added")
    });
  }
}
